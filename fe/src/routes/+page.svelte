<script>
	import { goto } from '$app/navigation';
	import { user } from '../stores/user';
	import { Notyf } from 'notyf';
	import { onMount } from 'svelte';

	let notyf;

	let account = $user;
	let pwd = "";

	let login = () => {
		fetch(`/api/login?name=${account}&pwd=${pwd}`)
			.then((v) => {
				if (!v.ok) {
					throw new Error("服务端异常");
				}
				return v.json();
			})
			.then((v) => {
				if (v && v.status !== 0) {
					console.log(v, decodeURIComponent(v.msg).replace(/\+/g, ' '));
					notyf.error('密码错误');
					return;
				}
				notyf.success('登录成功');
				user.set(account);
				goto('/home');
			})
			.catch((e) => {
				notyf.error(e.message);
			});
	};

	onMount(() => {
		notyf = new Notyf({
			duration: 3000,
			className: 'x-notification',
			position: { x: 'right', y: 'top' }
		});
	});
</script>

<div class="main">
	<div class="title h1">登录</div>
	<div>
		用户名: <input class="input" bind:value={account} type="text" />
	</div>
	<div>
		密码:&nbsp;&nbsp;&nbsp; <input class="input" bind:value={pwd} type="password" />
	</div>
	<a href="\register">还没账号?点我注册</a>
	<div>
		<button class="button is-primary" disabled={!account || !pwd} onclick={login}>登录</button>
	</div>
</div>

<style lang="scss">
	@import 'notyf/notyf.min.css';
	@import 'bulma/css/bulma.css';

	.input {
		width: unset;
		max-width: unset;
	}

	.main {
		display: flex;
		flex-direction: column;
		align-items: center;

		div {
			padding: 1em;
		}
	}

	a{
		margin-left: 150px;
		font-size: 13px;
		color: red;
	}
</style>
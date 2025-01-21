<script>
	import { goto } from '$app/navigation';
	import { Notyf } from 'notyf';
	import { onMount } from 'svelte';

	let notyf;
	let role = "user";  // 设置默认值为 'user'

	let account = "";
	let pwd = "";

	let login = () => {
		fetch(`/api/login`, {
			method: "POST",
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify({
				account,
				pwd,
				role
			})
		})
		.then((v) => {
			if (!v.ok) {
				throw new Error("服务端异常");
			}
			return v.json();
		})
		.then((v) => {
			if (v && v.status !== 0) {
				notyf.error(v.msg);
				return;
			}
			notyf.success('登录成功');
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
		<label>
		  <input type="radio" bind:group={role} value="admin" /> 管理员
		</label>
		<label>	
		  <input type="radio" bind:group={role} value="user" /> 普通用户
		</label>
	</div>
	<div>
		<button class="button is-primary" onclick={login}>登录</button>
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
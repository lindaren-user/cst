<script>
	import { Notyf } from 'notyf';
	import { onMount } from 'svelte';

	let notyf;

	let name = $state('');
	let pwd = $state('');

	let login = (e) => {
		fetch(`/api/login?name=${name}&pwd=${pwd}`)
			.then((v) => {
				if (!v.ok) {
					console.log(v);
					return v;
				}

				return v.json();
			})
			.then((v) => {
				if (v.status !== 0) {
					console.log(v, decodeURIComponent(v.msg).replace(/\+/g, ' '));
					notyf.error('登录失败');
					return;
				}

				notyf.success('登录成功');
			})
			.catch((e) => {
				console.log(e);
				if (e.message) {
					notyf.error(e.message);
				}
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
		用户名: <input class="input" bind:value={name} type="text" />
	</div>
	<div>
		密码: <input class="input" bind:value={pwd} type="password" />
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
</style>
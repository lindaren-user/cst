<script>
	import { goto } from '$app/navigation';
	import { Notyf } from 'notyf';
	import { onMount } from 'svelte';
	import { user } from '../../stores/user';

	let notyf;

	let inputAcc;

	let account = "";
	let pwd = "";
	let name = "";
	let mobilePhone = "";

	let checkInput;

	let register = () => {
		if(account === "" || pwd === "" || name === "" || mobilePhone === ""){
			notyf.error("请输入完整信息");
			return;
		}

		fetch(`/api/register?account=${account}&pwd=${pwd}&name=${name}&mobilePhone=${mobilePhone}`)
			.then((v) => {
				if(!v.ok){
					notyf.error("服务器错误，请稍后再试");
					return;
				}

				return v.json();
			})
			.then((v) => {
				if (v && v.status !== 0) {
					notyf.error(v.msg || "注册出错");
					return;
				}

				notyf.success("注册成功!");
				user.set(account);
				goto("/");
			})
			.catch((e) => {
				console.log(e);
				notyf.error("注册错误");
			});
	}

	onMount(() => {
		notyf = new Notyf({
			duration: 3000,
			className: 'x-notification',
			position: { x: 'right', y: 'top' }
		});

		// 检查输入框的内容
		checkInput = () => {
			fetch(`/api/checkAccount?account=${account}`)
				.then((v) => {
					if (!v.ok) {
						throw new Error("服务端异常");
					}
					return v.json();
				})
				.then((v) => {
					if (v && v.status !== 0) {
						notyf.error(v.msg || "用户名已存在");
						inputAcc.focus();
					}
				})
				.catch((e) => {
					notyf.error(e.message);
				});
		}
	});
</script>

<div class="main">
	<div class="title h1">注册</div>
	<div>
		用户名: <input class="input" bind:value={account} type="text" on:blur={checkInput} bind:this={inputAcc}/>
	</div>
	<div>
		密码:&nbsp;&nbsp;&nbsp; <input class="input" bind:value={pwd} type="password" />
	</div>
	<div>
		姓名:&nbsp;&nbsp;&nbsp; <input class="input" bind:value={name} type="text"  />
	</div>
	<div>
		电话:&nbsp;&nbsp;&nbsp; <input class="input" bind:value={mobilePhone} type="text" />
	</div>
	<div>
		<button class="button is-primary" on:click={register}>注册</button>
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

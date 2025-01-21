<script>
	import { goto } from '$app/navigation';
	import { Notyf } from 'notyf';
	import { onMount } from 'svelte';
	import { users } from '../../stores/users';

	export let data;

	let func = data.role === "admin" ? "创建" : "注册";

	let notyf;

	let inputAcc;

	let account = "";
	let pwd = "";
	let name = "";
	let mobile_phone = "";
	let role = "user";

	let checkInput;

	let register = () => {
		if(account === "" || pwd === "" || name === "" || mobile_phone === ""){
			notyf.error("请输入完整信息");
			return;
		}

		fetch(`/api/register`, {
			method: "POST",
			headers: {
          		'Content-Type': 'application/json',
			},
			body: JSON.stringify({
				account,
				pwd,
				name,
				mobilePhone: mobile_phone,
				role
			}),
		})
		.then((v) => {
			if(!v.ok){
				throw new Error("服务器错误");
			}
			return v.json();
		})
		.then((v) => {
			if (v && v.status !== 0) {
				notyf.error(v.msg);
				return;
			}
			if(data.role === 'admin') {
				let usersTemp = $users;
				usersTemp.unshift({
					id: v.id,
					account,
					name,
					mobile_phone
				});
				users.set(usersTemp);
				goto("/home/accountAdmin");
			}
			else goto("/");
			notyf.success(`注册成功!`);
		})
		.catch((e) => {
			notyf.error(e.message);
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
						notyf.error(v.msg);
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
	<div class="title h1">{func}</div>
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
		电话:&nbsp;&nbsp;&nbsp; <input class="input" bind:value={mobile_phone} type="text" />
	</div>
	{#if data.role === 'admin'}
	<div>	
		<label>
		  <input type="radio" bind:group={role} value="admin" /> 管理员
		</label>
		<label>	
		  <input type="radio" bind:group={role} value="user" /> 普通用户
		</label>
	</div>
	{:else}
	<p>注意只是普通用户</p>
	{/if}
	<div>
		<button class="button is-primary" on:click={register}>{func}</button>
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

	p{
		margin-left: 150px;
		font-size: 13px;
		color: red;
	}
</style>

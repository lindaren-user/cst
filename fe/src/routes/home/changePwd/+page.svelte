<script>
	import { goto } from '$app/navigation';
	import { Notyf } from 'notyf';
	import { onMount } from 'svelte';

	let notyf;

	let oldPwd = "";
	let newPwd = "";
    let secPwd = "";

    let isNotMatched = false;

    $: {
        if(newPwd && secPwd) {
            if(newPwd !== secPwd) {
                isNotMatched = true;
            } else {
                isNotMatched = false;
            }
        }
        else if(!newPwd && !secPwd){
            isNotMatched = false;
        }
    }

    let changePwd = () => {
        if(isNotMatched) {
            notyf.error("注意前后密码不匹配");
            return;
        }

        if(!oldPwd || !newPwd) {
            notyf.error("请输入完整的密码");
            return;
        }
        
        fetch("/api/changePwd", {
            method: "PUT",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                oldPwd, newPwd
            })
        })
        .then((v) => {
            if(!v.ok){
                throw new Error("服务端异常");
            }
            return v.json();
        })
        .then((v) => {
            if(v && v.status !== 0){
                notyf.error(v.msg || "原密码错误");
                return;
            }
            notyf.success("密码修改成功");
            goto("/");
        })
        .catch((e) => {
            notyf.error(e.message || "服务端异常");
        });
    }

	onMount(() => {
		notyf = new Notyf({
			duration: 3000,
			className: 'x-notification',
			position: { x: 'right', y: 'top' }
		});
	});
</script>

<div class="main">
	<div class="title h1">修改密码</div>
	<div>
		原密码:&nbsp;&nbsp;&nbsp; <input class="input" bind:value={oldPwd} type="password" />
	</div>
	<div>
		新密码:&nbsp;&nbsp;&nbsp; <input class="input" bind:value={newPwd} type="password"  />
	</div>
    <div>
		再次输入<input class="input" bind:value={secPwd} type="password"  />
	</div>
    {#if isNotMatched}
        <p>前后新密码不一致</p>
    {/if}
	<div>
		<button class="button is-primary" on:click={changePwd}>修改</button>
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

<script>
	import { page } from '$app/stores'; // 导入 page store
	import { onMount } from "svelte";
	import { users } from "../../../stores/users";
	import { Notyf } from 'notyf';
    import { goto } from "$app/navigation";

    let id;
    let notyf;
    let userT = {
        account: "",
        name: "",
        pwd: ""
    };
    let pwd = "";
    let modify;

    onMount(() => {
        notyf = new Notyf({
            duration: 3000,
            className: 'x-notification',
            position: { x: 'right', y: 'top' }
        });

        id = $page.url.searchParams.get('id');
        let modifyingUser = $users.find(user => user.id == id);

        userT = {
            account: modifyingUser.account,
            name: modifyingUser.name,
            pwd: ""
        }

        modify = () => {
            fetch(`/api/modifyInfo/${id}`, { 
                method: "PUT",
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    account: userT.account,
                    pwd
                })
            })
            .then((v) => {
                if (!v.ok) {
                    throw new Error("服务器出错");
                }
                return v.json();
            })
            .then((v) => {
                if (v && v.status !== 0) {
                    notyf.error(v.msg);
                    return;
                }
                users.set($users.map(u => u.id === id ? {
                    ...u,
                    account: userT.account,
                    pwd
                }: u));
                notyf.success(v.msg);
            })
            .catch((e) => {
                notyf.error(e.message);
            });
        }
    })
</script>

<div class="main">
	<div class="title h1">修改{userT.name}的相关信息</div>
	<div>
		用户名: <input class="input" bind:value={userT.account} type="text"/>
	</div>
	<div>
		密码:&nbsp;&nbsp;&nbsp; <input class="input" bind:value={pwd} type="password" />
	</div>
	<div>
		<button class="button is-primary" on:click={modify}>修改</button>
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

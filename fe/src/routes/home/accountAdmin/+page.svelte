<script>
	import { Notyf } from "notyf";
	import { onMount } from "svelte";
	import { users } from "../../../stores/users";
	import { goto } from "$app/navigation";

    let notyf;
    let deleteUser;
    
    onMount(() => {
        notyf = new Notyf({
            duration: 3000,
            className: 'x-notification',
            position: { x: 'right', y: 'top' }
        });

        fetch("/api/getAllAccounts")
        .then((v) => {
            if(!v.ok) {
                throw new Error("服务器异常")
            }
            return v.json();
        })
        .then((v) => {
            if(v && v.status !== 0) {
                notyf.error(v.msg);
                return;
            }
            users.set(v.users)
        })
        .catch((e) => {
            notyf.error(e.message)
        });

        deleteUser = (id) => {
            fetch(`/api/deleteAccount/${id}`)
            .then((v) => {
                if(!v.ok) {
                    throw new Error("服务器异常")
                }
                return v.json();
            })
            .then((v) => {
                if(v && v.status !== 0) {
                    notyf.error(v.msg);
                    return;
                }
                notyf.success(v.msg);  
                users.set($users.filter((user) => user.id !== id));
            })
            .catch((e) => {
                notyf.error(e.message)
            });
        }
    })
</script>

<button on:click={() => { goto("/register") }} class="create-btn">创建账号</button>    
<h2>所有用户账号</h2>
{#if !$users?.length}
    <p>暂无用户</p>
{:else}
    <ul>
        {#each $users as user (user.id)}
            <li class="user-item" id={user.id}>
                <div class="user-title">账号：{user.account}</div>
                <div class="user-author">名字：{user.name}</div>
                <div class="user-mobilePhone">电话：{user.mobile_phone}</div>
                <div class="user-actions">
                    <button on:click={() => { goto(`/home/modifyInfoAdmin?id=${user.id}`) }} class="edit-btn">修改</button>
                    <button on:click={() => deleteUser(user.id)} class="delete-btn">删除</button>
                </div>
            </li>
        {/each}
    </ul>
{/if}

<style>
    @import 'notyf/notyf.min.css';
    @import 'bulma/css/bulma.css';

    /* 用户列表样式 */
    .user-item {
        border: 1px solid #ccc;
        padding: 16px;
        margin: 10px 0;
        border-radius: 8px;
        background-color: #f9f9f9;
        position: relative;
        transition: background-color 0.3s ease;
    }

    .user-item:hover {
        background-color: #f1f1f1;  /* 鼠标悬停时的背景色变化 */
    }

    .user-title,
    .user-author,
    .user-mobilePhone {
        font-size: 18px;
        color: #333;
        margin-bottom: 8px;
    }

    .user-title {
        font-weight: bold;
    }

    .user-author {
        font-weight: normal;
    }

    .user-actions {
        position: absolute;
        top: 16px;
        right: 16px;
        display: flex;
        gap: 8px;
    }

    /* 按钮样式 */
    .create-btn,
    .edit-btn,
    .delete-btn {
        margin-left: 8px;
        padding: 8px 16px;
        font-size: 14px;
        border: none;
        border-radius: 6px;
        cursor: pointer;
        color: white;
        transition: background-color 0.3s ease;
    }

    /* 创建按钮样式 */
    .create-btn {
        background-color: green;
    }

    .create-btn:hover {
        background-color: #45a049; /* 悬停时颜色变深 */
    }

    /* 编辑按钮样式 */
    .edit-btn {
        background-color: #4CAF50; /* Green */
    }

    .edit-btn:hover {
        background-color: #45a049;
    }

    /* 删除按钮样式 */
    .delete-btn {
        background-color: #f44336; /* Red */
    }

    .delete-btn:hover {
        background-color: #e53935; /* 悬停时颜色变深 */
    }

    /* 用户列表容器 */
    ul {
        list-style-type: none;
        padding: 0;
    }

    /* 如果没有用户时的提示 */
    p {
        font-size: 16px;
        color: #777;
        text-align: center;
    }

</style>

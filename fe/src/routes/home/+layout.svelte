<script>
    import { Notyf } from 'notyf';
    import { onMount } from 'svelte';
    import { goto } from '$app/navigation';
    import { blogs } from '../../stores/blogs';
    import { users } from '../../stores/users';
  
    export let data;
  
    let notyf;

    let logout = () => {
      users.set([]);
      blogs.set([]);
  
      fetch('/api/logout')
      .then((v) => {
        if (!v.ok) {
          throw new Error("服务端异常");
        }
        return v.json();
      })
      .then((v) => {
        if (v && v.status !== 0) {
          notyf.error("退出错误");
          return;
        }
        notyf.success("退出成功");
        goto("/");
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
  
<div class="navbar">
    <div class="navbar-left">
        <a href="/home" class="navbar-item">首页</a>
        {#if data.role ===  'user'}
            <a href="/home/blogs" class="navbar-item">我的博客</a>
        {:else if data.role === 'admin'} 
            <a href="/home/blogs" class="navbar-item">博客管理</a>
            <a href="/home/accountAdmin" class="navbar-item">账号管理</a>
        {/if}
    </div>
    <div class="navbar-right">
        {#if data.account}
            <p class="navbar-account">你好！{data.account} {data.role}</p>
            <div class="dropdown">
                <button class="dropdown-button" on:click={() => {goto("/home/changePwd")}}>
                <i class="fa-solid fa-gear" ></i>修改密码
                </button>
            </div>
            <button class="navbar-item navbar-logout" on:click={logout}>退出登录</button>
        {:else}
            <p class="navbar-account">未登录</p>  
            <a href="/" class="navbar-item navbar-logout">前往登录</a>
        {/if}
    </div>
</div>
<div class="main">
    {#if data.account}
        <slot></slot>
    {:else}
        请先登录
    {/if}
</div>

<style>
    @import 'notyf/notyf.min.css';
    @import 'bulma/css/bulma.css';
    /* 固定在顶部的导航栏样式 */
    .navbar {
        position: fixed;
        top: 0;
        left: 0;
        right: 0;
        background-color: #333;
        color: white;
        padding: 1rem;
        z-index: 1000;
        display: flex;
        justify-content: space-between;
        align-items: center;
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    }

    /* 左侧导航栏内容 */
    .navbar-left {
        display: flex;
        align-items: center;
    }

    /* 右侧导航栏内容 */
    .navbar-right {
        display: flex;
        align-items: center;
        gap: 20px; /* 控制右侧内容的间距 */
    }

    /* 导航栏项的样式 */
    .navbar-item {
        color: white;
        text-decoration: none;
        padding: 0.5rem 1rem;
        font-size: 1.1rem;
        transition: background-color 0.3s ease;
    }

    /* 鼠标悬停时改变背景颜色 */
    .navbar-item:hover {
        background-color: #555;
    }

    /* 账号显示样式 */
    .navbar-account {
        font-size: 1.1rem;
        font-weight: bold;
        color: #ffcc00;
    }

    /* 主体内容区域的样式 */
    .main {
        margin-top: 80px;
        padding: 20px;
        font-size: 1.2rem;
    }

    /* 设置按钮样式 */
    .dropdown {
        position: relative;
    }

    .dropdown-button {
        background: none;
        color: white;
        border: none;
        font-size: 1.1rem;
        cursor: pointer;
        display: flex;
        align-items: center;
        gap: 5px;
    }

    .dropdown-button:hover {
        color: #ffcc00;
    }
</style>

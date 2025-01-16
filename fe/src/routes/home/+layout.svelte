<script>
  import { user } from '../../stores/user';
  import { Notyf } from 'notyf';
  import { onMount } from 'svelte';
  import { goto } from '$app/navigation';

  let notyf;
  let showDropdown = false; // 控制下拉菜单的显示状态

  let dropdownButton; // 获取设置按钮的 DOM 元素
  let menuTop = 0; // 菜单的垂直位置
  let menuLeft = 0; // 菜单的水平位置

  const toggleDropdown = () => {
    console.log("显示")
    showDropdown = !showDropdown;

    if (showDropdown && dropdownButton) {
      const rect = dropdownButton.getBoundingClientRect();
      menuTop = rect.bottom + 20; // 按钮底部的位置
      menuLeft = rect.left;  // 按钮左侧的位置

      console.log("yp")
    }
};
 
  let logout = () => {
    user.set(null);

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
    <a href="/home/blogs" class="navbar-item">我的博客</a>
  </div>
  <div class="navbar-right">
    <p class="navbar-account">你好！{$user}</p>
    <div class="dropdown">
      <div class="dropdown">
        <button class="dropdown-button" on:click={toggleDropdown} bind:this={dropdownButton}>
          <i class="fa-solid fa-gear"></i> 设置
        </button>
        {#if showDropdown}
          <div class="dropdown-menu" style="top: {menuTop}px; left: {menuLeft}px;">
            
            <a href="#" class="dropdown-item">其他</a>
          </div>
        {/if}
      </div>
    </div>
    <a href="#" class="navbar-item navbar-logout" on:click={logout}>退出登录</a>
  </div>
</div>
<div class="main">
  <slot></slot>
  <a href="/changePwd" class="dropdown-item">修改密码</a>
</div>

<style>
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

  /* 下拉菜单样式 */
  /* 下拉菜单样式 */
.dropdown-menu {
  position: fixed; /* 固定菜单位置 */
  background-color: #444;
  border: 1px solid #555;
  border-radius: 5px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  z-index: 1000;
  width: 150px;
}

.dropdown-item {
  color: white;
  text-decoration: none;
  display: block;
  padding: 10px;
  font-size: 1rem;
  transition: background-color 0.3s ease;
}

.dropdown-item:hover {
  background-color: #555;
}

</style>

<script>
  import { user } from '../../stores/user'
</script>

<div class="navbar">
  <a href="/home" class="navbar-item">首页</a>
  <a href="/home/blogs" class="navbar-item">我的博客</a>
  <p class="navbar-account">你好！{$user}</p>
  <a href="/" class="navbar-item navbar-logout">退出登录</a>
</div>  

<div class="main">
  <slot></slot>
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
    z-index: 1000; /* 保证导航栏在最上层 */
    display: flex;
    justify-content: space-between; /* 调整项的对齐方式 */
    align-items: center; /* 垂直居中 */
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1); /* 给导航栏增加阴影效果 */
  }

  /* 导航栏项的样式 */
  .navbar-item {
    color: white;
    text-decoration: none;
    padding: 0.5rem 1rem;
    font-size: 1.1rem; /* 增加字体大小 */
    transition: background-color 0.3s ease;
  }

  /* 鼠标悬停时改变背景颜色 */
  .navbar-item:hover {
    background-color: #555;
  }

  /* 账号显示样式 */
  .navbar-account {
    font-size: 1.1rem; /* 增加字体大小 */
    font-weight: bold; /* 加粗显示账号 */
    color: #ffcc00; /* 设置特殊颜色 */
  }

  /* 退出登录按钮样式 */
  .navbar-logout {
    margin-left: 10px; /* 控制"退出登录"与账号之间的间距 */
  }

  /* 主体内容区域的样式 */
  .main {
    margin-top: 80px; /* 预留导航栏空间，避免遮挡 */
    padding: 20px;
    font-size: 1.2rem; /* 设置页面的字体大小 */
  }
</style>

<script>
</script>

<div class="main">
    <a href="/home/blogs/create">创建文章</a>
    <a href="/home/blogs/edit">修改文章</a>
    <a href="/home/blogs/show">我的文章</a>
</div>

<style>   
    /* 整体容器样式 */
    .main {
        display: flex; /* 使用 flexbox 排列子元素 */
        justify-content: center; /* 水平居中对齐 */
        align-items: center; /* 垂直居中对齐 */
        flex-direction: column; /* 垂直排列 */
        height: 100vh; /* 使容器充满整个视口高度 */
        background-color: #f9f9f9; /* 背景色 */
        font-family: 'Arial', sans-serif; /* 字体 */
    }

    /* 链接样式 */
    .main a {
        display: inline-block; /* 将链接转换为块级元素，使得能够设置宽高等 */
        text-decoration: none; /* 去掉默认的下划线 */
        color: #fff; /* 文字颜色 */
        background-color: #007bff; /* 背景色 */
        padding: 15px 30px; /* 内边距，增加点击区域 */
        margin: 10px 0; /* 上下间距 */
        border-radius: 5px; /* 圆角边框 */
        font-size: 16px; /* 字体大小 */
        transition: all 0.3s ease; /* 添加过渡效果 */
    }

    /* 链接悬停效果 */
    .main a:hover {
        background-color: #0056b3; /* 悬停时改变背景色 */
        transform: translateY(-5px); /* 向上移动 */
    }

    /* 链接聚焦时的样式 */
    .main a:focus {
        outline: none; /* 去掉焦点时的外边框 */
        box-shadow: 0 0 5px rgba(0, 123, 255, 0.5); /* 聚焦时的阴影效果 */
    }

    /* 为响应式设计提供支持 */
    @media (max-width: 600px) {
        .main a {
            padding: 10px 20px; /* 在小屏幕上减少内边距 */
            font-size: 14px; /* 减小字体 */
        }
    }
</style>

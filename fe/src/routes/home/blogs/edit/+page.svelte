<script>
    import { goto } from '$app/navigation';
    import { Notyf } from 'notyf';
    import RichTextEditor from '$lib/components/RichTextEditor.svelte'; // 导入富文本编辑器组件
	import { onMount } from 'svelte';
    import { blogs } from "../../../../stores/blogs"
    import { page } from '$app/stores'; // 导入 page store
  
    // 从 page store 中获取 URL 查询参数
    let id;
    let blog;
    let article = {
          title: '',
          content: '',
        };
    let notyf;

    onMount(() => {
        notyf = new Notyf({
            duration: 3000,
            className: 'x-notification',
            position: { x: 'right', y: 'top' }
        });

        id = $page.url.searchParams.get('id'); // 得到的是字符串

        if (id) {
            blog = $blogs.find(b => b.id == id);
        }

        if(blog) {
            article.title = blog.title;
            article.content = blog.content;
        }
    });
  
    // 提交表单
    let edit = () => {
      fetch(`/api/editBlog/${id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(article),
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
        notyf.success(v.msg);

        blogs.set($blogs.map(blog => blog.id == id ? {
            id: id,
            ...article
        } : blog));

        goto(`/home/blogs/show?id=${id}`);
      })
      .catch((e) => {
        notyf.error(e.message);
      });
    };
</script>
  
<h2>编辑文章</h2>
<form on:submit|preventDefault={edit}>
<label for="title">标题:</label>
<input type="text" id="title" bind:value={article.title} required />

<!-- 使用 RichTextEditor 组件 -->
<RichTextEditor bind:value={article.content} onChange={(newContent) => article.content = newContent} />

<button type="submit">完成</button>
</form>
  
<style>
    @import 'notyf/notyf.min.css';
	@import 'bulma/css/bulma.css';
    input[type="text"] {
        width: 100%;
        padding: 8px;
        margin-bottom: 10px;
        border: 1px solid #ccc;
        border-radius: 4px;
    }

    button {
        padding: 10px 20px;
        background-color: #4CAF50;
        color: white;
        border: none;
        border-radius: 5px;
        cursor: pointer;
        margin-top: 10px;
    }

    button:hover {
        background-color: #45a049;
    }
</style>
  
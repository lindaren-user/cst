<script>
    import { goto } from '$app/navigation';
    import { Notyf } from 'notyf';
    import RichTextEditor from '$lib/components/RichTextEditor.svelte'; // 导入富文本编辑器组件
	  import { onMount } from 'svelte';
  
    let title = "";
    let content = ''; // 用于存储编辑器内容
    let notyf;
  
    // 提交表单
    let create = () => {
      fetch('/api/createBlogs', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          title: title,
          content: content,
        }),
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
        goto("/home/blogs");
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
    })
</script>
  
<h2>创建文章</h2>
<form on:submit|preventDefault={create}>
<label for="title">标题:</label>
<input type="text" id="title" bind:value={title} required />

<!-- 使用 RichTextEditor 组件 -->
<RichTextEditor bind:value={content} onChange={(newContent) => content = newContent} />

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
  
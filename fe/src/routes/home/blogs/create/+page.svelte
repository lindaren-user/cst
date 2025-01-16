<script>
    import { onMount } from 'svelte';
    import { Notyf } from 'notyf';
	import { goto } from '$app/navigation';

    let Quill;
    let notyf;
    let title = "";
    let formData = {
      content: '' // 初始为空
    };

    // 工具栏选项
    const toolbarOptions = [
      ['bold', 'italic', 'underline', 'strike'],
      ['blockquote', 'code-block'],
      [{ 'list': 'ordered' }, { 'list': 'bullet' }],
      [{ 'size': ['small', false, 'large', 'huge'] }],
      [{ 'header': [1, 2, 3, 4, 5, 6, false] }],
      [{ 'color': [] }, { 'background': [] }],
      [{ 'font': ['SimSun', 'SimHei', 'Microsoft-YaHei', 'KaiTi', 'FangSong', 'Arial'] }],
      [{ 'align': [] }],
      ['clean'],
      ['link', 'image'],
    ];

    // 编辑器初始化
    onMount(async () => {
      if (typeof window !== 'undefined') {
        // 动态导入 Quill 编辑器，仅在浏览器中加载
        const { default: QuillInstance } = await import('quill');
        Quill = QuillInstance;

        const quill = new Quill('#editor', {
          theme: 'snow',
          modules: {
            toolbar: {
              container: toolbarOptions,
            },
          },
        });

        quill.on('text-change', () => {
          formData.content = quill.root.innerHTML;
        });
      }

      notyf = new Notyf({
        duration: 3000,
        className: 'x-notification',
        position: { x: 'right', y: 'top' }
      });
    });

    // 提交表单
    let create = () => {
        fetch('/api/createBlogs', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                title: title,
                content: formData.content,
            }),
        })
        .then((v) => {
            if(!v.ok) {
                throw new Error("服务器出错");
            }
            return v.json();
        })
        .then((v) => {
            if(v && v.status !== 0) {
                notyf.error(v.msg);
                return;
            }
            notyf.success(v.msg);
            goto("/home/blogs");
        })
        .catch((e) => {
            notyf.error(e.msg);
        });
    }
</script>

<h2>创建文章</h2>
<form on:submit|preventDefault={create}>
  <label for="title">标题:</label>
  <input type="text" id="title" bind:value={title} required />

  <!-- Quill 编辑器容器 -->
  <div id="editor"></div>

  <button type="submit">完成</button>
</form>

<style>
  /* 加载 Quill 编辑器默认样式 */
  @import 'quill/dist/quill.snow.css'; /* 确保 Quill 样式被正确引入 */

  #editor {
    height: 300px;
    border: 1px solid #ccc;
    border-radius: 5px;
    padding: 10px;
    background-color: #fff; /* 可选的背景颜色 */
  }

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

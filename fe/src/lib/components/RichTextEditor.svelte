<script>
  import { onMount } from 'svelte';
  
  export let value = '';
  export let onChange = () => {}; // 用于监听内容变化

  let Quill;

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

  onMount(async () => {
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
        value = quill.root.innerHTML;
        onChange(value);
      });

      quill.root.innerHTML = value; // 设置初始值
  });
</script>

<div id="editor"></div>

<style>
  @import 'quill/dist/quill.snow.css'; 

  #editor {
    height: 300px;
    border: 1px solid #ccc;
    border-radius: 5px;
    padding: 10px;
    background-color: #fff;
  }
</style>

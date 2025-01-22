<script>
    import { blogs } from "../../../../stores/blogs"
    import { page } from '$app/stores'; // 导入 page store
	import { onMount } from "svelte";
	import { goto } from "$app/navigation";

    export let data;
    // 从 page store 中获取 URL 查询参数
    let id;
    let blog;
    onMount(() => {
        id = $page.url.searchParams.get('id'); // 得到的是字符串

        if (id) {
            blog = $blogs.find(b => b.id == id);
        }
  });
</script>

<button on:click={ () => { goto("/home/blogs")}} class="back-btn">返回</button>
<br><br>
{#if blog}
    <div class="blog-title">{blog.title}</div>
    {#if data.role === 'admin'}
        <div class="blog-author">{blog.author}</div>
    {/if}
    <div class="blog-content">{@html blog.content}</div>
{:else}
    没有文章
{/if}

<style>
    .back-btn {
        margin-left: 8px;
        padding: 6px 12px;
        font-size: 14px;
        border: none;
        border-radius: 4px;
        cursor: pointer;
        color: white;
        background-color: green;
    }

    .blog-title {
        font-size: 30px;
        font-weight: bold;
        color: #333;
        margin-bottom: 10px;
    }

    .blog-author {
        font-size: 18px;
        font-weight: bold;
        color: #333;
        margin-bottom: 10px;
    }

    .blog-content {
        font-size: 14px;
        color: #555;
        line-height: 1.6;
    }
</style>
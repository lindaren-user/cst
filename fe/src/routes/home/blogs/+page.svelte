<script>
	import { onMount } from "svelte";
    import { Notyf } from "notyf";
    import { goto } from '$app/navigation';
    import { blogs } from '../../../stores/blogs'

    export let data;

    let notyf;

    onMount(() => {
        notyf = new Notyf({
            duration: 3000,
            className: 'x-notification',
            position: { x: 'right', y: 'top' }
		});

        fetch("/api/showAllBlogs")
        .then((v) => {
            if(!v.ok) {
                throw new Error("服务端异常");
            }
            return v.json(); // 已经将 blogs 解析
        })
        .then((v) => {
            if(v && v.status !== 0) {
                notyf.error(v.msg);
                return;
            }
            blogs.set(v.blogs);
        })
        .catch((e) => {
            notyf.error(e.message);
        });
    });

    let deleteBlog = (id) => {
        fetch(`/api/deleteBlog/${id}`, { method: "DELETE" })
            .then((v) => {
                if(!v.ok) {
                    throw new Error("服务器错误");
                }
                return v.json();
            })
            .then((v) => {
                if(v && v.status !== 0) {
                    notyf.error(v.msg);
                    return;
                }
                notyf.success(v.msg);  
                blogs.set($blogs.filter(blog => blog.id !== id));
            })  
            .catch((e) => {
                notyf.error(e.message);
            });            
    };
</script>

{#if data.role === 'user'}
    <h2>我的文章</h2>
    <button on:click={() => { goto("/home/blogs/create") }} class="create-btn">点击创作</button>    
{:else}
    <h2>所有文章</h2>
{/if}
{#if !$blogs?.length}
    <p>暂无文章</p>
{:else}
    <ul>
        {#each $blogs as blog (blog.id)}
            <li class="blog-item" id={blog.id}>
                <div class="blog-title">标题：{blog.title}</div>
                {#if data.role === 'admin'}
                    <div class="blog-author">作者：{blog.author}</div>
                {/if}
                <div class="blog-content">{@html blog.content}</div>
                <div class="blog-actions">
                    <button on:click={() => { goto(`/home/blogs/show?id=${blog.id}`) }} class="show-btn">阅读</button>
                    {#if data.role === 'user'}
                        <button on:click={() => { goto(`/home/blogs/edit?id=${blog.id}`) }} class="edit-btn">编辑</button>
                    {/if}
                    <button on:click={() => deleteBlog(blog.id)} class="delete-btn">删除</button>
                </div>
            </li>
        {/each}
    </ul>
{/if}

<style>
    @import 'notyf/notyf.min.css';
	@import 'bulma/css/bulma.css';
    .blog-item {
        border: 1px solid #ccc;
        padding: 16px;
        margin: 10px 0;
        border-radius: 8px;
        background-color: #f9f9f9;
        position: relative;
    }

    .blog-title {
        font-size: 18px;
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
        font-size: 18px;
        color: #555;    
        line-height: 1.6;
        height: 50px;
        overflow: hidden;
    }

    .blog-actions {
        position: absolute;
        top: 16px;
        right: 16px;
    }

    .create-btn,
    .show-btn,
    .edit-btn,
    .delete-btn {
        margin-left: 8px;
        padding: 6px 12px;
        font-size: 14px;
        border: none;
        border-radius: 4px;
        cursor: pointer;
        color: white;
    }

    .create-btn {
        background-color: green;
    }

    .show-btn {
        background-color: #4CAF50;
    }   

    .edit-btn {
        background-color: blue;
    }
    
    .delete-btn {
        background-color: #f44336;
    }
</style>

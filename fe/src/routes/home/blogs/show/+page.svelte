<script>
	import { onMount } from "svelte";
    import { Notyf } from "notyf"

    let notyf;
    let blogs = [];

    onMount(() => {
        fetch("/api/showBlogs")
            .then((v) => {
                if(!v.ok) {
                    throw new Error("服务端异常");
                }
                return v.json();
            })
            .then((v) => {
                if(v && v.status !== 0) {
                    notyf.error(v.msg);
                    return;
                }
                blogs = v.blogs;
            })
            .catch((e) => {
                notyf.error(e.message);
            });

        notyf = new Notyf({
            duration: 3000,
            className: 'x-notification',
            position: { x: 'right', y: 'top' }
		});
    })
</script>

<h2>我的文章</h2>
{#if blogs.length === 0}
    <p>暂无文章</p>
{:else}
    <ul>
        {#each blogs as blog}
            <li>
                <div>{blog.title}</div>
                <div>{@html blog.content}</div>
            </li>
        {/each}
    </ul>
{/if}

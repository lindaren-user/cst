import adapter from '@sveltejs/adapter-static';
import preprocess from 'svelte-preprocess';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// 请查阅 https://github.com/sveltejs/svelte-preprocess
	// 获取更多关于预处理器的信息
	preprocess: preprocess(),

	onwarn: (warning, handler) => {
		if (warning.code.startsWith('a11y-')) {
			return;
		}
		handler(warning);
	},

	kit: {
		adapter: adapter({
			// 默认选项如下。在某些平台上这些选项会自动设置 —— 请查看下文
			pages: 'build',
			assets: 'build',
			fallback: 'index.html',
			precompress: true,
			strict: true,
		}),
	}
};

export default config;
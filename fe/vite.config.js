import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import path from 'path';

export default defineConfig({
	resolve: {
		alias: {
			'@src': path.resolve('./src')
		}
	},
	plugins: [sveltekit()],

	build: {
		sourcemap: true,
	},

	server: {
		host: 'localhost',
		port: 8080,
		open: false, // 禁止自动打开浏览器窗口
		proxy: {
			'/api': 'http://localhost:9090/',
		},
	}
});
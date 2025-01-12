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
		open: true,
		proxy: {
			'/api': 'http://localhost:6616/',
		},
	}
});
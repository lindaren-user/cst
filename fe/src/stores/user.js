import { writable } from 'svelte/store';

export const user = writable(null); // 创建一个 store，默认值为 null

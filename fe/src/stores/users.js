import { writable } from 'svelte/store';

// 创建一个持久化的 writable store
function persistentWritable(key, initialValue) {
    // 判断是否在浏览器端
    if (typeof window !== 'undefined' && window.localStorage) {
        const storedValue = localStorage.getItem(key);
        const value = storedValue ? JSON.parse(storedValue) : initialValue;

        const store = writable(value);

        // 订阅 store 变化并更新 localStorage
        store.subscribe(($store) => {
            localStorage.setItem(key, JSON.stringify($store));
        });

        return store;
    } else {
        // 如果在服务器端，返回默认值
        return writable(initialValue);
    }
}

// 使用 persistentWritable 创建一个持久化的 store
export const users = persistentWritable('users', []);

import { writable } from 'svelte/store';

function persistentWritable(key, initialValue) {

    if (typeof window !== 'undefined' && window.localStorage) {
        const storedValue = localStorage.getItem(key);
        const value = storedValue ? JSON.parse(storedValue) : initialValue;

        const store = writable(value);

        store.subscribe(($store) => {
            localStorage.setItem(key, JSON.stringify($store));
        });

        return store;
    } else {
        return writable(initialValue);
    }
}

export const blogs = persistentWritable('blogs', []);

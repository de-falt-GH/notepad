import { writable } from 'svelte/store'
import { browser } from '$app/environment'

export function createLocalStorageStore<T>(key: string, defaultValue: T) {
	const { subscribe, set } = writable(defaultValue, setInitial => {
		if (browser) {
			const item = localStorage.getItem(key)
			if (item) {
				setInitial(JSON.parse(item))
			}
		}
	})

	return {
		subscribe,

		set(value: T) {
			if (browser) {
				localStorage.setItem(key, JSON.stringify(value))
			}

			set(value)
		},
	}
}

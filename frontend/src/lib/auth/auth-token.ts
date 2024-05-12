import { createLocalStorageStore } from '../utils'

export const authToken = createLocalStorageStore<string | null>(
	'auth.access',
	null,
)

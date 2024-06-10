import { derived } from 'svelte/store'
import { createLocalStorageStore } from '../utils'

export const authToken = createLocalStorageStore<string | null>(
	'auth.access',
	null,
)

export const authJwtData = derived(authToken, $token =>
	$token !== null ? (parseJwt($token) as { exp: number; uid: number }) : null,
)

export const currentUserId = derived(authJwtData, $jwt => $jwt?.uid ?? null)

function parseJwt(token: string) {
	const base64Url = token.split('.')[1]
	const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/')

	const jsonPayload = decodeURIComponent(
		atob(base64)
			.split('')
			.map(c => '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2))
			.join(''),
	)

	return JSON.parse(jsonPayload)
}

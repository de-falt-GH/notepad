import { get } from 'svelte/store'
import axios from 'axios'
import { authToken } from '../auth'

declare const __BACKEND_PORT__: string

const apiBaseUrl = new URL(location.origin)
apiBaseUrl.port = __BACKEND_PORT__

const axiosInstance = axios.create({ baseURL: apiBaseUrl.toString() })

axiosInstance.interceptors.request.use(request => {
	const currentToken = get(authToken)

	if (currentToken) {
		request.headers['Authorization'] = currentToken
	}

	return request
})

type RegisterParameters = {
	login: string
	password: string
	email: string
	name: string
	info: string
}

type Profile = {
	login: string
	email: string
	name: string
	info: string
}

export type NoteFeedItem = {
	id: number
	name: string
	author_name: string // TODO: лень под camelCase переделать
	updated: string
}

type FetchNotesParameters = {
	search?: string
	limit?: number
	skip?: number
}

export const apiClient = {
	async register(params: RegisterParameters): Promise<string | null> {
		try {
			const {
				data: { token },
			} = await axiosInstance.post('/register', params)
			return token
		} catch {
			return null
		}
	},

	async authenticate(
		login: string,
		password: string,
	): Promise<string | null> {
		try {
			const {
				data: { token },
			} = await axiosInstance.post('/login', { login, password })

			return token
		} catch {
			return null
		}
	},

	async fetchProfile(): Promise<Profile> {
		const { data: profile } = await axiosInstance.get('/profile')
		return profile
	},

	async fetchPublicNotes({
		search = '',
		limit = 20,
		skip = 0,
	}: FetchNotesParameters = {}): Promise<{ notes: NoteFeedItem[] }> {
		const {
			data: { notes },
		} = await axiosInstance.get('/notes/public', {
			params: { search, limit, skip },
		})
		return { notes }
	},

	async fetchPrivateNotes({
		search = '',
		limit = 20,
		skip = 0,
	}: FetchNotesParameters = {}): Promise<{ notes: NoteFeedItem[] }> {
		const {
			data: { notes },
		} = await axiosInstance.get('/notes/private', {
			params: { search, limit, skip },
		})
		return { notes }
	},
}

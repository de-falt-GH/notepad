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

export const apiClient = {
	async register(params: RegisterParameters) {
		console.log(await axiosInstance.post('/register', params))
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
}

import { sveltekit } from '@sveltejs/kit/vite'
import { type UserConfigFn, defineConfig, loadEnv } from 'vite'

const config: UserConfigFn = ({ mode }) => {
	const { BACKEND_PORT: backendPortEnvFile } = loadEnv(
		mode,
		process.cwd() + '/../build',
		'',
	)
	const { BACKEND_PORT: backendPortFromProcess } = process.env

	const backendPort = backendPortEnvFile ?? backendPortFromProcess

	console.log({ backendPort })

	return defineConfig({
		plugins: [sveltekit()],

		define: {
			__BACKEND_PORT__: backendPort,
		},
	})
}

export default config

<script>
	import { Button } from '$lib/components/ui/button'
	import { authToken } from '$lib/auth'
	import { apiClient } from '$lib/api'

	function onRegister() {
		apiClient
			.register({
				login: 'picalines',
				password: '12345',
				email: 'f@off.com',
				name: 'petya',
				info: 'not much',
			})
			.then(console.log)
	}

	function onLogin() {
		apiClient.authenticate('123', '123').then(console.log)
	}
</script>

Auth token: {$authToken}

<Button variant="outline" on:click={onRegister}>Register</Button>
<Button variant="outline" on:click={onLogin}>Login</Button>

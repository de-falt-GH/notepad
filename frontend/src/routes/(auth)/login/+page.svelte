<script lang="ts">
	import * as Card from '@/lib/components/ui/card'
	import { Input } from '@/lib/components/ui/input'
	import { Button } from '@/lib/components/ui/button'
	import { Label } from '@/lib/components/ui/label'
	import { apiClient } from '@/lib/api'
	import { authToken } from '@/lib/auth'
	import { goto } from '$app/navigation'

	let submitting = false
	let login = ''
	let password = ''
	let errorMessage: string | null = null

	function onChange() {
		errorMessage = null
	}

	async function onSubmit() {
		submitting = true
		const token = await apiClient.authenticate(login, password)
		submitting = false

		if (token) {
			authToken.set(token)
			await goto('/explore')
		} else {
			errorMessage = 'invalid credentials'
		}
	}
</script>

<form on:submit={onSubmit}>
	<Card.Root>
		<Card.Header>
			<Card.Title>Login</Card.Title>
		</Card.Header>
		<Card.Content class="space-y-4">
			<div class="space-y-2">
				<Label>Login</Label>
				<Input
					bind:value={login}
					placeholder="login"
					disabled={submitting}
					on:change={onChange}
				/>
			</div>
			<div class="space-y-2">
				<Label>Password</Label>
				<Input
					bind:value={password}
					placeholder="password"
					type="password"
					disabled={submitting}
					on:change={onChange}
				/>
			</div>
			{#if errorMessage}
				<div class="text-destructive">{errorMessage}</div>
			{/if}
		</Card.Content>
		<Card.Footer>
			<Button
				class="w-full disabled:animate-pulse"
				type="submit"
				disabled={submitting}
			>
				{#if submitting}
					...
				{:else}
					Login
				{/if}
			</Button>
		</Card.Footer>
	</Card.Root>
</form>

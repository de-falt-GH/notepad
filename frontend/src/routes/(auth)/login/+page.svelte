<script lang="ts">
	import * as Card from '@/lib/components/ui/card'
	import { Input } from '@/lib/components/ui/input'
	import { Button, buttonVariants } from '@/lib/components/ui/button'
	import { Label } from '@/lib/components/ui/label'
	import { apiClient } from '@/lib/api'
	import { authToken } from '@/lib/auth'
	import { goto } from '$app/navigation'
	import { ArrowLeftIcon, Loader2Icon } from 'lucide-svelte'
	import { Separator } from '@/lib/components/ui/separator'

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

<div class="flex flex-col gap-4">
	<form on:submit={onSubmit}>
		<Card.Root class="border-8 border-yellow-300 drop-shadow-lg">
			<Card.Header>
				<Card.Title class="text-xl font-bold">Login</Card.Title>
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
						<Loader2Icon class="animate-spin" />
					{:else}
						Login
					{/if}
				</Button>
			</Card.Footer>
		</Card.Root>
	</form>
	<div class="flex w-full flex-row items-center gap-2 opacity-50">
		<Separator class="flex-1 bg-black" />
		<span class="text-black">OR</span>
		<Separator class="flex-1 bg-black" />
	</div>
	<a href="/register" class={buttonVariants({ variant: 'outline' })}>
		Register
	</a>
</div>

<a
	href="/explore"
	class={buttonVariants({
		variant: 'outline',
		className: 'absolute left-2 top-2',
	})}
>
	<ArrowLeftIcon />
</a>

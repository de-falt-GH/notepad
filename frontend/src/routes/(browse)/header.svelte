<script lang="ts">
	import { Button, buttonVariants } from '$lib/components/ui/button'
	import { authToken } from '$lib/auth'
	import { cn } from '@/lib/utils'
	import { apiClient } from '@/lib/api'
	import * as Avatar from '@/lib/components/ui/avatar'
	import { PencilIcon, LogOutIcon } from 'lucide-svelte/icons'
	import { goto } from '$app/navigation'

	let className = ''
	export { className as class }

	let name = ''

	$: nameAbbr = name
		.split(' ')
		.map(w => w[0])
		.join('')
		.slice(0, 2)
		.toUpperCase()

	$: if ($authToken)
		(async () => {
			try {
				const profile = await apiClient.fetchProfile()
				name = profile.name
			} catch {
				authToken.set(null)
			}
		})()

	async function onLogOut() {
		authToken.set('')
		await goto('/')
	}
</script>

<div
	class={cn(
		'flex h-min w-full flex-row gap-2 border-b bg-yellow-300 p-2 shadow-lg',
		className,
	)}
>
	{#if $authToken}
		<Avatar.Root class="outline outline-1 outline-border">
			<Avatar.Fallback class="bg-yellow-200 font-bold"
				>{nameAbbr}</Avatar.Fallback
			>
		</Avatar.Root>
		<a
			class={buttonVariants({
				variant: 'outline',
				className: 'gap-1 bg-yellow-200 p-2',
			})}
			href="/notes/new"
		>
			<PencilIcon class="h-6 w-6" />
		</a>
	{:else}
		<a class={buttonVariants({ variant: 'outline' })} href="/login">
			Login
		</a>
	{/if}

	<div class="flex flex-grow">
		<span
			class="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 font-bold"
		>
			notepad
		</span>
	</div>

	{#if $authToken}
		<Button variant="outline" class="bg-yellow-200 p-2" on:click={onLogOut}>
			<LogOutIcon class="h-6 w-6" />
		</Button>
	{/if}
</div>

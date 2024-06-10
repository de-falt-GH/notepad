<script lang="ts">
	import { apiClient } from '@/lib/api'
	import { Loader2Icon, PencilIcon } from 'lucide-svelte'
	import { authToken } from '@/lib/auth'
	import type { NoteFeedItem } from '@/lib/api/client'
	import NoteLink from './NoteLink.svelte'
	import type { ComponentProps } from 'svelte'
	import { Button } from '@/lib/components/ui/button'
	import { goto } from '$app/navigation'

	type NoteLinkData = ComponentProps<NoteLink>

	const mapToLinkData = ({ id, name, author_name, updated }: NoteFeedItem) =>
		({
			id,
			name,
			authorName: author_name,
			updated: new Date(updated),
		}) satisfies NoteLinkData

	$: publicNotes = apiClient
		.fetchPublicNotes()
		.then(({ notes }) => notes.map(mapToLinkData))

	$: privateNotes =
		$authToken !== null
			? apiClient
					.fetchPrivateNotes()
					.then(({ notes }) => notes.map(mapToLinkData))
			: null

	const onCreateClick = async () => {
		const noteId = await apiClient.createNote({
			name: 'Untitled',
			data: '# Untitled\n\n',
			public: false,
		})

		if (noteId !== null) {
			await goto(`notes/${noteId}`)
		}
	}
</script>

{#if privateNotes !== null}
	<h1 class="flex items-center gap-2 text-lg font-bold">
		<span>My Notes</span>
		<Button variant="outline" class="h-6 w-6 p-1" on:click={onCreateClick}>
			<PencilIcon />
		</Button>
	</h1>
	{#await privateNotes}
		<div
			class="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2"
		>
			<Loader2Icon class="animate-spin" />
		</div>
	{:then privateNotes}
		<div
			class="grid grid-cols-1 items-stretch gap-4 pt-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4"
		>
			{#each privateNotes as note}
				<NoteLink {...note} />
			{/each}
		</div>
		{#if privateNotes.length == 0}
			<div
				class="flex h-12 w-full items-center justify-center rounded-md border border-dashed border-black"
			>
				Create your first note!
			</div>
		{/if}
	{/await}

	<h1 class="mt-4 text-lg font-bold">Public Notes</h1>
{/if}

{#await publicNotes}
	<div class="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2">
		<Loader2Icon class="animate-spin" />
	</div>
{:then publicNotes}
	<div
		class="grid grid-cols-1 items-stretch gap-4 pt-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4"
	>
		{#each publicNotes as note}
			<NoteLink {...note} />
		{/each}
	</div>
{/await}

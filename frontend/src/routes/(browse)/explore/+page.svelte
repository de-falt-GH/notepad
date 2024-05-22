<script lang="ts">
	import * as Card from '@/lib/components/ui/card'
	import { apiClient } from '@/lib/api'
	import { getRelativeTimeString } from '@/lib/utils'
	import { Loader2Icon } from 'lucide-svelte'

	type NoteLinkData = {
		id: number
		name: string
		authorName: string
		updated: Date
	}

	$: notes = apiClient.fetchPublicNotes().then(({ notes: fetchedNotes }) =>
		fetchedNotes.map(
			({ id, name, author_name, updated }) =>
				({
					id,
					name,
					authorName: author_name,
					updated: new Date(updated),
				}) satisfies NoteLinkData,
		),
	)
</script>

<div
	class="grid grid-cols-1 items-stretch gap-4 pt-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4"
>
	{#await notes}
		<div
			class="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2"
		>
			<Loader2Icon class=" animate-spin" />
		</div>
	{:then notes}
		{#each notes as { id, name, authorName, updated }}
			<a href="notes/{id}">
				<Card.Root
					class="h-full w-full drop-shadow-sm transition hover:-translate-y-1 hover:drop-shadow-md"
				>
					<Card.Header>
						<Card.Title>{name}</Card.Title>
						<Card.Description>
							<div>by {authorName}</div>
							<div>updated {getRelativeTimeString(updated)}</div>
						</Card.Description>
					</Card.Header>
				</Card.Root>
			</a>
		{/each}
	{/await}
</div>

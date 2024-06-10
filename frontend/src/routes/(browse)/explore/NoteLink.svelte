<script lang="ts">
	import { getRelativeTimeString } from '@/lib/utils'
	import * as Card from '@/lib/components/ui/card'

	export let id: number

	export let name: string

	export let authorName: string

	export let updated: Date
</script>

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

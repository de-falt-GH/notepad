<script>
	import { goto } from '$app/navigation'
	import { page } from '$app/stores'
	import { apiClient } from '@/lib/api'
	import { currentUserId } from '@/lib/auth'
	import { Switch } from '@/lib/components/ui/switch'
	import { debounce } from '@/lib/utils'
	import { code } from '@cartamd/plugin-code'
	import { math } from '@cartamd/plugin-math'
	import { Carta, Markdown, MarkdownEditor } from 'carta-md'
	import 'carta-md/default.css'
	import 'katex/dist/katex.css'
	import { ArrowLeftIcon, Loader2Icon } from 'lucide-svelte'
	import { onMount } from 'svelte'

	$: noteId = Number($page.params.id)

	let fetchingNote = true

	let isEditable = false
	let noteContents = ''
	let isPublic = false

	$: noteTitle = /^\s*#\s*(.*)$/m.exec(noteContents)?.[1] ?? 'Untitled'

	onMount(async () => {
		fetchingNote = true
		const note = await apiClient.fetchNote(noteId)
		fetchingNote = false
		if (!note) {
			await goto('/explore')
			return
		}

		isEditable = note.author_id == $currentUserId
		noteContents = note.data
		isPublic = note.public
	})

	const saveNote = debounce(async () => {
		if (!isEditable) {
			return
		}

		await apiClient.updateNote(noteId, {
			name: noteTitle,
			data: noteContents,
			public: isPublic,
		})
	})

	$: {
		noteTitle, noteContents, isPublic
		saveNote()
	}

	const carta = new Carta({
		sanitizer: false, // TODO: add the sanitizer
		extensions: [code(), math()],
	})
</script>

{#if fetchingNote}
	<div class="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2">
		<Loader2Icon class="animate-spin" />
	</div>
{:else}
	<title>{noteTitle}</title>

	{#if isEditable}
		<MarkdownEditor mode="split" bind:value={noteContents} {carta} />
	{:else}
		<div class="h-8 w-full bg-yellow-300" />
		<div class="p-2">
			<Markdown value={noteContents} {carta} />
		</div>
	{/if}

	<div class="absolute left-1 top-1 flex flex-row gap-2">
		<a href="/explore" class="h-6 w-6">
			<ArrowLeftIcon />
		</a>
		{#if isEditable}
			<div class="flex items-center gap-2">
				<span>Public:</span>
				<Switch bind:checked={isPublic} />
			</div>
		{/if}
	</div>
{/if}

<style>
	:global(.carta-editor) {
		@apply h-full w-full;
	}

	:global(.carta-wrapper) {
		@apply flex-grow;
	}

	:global(.carta-container) {
		@apply !h-full;
	}

	:global(.carta-input) {
		@apply !h-[calc(100%-2rem)];
	}
</style>

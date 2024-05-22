<script>
	import Header from './header.svelte'
</script>

<div class="flex h-[100svh] flex-col">
	<Header class="sticky top-0 z-10" />
	<main class="flex-grow bg-neutral-200 p-2">
		<slot />
	</main>
</div>

<script>
	import { buttonVariants } from '@/lib/components/ui/button'
</script>

<main class="h-[100svh] w-full bg-neutral-200">
	<div
		class="absolute left-1/2 top-1/2 flex w-fit -translate-x-1/2 -translate-y-1/2 flex-col items-center gap-20 rounded-md bg-white p-20 outline outline-8 outline-yellow-300 drop-shadow-lg transition hover:-translate-y-[calc(50%+5px)] hover:drop-shadow-xl"
	>
		<div
			class="absolute left-0 top-0 flex w-full -translate-y-1/2 flex-row items-center justify-around"
		>
			{#each { length: 6 } as _}
				<div class="h-8 w-2 rounded-full bg-neutral-400" />
			{/each}
		</div>

		<div
			class="relative flex flex-col items-center justify-center gap-2 text-nowrap"
		>
			<div class="text-2xl font-bold">notepad</div>
			<div class="text-lg italic">you note - we pad</div>
		</div>

		<a
			class={buttonVariants({
				class: 'w-min shadow-lg shadow-transparent transition hover:-translate-y-1 hover:shadow-primary',
			})}
			href="/explore"
		>
			Explore
		</a>
	</div>
</main>

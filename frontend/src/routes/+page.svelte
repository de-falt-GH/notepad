<script>
	import { buttonVariants } from '@/lib/components/ui/button'
</script>

<main class="flex h-[100svh] w-full flex-col items-center">
	<div
		class="relative flex h-[400px] flex-col items-center justify-center gap-2"
	>
		<div class="text-lg font-bold">notepad</div>
		<div class="italic">you note - we pad</div>
	</div>

	<a
		class={buttonVariants({ variant: 'outline', class: 'w-min' })}
		href="/explore"
	>
		Explore
	</a>
</main>

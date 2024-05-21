<script lang="ts">
	import { buttonVariants } from '$lib/components/ui/button'
	import { authToken } from '$lib/auth'
</script>

<div class="h-min w-full border-b p-2">
	{#if $authToken}
		{$authToken}
	{:else}
		<a class={buttonVariants({ variant: 'outline' })} href="/login">
			Login
		</a>
	{/if}
</div>

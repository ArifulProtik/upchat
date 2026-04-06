<script lang="ts">
	import type { Snippet } from 'svelte';
	import { useSession } from '$lib/api/auth';
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { removeToken } from '$lib/api/client';

	let { children }: { children: Snippet } = $props();

	const session = useSession();

	$effect(() => {
		if (session.isError) {
			removeToken();
			goto(resolve('/signin'));
		}
	});
</script>

{#if session.isLoading}
	<div class="flex h-screen items-center justify-center">
		<div class="flex flex-col items-center gap-3">
			<div
				class="size-8 animate-spin rounded-full border-4 border-muted border-t-primary"
			></div>
			<p class="text-sm text-muted-foreground">Loading...</p>
		</div>
	</div>
{:else if session.isError}
	<div class="flex h-screen items-center justify-center">
		<p class="text-sm text-destructive">Session expired. Redirecting...</p>
	</div>
{:else}
	{@render children()}
{/if}

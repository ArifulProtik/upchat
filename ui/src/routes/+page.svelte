<script lang="ts">
	import { resolve } from '$app/paths';

	let count = $state(0);
</script>

<svelte:head>
	<title>Home - UpChat</title>
</svelte:head>

<button class="text-2xl font-bold text-red-800" onclick={() => count++}>
	Up
</button>
{count}
<a href={resolve('/signin')}>Sign In</a>

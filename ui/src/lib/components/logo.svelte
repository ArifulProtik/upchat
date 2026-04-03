<script>
	import ChartCircles from '@tabler/icons-svelte/icons/chart-circles';
	import { resolve } from '$app/paths';
</script>

<a href={resolve('/')}>
	<div class="flex items-center gap-2">
		<div
			class="flex h-10 w-10 items-center justify-center rounded-full bg-primary"
		>
			<span class="text-xl font-bold text-primary-foreground"
				><ChartCircles /></span
			>
		</div>
		<span class="text-xl font-bold">UpChat</span>
	</div>
</a>

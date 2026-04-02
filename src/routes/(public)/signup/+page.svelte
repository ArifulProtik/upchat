<script lang="ts">
	import Logo from '$lib/components/logo.svelte';
	import SignupForm from '$lib/components/signup-form.svelte';
</script>

<svelte:head>
	<title>Signup - Upchat</title>
</svelte:head>

<div class="flex h-screen items-center justify-center bg-muted/60">
	<div class="flex w-full flex-col items-center justify-center gap-4">
		<div class="text-center">
			<Logo />
		</div>
		<SignupForm />
	</div>
</div>

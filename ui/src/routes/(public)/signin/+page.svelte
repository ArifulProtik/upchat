<script lang="ts">
	import Logo from '$lib/components/logo.svelte';
	import SigninForm from '$lib/components/signin-form.svelte';
</script>

<svelte:head>
	<title>Sign In - UpChat</title>
</svelte:head>

<div class="flex h-screen items-center justify-center bg-muted/60">
	<div class="flex w-full flex-col items-center justify-center gap-4">
		<div class="text-center">
			<Logo />
		</div>
		<SigninForm />
	</div>
</div>

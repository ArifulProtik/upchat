<script lang="ts">
	import { Button } from '$lib/components/ui/button/index.js';
	import { Label } from '$lib/components/ui/label/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import * as Card from '$lib/components/ui/card/index.js';
	import { resolve } from '$app/paths';
	import { signupSchema, type SignupSchema } from '$lib/types/auth-schema';
	import { z } from 'zod';
	import { useSignup } from '$lib/api/auth';
	import { ApiError } from '$lib/api/client';
	import { toast } from 'svelte-sonner';
	import { goto } from '$app/navigation';

	let formData: SignupSchema = $state({
		name: '',
		email: '',
		password: ''
	});

	let touchedField = $state({
		name: false,
		email: false,
		password: false
	});

	let serverError = $state('');

	let errors = $derived.by(() => {
		const result = signupSchema.safeParse(formData);
		if (result.success) return {};
		const err = z.flattenError(result.error).fieldErrors;
		return {
			name: touchedField.name ? err.name?.[0] : undefined,
			email: touchedField.email ? err.email?.[0] : undefined,
			password: touchedField.password ? err.password?.[0] : undefined
		};
	});

	const signup = useSignup();

	const handleSubmit = (e: SubmitEvent) => {
		e.preventDefault();
		touchedField.name = true;
		touchedField.email = true;
		touchedField.password = true;
		serverError = '';

		const result = signupSchema.safeParse(formData);
		if (!result.success) return;

		signup.mutate(result.data, {
			onSuccess: () => {
				toast.success('Account created! Please sign in.');
				goto(resolve('/signin'));
			},
			onError: (error) => {
				if (error instanceof ApiError) {
					serverError = error.message;
				} else {
					toast.error('Something went wrong. Please try again.');
				}
			}
		});
	};
</script>

<Card.Root class="w-full max-w-sm">
	<Card.Header>
		<Card.Title>Create an account</Card.Title>
		<Card.Description
			>Enter your credentials below to create an account</Card.Description
		>
	</Card.Header>
	<Card.Content>
		<form onsubmit={handleSubmit}>
			<div class="flex flex-col gap-6">
				{#if serverError}
					<p class="text-center text-sm text-destructive">
						{serverError}
					</p>
				{/if}
				<div class="grid gap-2">
					<Label for="name">Name</Label>
					<Input
						id="name"
						type="text"
						placeholder="John Doe"
						bind:value={formData.name}
						oninput={() => (touchedField.name = true)}
						disabled={signup.isPending}
					/>
					{#if errors.name}
						<p class="text-destructive">{errors.name}</p>
					{/if}
				</div>
				<div class="grid gap-2">
					<Label for="email">Email</Label>
					<Input
						id="email"
						type="email"
						placeholder="m@example.com"
						bind:value={formData.email}
						oninput={() => (touchedField.email = true)}
						disabled={signup.isPending}
					/>
					{#if errors.email}
						<p class="text-destructive">{errors.email}</p>
					{/if}
				</div>
				<div class="grid gap-2">
					<Label for="password">Password</Label>
					<Input
						id="password"
						type="password"
						bind:value={formData.password}
						oninput={() => (touchedField.password = true)}
						disabled={signup.isPending}
					/>
					{#if errors.password}
						<p class="text-destructive">{errors.password}</p>
					{/if}
				</div>
				<Button type="submit" class="w-full" disabled={signup.isPending}>
					{#if signup.isPending}
						Creating account...
					{:else}
						Sign Up
					{/if}
				</Button>
			</div>
		</form>
	</Card.Content>
	<Card.Footer class="flex-col gap-3">
		<p class="text-center text-sm text-muted-foreground">
			Already have an account? <a
				href={resolve('/signin')}
				class="font-medium underline underline-offset-4 hover:underline"
				>Login</a
			>
		</p>

		<p class="mt-4 text-center text-sm text-muted-foreground">
			By signing up, you agree to our <a
				href="##"
				class="font-medium underline underline-offset-4 hover:underline"
				>Terms of Service</a
			>
			and
			<a
				href="##"
				class="font-medium underline underline-offset-4 hover:underline"
				>Privacy Policy</a
			>
		</p>
	</Card.Footer>
</Card.Root>

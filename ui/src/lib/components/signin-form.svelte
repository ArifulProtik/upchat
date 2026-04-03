<script lang="ts">
	import { Button } from '$lib/components/ui/button/index.js';
	import { Label } from '$lib/components/ui/label/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import * as Card from '$lib/components/ui/card/index.js';
	import { resolve } from '$app/paths';
	import { signinSchema, type SigninSchema } from '$lib/types/auth-schema';
	import { z } from 'zod';

	let formData: SigninSchema = $state({
		email: '',
		password: ''
	});

	let touchedField = $state({
		email: false,
		password: false
	});
	let errors = $derived.by(() => {
		const result = signinSchema.safeParse(formData);
		if (result.success) return {};
		const err = z.flattenError(result.error).fieldErrors;
		return {
			email: touchedField.email ? err.email?.[0] : undefined,
			password: touchedField.password ? err.password?.[0] : undefined
		};
	});

	const handleSubmit = (e: SubmitEvent) => {
		e.preventDefault();
		touchedField.email = true;
		touchedField.password = true;
		const result = signinSchema.safeParse(formData);
		if (result.success) {
			console.log(result.data);
		}
	};
</script>

<Card.Root class="w-full max-w-sm">
	<Card.Header>
		<Card.Title>Login to your account</Card.Title>
		<Card.Description
			>Enter your email below to login to your account</Card.Description
		>
	</Card.Header>
	<Card.Content>
		<form onsubmit={handleSubmit}>
			<div class="flex flex-col gap-6">
				<div class="grid gap-2">
					<Label for="email">Email</Label>
					<Input
						id="email"
						type="email"
						placeholder="m@example.com"
						bind:value={formData.email}
						oninput={() => (touchedField.email = true)}
					/>
					{#if errors.email}
						<p class="text-destructive">{errors.email}</p>
					{/if}
				</div>
				<div class="grid gap-2">
					<div class="flex items-center">
						<Label for="password">Password</Label>
						<a
							href="##"
							class="ms-auto inline-block text-sm underline-offset-4 hover:underline"
						>
							Forgot your password?
						</a>
					</div>
					<Input
						id="password"
						type="password"
						bind:value={formData.password}
						oninput={() => (touchedField.password = true)}
					/>
					{#if errors.password}
						<p class="text-destructive">{errors.password}</p>
					{/if}
				</div>
				<Button type="submit" class="w-full">Login</Button>
			</div>
		</form>
	</Card.Content>
	<Card.Footer class="flex-col gap-3">
		<p
			class="mt-8 w-full border-t py-2 text-center text-sm text-muted-foreground"
		>
			Don't have an account? <a
				href={resolve('/signup')}
				class="font-medium hover:text-primary">Sign Up</a
			>
		</p>
	</Card.Footer>
</Card.Root>

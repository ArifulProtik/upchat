import {
	createQuery,
	createMutation,
	useQueryClient
} from '@tanstack/svelte-query';
import { goto } from '$app/navigation';
import { resolve } from '$app/paths';
import { apiFetch, setToken, removeToken } from '$lib/api/client';
import type {
	LoginResponse,
	UserResponse,
	GetSessionResponse
} from '$lib/types/api';
import type { SigninSchema } from '$lib/types/auth-schema';
import type { SignupSchema } from '$lib/types/auth-schema';

export const sessionQueryKey = ['session'] as const;

export function useSession() {
	return createQuery(() => ({
		queryKey: sessionQueryKey,
		queryFn: () => apiFetch<GetSessionResponse>('/auth/get-session'),
		retry: false,
		staleTime: 5 * 60 * 1000
	}));
}

export function useLogin() {
	const queryClient = useQueryClient();

	return createMutation<LoginResponse, Error, SigninSchema>(() => ({
		mutationFn: (data: SigninSchema) =>
			apiFetch<LoginResponse>('/auth/login', {
				method: 'POST',
				body: JSON.stringify(data)
			}),
		onSuccess: (data: LoginResponse) => {
			setToken(data.token);
			queryClient.setQueryData(sessionQueryKey, {
				user: data.user,
				account: null
			});
			goto(resolve('/chat'));
		}
	}));
}

export function useSignup() {
	return createMutation<UserResponse, Error, SignupSchema>(() => ({
		mutationFn: (data: SignupSchema) =>
			apiFetch<UserResponse>('/auth/signup', {
				method: 'POST',
				body: JSON.stringify(data)
			})
	}));
}

export function useLogout() {
	const queryClient = useQueryClient();

	return createMutation<{ message: string }, Error, void>(() => ({
		mutationFn: () =>
			apiFetch<{ message: string }>('/auth/logout', {
				method: 'POST'
			}),
		onSettled: () => {
			removeToken();
			queryClient.setQueryData(sessionQueryKey, null);
			queryClient.invalidateQueries({ queryKey: sessionQueryKey });
			goto(resolve('/signin'));
		}
	}));
}

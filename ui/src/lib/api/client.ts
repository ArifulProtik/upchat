import type { ApiErrorBody } from '$lib/types/api';

const TOKEN_KEY = 'auth_token';

export function getToken(): string | null {
	if (typeof document === 'undefined') return null;
	const match = document.cookie.match(
		new RegExp(`(?:^|; )${TOKEN_KEY}=([^;]*)`)
	);
	return match ? decodeURIComponent(match[1]) : null;
}

export function setToken(token: string): void {
	const maxAge = 60 * 60 * 24 * 30; // 30 days
	document.cookie = `${TOKEN_KEY}=${encodeURIComponent(token)}; path=/; max-age=${maxAge}; SameSite=Lax`;
}

export function removeToken(): void {
	document.cookie = `${TOKEN_KEY}=; path=/; max-age=0; path=/`;
}

export class ApiError extends Error {
	status: number;
	body: ApiErrorBody;

	constructor(status: number, body: ApiErrorBody) {
		const msg =
			typeof body.error === 'string' ? body.error : 'Validation failed';
		super(msg);
		this.name = 'ApiError';
		this.status = status;
		this.body = body;
	}
}

export async function apiFetch<T>(
	path: string,
	options?: RequestInit
): Promise<T> {
	const token = getToken();
	const headers = new Headers(options?.headers);

	if (token) {
		headers.set('Authorization', `Bearer ${token}`);
	}

	if (options?.body && !headers.has('Content-Type')) {
		headers.set('Content-Type', 'application/json');
	}

	const res = await fetch(`/api${path}`, {
		...options,
		headers
	});

	if (!res.ok) {
		const body = await res.json().catch(() => ({
			status: res.status,
			error: res.statusText
		}));
		throw new ApiError(res.status, body);
	}

	return res.json();
}

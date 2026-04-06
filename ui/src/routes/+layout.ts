import type { LayoutLoad } from './$types';
import { QueryClient } from '@tanstack/svelte-query';
import { browser } from '$app/environment';

export const load: LayoutLoad = async () => {
	const queryClient = new QueryClient({
		defaultOptions: {
			queries: {
				enabled: browser,
				staleTime: 60 * 1000
			}
		}
	});

	return { queryClient };
};

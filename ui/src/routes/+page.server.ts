import { redirect } from '@sveltejs/kit';
import { resolve } from '$app/paths';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ locals }) => {
	if (locals.token) {
		redirect(302, resolve('/chat'));
	} else {
		redirect(302, resolve('/signin'));
	}
};

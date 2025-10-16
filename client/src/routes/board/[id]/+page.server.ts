import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ parent, params }) => {
	const { user } = await parent();

	if (!user) {
		throw redirect(303, '/login');
	}

	return {
		boardId: params.id,
		user
	};
};

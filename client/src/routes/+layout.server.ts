import type { LayoutServerLoad } from './$types';
import { adminAuth } from '$lib/firebase/admin';

export const load: LayoutServerLoad = async ({ cookies }) => {
	const sessionCookie = cookies.get('session');

	if (!sessionCookie) {
		return { user: null };
	}

	try {
		const decodedToken = await adminAuth.verifySessionCookie(sessionCookie, true);
		const user = await adminAuth.getUser(decodedToken.uid);

		return {
			user: {
				uid: user.uid,
				email: user.email ?? null,
				displayName: user.displayName ?? null,
				photoURL: user.photoURL ?? null
			}
		};
	} catch (error) {
		cookies.delete('session', { path: '/' });
		return { user: null };
	}
};

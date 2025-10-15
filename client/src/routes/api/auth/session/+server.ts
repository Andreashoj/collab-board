import { json, type RequestHandler } from '@sveltejs/kit';
import { adminAuth } from '$lib/firebase/admin';

export const POST: RequestHandler = async ({ request, cookies }) => {
	const { idToken } = await request.json();

	if (!idToken) {
		return json({ error: 'Missing ID token' }, { status: 400 });
	}

	try {
		// Verify the ID token
		const decodedToken = await adminAuth.verifyIdToken(idToken);

		// Create session cookie (expires in 5 days)
		const expiresIn = 60 * 60 * 24 * 5 * 1000; // 5 days
		const sessionCookie = await adminAuth.createSessionCookie(idToken, { expiresIn });

		// Set cookie
		cookies.set('session', sessionCookie, {
			maxAge: expiresIn,
			httpOnly: true,
			secure: true,
			path: '/',
			sameSite: 'lax'
		});

		return json({ status: 'success', uid: decodedToken.uid });
	} catch (error) {
		console.error('Session creation error:', error);
		return json({ error: 'Failed to create session' }, { status: 500 });
	}
};

export const DELETE: RequestHandler = async ({ cookies }) => {
	cookies.delete('session', { path: '/' });
	return json({ status: 'success' });
};

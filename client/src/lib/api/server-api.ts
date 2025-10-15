import axios from 'axios';
import { adminAuth } from '$lib/firebase/admin';
import type { Cookies } from '@sveltejs/kit';
import { VITE_API_URL } from '$env/static/private';

export async function createServerApiClient(cookies: Cookies) {
	const sessionCookie = cookies.get('session');
	
	if (!sessionCookie) {
		throw new Error('No session cookie found');
	}
	const decodedToken = await adminAuth.verifySessionCookie(sessionCookie, true);
	
	const api = axios.create({
		baseURL: VITE_API_URL || 'http://localhost:8080',
		headers: {
			'Content-Type': 'application/json',
			'X-Firebase-UID': decodedToken.uid
		}
	});

	return { api, uid: decodedToken.uid };
}

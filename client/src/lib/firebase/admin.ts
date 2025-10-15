import { getAuth } from 'firebase-admin/auth';
import { getApps, initializeApp, cert } from 'firebase-admin/app';
import { FIREBASE_SERVICE_ACCOUNT } from '$env/static/private';

function getFirebaseAdmin() {
	if (getApps().length === 0) {
		const serviceAccount = JSON.parse(FIREBASE_SERVICE_ACCOUNT);

		initializeApp({
			credential: cert(serviceAccount)
		});
	}

	return getAuth();
}

export const adminAuth = getFirebaseAdmin();

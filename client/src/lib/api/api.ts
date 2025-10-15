import axios from 'axios';
import { browser } from '$app/environment';

// API client for Go backend - includes Firebase auth token
const api = axios.create({
	baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8080',
	headers: {
		'Content-Type': 'application/json'
	},
	withCredentials: true
});

// Request interceptor - automatically add Firebase ID token
api.interceptors.request.use(
	async (config) => {
		// Only run in browser (not during SSR)
		if (browser) {
			try {
				// Wait for auth state to be ready
				const { waitForAuth } = await import('$lib/firebase/client');
				const user = await waitForAuth();
				
				if (user) {
					const idToken = await user.getIdToken();
					config.headers.Authorization = `Bearer ${idToken}`;
				} else {
					console.warn('No authenticated user - request will be sent without auth');
				}
			} catch (error) {
				console.error('Failed to get ID token:', error);
			}
		}
		return config;
	},
	(error) => {
		return Promise.reject(error);
	}
);

// Response interceptor
api.interceptors.response.use(
	(response) => {
		return response;
	},
	(error) => {
		// Handle common errors
		console.log('hello')
		if (error.response?.status === 401) {
			// Unauthorized - could redirect to login
			console.error('Unauthorized request');
		}
		return Promise.reject(error);
	}
);

// Local API client for SvelteKit endpoints (auth, etc.) - no auth token needed
export const localApi = axios.create({
	baseURL: '/api',
	headers: {
		'Content-Type': 'application/json'
	},
	withCredentials: true
});

export default api;

<script lang="ts">
	import { auth } from '$lib/firebase/client';
	import {
		signInWithEmailAndPassword,
		GoogleAuthProvider,
		signInWithPopup
	} from 'firebase/auth';
	import { goto } from '$app/navigation';
	import { localApi } from '$lib/api/api';

	let email = $state('');
	let password = $state('');
	let error = $state('');
	let loading = $state(false);

	async function handleEmailLogin() {
		if (!email || !password) {
			error = 'Please fill in all fields';
			return;
		}

		loading = true;
		error = '';

		try {
			const userCredential = await signInWithEmailAndPassword(auth, email, password);
			const idToken = await userCredential.user.getIdToken();

			await localApi.post('/auth/session', { idToken });
			goto('/');
		} catch (err: any) {
			error = err.message || 'Failed to sign in';
		} finally {
			loading = false;
		}
	}

	async function handleGoogleLogin() {
		loading = true;
		error = '';

		try {
			const provider = new GoogleAuthProvider();
			const userCredential = await signInWithPopup(auth, provider);
			const idToken = await userCredential.user.getIdToken();

			await localApi.post('/auth/session', { idToken });
			goto('/');
		} catch (err: any) {
			error = err.message || 'Failed to sign in with Google';
		} finally {
			loading = false;
		}
	}
</script>

<div class="flex min-h-screen items-center justify-center">
	<div class="w-full max-w-md space-y-8 rounded-lg border p-8">
		<h1 class="text-center text-3xl font-bold">Sign In</h1>

		{#if error}
			<div class="rounded bg-red-50 p-4 text-red-800">
				{error}
			</div>
		{/if}

		<form class="space-y-6" onsubmit={(e) => { e.preventDefault(); handleEmailLogin(); }}>
			<div>
				<label for="email" class="block text-sm font-medium">Email</label>
				<input
					id="email"
					type="email"
					bind:value={email}
					disabled={loading}
					class="mt-1 w-full rounded border px-3 py-2"
					required
				/>
			</div>

			<div>
				<label for="password" class="block text-sm font-medium">Password</label>
				<input
					id="password"
					type="password"
					bind:value={password}
					disabled={loading}
					class="mt-1 w-full rounded border px-3 py-2"
					required
				/>
			</div>

			<button
				type="submit"
				disabled={loading}
				class="w-full rounded bg-blue-600 px-4 py-2 text-white hover:bg-blue-700 disabled:opacity-50"
			>
				{loading ? 'Signing in...' : 'Sign In'}
			</button>
		</form>

		<div class="relative">
			<div class="absolute inset-0 flex items-center">
				<div class="w-full border-t"></div>
			</div>
			<div class="relative flex justify-center text-sm">
				<span class="bg-white px-2 text-gray-500">Or continue with</span>
			</div>
		</div>

		<button
			onclick={handleGoogleLogin}
			disabled={loading}
			class="w-full rounded border border-gray-300 bg-white px-4 py-2 hover:bg-gray-50 disabled:opacity-50"
		>
			Sign in with Google
		</button>

		<p class="text-center text-sm">
			Don't have an account? <a href="/signup" class="text-blue-600 hover:underline">Sign up</a>
		</p>
	</div>
</div>

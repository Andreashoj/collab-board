<script lang="ts">
	import { auth } from '$lib/firebase/client';
	import { createUserWithEmailAndPassword, updateProfile } from 'firebase/auth';
	import { goto } from '$app/navigation';
	import { localApi } from '$lib/api/api';

	let email = $state('');
	let password = $state('');
	let displayName = $state('');
	let error = $state('');
	let loading = $state(false);

	async function handleSignup() {
		if (!email || !password || !displayName) {
			error = 'Please fill in all fields';
			return;
		}

		if (password.length < 6) {
			error = 'Password must be at least 6 characters';
			return;
		}

		loading = true;
		error = '';

		try {
			const userCredential = await createUserWithEmailAndPassword(auth, email, password);
			
			await updateProfile(userCredential.user, { displayName });

			const idToken = await userCredential.user.getIdToken();

			await localApi.post('/auth/session', { idToken });
			goto('/dashboard');
		} catch (err: any) {
			error = err.message || 'Failed to create account';
		} finally {
			loading = false;
		}
	}
</script>

<div class="flex min-h-screen items-center justify-center">
	<div class="w-full max-w-md space-y-8 rounded-lg border p-8">
		<h1 class="text-center text-3xl font-bold">Create Account</h1>

		{#if error}
			<div class="rounded bg-red-50 p-4 text-red-800">
				{error}
			</div>
		{/if}

		<form class="space-y-6" onsubmit={(e) => { e.preventDefault(); handleSignup(); }}>
			<div>
				<label for="displayName" class="block text-sm font-medium">Display Name</label>
				<input
					id="displayName"
					type="text"
					bind:value={displayName}
					disabled={loading}
					class="mt-1 w-full rounded border px-3 py-2"
					required
				/>
			</div>

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
				{loading ? 'Creating account...' : 'Sign Up'}
			</button>
		</form>

		<p class="text-center text-sm">
			Already have an account? <a href="/login" class="text-blue-600 hover:underline">Sign in</a>
		</p>
	</div>
</div>

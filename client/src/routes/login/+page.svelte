<script lang="ts">
	import { auth } from '$lib/firebase/client';
	import {
		signInWithEmailAndPassword,
		GoogleAuthProvider,
		signInWithPopup
	} from 'firebase/auth';
	import { goto } from '$app/navigation';
	import { localApi } from '$lib/api/api';
	import { Button } from '$lib/components/ui/button';
	import * as Card from '$lib/components/ui/card';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import * as Alert from '$lib/components/ui/alert';
	import { Separator } from '$lib/components/ui/separator';

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

<div class="flex min-h-screen items-center justify-center p-4">
	<Card.Root class="w-full max-w-md">
		<Card.Header>
			<Card.Title class="text-center text-2xl">Sign In</Card.Title>
			<Card.Description class="text-center">Enter your credentials to access your account</Card.Description>
		</Card.Header>
		<Card.Content class="space-y-6">
			{#if error}
				<Alert.Root variant="destructive">
					<Alert.Description>{error}</Alert.Description>
				</Alert.Root>
			{/if}

			<form class="space-y-4" onsubmit={(e) => { e.preventDefault(); handleEmailLogin(); }}>
				<div class="space-y-2">
					<Label for="email">Email</Label>
					<Input
						id="email"
						type="email"
						bind:value={email}
						disabled={loading}
						placeholder="you@example.com"
						required
					/>
				</div>

				<div class="space-y-2">
					<Label for="password">Password</Label>
					<Input
						id="password"
						type="password"
						bind:value={password}
						disabled={loading}
						placeholder="••••••••"
						required
					/>
				</div>

				<Button type="submit" class="w-full" disabled={loading}>
					{loading ? 'Signing in...' : 'Sign In'}
				</Button>
			</form>

			<div class="relative">
				<div class="absolute inset-0 flex items-center">
					<Separator />
				</div>
				<div class="relative flex justify-center text-xs uppercase">
					<span class="bg-background px-2 text-muted-foreground">Or continue with</span>
				</div>
			</div>

			<Button
				variant="outline"
				class="w-full"
				onclick={handleGoogleLogin}
				disabled={loading}
			>
				Sign in with Google
			</Button>

			<p class="text-center text-sm text-muted-foreground">
				Don't have an account? <a href="/signup" class="text-primary hover:underline">Sign up</a>
			</p>
		</Card.Content>
		</Card.Root>
</div>

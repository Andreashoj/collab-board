<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import type { PageData } from './$types';
	import api, { localApi } from '$lib/api/api';

	let { data }: { data: PageData } = $props();

	async function handleLogout() {
		await localApi.post('/auth/logout');
		goto('/login');
	}

	async function getUser() {
		const user = await api.get("/user")
		console.log(user)
	}

	// Only run in browser after component mounts
	onMount(() => {
		getUser();
	});
</script>

<div class="min-h-screen p-8">
	<div class="mx-auto max-w-4xl">
		<div class="mb-8 flex items-center justify-between">
			<h1 class="text-3xl font-bold">Dashboard</h1>
			<button
				onclick={handleLogout}
				class="rounded border border-red-600 px-4 py-2 text-red-600 hover:bg-red-50"
			>
			Logout
			</button>
		</div>

		<div class="rounded-lg border p-6">
			<h2 class="mb-4 text-xl font-semibold">Welcome, {data.user?.displayName || 'User'}!</h2>
			<div class="space-y-2 text-sm">
				<p><strong>UID:</strong> {data.user?.uid}</p>
				<p><strong>Email:</strong> {data.user?.email}</p>
			</div>
		</div>
	</div>
</div>

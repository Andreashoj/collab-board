<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import type { PageData } from './$types';
	import { localApi } from '$lib/api/api';
	import { Button } from '$lib/components/ui/button';
	import * as Card from '$lib/components/ui/card';
	import CreateBoardDialog from '$lib/components/dialogs/create-board-dialog.svelte';
	import BoardListSkeleton from '$lib/components/loaders/board-list-skeleton.svelte';
	import { boardsStore, boards, boardsLoading } from '$lib/stores/boards';

	let { data }: { data: PageData } = $props();

	let showCreateDialog = $state(false);

	async function handleLogout() {
		await localApi.post('/auth/logout');
		boardsStore.reset();
		goto('/login');
	}

	async function handleCreateBoard(name: string) {
		await boardsStore.createBoard(name);
	}

	onMount(() => {
		boardsStore.fetchBoards();
	});
</script>

<div class="min-h-screen p-8">
	<div class="mx-auto max-w-4xl">
		<div class="mb-8 flex items-center justify-between">
			<h1 class="text-3xl font-bold">Dashboard</h1>
			<Button variant="destructive" onclick={handleLogout}>
				Logout
			</Button>
		</div>

		<Card.Root>
			<Card.Header>
				<Card.Title>Welcome, {data.user?.displayName || 'User'}!</Card.Title>
			</Card.Header>
			<Card.Content>
				<div class="space-y-2 text-sm">
					<p><strong>UID:</strong> {data.user?.uid}</p>
					<p><strong>Email:</strong> {data.user?.email}</p>
				</div>
			</Card.Content>
		</Card.Root>

		<Card.Root class="mt-8">
			<Card.Header>
				<div class="flex items-center justify-between">
					<Card.Title>My Boards</Card.Title>
					<Button onclick={() => showCreateDialog = true}>
						Create Board
					</Button>
				</div>
			</Card.Header>
			<Card.Content>
				{#if $boardsLoading}
					<BoardListSkeleton />
				{:else if $boards.length === 0}
					<p class="text-muted-foreground">No boards yet. Create one to get started!</p>
				{:else}
					<div class="space-y-3">
						{#each $boards as board (board.id)}
							<a href="/board/{board.id}">
								<Card.Root class="transition-colors hover:bg-accent cursor-pointer mb-4">
									<Card.Header>
										<Card.Title class="text-lg">{board.name}</Card.Title>
										<Card.Description>
											Created: {new Date(board.created_at).toLocaleDateString()}
										</Card.Description>
									</Card.Header>
								</Card.Root>
							</a>
						{/each}
					</div>
				{/if}
			</Card.Content>
		</Card.Root>
	</div>
</div>

<CreateBoardDialog
	bind:open={showCreateDialog}
	onConfirm={handleCreateBoard}
/>

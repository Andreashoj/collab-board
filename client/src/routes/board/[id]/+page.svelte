<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import { goto } from '$app/navigation';
	import type { PageData } from './$types';
	import { Button } from '$lib/components/ui/button';
	import * as Card from '$lib/components/ui/card';
	import { Badge } from '$lib/components/ui/badge';
	import RenameDialog from '$lib/components/dialogs/rename-dialog.svelte';
	import ConfirmDialog from '$lib/components/dialogs/confirm-dialog.svelte';
	import BoardSkeleton from '$lib/components/loaders/board-skeleton.svelte';
	import { 
		boardsStore, 
		currentBoard, 
		currentMembers, 
		currentLogs, 
		boardsLoading, 
		boardsError 
	} from '$lib/stores/boards';

	let { data }: { data: PageData } = $props();

	let showRenameDialog = $state(false);
	let showDeleteDialog = $state(false);

	async function handleRename(newName: string) {
		await boardsStore.updateBoard(data.boardId, newName);
	}

	async function handleDelete() {
		await boardsStore.deleteBoard(data.boardId);
		goto('/dashboard');
	}

	onMount(() => {
		boardsStore.fetchBoard(data.boardId);
	});

	onDestroy(() => {
		boardsStore.clearCurrentBoard();
	});
</script>

<div class="min-h-screen p-8">
	<div class="mx-auto max-w-6xl">
		<!-- Header -->
		<div class="mb-8">
			<Button variant="outline" href="/dashboard">
				← Back to Dashboard
			</Button>
		</div>

		{#if $boardsLoading}
			<BoardSkeleton />
		{:else if $boardsError}
			<Card.Root class="border-destructive">
				<Card.Header>
					<Card.Title class="text-destructive">Error</Card.Title>
				</Card.Header>
				<Card.Content>
					<p class="text-destructive">{$boardsError}</p>
					<Button href="/dashboard" class="mt-4" variant="destructive">
						Go to Dashboard
					</Button>
				</Card.Content>
			</Card.Root>
		{:else if $currentBoard}
			<!-- Board Header -->
			<Card.Root class="mb-6">
				<Card.Header>
					<div class="flex items-start justify-between">
						<div>
							<Card.Title class="text-3xl">{$currentBoard.name}</Card.Title>
							<Card.Description class="mt-2">
								Created: {new Date($currentBoard.created_at).toLocaleString()}
							</Card.Description>
							<code class="text-xs text-muted-foreground mt-1 block">{$currentBoard.id}</code>
						</div>
						<div class="flex gap-2">
							<Button onclick={() => showRenameDialog = true}>
								Rename
							</Button>
							<Button variant="destructive" onclick={() => showDeleteDialog = true}>
								Delete
							</Button>
						</div>
					</div>
				</Card.Header>
			</Card.Root>

			<div class="grid grid-cols-1 gap-6 lg:grid-cols-2">
				<!-- Members Section -->
				<Card.Root>
					<Card.Header>
						<Card.Title>Members ({$currentMembers.length})</Card.Title>
					</Card.Header>
					<Card.Content>
						{#if $currentMembers.length === 0}
							<p class="text-muted-foreground">No members yet</p>
						{:else}
							<div class="space-y-3">
								{#each $currentMembers as member}
									<div class="flex items-center justify-between rounded-lg border p-3">
										<div>
											<p class="font-medium">{member.user?.email || 'Unknown'}</p>
											<p class="text-sm text-muted-foreground">Role: {member.role}</p>
										</div>
										<Badge variant="secondary">{member.role}</Badge>
									</div>
								{/each}
							</div>
						{/if}
					</Card.Content>
				</Card.Root>

				<!-- Activity Log Section -->
				<Card.Root>
					<Card.Header>
						<Card.Title>Activity Log ({$currentLogs.length})</Card.Title>
					</Card.Header>
					<Card.Content>
						{#if $currentLogs.length === 0}
							<p class="text-muted-foreground">No activity yet</p>
						{:else}
							<div class="space-y-2 max-h-96 overflow-y-auto">
								{#each $currentLogs as log}
									<div class="rounded-lg border p-3">
										<p class="text-sm">{log.change}</p>
										<p class="mt-1 text-xs text-muted-foreground">
											{new Date(log.created_at).toLocaleString()}
										</p>
									</div>
								{/each}
							</div>
						{/if}
					</Card.Content>
				</Card.Root>
			</div>

			<!-- Board Canvas - Placeholder for collaboration features -->
			<Card.Root class="mt-6">
				<Card.Header>
					<Card.Title>Board Canvas</Card.Title>
				</Card.Header>
				<Card.Content>
					<div class="flex h-96 items-center justify-center rounded-lg bg-muted">
						<p class="text-muted-foreground">Collaboration canvas coming soon...</p>
					</div>
				</Card.Content>
			</Card.Root>
		{/if}
	</div>
</div>

{#if $currentBoard}
	<RenameDialog
		bind:open={showRenameDialog}
		currentName={$currentBoard.name}
		onConfirm={handleRename}
	/>

	<ConfirmDialog
		bind:open={showDeleteDialog}
		title="Delete Board"
		description="Are you sure you want to delete this board? This action cannot be undone and will remove all associated data."
		confirmText="Delete"
		onConfirm={handleDelete}
	/>
{/if}

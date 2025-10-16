<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog';
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';

	let { 
		open = $bindable(false),
		onConfirm
	}: {
		open?: boolean;
		onConfirm: (name: string) => Promise<void>;
	} = $props();

	let name = $state('');
	let loading = $state(false);
	let error = $state<string | null>(null);

	async function handleSubmit() {
		if (!name.trim()) {
			error = 'Board name cannot be empty';
			return;
		}

		loading = true;
		error = null;

		try {
			await onConfirm(name);
			open = false;
			name = '';
		} catch (err: any) {
			error = err.message || 'Failed to create board';
		} finally {
			loading = false;
		}
	}

	// Reset when dialog opens/closes
	$effect(() => {
		if (!open) {
			name = '';
			error = null;
		}
	});
</script>

<Dialog.Root bind:open>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Create New Board</Dialog.Title>
			<Dialog.Description>
				Give your board a name to get started.
			</Dialog.Description>
		</Dialog.Header>

		<form onsubmit={(e) => { e.preventDefault(); handleSubmit(); }} class="space-y-4">
			<div class="space-y-2">
				<Label for="board-name">Board Name</Label>
				<Input
					id="board-name"
					bind:value={name}
					placeholder="My Awesome Board"
					disabled={loading}
					autofocus
				/>
				{#if error}
					<p class="text-sm text-destructive">{error}</p>
				{/if}
			</div>

			<Dialog.Footer>
				<Button
					type="button"
					variant="outline"
					onclick={() => open = false}
					disabled={loading}
				>
					Cancel
				</Button>
				<Button type="submit" disabled={loading}>
					{loading ? 'Creating...' : 'Create Board'}
				</Button>
			</Dialog.Footer>
		</form>
	</Dialog.Content>
</Dialog.Root>

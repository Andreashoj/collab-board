<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog';
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';

	let { 
		open = $bindable(false),
		currentName,
		onConfirm
	}: {
		open?: boolean;
		currentName: string;
		onConfirm: (newName: string) => Promise<void>;
	} = $props();

	let name = $state(currentName);
	let loading = $state(false);
	let error = $state<string | null>(null);

	async function handleSubmit() {
		if (!name.trim()) {
			error = 'Name cannot be empty';
			return;
		}

		if (name === currentName) {
			open = false;
			return;
		}

		loading = true;
		error = null;

		try {
			await onConfirm(name);
			open = false;
		} catch (err: any) {
			error = err.message || 'Failed to update name';
		} finally {
			loading = false;
		}
	}

	// Reset when dialog opens/closes
	$effect(() => {
		if (open) {
			name = currentName;
			error = null;
		}
	});
</script>

<Dialog.Root bind:open>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Rename</Dialog.Title>
			<Dialog.Description>
				Enter a new name for this item.
			</Dialog.Description>
		</Dialog.Header>

		<form onsubmit={(e) => { e.preventDefault(); handleSubmit(); }} class="space-y-4">
			<div class="space-y-2">
				<Label for="name">Name</Label>
				<Input
					id="name"
					bind:value={name}
					placeholder="Enter name..."
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
					{loading ? 'Saving...' : 'Save'}
				</Button>
			</Dialog.Footer>
		</form>
	</Dialog.Content>
</Dialog.Root>

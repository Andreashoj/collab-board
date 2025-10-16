<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog';
	import { Button } from '$lib/components/ui/button';

	let { 
		open = $bindable(false),
		title = 'Are you sure?',
		description,
		confirmText = 'Confirm',
		cancelText = 'Cancel',
		variant = 'destructive',
		onConfirm
	}: {
		open?: boolean;
		title?: string;
		description: string;
		confirmText?: string;
		cancelText?: string;
		variant?: 'default' | 'destructive' | 'outline' | 'secondary' | 'ghost' | 'link';
		onConfirm: () => Promise<void>;
	} = $props();

	let loading = $state(false);

	async function handleConfirm() {
		loading = true;
		try {
			await onConfirm();
			open = false;
		} catch (err) {
			console.error('Confirmation action failed:', err);
		} finally {
			loading = false;
		}
	}
</script>

<Dialog.Root bind:open>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>{title}</Dialog.Title>
			<Dialog.Description>
				{description}
			</Dialog.Description>
		</Dialog.Header>

		<Dialog.Footer>
			<Button
				type="button"
				variant="outline"
				onclick={() => open = false}
				disabled={loading}
			>
				{cancelText}
			</Button>
			<Button
				type="button"
				{variant}
				onclick={handleConfirm}
				disabled={loading}
			>
				{loading ? 'Processing...' : confirmText}
			</Button>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>

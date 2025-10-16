<script lang="ts">
	import { Skeleton } from '$lib/components/ui/skeleton';
	import * as Card from '$lib/components/ui/card';

	let { count = 3 }: { count?: number } = $props();
</script>

<div class="space-y-3">
	{#each Array(count) as _, i}
		<Card.Root>
			<Card.Header>
				<Skeleton class="h-6 w-3/4" />
				<Skeleton class="h-4 w-1/2 mt-2" />
			</Card.Header>
		</Card.Root>
	{/each}
</div>

<script lang="ts">
	import { Skeleton } from '$lib/components/ui/skeleton';
	import * as Card from '$lib/components/ui/card';
</script>

<div class="space-y-6">
	<!-- Header Skeleton -->
	<Card.Root>
		<Card.Header>
			<div class="flex items-start justify-between">
				<div class="space-y-2 flex-1">
					<Skeleton class="h-8 w-1/2" />
					<Skeleton class="h-4 w-1/3" />
					<Skeleton class="h-3 w-1/4" />
				</div>
				<div class="flex gap-2">
					<Skeleton class="h-10 w-20" />
					<Skeleton class="h-10 w-20" />
				</div>
			</div>
		</Card.Header>
	</Card.Root>

	<!-- Content Grid Skeleton -->
	<div class="grid grid-cols-1 gap-6 lg:grid-cols-2">
		{#each [1, 2] as _ (_)}
			<Card.Root>
				<Card.Header>
					<Skeleton class="h-6 w-1/3" />
				</Card.Header>
				<Card.Content>
					<div class="space-y-3">
						{#each [1, 2, 3] as _ (_)}
							<div class="rounded-lg border p-3">
								<Skeleton class="h-4 w-3/4" />
								<Skeleton class="h-3 w-1/2 mt-2" />
							</div>
						{/each}
					</div>
				</Card.Content>
			</Card.Root>
		{/each}
	</div>

	<!-- Canvas Skeleton -->
	<Card.Root>
		<Card.Header>
			<Skeleton class="h-6 w-1/4" />
		</Card.Header>
		<Card.Content>
			<Skeleton class="h-96 w-full" />
		</Card.Content>
	</Card.Root>
</div>

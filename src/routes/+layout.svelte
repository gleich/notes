<script lang="ts">
	import type { Snippet } from 'svelte';
	import '@gleich/ui/styles.css';
	import '../global.css';

	const { children }: { children: Snippet } = $props();
</script>

<main>
	<div class="container">
		{@render children()}
	</div>
</main>

<style>
	main {
		min-height: 100vh;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.container {
		display: flex;
		flex-direction: column;
		max-width: 1000px;
	}
</style>

<script lang="ts">
	import type { Snippet } from 'svelte';
	import '@gleich/ui/styles.css';
	import '../global.css';
	import { Copyright } from '@gleich/ui';

	const { children }: { children: Snippet } = $props();
</script>

<main>
	<div class="container">
		{@render children()}
		<div class="copyright">
			<Copyright repo="gleich/notes" />
		</div>
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
		width: 100%;
		flex: 1;
		padding: 10px;
	}

	.copyright {
		margin-top: 30px;
	}
</style>

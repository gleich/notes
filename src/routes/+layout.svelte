<script lang="ts">
	import type { Snippet } from 'svelte';
	import '@gleich/ui/styles.css';
	import '../global.css';
	import { Layout } from '@gleich/ui';

	const { children }: { children: Snippet } = $props();
</script>

<Layout repo="gleich/notes">
	{@render children()}
</Layout>

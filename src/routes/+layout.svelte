<script lang="ts">
	import type { Snippet } from 'svelte';
	import '@gleich/ui/styles.css';
	import '../global.css';
	import { Layout, siteName } from '@gleich/ui';

	siteName.set('notes.mattglei.ch');

	const { children }: { children: Snippet } = $props();
</script>

<Layout repo="gleich/notes">
	{@render children()}
</Layout>

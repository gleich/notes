<script lang="ts">
	import { NavLogo } from '@gleich/ui';
	import notes from './notes.json';
</script>

<div class="header">
	<NavLogo width="55px" />
	<h1>Matt Gleich's Notes</h1>
</div>

{#each notes as note (note.slug)}
	<a href={note.slug}>{note.title}</a>
{/each}

<style>
	.header {
		display: flex;
		gap: 10px;
	}
</style>

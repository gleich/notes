<script lang="ts">
	import { Card, DynamicHead, NavLogo } from '@gleich/ui';
	import notes from './notes.json';
</script>

<DynamicHead
	title="Notes"
	description="Matt Gleich's personal notes from college and life in general"
/>

<div class="header">
	<NavLogo width="55px" />
	<h1>Matt Gleich's Notes</h1>
</div>

<div class="notes">
	{#each notes as note (note.slug)}
		<Card>
			<a href={note.slug} class="note">
				<h4>{note.title}</h4>
				<div class="note-data">
					<p>{note.slug.replaceAll('/', ' → ')}</p>
				</div>
			</a>
		</Card>
	{/each}
</div>

<style>
	.header {
		display: flex;
		gap: 10px;
	}

	.notes {
		margin-top: 30px;
		display: grid;
		grid-template-columns: 1fr 1fr;
	}

	.note {
		color: inherit;
		text-decoration: none;
		display: flex;
		flex-direction: column;
		justify-content: space-between;
	}

	.note-data {
		color: grey;
	}
</style>

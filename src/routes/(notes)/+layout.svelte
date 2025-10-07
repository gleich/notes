<script lang="ts">
	import { page } from '$app/state';
	import type { Snippet } from 'svelte';
	import notes from '../notes.json';
	import { error } from '@sveltejs/kit';
	import { Card, DynamicHead, NavLogo, Scrolling } from '@gleich/ui';
	import dayjs from 'dayjs';
	import advancedFormat from 'dayjs/plugin/advancedFormat';

	dayjs.extend(advancedFormat);

	const { children }: { children: Snippet } = $props();
	const note = notes.find((n) => n.slug === page.url.pathname.slice(1));
	const folderPath = note ? dayjs(note.date).format('dddd, MMMM Do, YYYY') : 'Note not found';
	if (!note) {
		error(404);
	}
</script>

<DynamicHead title={note.title} description={folderPath} />

<div class="header">
	<a href="/" class="title">
		<div class="left">
			<NavLogo />
			<Scrolling>
				<h1>{note.title}</h1>
			</Scrolling>
		</div>
		<div class="right">
			<div class="data">
				<p>{folderPath}</p>
				<p>{dayjs(note.date).format('dddd, MMMM Do, YYYY')}</p>
			</div>
		</div>
	</a>
</div>

<Card padding="0">
	<div class="note-body">
		{@render children()}
	</div>
</Card>

<style>
	.header {
		margin-bottom: 20px;
	}

	.title {
		display: flex;
		gap: 10px;
		justify-content: space-between;
		color: inherit;
		text-decoration: inherit;
	}

	.left {
		display: flex;
		gap: 10px;
		align-items: center;
		justify-content: center;
	}

	.data {
		color: grey;
		display: flex;
		flex-direction: column;
		align-items: flex-end;
	}

	.note-body {
		padding: 10px;
	}

	@media (max-width: 800px) {
		.title {
			flex-direction: column;
			align-items: flex-start;
		}

		.data {
			align-items: flex-start;
		}
	}

	:global(
		.note-body h1,
		.note-body h2,
		.note-body h3,
		.note-body h4,
		.note-body h5,
		.note-body h6
	) {
		margin-top: 40px;
		margin-bottom: 10px;
		color: var(--green-foreground);
		background-color: var(--green-background);
		padding: 5px 10px;
		border-radius: 3px;
	}

	:global(.note-body > :is(h1, h2, h3, h4, h5, h6):first-child) {
		margin-top: 0;
	}
</style>

<script lang="ts">
	import { page } from '$app/state';
	import type { Snippet } from 'svelte';
	import notes from '../notes.json';
	import { error } from '@sveltejs/kit';
	import { Card, DynamicHead, NavLogo } from '@gleich/ui';
	import dayjs from 'dayjs';
	import advancedFormat from 'dayjs/plugin/advancedFormat';

	dayjs.extend(advancedFormat);

	const { children }: { children: Snippet } = $props();
	const note = notes.find((n) => n.slug === page.url.pathname.slice(1));
	if (!note) {
		error(404);
	}
</script>

<DynamicHead title={note.title} description="foo bar" />

<div class="header">
	<div class="title">
		<NavLogo width="50px" />
		<h1>{note.title}</h1>
	</div>
	<div class="data">
		<p>{note.slug.replaceAll('/', ' → ')}</p>
		<p>{dayjs(note.date).format('dddd, MMMM Do, YYYY')}</p>
	</div>
</div>

<Card padding="0">
	<div class="body">
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
		align-items: center;
	}

	.data {
		color: grey;
	}

	.body {
		padding: 10px;
	}
</style>

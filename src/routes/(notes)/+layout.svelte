<script lang="ts">
	import { page } from '$app/state';
	import type { Snippet } from 'svelte';
	import notes from '../notes.json';
	import { error } from '@sveltejs/kit';
	import { Card, DynamicHead, NavLogo, Scrolling } from '@gleich/ui';
	import dayjs from 'dayjs';
	import advancedFormat from 'dayjs/plugin/advancedFormat';
	import { resolve } from '$app/paths';

	dayjs.extend(advancedFormat);

	const { children }: { children: Snippet } = $props();

	const note = notes.find((n) => n.slug === page.url.pathname.slice(1));
	const paths: string[] = [];
	if (note) {
		const parts = note.slug.split('/').slice(0, -1);
		let path = '/';
		for (const part of parts) {
			path += part + '/';
			paths.push(path);
		}
	} else {
		error(404);
	}
	const date = dayjs(note.date).format('MMMM Do, YYYY [at] h:mma');
</script>

<DynamicHead title={note.title} description={note ? `${note.slug}, ${date}` : 'Note not found'} />

<div class="header">
	<div class="title">
		<a href={resolve('/')} class="left">
			<NavLogo width="50px" height="50px" />
			<Scrolling>
				<h1>{note.title}</h1>
			</Scrolling>
		</a>
		<div class="right">
			<div class="data">
				{#if note}
					<p>
						{#each paths as path, i (path)}
							<!-- eslint-disable-next-line svelte/no-navigation-without-resolve -->
							{#if i > 0}<span>/</span>{/if}<a href={path} class="path-url"
								>{path.split('/').at(-2)}</a
							>
						{/each}
					</p>
					<p>{date}</p>
				{/if}
			</div>
		</div>
	</div>
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
	}

	.left {
		display: flex;
		gap: 10px;
		align-items: center;
		justify-content: center;
		color: inherit;
		text-decoration: inherit;
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

	.path-url {
		color: inherit;
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
		margin-top: 30px;
		margin-bottom: 10px;
		color: var(--green-foreground);
		background-color: var(--green-background);
		padding: 5px 10px;
		border-radius: 3px;
	}

	:global(.note-body h1) {
		font-size: 23px;
	}

	:global(.note-body h2) {
		font-size: 20px;
	}

	:global(.note-body h3) {
		font-size: 15px;
	}

	:global(.note-body h4, .note-body h5, .note-body h6) {
		font-size: 13px;
		font-style: normal;
	}

	:global(.note-body > :is(h1, h2, h3, h4, h5, h6):first-child) {
		margin-top: 0;
	}

	:global(.note-body .drawing) {
		border: 1px solid var(--border);
		box-shadow: var(--box-shadow);
		border-radius: var(--border-radius);
		padding: 50px 30px;
		display: flex;
		justify-content: center;
		max-width: 100%;
	}

	:global(.note-body .drawing .drawing-scale) {
		max-width: 100%;
		width: calc(100% / 2.2);
		display: flex;
		align-content: center;
		justify-content: center;
	}

	:global(.note-body .drawing .drawing-scale img) {
		max-width: 100%;
		height: auto;

		transform: scale(2.2);
		transform-origin: center;
	}

	/* CODE RELATED STYLES */

	:global(.note-body pre) {
		border: 1px solid var(--border);
		padding: 10px;
		box-shadow: var(--box-shadow);
		border-radius: var(--border-radius);
		margin-top: 10px;
		overflow-x: auto;
		font-size: 14px;
	}

	:global(.note-body code:not(pre code)) {
		font-size: 13.5px;
		border-radius: 3px;
		border: 0.5px solid var(--border);
		padding: 0 3px;
		font-family: 'IBM Plex Mono';
		background-color: var(--section-name-background-color);
	}

	@media (prefers-color-scheme: dark) {
		:global(.note-body .drawing img) {
			filter: invert(1);
		}
	}

	@media (max-width: 500px) {
		:global(.note-body .drawing) {
			padding: 35px 10px;
		}
	}
</style>

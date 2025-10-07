import { mdsvex } from 'mdsvex';
import adapter from '@sveltejs/adapter-cloudflare';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';
import { createHighlighter } from 'shiki';
import { escapeSvelte } from 'mdsvex';

const theme = 'gruvbox-dark-soft';
const highlighter = await createHighlighter({
	themes: [theme],
	langs: ['javascript', 'typescript', 'python', 'c']
});

const config = {
	preprocess: [
		vitePreprocess(),
		mdsvex({
			smartypants: true,
			extension: '.md',
			highlight: {
				highlighter: async (code, lang = 'text') => {
					const html = escapeSvelte(highlighter.codeToHtml(code, { lang, theme }));
					return `{@html \`${html}\` }`;
				}
			}
		})
	],
	kit: { adapter: adapter() },
	extensions: ['.svelte', '.md']
};

export default config;

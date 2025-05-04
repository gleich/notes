# move script

Takes the files located in `notes/`:

```txt
.
└── college
    └── sofa-103
        ├── README.md
        └── week-13
            └── displays.md
```

And creates the following directory structure in `src/routes/(notes)` for SvelteKit to prerender:

```txt
.
└── college
    └── sofa-103
        └── week-13
            └── displays
                └── +page.md
```

Also creates a JSON file with information about all of these routes located in `src/routes/notes.json`:

```json
[
	{
		"title": "Displays",
		"slug": "college/sofa-103/week-13/displays",
		"date": "2025-04-28T00:00:00Z"
	}
]
```

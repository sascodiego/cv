# web

Astro 6.4 static site that renders the processed CV from `src/data/cv-processed.json`.

See the [root README](../README.md) for the full project overview.

## Commands

| Command | What it does |
|---------|-------------|
| `pnpm dev` | Start dev server at `localhost:4321` |
| `pnpm build` | Build static site to `dist/` |
| `pnpm preview` | Preview the built site |
| `pnpm export-pdf` | Build + generate PDF with private data via Puppeteer |

## Private Data

Set `PRIVATE_PHONE`, `PRIVATE_ADDRESS`, and `PRIVATE_CV_NOTE` in `web/.env` for local PDF export. These variables are never set in CI, so the public GitHub Pages build omits them.

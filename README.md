# cv-pipeline

Single-source-of-truth CV system for **Diego Sasco**. One YAML file feeds a public web CV (GitHub Pages) and a local PDF with private data — no sensitive information ever leaves your machine.

```
data/cv-source.yaml
        │
        │  cd pipeline && go run .
        ▼
web/src/data/cv-processed.json      ← skills enriched with overlap-aware experience
        │
        │  cd web && pnpm build
        ▼
web/dist/                           ← public static site → GitHub Pages
        │
        │  pnpm export-pdf          ← injects PRIVATE_* from .env
        ▼
Diego_Sasco_CV_Privado.pdf          ← local-only PDF with full contact info
```

## Quick Start

### Prerequisites

| Tool | Version |
|------|---------|
| Go | 1.26+ |
| Node.js | ≥ 22.12 |
| pnpm | 10+ |

### Run the pipeline

```bash
# 1. Process YAML → JSON (from repo root)
cd pipeline && go run .

# 2. Preview the web CV locally
cd web && pnpm install && pnpm dev
```

### Generate the private PDF

```bash
# Create web/.env with your private data (never committed)
cd web
cp .env.example .env   # then fill in your values

# Build + export
pnpm export-pdf
```

Output: `Diego_Sasco_CV_Privado.pdf` at the repo root.

## Architecture

| Layer | Responsibility |
|-------|---------------|
| `data/cv-source.yaml` | Canonical source — jobs, skills, projects, education (no personal contact data) |
| `pipeline/` (Go) | YAML → JSON processor; computes per-skill experience with interval merging |
| `web/` (Astro) | Static site renderer; public contact via Google Form button only |
| `web/scripts/generate-pdf.js` | Local-only PDF export — injects private data via Puppeteer DOM manipulation |
| `.github/workflows/` | CI: builds public site + deploys to GitHub Pages |

### Security model

- **No PII in source:** `cv-source.yaml` contains professional content only (name + title). No email, phone, address, or social links.
- **No private data in Astro:** The web app is purely public. It renders a "Contact" button linked to a Google Form. Zero knowledge of PDF or private data.
- **Public build (GitHub Actions):** `PUBLIC_GOOGLE_FORM_URL` is injected from GitHub Secrets. No other env vars are set.
- **Local PDF export:** `web/.env` contains all private data. `generate-pdf.js` uses Puppeteer to inject it into the DOM before capturing. The `.env` file is gitignored and never leaves your notebook.

## Project Structure

```
cv-pipeline/
├── data/
│   └── cv-source.yaml            # Single source of truth
├── pipeline/                     # Go: YAML → processed JSON
│   ├── main.go                   # CLI entrypoint
│   ├── models.go                 # Data types (CVDate, CVRaw, CVProcessed)
│   ├── calculator.go             # Duration math, interval merging, experience calc
│   └── calculator_test.go        # Unit tests
├── web/                          # Astro: static site + PDF export
│   ├── src/
│   │   ├── pages/index.astro     # Single-page CV renderer
│   │   ├── data/cv-processed.json
│   │   └── styles/global.css     # Dark premium theme
│   ├── scripts/generate-pdf.js   # Puppeteer PDF generator
│   ├── astro.config.mjs
│   └── package.json
├── docs/
│   └── init.md                   # Full technical proposal (ES)
├── .github/workflows/deploy.yml  # GitHub Pages deployment
└── README.md                     # You are here
```

## Tech Stack

| Area | Technology |
|------|-----------|
| Data processing | Go 1.26, `gopkg.in/yaml.v3` |
| Web frontend | Astro 6.4, plain TypeScript, CSS custom properties |
| PDF generation | Puppeteer (headless Chrome) |
| Deployment | GitHub Actions → GitHub Pages |
| Fonts | Inter, Outfit (Google Fonts) |

## Available Commands

| Command | Where | What it does |
|---------|-------|-------------|
| `go run .` | `pipeline/` | Process `cv-source.yaml` → `cv-processed.json` |
| `go test` | `pipeline/` | Run unit tests |
| `pnpm dev` | `web/` | Start Astro dev server |
| `pnpm build` | `web/` | Build static site to `web/dist/` |
| `pnpm preview` | `web/` | Preview built site locally |
| `pnpm export-pdf` | `web/` | Build + generate private PDF via Puppeteer |

## GitHub Pages Setup

The site deploys automatically on every push to `main`. To enable it:

1. Go to **Settings → Pages** in your GitHub repository.
2. Under **Source**, select **GitHub Actions**.
3. Push to `main` — the workflow builds and deploys automatically.

The public site URL will be:
```
https://<YOUR_USERNAME>.github.io/cv-pipeline/
```

## Editing Your CV

All content lives in `data/cv-source.yaml`. Edit that file, then re-run the pipeline and/or rebuild the site.

Key sections in the YAML:

- `personal_info` — name and title only (contact data lives in `.env`)
- `skills` — array of `{ id, name, category }` entries
- `work_experience` — jobs with date ranges and technology references
- `projects` — linked to jobs via `id`; orphan projects appear independently
- `education` — degrees and certifications
- `recommendations` — quotes from colleagues

## License

Private repository. All rights reserved.

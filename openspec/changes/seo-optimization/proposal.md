# Proposal: SEO Optimization for CV Site

## Change ID
`seo-optimization`

## Problem Statement
The CV site at https://sascodiego.github.io/cv/ scores 4/10 on SEO. It lacks critical technical SEO infrastructure (robots.txt, sitemap, social sharing meta), has incomplete structured data, and includes an unnecessary "Built with Astro" footer. This limits search discoverability and social sharing quality.

## Intent
Maximize SEO score and social sharing quality for the single-page CV site through technical SEO, on-page optimization, structured data, and social meta tags.

## Goals
- Remove "Built with Astro" footer
- Add robots.txt and auto-generated sitemap via `@astrojs/sitemap`
- Add Twitter Card and complete Open Graph meta tags
- Expand Schema.org Person with contact info, social profiles, and location
- Optimize title tag with specialties and location
- Add `og:image` support (placeholder until real image is designed)
- Add web manifest for PWA

## Non-Goals
- Do NOT change the visual design or layout
- Do NOT add multi-language/i18n pages
- Do NOT change cv-processed.json data schema
- Do NOT touch PDF-related files or cv-pdf.astro
- Do NOT add analytics or tracking scripts
- Real OG image design is out of scope (use placeholder SVG)

## Scope
- `web/src/pages/index.astro` — meta tags, footer, structured data
- `web/astro.config.mjs` — add sitemap integration
- `web/package.json` — add @astrojs/sitemap
- `web/public/robots.txt` — new file
- `web/public/manifest.json` — new file

## Estimated Lines
~80 lines changed across 5 files. Well under 800-line review budget.

## Conflict Risk
Low. PDF work adds new `cv-pdf.astro`; SEO touches only `index.astro` and config files.

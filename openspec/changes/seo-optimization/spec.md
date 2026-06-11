# Spec: SEO Optimization — `seo-optimization`

## Functional Requirements

### FR-1: Remove "Built with Astro" footer
The `<footer class="site-footer">` containing the Astro attribution link must be removed from `index.astro`.

### FR-2: robots.txt
Create `web/public/robots.txt`:
- Allow all crawlers
- Disallow nothing (single public page)
- Reference sitemap at `https://sascodiego.github.io/cv/sitemap-index.xml`

### FR-3: Auto-generated sitemap
- Install `@astrojs/sitemap` as dependency
- Add `sitemap()` integration to `astro.config.mjs`
- This auto-generates `sitemap-index.xml` at build time

### FR-4: Open Graph image
- Create `web/public/og-image.svg` as a minimal placeholder (branded text-based card)
- Add `og:image` and `og:image:width`, `og:image:height`, `og:image:alt` meta tags
- Use SVG for now; real PNG can replace later

### FR-5: Twitter Card meta tags
Add to `<head>`:
- `twitter:card` = `summary_large_image`
- `twitter:title` (reuse title)
- `twitter:description` (reuse description)
- `twitter:image` (reuse og:image)

### FR-6: Optimized title tag
Change from `{name} | {title}` to:
`{name} — {title} | Distributed Systems & IoT`

This adds keyword context without being spammy.

### FR-7: Expanded Schema.org Person
Add to existing JSON-LD:
- `address` with `PostalAddress` (Montevideo, UY)
- `sameAs` array with LinkedIn and GitHub profile URLs
- `alumniOf` with university info from education data
- `knowsAbout` expanded from cv-processed.json skills

### FR-8: Web manifest
Create `web/public/manifest.json`:
- `name`, `short_name`, `start_url`, `display: standalone`
- `background_color`, `theme_color` matching site design
- `icons` referencing existing favicon.svg

## Non-Functional Requirements

### NFR-1: Lighthouse SEO target ≥ 90
After implementation, Lighthouse SEO audit should score ≥ 90.

### NFR-2: No visual regression
All changes are in `<head>`, config, or new public files. No CSS or layout changes.

### NFR-3: Zero JS added
SEO changes must not add client-side JavaScript.

## Acceptance Criteria
- [ ] robots.txt returns valid content at deployed URL
- [ ] sitemap-index.xml is auto-generated and accessible
- [ ] All OG tags present (title, description, type, url, locale, image)
- [ ] All Twitter Card tags present
- [ ] Schema.org Person validates at validator.schema.org
- [ ] Title tag contains specialties
- [ ] "Built with Astro" footer removed
- [ ] manifest.json is valid JSON
- [ ] `pnpm build` succeeds
- [ ] No visual changes to the page

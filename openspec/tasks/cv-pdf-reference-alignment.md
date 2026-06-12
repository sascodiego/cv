# Tasks: CV PDF Reference Alignment

**Change ID:** `cv-pdf-reference-alignment`

## Review Workload Forecast

| Field | Value |
|-------|-------|
| Estimated changed lines | 250–350 |
| 400-line budget risk | Low |
| Chained PRs recommended | No |
| Suggested split | Single PR |
| Delivery strategy | auto-chain |
| Chain strategy | size-exception |

Decision needed before apply: No
Chained PRs recommended: No
Chain strategy: size-exception
400-line budget risk: Low

---

## Phase 1: Technology Label Mapper (Low Risk)

### 1.1 Add technology label mapper object
- [ ] In `web/scripts/generate-pdf.js`, add `TECH_LABEL_MAP` constant at file top
- [ ] Populate with common technology mappings (csharp→C#, typescript→TypeScript, javascript→JavaScript, python→Python, go→Go, rust→Rust, java→Java, kotlin→Kotlin, swift→Swift, c→C, cpp→C++, ruby→Ruby, php→PHP, scala→Scala, haskell→Haskell, elixir→Elixir, erlang→Erlang)
- [ ] Add framework mappings (react→React, vue→Vue.js, angular→Angular, svelte→Svelte, nextjs→Next.js, nuxt→Nuxt.js, express→Express.js, fastapi→FastAPI, django→Django, flask→Flask, spring→Spring, springboot→Spring Boot, aspnet→ASP.NET, laravel→Laravel)
- [ ] Add database mappings (postgresql→PostgreSQL, mysql→MySQL, sqlite→SQLite, mongodb→MongoDB, redis→Redis, elasticsearch→Elasticsearch)
- [ ] Add cloud and infrastructure mappings (aws→AWS, azure→Azure, gcp→GCP, kubernetes→Kubernetes, docker→Docker, terraform→Terraform, ansible→Ansible)
- [ ] Sort all entries alphabetically by key for maintainability

### 1.2 Implement normalization function
- [ ] Add `normalizeTechLabel(tag)` function in `web/scripts/generate-pdf.js`
- [ ] Implement null/undefined guard: return empty string for invalid input
- [ ] Convert tag to lowercase and trim whitespace
- [ ] Check explicit mapping in `TECH_LABEL_MAP` and return mapped value if found
- [ ] Implement fallback: capitalize first character and keep rest as-is for unmapped tags
- [ ] Add JSDoc comment documenting function behavior and return values

### 1.3 Integrate mapper with DOM injection
- [ ] Locate `injectPrivateData()` function in `web/scripts/generate-pdf.js`
- [ ] Modify function to accept `TECH_LABEL_MAP` as second argument to `page.evaluate()`
- [ ] Add DOM manipulation code to find all skill list items (`.skill-list li`)
- [ ] Implement regex to replace technology names before parentheses: `/([a-zA-Z0-9_\-]+)\s*\(/g`
- [ ] Apply `normalizeTechLabel()` to matched technology names
- [ ] Replace text content of each skill item with normalized labels
- [ ] Add error handling for DOM manipulation failures

### 1.4 Test mapper in isolation
- [ ] Run `cd /src/cv-pipeline/web && pnpm export-pdf`
- [ ] Inspect generated PDF skills section for normalized labels
- [ ] Verify `csharp` → `C#`, `typescript` → `TypeScript`, `javascript` → `JavaScript`
- [ ] Verify fallback capitalization works for unmapped tags
- [ ] Confirm `cv-processed.json` source data unchanged (still contains lowercase IDs)

---

## Phase 2: Typography Scale (Medium Risk)

### 2.1 Update font sizes in print CSS
- [ ] In `web/src/pages/cv-pdf.astro`, locate `@media print` block
- [ ] Update body font size: `9.25pt` → `9.75pt`
- [ ] Update `.cv-header h1` font size: `14pt` → `14.5pt`
- [ ] Update `.professional-title` font size: `9.5pt` → `10.5pt`
- [ ] Update `h2.section-header` font size: `10pt` → `10.75pt`
- [ ] Update company/institution name font size: confirm `9.8pt` for `h3.company-name` and `h3.institution-name`
- [ ] Update role/degree font size: confirm `9pt` for `.role-title` and `.degree-title`
- [ ] Update `.entry-bullets li` font size: `8.95pt` → `9.5pt`
- [ ] Update `.skill-list li` font size: `8.95pt` → `9.5pt`
- [ ] Update contact items font size: confirm `8.75pt` for `.contact-item`
- [ ] Update `.summary-text` font size: `9pt` → `9.5pt`

### 2.2 Update line heights in print CSS
- [ ] Update body line-height: `1.16` → `1.3`
- [ ] Update `.cv-header h1` line-height: `1.05` → `1.06`
- [ ] Update `.professional-title` line-height: confirm `1.2`
- [ ] Update `h2.section-header` line-height: add explicit `1.15`
- [ ] Update `.entry-bullets li` line-height: `1.15` → `1.2`
- [ ] Update `.skill-list li` line-height: `1.15` → `1.2`
- [ ] Update `.summary-text` line-height: `1.18` → `1.3`

### 2.3 Update letter spacing in print CSS
- [ ] Confirm `.cv-header h1` letter-spacing at `1pt`
- [ ] Update `h2.section-header` letter-spacing: `0.35pt` → `0.4pt`
- [ ] Verify all other elements use default letter-spacing

### 2.4 Test typography changes
- [ ] Run `cd /src/cv-pipeline/web && pnpm export-pdf`
- [ ] Open generated PDF and verify text is more legible than before
- [ ] Compare with reference PDF for proportional similitude
- [ ] Verify page count is 1–2 pages (accept slight expansion)
- [ ] Confirm no text overflow or layout breaks

---

## Phase 3: Whitespace and Rhythm (Medium Risk)

### 3.1 Verify page margins
- [ ] Confirm `@page` margin values: `16mm 15mm 16mm 15mm` (top, right, bottom, left)
- [ ] Verify Puppeteer margin config in `generate-pdf.js` matches CSS

### 3.2 Update section spacing
- [ ] Update `.cv-header` padding-bottom: `6pt` → `8pt`
- [ ] Update `.cv-header` margin-bottom: `8pt` → `10pt`
- [ ] Update `h2.section-header` margin: `10pt 0 4pt` → `12pt 0 6pt`
- [ ] Update `h2.section-header` padding-bottom: `0 0 2pt` → `0 0 3pt`
- [ ] Add explicit margin-top to `.section`: `0` (no change, verify consistency)

### 3.3 Update entry spacing
- [ ] Update `.entry` margin: `0 0 8pt` → `0 0 10pt`
- [ ] Update `.entry-header` margin: `0 0 2pt` → `0 0 4pt`
- [ ] Update `.role-title` margin: `0 0 4pt` → `0 0 5pt`
- [ ] Update `.degree-title` margin: `0 0 4pt` → `0 0 5pt`
- [ ] Update `.entry-bullets` margin-top: `2pt` → `4pt`
- [ ] Update `.entry-bullets li` margin-bottom: `0 0 2pt` → `0 0 3pt`

### 3.4 Update header block spacing
- [ ] Update `.professional-title` margin: `2pt 0 4pt` → `3pt 0 5pt`
- [ ] Update `.contact-info` margin-top: confirm `5pt` from title
- [ ] Update `.summary-text` margin-bottom: `0 0 8pt` → `0 0 10pt`

### 3.5 Update skill and reference spacing
- [ ] Update `.skill-list li` margin-bottom: confirm `3pt`
- [ ] Update `.references-note` margin-top: `10pt 0 0` → `12pt 0 0`
- [ ] Update `.education-status` margin-top: confirm `3pt`

### 3.6 Test rhythm changes
- [ ] Run `cd /src/cv-pipeline/web && pnpm export-pdf`
- [ ] Review PDF for clear visual separation between sections
- [ ] Check for orphaned content or awkward page breaks
- [ ] Verify margins provide adequate framing
- [ ] Confirm whitespace feels intentional and professional

---

## Phase 4: Contact Block Refinement (Low Risk)

### 4.1 Verify contact block treatment
- [ ] Confirm all icons are removed (check for any SVG or icon classes)
- [ ] Verify contact items are text-only (no hyperlinks in PDF)
- [ ] Confirm separator character is middle dot `·` (U+00B7)
- [ ] Verify `.contact-separator` padding is `0 4pt`

### 4.2 Refine contact block spacing
- [ ] Verify `.professional-title` to `.contact-info` spacing: `5pt`
- [ ] Verify `.contact-info` to `.summary-text` spacing: via `.cv-header` margin-bottom `10pt`
- [ ] Confirm `.contact-info` uses `flex` layout with `justify-content: center`
- [ ] Verify all contact items have `white-space: nowrap`

### 4.3 Verify contact block ordering
- [ ] Confirm ordering: address, LinkedIn, phone, email, portfolio, GitHub
- [ ] Test with actual private data injection to verify rendering

### 4.4 Test contact block
- [ ] Run `cd /src/cv-pipeline/web && pnpm export-pdf`
- [ ] Review contact line for correct ordering and spacing
- [ ] Check for line wrapping issues on narrow previews
- [ ] Verify middle dot separators are visible and correctly spaced

---

## Phase 5: Build and Export Verification

### 5.1 Verify web build passes
- [ ] Run `cd /src/cv-pipeline/web && pnpm build`
- [ ] Confirm build completes without errors or warnings
- [ ] Verify output directory contains generated files

### 5.2 Verify PDF export succeeds
- [ ] Run `cd /src/cv-pipeline/web && pnpm export-pdf`
- [ ] Confirm PDF generation completes without errors
- [ ] Verify output PDF file exists in correct location
- [ ] Check PDF file size is reasonable (< 500KB)

### 5.3 Verify web page untouched
- [ ] Run `git diff web/src/pages/index.astro` — confirm no changes
- [ ] Run `git diff web/src/styles/` — confirm no changes to site styles
- [ ] Verify public site deployment remains unaffected
- [ ] Check that only `cv-pdf.astro` and `generate-pdf.js` were modified

### 5.4 Verify data source unchanged
- [ ] Inspect `cv-processed.json` — confirm technology tags are still lowercase IDs
- [ ] Verify no changes to data schema or structure
- [ ] Confirm Go pipeline output remains unchanged
- [ ] Check that `cv-source.yaml` is unmodified

---

## Phase 6: Final Visual Verification

### 6.1 Typography verification
- [ ] Compare generated PDF with reference PDF side-by-side
- [ ] Verify font sizes feel larger and more legible than before
- [ ] Confirm line heights provide comfortable reading rhythm
- [ ] Check that uppercase headers have appropriate letter-spacing
- [ ] Verify font hierarchy is clear: name > headers > body > bullets > dates

### 6.2 Document rhythm verification
- [ ] Verify margins provide adequate framing (16mm top/bottom, 15mm left/right)
- [ ] Confirm section spacing creates clear visual separation
- [ ] Check that entry spacing distinguishes individual jobs/degrees
- [ ] Verify whitespace feels intentional, not accidental
- [ ] Confirm no cramped or orphaned content

### 6.3 Technology label verification
- [ ] **Critical**: Verify `csharp` → `C#` normalization in skills section
- [ ] Verify `typescript` → `TypeScript`, `javascript` → `JavaScript`, `python` → `Python`
- [ ] Confirm fallback capitalization works for unmapped tags
- [ ] Check that technology labels are human-friendly and professional

### 6.4 Contact block verification
- [ ] Confirm all icons are absent (text-only presentation)
- [ ] Verify middle dot separators with 4pt padding
- [ ] Check correct ordering: address, LinkedIn, phone, email, portfolio, GitHub
- [ ] Confirm contact items are readable and not cramped

### 6.5 Visual parity verification
- [ ] Compare overall layout with reference PDF
- [ ] Verify proportions feel similar (not pixel-perfect)
- [ ] Confirm document rhythm matches reference flow
- [ ] Check that visual hierarchy guides reader through content
- [ ] Verify formality and professionalism match reference

### 6.6 ATS compatibility verification
- [ ] Open PDF in viewer and select text — confirm all text is selectable
- [ ] Verify no content is rendered as images
- [ ] Confirm single-column layout is maintained
- [ ] Check that semantic HTML structure is preserved (h1, h2, h3, ul, li)
- [ ] Verify bullets use standard list markers

### 6.7 Page format verification
- [ ] Verify PDF is A4 size (210mm × 297mm)
- [ ] Confirm background is white (#fff)
- [ ] Verify page count is 1–2 pages (accept slight expansion for readability)
- [ ] Check that page breaks respect content structure (no orphaned headers)

---

## Phase 7: Regression and Rollback Validation

### 7.1 Regression checks
- [ ] Run `cd /src/cv-pipeline/web && pnpm build` — confirm passes
- [ ] Run `cd /src/cv-pipeline/web && pnpm export-pdf` — confirm succeeds
- [ ] Verify public web page is unchanged (check `index.astro` has no git diff)
- [ ] Confirm no changes to data pipeline or source files
- [ ] Check that private data injection still works correctly

### 7.2 Rollback validation (prepare rollback steps)
- [ ] Document rollback command: `git checkout web/src/pages/cv-pdf.astro`
- [ ] Document rollback command: `git checkout web/scripts/generate-pdf.js`
- [ ] Verify rollback commands are accurate
- [ ] Confirm changes fit within 400-line review budget (count changed lines)

---

## Summary

**Total tasks:** 58
**Files modified:** 2
- `web/src/pages/cv-pdf.astro` (~150–200 lines changed)
- `web/scripts/generate-pdf.js` (~100–150 lines changed)

**Estimated changed lines:** 250–350
**Risk level:** Low-Medium
**Delivery strategy:** Single PR (fits within 400-line budget)

**Key verification points:**
- Technology labels normalized: `csharp` → `C#`
- Public web untouched: no changes to `index.astro` or site styles
- Data source unchanged: `cv-processed.json` still contains lowercase IDs
- Typography improved: larger fonts, better line heights
- Document rhythm enhanced: better margins and spacing
- Visual parity achieved: similitude with reference PDF
- ATS compatibility maintained: selectable text, single-column, semantic HTML
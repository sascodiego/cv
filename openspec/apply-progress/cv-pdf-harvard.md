# Apply Progress: cv-pdf-harvard

## Status: Completed

## Changes Made

### New File: `web/src/pages/cv-pdf.astro` (~520 lines)
- Dedicated Astro template for PDF rendering at route `/cv-pipeline/cv-pdf`
- Frontmatter logic: skill category mapping (9→4 Harvard), project extraction from work_experience, summary truncation, bullet splitting with char limits
- 7 sections: Header, Professional Summary, Technical Skills, Professional Experience, Key Projects, Education, References
- Inline `@media print` CSS: Times New Roman, 9pt body, white background, black text, page-break control
- Inline `@media screen` CSS: Arial preview with shadow container
- DOM selectors for Puppeteer: `.cv-header h1`, `.contact-info`, `.recommendation-card`, `.pdf-note`
- Contact item CSS (`.contact-item svg`) constraining SVG icons to 9pt to prevent overflow

### Modified File: `web/scripts/generate-pdf.js` (4 changes)
1. Line 183: URL changed to `http://localhost:4321/cv/cv-pdf` (was `/cv-pipeline/`, corrected for base path `/cv/`)
2. Line 51: Selector changed to `.cv-header h1` (was `.hero h1`)
3. Line 230: `printBackground: false` (was `true`)
4. Line 129: `.cv-container` selector (was `main`) for pdf-note injection

## Issues Resolved During Apply
1. **5 pages → 2 pages**: SVG contact icons (viewBox 512x512) rendered without CSS constraints, inflating page count. Fixed by adding `.contact-item svg { width: 9pt; height: 9pt; }` CSS.
2. **Summary truncation**: `truncateSummary` regex split on "/." in "C#/.NET". Fixed by using lookbehind split on `. ` followed by uppercase letter.
3. **Page selector**: Template uses `<div class="cv-container">` not `<main>`. Fixed `querySelector("main")` → `querySelector(".cv-container")`.
4. **Base path**: Astro config uses `base: "/cv/"` not `/cv-pipeline/`. Fixed URL in generate-pdf.js.
5. **Font size/spacing**: Original 11pt with generous spacing produced 4 pages even after SVG fix. Reduced to 9pt body, tighter margins (10mm top/bottom, 12mm sides), compressed spacing throughout.

## Verification Results
- `pnpm build`: ✅ Passes (2 pages built: index.html, cv-pdf/index.html)
- `pnpm export-pdf`: ✅ Generates valid PDF
- PDF pages: 2
- PDF format: A4 (595.92 × 841.92 pts)
- Text: Selectable (ATS-friendly)
- Private data: Contact info + recommender names injected correctly
- All 7 sections present and correctly ordered

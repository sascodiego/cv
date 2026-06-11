# Proposal: Harvard-Style CV PDF Generator

## Change ID

`cv-pdf-harvard`

## Title

Dedicated Harvard-style CV PDF template optimized for hiring and ATS compatibility

## Problem Statement

The current PDF generator (`web/scripts/generate-pdf.js`) uses Puppeteer to print the dark-themed web page as-is. The result is NOT a professional CV — it's a styled web dump with dark backgrounds, zinc color scheme, visual effects, and web-specific layout patterns that are inappropriate for traditional hiring workflows and ATS (Applicant Tracking System) processing.

This creates a product problem: users cannot generate a print-ready, ATS-optimized CV that follows established academic and industry standards. The web page's editorial dark design does not translate to a professional document suitable for job applications, academic submissions, or formal résumé distribution.

## Intent

Create a dedicated Harvard-style CV PDF template that produces a professional, print-ready document optimized for hiring workflows and ATS compatibility, while maintaining the single data source (`cv-processed.json`) with the web page.

## Goals

- Create a dedicated Astro template (`web/src/pages/cv-pdf.astro`) for PDF rendering only
- Share `cv-processed.json` as single data source with the web page
- Implement independent layout, typography, section order, and formatting from the web design
- Follow Harvard GSAS/FAS CV standards: serif/sans-serif fonts, 10-12pt, equal margins, no shading/text boxes
- Use action-verb bullets with quantifiable results (CAR framework)
- Produce 1-2 pages, A4 format, white background, professional print quality
- Inject private data from `.env` (email, phone, address, LinkedIn, GitHub, portfolio, recommenders)
- Ensure ATS-friendly output: text-based, single-column, no decorative elements
- Maintain no JavaScript interactivity in the PDF template

## Non-Goals

- Do NOT redesign the web page
- Do NOT change the data schema (`cv-processed.json`) unless a styling blocker is found
- Do NOT add JavaScript interactivity to the PDF template
- Do NOT generate a dark-themed PDF
- Do NOT change the web page's existing visual design or functionality

## User and Product Impact

- Job applicants receive a professional, ATS-optimized CV that meets established academic and industry standards
- Recruiters and hiring managers receive a clean, scannable document optimized for traditional CV workflows
- ATS systems can reliably parse and extract structured information from the PDF
- The CV owner maintains a single source of truth (`cv-processed.json`) while supporting dual output formats
- Private contact data remains secure in `.env` and is only injected during PDF generation

## Scope

### In Scope

- **New template**: `web/src/pages/cv-pdf.astro`
  - Dedicated PDF-only rendering with inline print CSS
  - Harvard-style layout: single-column, white background, professional typography
  - Font stack: Times New Roman or Arial, 10-12pt
  - Equal margins: at least 0.5 inch (recommended 15mm)
  - No text boxes, no underlining, no shading
  - Section headers: bold, uppercase or title case, consistent size, clear hierarchy

- **Section order** (Harvard hybrid for industry + technical):
  1. Header: Full name, professional title, contact row
  2. Professional Summary: 2-3 concise lines
  3. Technical Skills: Grouped by category
  4. Professional Experience: Reverse chronological, action-verb bullets
  5. Key Projects: Selected, with tech stack
  6. Education: Institution, degree, period, status
  7. References: Named recommenders from .env or "Available on request"

- **Modified script**: `web/scripts/generate-pdf.js`
  - Change navigation from `/cv-pipeline/` to `/cv-pipeline/cv-pdf`
  - Keep private data injection logic (contact info, recommenders)
  - Keep A4 format and 15mm margins
  - Ensure `printBackground: false` (white background)

- **Print CSS**: Inline in `cv-pdf.astro`
  - `@media print` rules for professional document output
  - Page-break control to avoid orphaned sections
  - Ensure selectable text (no content as images)
  - Remove dark theme remnants completely

### Out of Scope

- Changes to `web/src/pages/index.astro` (web page)
- Changes to `web/src/styles/global.css` (web styles)
- New data fields in `cv-processed.json`
- Light/dark theme toggles
- Interactive elements (buttons, hover effects, animations)
- Multiple PDF templates or variants
- Changes to the Go data pipeline

## Affected Areas

- `web/src/pages/cv-pdf.astro` — new file
- `web/scripts/generate-pdf.js` — single line change (URL)
- `web/src/data/cv-processed.json` — used as-is (no changes)
- `.env` — private data source (no changes)

## Design Decisions

- **Typography**: System font stack prioritizing Times New Roman (serif) or Arial (sans-serif) for maximum ATS compatibility
- **Font sizes**: 10-12pt body text, larger for headers, consistent hierarchy
- **Margins**: 15mm on all sides (approximately 0.6 inch), exceeding minimum 0.5 inch
- **Layout**: Strict single-column, no grids or multi-section arrangements
- **Colors**: Black text on white background, no decorative colors or backgrounds
- **Section headers**: Bold, uppercase, no underlining or text boxes
- **Bullets**: Standard bullets or simple dashes, action-verb driven, 1-3 lines per role
- **Contact row**: Single line with email, phone, location, LinkedIn, GitHub, portfolio
- **References**: Full name, role, and relation for each recommender from `.env`, or "References available on request" if not provided

## Section Order and Formatting Standards

### 1. Header
- **Name**: Bold, centered or left-aligned, larger font (14-16pt)
- **Professional title**: Regular weight, below name
- **Contact row**: Single line, comma-separated, smaller font (10pt)
  - Email, phone, address, LinkedIn, GitHub, portfolio
  - Icons optional (SVG or Unicode), kept minimal

### 2. Professional Summary
- 2-3 concise lines
- Present tense, no first-person pronouns
- Focus on years of experience, key domains, technical strengths

### 3. Technical Skills
- Grouped by category (Programming Languages, Frameworks, Tools, etc.)
- Each category with bold header
- Skills as comma-separated list or simple bullets
- Include months/years of experience if available

### 4. Professional Experience
- Reverse chronological order
- Each entry:
  - **Company name**: Bold, left-aligned
  - **Role title**: Regular weight, below company name
  - **Period**: Right-aligned or on same line as role (e.g., "Jan 2020 – Present")
  - **Location**: Optional, below role or on same line
  - **Bullets**: 1-3 per role, action verbs, CAR framework, measurable outcomes

### 5. Key Projects
- Selected projects (3-5 maximum)
- Each entry:
  - **Project name**: Bold, left-aligned
  - **Period**: Right-aligned
  - **Tech stack**: Comma-separated, smaller font
  - **Bullets**: 1-2 per project, action verbs, measurable outcomes

### 6. Education
- Reverse chronological order
- Each entry:
  - **Institution**: Bold, left-aligned
  - **Degree**: Regular weight, below institution
  - **Period**: Right-aligned
  - **Status**: Below degree (e.g., "Completed", "In Progress")

### 7. References
- Option A: Full recommender details (if all `.env` vars provided)
  - Name, role, relation for each recommender
  - 3 maximum
- Option B: "References available on request" (if any missing data)

## Technical Architecture

### Data Flow

```
cv-source.yaml → Go pipeline → cv-processed.json
                                          ↓
                                    (shared data)
                                          ↓
                ┌─────────────────────────┴─────────────────────────┐
                ↓                                                   ↓
        web/src/pages/index.astro                           web/src/pages/cv-pdf.astro
        (web page, dark zinc)                                 (PDF, white, Harvard)
```

### Build Process

1. **Data generation**: Go pipeline produces `cv-processed.json` (unchanged)
2. **Web build**: `pnpm build` builds both pages (`index.astro` and `cv-pdf.astro`)
3. **PDF export**: `pnpm export-pdf` runs `generate-pdf.js`
   - Starts Astro preview server
   - Navigates to `/cv-pipeline/cv-pdf`
   - Injects private data from `.env` into DOM
   - Generates PDF via Puppeteer

### Private Data Injection

The existing `injectPrivateData()` function in `generate-pdf.js` will work with the new template:
- Updates hero name
- Builds contact info block
- Injects recommender attribution into `<footer class="recommendation-author">` elements

### Print CSS Strategy

- Inline `@media print` in `cv-pdf.astro`
- Override all web-specific styling
- Force white background: `background-color: white; color: black;`
- Ensure margins and page-break control
- Remove all dark theme variables and effects

## Acceptance Criteria

- **Build**: `cd /src/cv-pipeline/web && pnpm build` passes without errors
- **PDF generation**: `cd /src/cv-pipeline/web && pnpm export-pdf` generates valid PDF at `/src/cv-pipeline/Diego_Sasco_CV_Privado.pdf`
- **Format**: PDF is 1-2 pages, A4, white background, professional layout
- **Visual**: No dark theme remnants, no text boxes, no shading, no underlining
- **Content**: All sections present and correctly ordered
- **Typography**: Times New Roman or Arial, 10-12pt, consistent throughout
- **Margins**: Equal on all sides, at least 0.5 inch (15mm recommended)
- **Private data**: Contact info and recommenders injected correctly from `.env`
- **ATS-friendly**: Selectable text (no images for content), single-column, parseable by ATS
- **Page breaks**: No orphaned sections, logical break points

## Risks

- **Schema blockers**: If `cv-processed.json` lacks required fields for Harvard formatting, may require minor schema changes
- **Font rendering**: System fonts may vary across platforms; test on Windows, macOS, and Linux
- **Private data missing**: If `.env` vars are incomplete, PDF may show placeholders instead of contact info
- **Page overflow**: Dense content may exceed 2 pages; may require content curation or font size adjustment
- **Print CSS conflicts**: Global CSS may interfere with print-specific rules; use specificity and !pragma appropriately

## Rollback Plan

- Revert `web/scripts/generate-pdf.js` to navigate back to `/cv-pipeline/`
- Delete `web/src/pages/cv-pdf.astro`
- No data migration or schema changes needed
- Web page and existing functionality remain untouched

## Success Criteria

- PDF visually matches Harvard GSAS/FAS standards
- ATS systems can parse all content sections reliably
- Recruiters receive a clean, scannable document optimized for traditional workflows
- Single data source (`cv-processed.json`) maintained across web and PDF outputs
- Implementation remains within the 800-line review budget
- No regression in web page functionality or appearance

## First-Slice Recommendation

Implement the first slice as a focused PDF template:

1. Create `web/src/pages/cv-pdf.astro` with basic Harvard structure
2. Implement all 7 sections in correct order
3. Add inline print CSS with typography, margins, and white background
4. Update `generate-pdf.js` to navigate to `/cv-pipeline/cv-pdf`
5. Verify PDF output meets all formatting standards
6. Test ATS compatibility with sample parsing tools

Avoid content rewriting, schema changes, or web page modifications in the first slice unless strictly necessary for PDF correctness.

## Review Workload Forecast

Estimated implementation diff: approximately 300–400 changed lines
- `web/src/pages/cv-pdf.astro`: ~250–300 lines (new file)
- `web/scripts/generate-pdf.js`: ~1 line (URL change)
- Total: within the configured 800-line review budget

## Open Questions

No blocking open questions remain. Optional product refinements for later review:

- Should contact icons be SVG, Unicode, or plain text for maximum ATS compatibility?
- Should project selection be automatic (first 5) or manual (flag field in data)?
- Should "References available on request" be the default if any recommender data is missing, or only if all are missing?

## Dependencies

- `cv-processed.json` must be generated before PDF export (existing pipeline)
- `.env` file with `PRIVATE_*` variables must exist for contact data injection
- Puppeteer must be installed (already in `web/package.json`)
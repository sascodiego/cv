# Harvard-Style CV PDF Template — Technical Specification

## Purpose

This specification defines the technical requirements for a dedicated Harvard-style CV PDF template that produces a professional, print-ready document optimized for hiring workflows and ATS (Applicant Tracking System) compatibility.

## Requirements

### Requirement: Template Structure and HTML Semantics

The system MUST provide a dedicated Astro template at `web/src/pages/cv-pdf.astro` with a semantic HTML structure optimized for PDF rendering and ATS parsing.

#### Scenario: Template loads and renders all sections

- GIVEN the user runs `pnpm build`
- WHEN the build process includes `web/src/pages/cv-pdf.astro`
- THEN the template MUST render at route `/cv-pipeline/cv-pdf`
- AND the template MUST include all 7 sections in the correct order: Header, Professional Summary, Technical Skills, Professional Experience, Key Projects, Education, References
- AND the template MUST use semantic HTML5 elements: `<header>`, `<section>`, `<article>`, `<footer>`, `<ul>`, `<li>`

#### Scenario: Template provides DOM selectors for Puppeteer injection

- GIVEN the template renders HTML
- WHEN the Puppeteer script in `generate-pdf.js` navigates to the page
- THEN the template MUST provide these DOM selectors:
  - `.cv-header h1` for the name element
  - `.contact-info` container for contact items (cleared and rebuilt by Puppeteer)
  - `.recommendation-card footer` for recommender attribution injection
- AND these selectors MUST exist even if initially empty or containing placeholder content

---

### Requirement: CSS and Print Stylesheet Specification

The system MUST provide inline `@media print` CSS in `cv-pdf.astro` that overrides all web-specific styling and produces a professional Harvard-style document.

#### Scenario: Print CSS enforces white background and black text

- GIVEN the page is rendered
- WHEN the browser applies `@media print` rules
- THEN the background MUST be pure white (`#ffffff`)
- AND all text MUST be black (`#000000`)
- AND no decorative backgrounds, gradients, or dark theme colors MUST appear

#### Scenario: Typography follows Harvard standards

- GIVEN the page is rendered
- WHEN the print CSS applies typography rules
- THEN the font stack MUST prioritize Times New Roman (serif) followed by Arial (sans-serif)
- AND body text MUST be 10-12pt
- AND the name in the header MUST be 14-16pt, bold
- AND section headers MUST be bold, uppercase, 12-14pt
- AND contact row text MUST be 10pt, regular weight

#### Scenario: Margins equal on all sides

- GIVEN the page is rendered
- WHEN the print CSS applies margin rules
- THEN margins MUST be 15mm on all sides (top, bottom, left, right)
- AND margins MUST NOT include decorative borders or text boxes

#### Scenario: Page-break control prevents orphaned sections

- GIVEN the page is rendered
- WHEN the print CSS applies page-break rules
- THEN sections MUST NOT break in the middle of a single entry
- AND section headers MUST NOT be orphaned at the bottom of a page
- AND page breaks MUST prefer breaks between major sections

#### Scenario: No decorative elements

- GIVEN the page is rendered
- WHEN the print CSS applies styling rules
- THEN there MUST be no text boxes
- AND there MUST be no underlining
- AND there MUST be no shading or background colors
- AND there MUST be no decorative borders or graphics (except minimal icons in contact row, optional)

---

### Requirement: Data Mapping from cv-processed.json

The system MUST map fields from `cv-processed.json` to the appropriate template sections without modifying the data schema.

#### Scenario: Personal info maps to header and summary

- GIVEN the `cv-processed.json` file contains `personal_info` object
- WHEN the template renders
- THEN `personal_info.name` MUST display in `.cv-header h1`
- AND `personal_info.title` MUST display below the name as professional title
- AND `personal_info.about_me` MUST display in the Professional Summary section (first 2-3 lines)

#### Scenario: Skills map to Technical Skills section by category

- GIVEN the `cv-processed.json` file contains `skills` array with 21 entries
- WHEN the template renders
- THEN skills MUST be grouped by `category` field
- AND each category MUST display as a bold header
- AND skills within each category MUST display as comma-separated list
- AND `experience_text` field MUST optionally display next to each skill (e.g., "Go (2 años)")
- AND categories from data MUST include: "Languages & Frameworks", "Databases & Messaging", "Tools & Platforms", "IoT & Embedded", and others present in data

#### Scenario: Work experience maps to Professional Experience section

- GIVEN the `cv-processed.json` file contains `work_experience` array with 6 entries
- WHEN the template renders
- THEN entries MUST display in reverse chronological order (most recent first)
- AND each entry MUST display:
  - Company name in bold
  - Role title below company name
  - Period text (from `period_text` field) right-aligned or on same line as role
  - Description as 1-3 action-verb bullets
  - Technologies as comma-separated list below bullets
- AND if `projects` array exists within an entry, those projects MUST be handled in the Key Projects section

#### Scenario: Projects from work_experience map to Key Projects section

- GIVEN the `cv-processed.json` file contains nested `projects` arrays within `work_experience` entries
- WHEN the template renders
- THEN projects MUST be selected from work experience entries (up to 5 total)
- AND each project MUST display:
  - Project name in bold
  - Period text (from `period_text` field) right-aligned
  - Tech stack as comma-separated list in smaller font
  - Description as 1-2 action-verb bullets
- AND if no projects exist, the Key Projects section MUST be omitted

#### Scenario: Education maps to Education section

- GIVEN the `cv-processed.json` file contains `education` array with 4 entries
- WHEN the template renders
- THEN entries MUST display in reverse chronological order (most recent first)
- AND each entry MUST display:
  - Institution name in bold
  - Degree title below institution
  - Period text right-aligned
  - Status below degree (e.g., "Completed", "In Progress")

#### Scenario: Recommendations map to References section

- GIVEN the `cv-processed.json` file contains `recommendations` array with 3 entries
- WHEN the template renders
- THEN the References section MUST display
- AND if all `.env` recommender variables are provided, full recommender details MUST display
- AND if any `.env` recommender variable is missing, the section MUST display "References available on request"
- AND recommender details MUST include: name, role, and relation from `.env`

---

### Requirement: Private Data Injection Contract

The system MUST support Puppeteer-based private data injection from `.env` via DOM manipulation as implemented in `generate-pdf.js`.

#### Scenario: Name injection via Puppeteer

- GIVEN the `.env` file contains `PRIVATE_NAME` variable
- WHEN Puppeteer calls `injectPrivateData(page)`
- THEN the script MUST find `.cv-header h1` selector
- AND the script MUST update the `textContent` with the private name value

#### Scenario: Contact info injection via Puppeteer

- GIVEN the `.env` file contains contact variables (`PRIVATE_EMAIL`, `PRIVATE_PHONE`, `PRIVATE_ADDRESS`, `PRIVATE_WEBSITE`, `PRIVATE_GITHUB`, `PRIVATE_LINKEDIN`)
- WHEN Puppeteer calls `injectPrivateData(page)`
- THEN the script MUST find `.contact-info` selector
- AND the script MUST clear all existing children of the container
- AND the script MUST create contact item elements with SVG icons and labels for each non-empty contact variable
- AND the script MUST append contact items in this order: Email, Phone, Address, Website, GitHub, LinkedIn

#### Scenario: Recommender injection via Puppeteer

- GIVEN the `.env` file contains recommender variables (`PRIVATE_RECOMMENDER_1_NAME`, `PRIVATE_RECOMMENDER_1_ROLE`, `PRIVATE_RECOMMENDER_1_RELATION`, etc. for 3 recommenders)
- WHEN Puppeteer calls `injectPrivateData(page)`
- THEN the script MUST find all `.recommendation-card` elements
- AND for each card (up to 3), the script MUST create a `<footer class="recommendation-author">` element
- AND the script MUST append `<cite class="author-name">` with recommender name
- AND if role exists, append `<span class="author-role">` with role
- AND if relation exists, append `<span class="author-relation">` with relation

---

### Requirement: generate-pdf.js Modifications

The system MUST modify `web/scripts/generate-pdf.js` to navigate to the new template route.

#### Scenario: URL change to new template

- GIVEN the existing `generate-pdf.js` navigates to `http://localhost:4321/cv-pipeline/`
- WHEN the script is modified
- THEN the navigation URL MUST change to `http://localhost:4321/cv-pipeline/cv-pdf`
- AND all other logic (Puppeteer launch, injection, PDF generation) MUST remain unchanged

#### Scenario: Print background disabled

- GIVEN the `page.pdf()` options in `generate-pdf.js`
- WHEN the script is modified
- THEN `printBackground` option MUST be set to `false` to ensure white background

---

### Requirement: Page Layout Specification

The system MUST produce a single-column, professional layout meeting Harvard GSAS/FAS standards.

#### Scenario: Single-column layout with consistent spacing

- GIVEN the page is rendered
- WHEN the print CSS applies layout rules
- THEN all content MUST flow in a single column
- AND spacing between sections MUST be consistent (e.g., 12-16pt or 0.5-0.75em)
- AND spacing between entries within sections MUST be smaller than section spacing (e.g., 8-12pt or 0.4-0.5em)

#### Scenario: Bullet styling for descriptions

- GIVEN the page is rendered
- WHEN descriptions display as bullet points
- THEN bullets MUST be standard solid bullets (`•`) or simple dashes
- AND bullet text MUST be 10-12pt, black, regular weight
- AND bullets MUST NOT use custom Unicode or decorative characters

#### Scenario: Section header formatting

- GIVEN the page is rendered
- WHEN section headers display
- THEN section headers MUST be bold, uppercase
- AND section headers MUST be left-aligned
- AND section headers MUST be followed by a consistent spacing rule (e.g., 4-6pt before next element)

---

### Requirement: ATS Compatibility Requirements

The system MUST ensure the PDF is parseable by ATS systems and follows established standards.

#### Scenario: Selectable text (no images for content)

- GIVEN the page is rendered
- WHEN the PDF is generated
- THEN all content MUST be selectable as text
- AND no content MUST be rendered as images
- AND text MUST NOT be embedded in canvas or SVG text elements

#### Scenario: Simple, hierarchical structure

- GIVEN the page is rendered
- WHEN ATS parsers analyze the document
- THEN the document MUST use clear, consistent formatting
- AND section headers MUST be clearly distinguished
- AND bullet lists MUST use standard HTML `<ul>`/`<li>` elements
- AND dates and periods MUST be in a consistent format

#### Scenario: No decorative elements that confuse parsers

- GIVEN the page is rendered
- WHEN ATS parsers analyze the document
- THEN there MUST be no text boxes
- AND there MUST be no multi-column layouts
- AND there MUST be no tables used for layout
- AND there MUST be no decorative backgrounds or borders
- AND icons in contact row MUST be optional or minimal

---

## Technical Decisions

### Font Stack Prioritization
The font stack prioritizes Times New Roman (serif) for maximum ATS compatibility, with Arial (sans-serif) as fallback. This aligns with Harvard GSAS/FAS standards and ensures readability across all ATS systems.

### Margins
15mm margins on all sides exceed the minimum 0.5 inch requirement and provide adequate whitespace for printing while maximizing content area.

### Section Order
The Harvard hybrid order (Header → Summary → Skills → Experience → Projects → Education → References) balances academic standards with technical industry expectations.

### Private Data Injection Strategy
Puppeteer DOM manipulation (rather than server-side rendering or API endpoints) keeps private data secure in `.env` and avoids exposing it to the web build or Git repository.

### Print CSS Inline in Template
Inline `@media print` rules in `cv-pdf.astro` (rather than separate CSS file) ensure portability and avoid conflicts with global web styles.

### Project Selection from Nested Data
Projects are extracted from `work_experience[].projects` rather than the top-level `projects` array (which is empty in current data), ensuring the Key Projects section populates from existing data structure.

---

## Known Issues and Constraints

### Existing .env Selector Mismatch
The current `injectPrivateData()` function uses `.hero h1` for name injection and expects a specific DOM structure from the web page. The new template MUST provide `.cv-header h1` as a replacement selector.

### Recommendation Section Fallback
If any recommender data is missing from `.env`, the entire References section SHOULD display "References available on request" rather than partial data, to maintain professionalism.

### Font Rendering Variability
System fonts may vary across platforms (Windows, macOS, Linux). Testing on all platforms is recommended to ensure consistent output.

### Page Overflow Risk
With dense content (21 skills, 6 work entries, 4 education entries, 3+ projects), the PDF may exceed 2 pages. Font size adjustment or content curation may be required.

---

## Success Criteria

- Build passes: `cd /src/cv-pipeline/web && pnpm build` succeeds without errors
- PDF generates: `pnpm export-pdf` produces valid PDF at `/src/cv-pipeline/Diego_Sasco_CV_Privado.pdf`
- Format: PDF is 1-2 pages, A4, white background, professional layout
- Visual: No dark theme remnants, no text boxes, no shading, no underlining
- Content: All 7 sections present and correctly ordered
- Typography: Times New Roman or Arial, 10-12pt body, 14-16pt name
- Margins: Equal 15mm on all sides
- Private data: Contact info and recommenders injected correctly from `.env`
- ATS-friendly: Selectable text, single-column, parseable by ATS systems
- Page breaks: No orphaned sections, logical break points
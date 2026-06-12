# Delta for CV PDF — Harvard Formalization

## Purpose

This delta specification defines modifications to the existing Harvard-style CV PDF to restore true Harvard GSAS/FAS formal standards by removing iconography and correcting typography, spacing, and document structure.

## ADDED Requirements

### Requirement: Text-Only Contact Block Presentation

The system MUST render contact information as text-only without SVG icons or graphical elements in the PDF.

#### Scenario: Contact items display as text-only with separators

- GIVEN the PDF is generated from `cv-pdf.astro`
- WHEN Puppeteer injects contact data from `.env`
- THEN contact items MUST display as text-only spans or anchor elements
- AND each contact item MUST use no SVG icons or graphical elements
- AND contact items MUST be separated by pipe characters (` | `)
- AND the final contact block MUST read as a single readable line (e.g., "email@example.com | +1 555-123-4567 | City, Country | portfolio.com")

#### Scenario: Contact info container uses text-only styling

- GIVEN the print CSS applies styling to `.contact-info`
- WHEN the contact info renders in the PDF
- THEN `.contact-info` MUST display as flex container with center alignment
- AND `.contact-item` elements MUST be inline-block or inline-flex
- AND `.contact-item svg` selector MUST be removed or set to `display: none`
- AND `.contact-item` MUST have no width/height constraints for icons

#### Scenario: Script creates text-only contact items

- GIVEN `generate-pdf.js` executes the `injectPrivateData()` function
- WHEN contact items are created and injected
- THEN `createContactItem()` MUST NOT accept `svgPath` parameter
- AND `createContactItem()` MUST NOT create or inject SVG elements
- AND `createContactItem()` MUST create `span` or `a` elements with `textContent` set to the label
- AND for URL-based items (email, website, GitHub, LinkedIn), the element MUST be an `<a>` tag with `href`
- AND for non-URL items (phone, address), the element MUST be a `<span>` tag

---

### Requirement: Harvard Formal Hierarchy Enforcement

The system MUST enforce a strict visual hierarchy that follows Harvard GSAS/FAS document conventions with specific font sizes, weights, and treatments.

#### Scenario: Name uses formal centering and uppercase treatment

- GIVEN the PDF header renders the name
- WHEN the print CSS applies to `.cv-header h1`
- THEN the name MUST be 14-16pt font size
- AND the name MUST be bold font weight
- AND the name MUST be text-transform uppercase
- AND the name MUST have letter-spacing of 1pt
- AND the name MUST be text-align center

#### Scenario: Professional title uses formal italic treatment

- GIVEN the PDF header renders the professional title
- WHEN the print CSS applies to `.professional-title`
- THEN the title MUST be 10-11pt font size
- AND the title MUST be normal font weight
- AND the title MUST be italic font style
- AND the title MUST be text-align center
- AND the title MUST have 2-4pt margin below the name

#### Scenario: Section headers use formal uppercase with bottom border

- GIVEN a section header renders (e.g., "PROFESSIONAL EXPERIENCE")
- WHEN the print CSS applies to `h2.section-header`
- THEN the header MUST be 10-12pt font size
- AND the header MUST be bold font weight
- AND the header MUST be text-transform uppercase
- AND the header MUST have letter-spacing of 0.3-0.5pt
- AND the header MUST have a bottom border of 1pt solid black
- AND the header MUST have 8-12pt margin above the section
- AND the header MUST have 2-4pt padding below the border

#### Scenario: Entry headers use formal bold treatment

- GIVEN an entry header renders (company name, institution name, project name)
- WHEN the print CSS applies to `h3.company-name`, `h3.institution-name`, or `h3.project-name`
- THEN the header MUST be 9.5-11pt font size
- AND the header MUST be bold font weight
- AND the header MUST be normal font style (not italic)
- AND the header MUST be left-aligned

#### Scenario: Metadata (period, role, degree) uses formal italic treatment

- GIVEN metadata renders (period text, role title, degree title)
- WHEN the print CSS applies to `.period`, `.role-title`, or `.degree-title`
- THEN the metadata MUST be 9-10pt font size
- AND the metadata MUST be normal font weight
- AND the metadata MUST be italic font style
- AND the period MUST be right-aligned in the entry header

---

### Requirement: Formal Spacing and Document Rhythm

The system MUST implement consistent spacing that creates visual rhythm and guides the reader through the document following academic document conventions.

#### Scenario: Section spacing creates clear visual separation

- GIVEN sections render sequentially
- WHEN the print CSS applies spacing rules
- THEN each section MUST have 8-12pt margin above its header
- AND each section MUST have 4-6pt margin below its last element
- AND the spacing between sections MUST be consistent throughout the document

#### Scenario: Entry spacing creates hierarchical rhythm

- GIVEN entries render within sections (experience, education, projects)
- WHEN the print CSS applies spacing rules
- THEN each entry MUST have 6-8pt margin below its last element
- AND entry spacing MUST be smaller than section spacing
- AND the spacing must be consistent for all entries of the same type

#### Scenario: Bullet spacing ensures readability

- GIVEN bullet lists render within entries
- WHEN the print CSS applies spacing rules
- THEN each bullet item MUST have 1-2pt margin below
- AND the bullet list MUST have 2-4pt margin above and below
- AND bullets MUST use standard disc markers (not custom characters)

#### Scenario: Margin and padding rules ensure document-first appearance

- GIVEN the document renders
- WHEN the print CSS applies margin/padding rules
- THEN the document MUST have equal margins on all sides (via Puppeteer PDF options)
- AND no element MUST use decorative padding or borders
- AND no element MUST use text boxes or background colors
- AND white space MUST be used intentionally for separation, not decoration

---

## MODIFIED Requirements

### Requirement: Private Data Injection Contract

(Previously: Contact items created with SVG icons and labels; icons allowed)

The system MUST support Puppeteer-based private data injection from `.env` via DOM manipulation, creating text-only contact elements without SVG icons.

#### Scenario: Contact info injection creates text-only elements

- GIVEN the `.env` file contains contact variables (`PRIVATE_EMAIL`, `PRIVATE_PHONE`, `PRIVATE_ADDRESS`, `PRIVATE_WEBSITE`, `PRIVATE_GITHUB`, `PRIVATE_LINKEDIN`)
- WHEN Puppeteer calls `injectPrivateData(page)`
- THEN the script MUST find `.contact-info` selector
- AND the script MUST clear all existing children of the container
- AND the script MUST create text-only contact item elements (no SVG icons)
- AND the script MUST append contact items in this order: Email, Phone, Address, Website, GitHub, LinkedIn
- AND the script MUST insert pipe separators (` | `) between items

#### Scenario: Script removes SVG icon creation logic

- GIVEN `generate-pdf.js` contains the `injectPrivateData()` function
- WHEN the script is modified for formalization
- THEN `createContactItem()` MUST NOT accept or use `svgPath` parameter
- AND all SVG icon constants (e.g., `svgEmail`, `svgPhone`) MUST be removed
- AND the SVG icon template creation logic MUST be removed
- AND the contact item creation MUST set only `textContent` or `href` attributes

---

### Requirement: CSS and Print Stylesheet Specification

(Previously: Icons allowed in contact row with `.contact-item svg` styling; spacing and typography may not meet Harvard standards)

The system MUST provide inline `@media print` CSS in `cv-pdf.astro` that removes all icon-related styling and enforces Harvard-standard typography and spacing.

#### Scenario: Print CSS removes icon-related rules

- GIVEN the print CSS applies to `.contact-item`
- WHEN the contact info renders in the PDF
- THEN `.contact-item svg` selector MUST be removed or set to `display: none !important`
- AND `.contact-item` MUST have no width/height constraints for icons
- AND `.contact-item` MUST use `display: inline-flex` with `gap: 3pt` for separator spacing

#### Scenario: Print CSS enforces Harvard-standard font sizes

- GIVEN the print CSS applies typography rules
- THEN body text MUST be 9-10pt (was 9pt, acceptable)
- AND the name MUST be 14-16pt (was 14pt, acceptable)
- AND section headers MUST be 10-12pt (was 10pt, acceptable)
- AND entry headers MUST be 9.5-11pt (was 9.5pt, acceptable)
- AND contact text MUST be 8.5-9pt (was 8.5pt, acceptable)

#### Scenario: Print CSS enforces Harvard-standard line heights

- GIVEN the print CSS applies line-height rules
- THEN body text MUST have line-height of 1.15-1.25 (was 1.15, acceptable)
- AND summary text MUST have line-height of 1.2-1.3 (was 1.25, acceptable)
- AND bullet items MUST have line-height of 1.15-1.25 (was 1.2, acceptable)

#### Scenario: Print CSS enforces Harvard-standard letter-spacing

- GIVEN the print CSS applies letter-spacing rules
- THEN the name MUST have letter-spacing of 0.8-1.2pt (was 1pt, acceptable)
- AND section headers MUST have letter-spacing of 0.3-0.5pt (was 0.3pt, acceptable)
- AND body text MUST have letter-spacing of 0pt (default, acceptable)

#### Scenario: Print CSS enforces Harvard-standard spacing

- GIVEN the print CSS applies spacing rules
- THEN section headers MUST have 8-12pt margin above (was 8pt, acceptable)
- AND entries MUST have 6-8pt margin below (was 6pt, acceptable)
- AND bullets MUST have 1-2pt margin below (was 1pt, acceptable)
- AND the header section MUST have 4-6pt margin below (was 6pt, may need adjustment)

---

## REMOVED Requirements

### Requirement: SVG Icon Display in Contact Row

(Reason: Icons are decorative elements inappropriate for formal Harvard-style CVs; they add visual clutter, distract from document-first presentation, and may cause ATS parsing issues)

(Migration: Contact information will display as text-only with pipe separators; no data migration needed; users with `.env` contact variables will see the same information in a more formal presentation)

---

## Technical Constraints

### Do NOT Modify Public Web Page

The `web/src/pages/index.astro` file and its associated styles MUST NOT be modified. This change is PDF-only scope.

### Do NOT Compromise Harvard Formality for Page Density

Page count expansion is acceptable in this first pass. Restoring formal spacing and typography takes priority over fitting content onto 2 pages. Page density optimization, if needed, will be addressed in a follow-up change.

### Preview Fidelity Is Not Required

The screen preview (`@media screen` in `cv-pdf.astro`) may look different from the PDF output. Only the final PDF must meet Harvard formality standards. Preview inconsistencies are acceptable.

### Single-Column Layout Only

The PDF MUST use a single-column layout. No multi-column or grid-based layouts are permitted.

### No Decorative Web-Style Elements

The PDF MUST NOT contain:
- SVG icons or graphical elements
- Text boxes or background colors
- Underlining (except for anchor links, which should be black without decoration)
- Shading or gradients
- Decorative borders (except the 1pt bottom border on section headers)
- Custom Unicode or decorative bullet characters

---

## Acceptance Criteria

- **Icon removal**: PDF contains zero icons; all contact information is text-only with pipe separators
- **Typography**: Font sizes, line heights, and letter-spacing match Harvard GSAS/FAS standards
- **Spacing**: Section, entry, and bullet spacing follow academic document conventions
- **Hierarchy**: Clear visual hierarchy (name > section headers > entry headers > body) with appropriate treatments
- **Document-first**: PDF feels like a formal academic document, not a web page
- **ATS-friendly**: Clean, parseable output without visual noise
- **Web untouched**: `web/src/pages/index.astro` remains unchanged
- **Build success**: `cd /src/cv-pipeline/web && pnpm build` passes without errors
- **PDF generation**: `cd /src/cv-pipeline/web && pnpm export-pdf` generates valid PDF
- **Format**: PDF is 1-2 pages (acceptable if 2+ pages due to formal spacing), A4, white background
- **Private data**: Contact info and recommenders injected correctly from `.env`

---

## Implementation Notes

### Separator Strategy

Contact items will use pipe characters (` | `) as separators. This is the most common convention for formal CVs and provides clear visual separation without adding clutter.

### Icon Removal in generate-pdf.js

The `createContactItem()` function will be simplified from:

```javascript
function createContactItem(svgPath, label, href) {
  const el = document.createElement(href ? "a" : "div");
  el.className = "contact-item";
  // ... inject SVG icon ...
  const span = document.createElement("span");
  span.textContent = label;
  el.appendChild(span);
  contactInfo.appendChild(el);
}
```

To:

```javascript
function createContactItem(label, href) {
  const el = document.createElement(href ? "a" : "span");
  el.className = "contact-item";
  if (href) {
    el.href = href;
    el.target = "_blank";
    el.rel = "noopener noreferrer";
  }
  el.textContent = label;
  contactInfo.appendChild(el);
}
```

### CSS Removal in cv-pdf.astro

The following CSS rules will be removed or modified:

```css
/* REMOVE THIS */
.contact-item svg {
  width: 9pt;
  height: 9pt;
  fill: #000000;
}

/* MODIFY THIS to remove icon-related constraints */
.contact-item {
  display: inline-flex;
  align-items: center;
  gap: 3pt;  /* Keep for separator spacing */
  font-size: 8.5pt;
  color: #000000;
}
```

### Typography Verification

After implementation, verify the following against the PDF output:
- Name: 14-16pt, bold, uppercase, centered
- Title: 10-11pt, italic, centered
- Section headers: 10-12pt, bold, uppercase, with bottom border
- Entry headers: 9.5-11pt, bold, left-aligned
- Body text: 9-10pt, regular, left-aligned
- Line heights: 1.15-1.25 throughout
- Letter-spacing: 0.8-1.2pt for name, 0.3-0.5pt for headers, 0pt for body

---

## Risks

- **Page density**: Restoring formal spacing may expand page count beyond 2 pages. This is acceptable in this first pass. Page density optimization can be addressed later if needed by slightly reducing font sizes or spacing within Harvard-acceptable ranges.

- **User preference**: Some users may prefer icons for visual identification. This change defers to formal Harvard standards, which prioritize document-first presentation over visual decoration.

- **ATS parsing**: While text-only contact info is more ATS-friendly than icons, there is a small risk that parsers may not correctly identify email, phone, and URL patterns without icon cues. Testing with ATS systems is recommended.

- **Preview inconsistency**: The screen preview in `cv-pdf.astro` may look different from the PDF output. This is acceptable per the decision that only the final PDF matters.

---

## Success Criteria

- PDF contains zero icons; all contact information is text-only with pipe separators
- Typography and spacing match Harvard GSAS/FAS formal standards
- Document feels formal and document-first, not web-derived
- Page count is acceptable (1-2 pages preferred, but formal quality takes priority)
- ATS systems can reliably parse all content sections
- Implementation remains within the 400-line review budget
- No regression in web page functionality or appearance
- Private data injection continues to work correctly
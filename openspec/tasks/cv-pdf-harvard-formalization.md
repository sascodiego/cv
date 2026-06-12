# Tasks: CV PDF Harvard Formalization

## Change ID

`cv-pdf-harvard-formalization`

## Overview

This task list implements the removal of all iconography from the Harvard-style CV PDF and restores formal typography, spacing, and document structure. The change is scoped to PDF-only files: `web/src/pages/cv-pdf.astro` and `web/scripts/generate-pdf.js`.

## Review Workload Forecast

| Field | Value |
|-------|-------|
| Estimated changed lines | 200–300 lines |
| 400-line budget risk | Low |
| Chained PRs recommended | No |
| Suggested split | Single PR |
| Delivery strategy | auto-chain |

```
Decision needed before apply: No
Chained PRs recommended: No
Chain strategy: single-pr
400-line budget risk: Low
```

---

## Implementation Tasks

### Phase 1: Script Preparation — Remove SVG Icon Logic from `generate-pdf.js`

- [ ] **1.1. Locate and review current SVG icon constants in `web/scripts/generate-pdf.js`**
  - Find all SVG icon constants (`svgEmail`, `svgPhone`, `svgLocation`, `svgWeb`, `svgGithub`, `svgLinkedin`)
  - Document current `createContactItem()` function signature and implementation
  - Verify the `injectPrivateData()` function contact item creation flow

- [ ] **1.2. Refactor `createContactItem()` function to remove SVG parameter**
  - Change function signature from `createContactItem(svgPath, label, href)` to `createContactItem(label, href)`
  - Remove SVG injection logic (template and `appendChild(template.content.firstChild)`)
  - Change element type from `div` (non-link) to `span` (inline)
  - Remove nested `span` element for label
  - Set `textContent` directly on the element
  - Keep `href`, `target="_blank"`, and `rel="noopener noreferrer"` for links

- [ ] **1.3. Create new `createSeparator()` helper function**
  - Create function that generates `span.contact-separator` elements
  - Set `textContent` to `" | "` (space, pipe, space)
  - Append to `contactInfo` container

- [ ] **1.4. Refactor contact item collection to use array-based approach**
  - Replace individual `if` blocks with array collection:
    ```javascript
    const items = [];
    if (data.email) items.push({ label: data.email, href: `mailto:${data.email}` });
    if (data.phone) items.push({ label: data.phone, href: null });
    if (data.address) items.push({ label: data.address, href: null });
    if (data.website) items.push({ label: "Portafolio", href: data.website });
    if (data.github) items.push({ label: "GitHub", href: data.github });
    if (data.linkedin) items.push({ label: "LinkedIn", href: data.linkedin });
    ```
  - Iterate through array and call `createContactItem()` for each
  - Call `createSeparator()` after each item except the last (`index < items.length - 1`)

- [ ] **1.5. Delete SVG icon constants from `generate-pdf.js`**
  - Remove `svgEmail`, `svgPhone`, `svgLocation`, `svgWeb`, `svgGithub`, `svgLinkedin` constants
  - Ensure no orphaned references remain

- [ ] **1.6. Verify unchanged `injectPrivateData()` logic**
  - Confirm name update logic remains unchanged
  - Confirm contact info container clearing remains unchanged
  - Confirm footer note injection remains unchanged
  - Confirm recommender attribution injection remains unchanged

### Phase 2: CSS Cleanup — Remove Icon-Related Styling from `cv-pdf.astro`

- [ ] **2.1. Locate and review current `@media print` block in `web/src/pages/cv-pdf.astro`**
  - Identify all icon-related CSS rules (`.contact-item svg`, `.contact-item svg path`)
  - Identify all icon-related width/height constraints
  - Identify all gap properties used for icon spacing (`gap: 3pt`, `gap: 8pt`)
  - Document current `.contact-info` and `.contact-item` rules

- [ ] **2.2. Remove all SVG-related CSS selectors from `@media print`**
  - Delete `.contact-item svg` block entirely
  - Delete any `.contact-item svg path` rules
  - Delete any icon-specific width/height constraints in contact context
  - Delete any icon-related display or flex properties

- [ ] **2.3. Update `.contact-info` CSS rule**
  - Set `gap: 0` (separators handle spacing, not gap)
  - Keep `display: flex`, `flex-wrap: wrap`, `justify-content: center`, `margin-top: 4pt`

- [ ] **2.4. Update `.contact-item` CSS rule**
  - Remove any width/height constraints for icons
  - Keep `display: inline-flex`, `align-items: center`, `font-size: 8.5pt`, `color: #000000`
  - Ensure no `gap` property on this rule (gap handled at container level)

- [ ] **2.5. Add new `.contact-separator` CSS rule**
  - Define `display: inline-flex`
  - Set `font-size: 8.5pt`
  - Set `color: #000000`

- [ ] **2.6. Add `.contact-item a` CSS rule for links**
  - Set `color: #000000`
  - Set `text-decoration: none` (ATS-friendly, no underlining)

- [ ] **2.7. Verify no icon remnants in `@media print`**
  - Grep `@media print` block for `svg` — expect zero matches
  - Grep `@media print` block for `gap: 3pt` or `gap: 8pt` — expect zero matches
  - Grep `@media print` block for icon-specific sizes (`width: 9pt`, `height: 9pt`) in contact context — expect zero matches

- [ ] **2.8. Verify unchanged CSS rules**
  - Confirm typography rules (name, title, section headers, entry headers) remain unchanged
  - Confirm spacing rules (section margins, entry margins, bullet margins) remain unchanged
  - Confirm page-break control rules remain unchanged

### Phase 3: Build and Verification

- [ ] **3.1. Build the web project to verify no syntax errors**
  - Run `cd /src/cv-pipeline/web && pnpm build`
  - Confirm build completes successfully with zero errors
  - Review any warnings (non-blocking, but note for resolution)

- [ ] **3.2. Generate PDF to verify contact injection works**
  - Run `cd /src/cv-pipeline/web && pnpm export-pdf`
  - Confirm PDF generation completes successfully
  - Verify no runtime errors in console output

- [ ] **3.3. Manual PDF inspection — verify zero icons**
  - Open generated PDF file
  - Visually inspect contact info header for any SVG icons
  - Use PDF text selection tool to verify all contact info is selectable text
  - Confirm PDF contains zero graphical elements or icons

- [ ] **3.4. Manual PDF inspection — verify contact presentation**
  - Verify contact info displays as text-only line (e.g., "email@example.com | +1 555-123-4567 | City, Country | portfolio.com")
  - Verify pipe separators (` | `) appear between contact items
  - Verify no trailing separator after the last contact item
  - Verify contact items are centered across the document width
  - Verify links are black with no underlining

- [ ] **3.5. Manual PDF inspection — verify typography and spacing**
  - Verify name is 14pt, bold, uppercase, centered
  - Verify title is 9.5pt, italic, centered
  - Verify section headers are 10pt, bold, uppercase, with bottom border
  - Verify entry headers are 9.5pt, bold, left-aligned
  - Verify body text is 9pt, regular, left-aligned
  - Verify section headers have 8pt margin above
  - Verify entries have 6pt margin below
  - Verify bullets have 1pt margin below

- [ ] **3.6. Manual PDF inspection — verify document-first appearance**
  - Verify PDF feels like a formal academic document, not a web page
  - Verify no decorative elements (shadows, gradients, text boxes)
  - Verify no content as images
  - Verify single-column layout
  - Verify standard disc bullets (not custom characters)

- [ ] **3.7. Verify ATS compatibility via text extraction test**
  - Copy all text from PDF and paste into text editor
  - Verify contact info is parseable as plain text with correct separators
  - Verify no hidden content or image-based text
  - Verify all sections and entries are selectable

- [ ] **3.8. Verify private data injection**
  - Confirm contact info injects correctly from `.env` (email, phone, address, website, GitHub, LinkedIn)
  - Confirm name injects correctly from `.env`
  - Confirm recommender details inject correctly
  - Confirm footer note injects correctly if present

### Phase 4: Web-Page Guardrail Verification

- [ ] **4.1. Verify `web/src/pages/index.astro` is unchanged**
  - Run `git diff` to confirm no changes to `index.astro`
  - Confirm public web page remains untouched

- [ ] **4.2. Verify public web site functionality**
  - Run `cd /src/cv-pipeline/web && pnpm dev` (optional, if user wants to verify)
  - Confirm web page renders correctly with no errors
  - Confirm web page retains its current visual design

### Phase 5: Final Visual Verification Criteria

- [ ] **5.1. Verify "zero icons in final PDF" criterion**
  - Open generated PDF and search for any graphical or icon elements
  - Use PDF selection tool to confirm all contact info is text, not images
  - Confirm `@media print` block contains no `svg` references (grep verification)

- [ ] **5.2. Verify "web untouched" criterion**
  - Run `git status` to confirm only `cv-pdf.astro` and `generate-pdf.js` are modified
  - Confirm `index.astro` is not in the changed files list
  - Confirm no changes to site styles, branding, or visual design files

- [ ] **5.3. Verify "formal Harvard PDF" criterion**
  - PDF uses Times New Roman throughout (or Arial for ATS compatibility)
  - Name: 14-16pt, bold, uppercase, centered
  - Title: 10-11pt, italic, centered
  - Section headers: 10-12pt, bold, uppercase, with 1pt bottom border
  - Entry headers: 9.5-11pt, bold, left-aligned
  - Body text: 9-10pt, regular, left-aligned
  - Line heights: 1.15-1.25 throughout
  - Letter-spacing: 0.8-1.2pt for name, 0.3-0.5pt for headers, 0pt for body
  - Section spacing: 8-12pt above each section header
  - Entry spacing: 6-8pt below each entry
  - Bullet spacing: 1-2pt below each bullet

- [ ] **5.4. Verify PDF format and page count**
  - Confirm PDF is A4 format with white background
  - Note page count (acceptable if 1-2 pages preferred, but formal quality takes priority)
  - If page count exceeds 2 pages, note for potential density optimization in follow-up change (not blocking)

- [ ] **5.5. Verify build reproducibility**
  - Delete generated PDF
  - Run `pnpm export-pdf` again
  - Confirm second PDF generation succeeds with identical output

---

## Task Dependencies

- **Phase 1 → Phase 2**: Script refactoring must complete before CSS cleanup (to avoid mismatched expectations)
- **Phase 2 → Phase 3**: CSS cleanup must complete before build and verification
- **Phase 3 → Phase 4**: Build and verification must pass before guardrail verification
- **Phase 4 → Phase 5**: Guardrail verification must pass before final visual verification

---

## Rollback Plan

If issues arise during implementation:

1. **Revert `web/scripts/generate-pdf.js`**:
   - Restore SVG icon constants (`svgEmail`, `svgPhone`, etc.)
   - Restore `createContactItem(svgPath, label, href)` signature
   - Restore SVG injection logic
   - Remove `createSeparator()` function
   - Restore individual `if` block approach for contact items

2. **Revert `web/src/pages/cv-pdf.astro`**:
   - Restore `.contact-item svg` CSS rule
   - Restore `gap: 3pt` and `gap: 8pt` on contact-related rules
   - Remove `.contact-separator` CSS rule
   - Remove `.contact-item a` rule

3. **No data migration needed**:
   - No changes to `.env` file structure
   - No changes to `cv-processed.json` schema
   - No changes to Go pipeline

4. **Web page unaffected**:
   - `index.astro` remains untouched
   - No impact on public website functionality

---

## Success Criteria

- [ ] PDF contains zero icons; all contact information is text-only with pipe separators
- [ ] Typography and spacing match Harvard GSAS/FAS formal standards
- [ ] Document feels formal and document-first, not web-derived
- [ ] Page count is acceptable (1-2 pages preferred, but formal quality takes priority)
- [ ] ATS systems can reliably parse all content sections
- [ ] Implementation remains within the 400-line review budget
- [ ] No regression in web page functionality or appearance
- [ ] Private data injection continues to work correctly
- [ ] `pnpm build` passes without errors
- [ ] `pnpm export-pdf` generates valid PDF

---

## Notes

- Page density tradeoffs are acceptable in this first pass. If the formal PDF exceeds 2 pages, density optimization can be addressed in a follow-up change by slightly reducing font sizes or spacing within Harvard-acceptable ranges.
- Screen preview in `cv-pdf.astro` may look different from the PDF output. This is acceptable per the decision that only the final PDF matters.
- Contact separators use pipe characters (` | `) as the most common convention for formal CVs. Can be changed to commas in a follow-up change if testing suggests better readability.
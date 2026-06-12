# Proposal: CV PDF Harvard Formalization

## Change ID

`cv-pdf-harvard-formalization`

## Title

Restore Harvard formality to the PDF by removing icons and optimizing typography, spacing, and document structure

## Problem Statement

The current PDF implementation from `cv-pdf-harvard` produces a document that fails to meet Harvard formal standards. While it successfully moved away from the dark web theme, the PDF still suffers from visual and structural problems that undermine its professional credibility:

1. **Oversized iconography**: Contact information uses SVG icons that create visual clutter, add unnecessary page weight, and distract from the document-first nature of a formal CV
2. **Inadequate hierarchy and spacing**: Typography, spacing, and section rhythm do not follow Harvard GSAS/FAS standards for document composition
3. **Layout inconsistencies**: The balance between density and readability is misaligned; the current compression focused on fitting content onto 2 pages at the expense of formal presentation
4. **Web-derived decorative treatment**: Residual web-specific styling decisions persist despite the Harvard template's intent

This creates a product problem: the PDF feels "wrong" to formal hiring contexts, academic committees, and traditional CV reviewers. It lacks the document-first, structured, and ATS-friendly presentation expected of a Harvard-style CV.

## Intent

Correct the PDF implementation to become truly document-first, formal, and Harvard-faithful by removing all iconography and restoring proper typography, hierarchy, spacing, and section rhythm. This is a formalization change that prioritizes formal Harvard standards over preview fidelity and page density.

## Goals

- Remove all icons from the formal PDF, replacing them with text-only contact presentation
- Restore Harvard GSAS/FAS typography standards: consistent font sizes, proper line heights, formal hierarchy
- Implement proper spacing and section rhythm that follows academic document conventions
- Ensure the PDF feels document-first, not web-derived or decorated
- Maintain ATS-friendly output with clean, structured, and parseable content
- Allow revisiting page density after formal structure is restored, if necessary

## Non-Goals

- Do NOT change the public web page (`index.astro`) or its visual design
- Do NOT redesign the site's overall style or branding
- Do NOT modify the Go data pipeline or data schema
- Do NOT prioritize preview fidelity over formal PDF output
- Do NOT compromise Harvard formality for page density constraints in this first pass

## User and Product Impact

- Job applicants receive a truly formal, Harvard-faithful CV that meets the expectations of traditional hiring contexts
- Academic committees and formal reviewers receive a document that respects established conventions
- ATS systems continue to parse the CV reliably without visual noise or decorative elements
- The CV owner maintains confidence that the PDF represents professional formal standards
- Preview discrepancies are acceptable as long as the final PDF meets Harvard formality requirements

## Scope

### In Scope

- **Modified template**: `web/src/pages/cv-pdf.astro`
  - Remove all SVG icons from contact information
  - Replace icon-based contact display with text-only presentation
  - Restore Harvard typography standards: font sizes, line heights, letter-spacing
  - Implement proper section spacing and rhythm
  - Ensure hierarchy follows academic document conventions
  - Optimize for document-first presentation, not web-derived styling

- **Modified script**: `web/scripts/generate-pdf.js`
  - Update `injectPrivateData()` to build text-only contact items
  - Remove SVG icon injection logic
  - Maintain private data security and injection mechanism

- **Print CSS refinement**: Inline in `cv-pdf.astro`
  - Remove icon-related styling rules
  - Reinforce formal typography and spacing
  - Ensure page-break control respects document structure

### Out of Scope

- Changes to `web/src/pages/index.astro` (public web page)
- Changes to web site styles, branding, or visual design
- Modifications to the Go data pipeline or `cv-processed.json`
- Preview fidelity optimization (screen styles in `cv-pdf.astro`)
- Page density optimization in this first pass (deferred if needed after formal structure)

## Affected Areas

- `web/src/pages/cv-pdf.astro` — remove icons, restore formal typography and spacing
- `web/scripts/generate-pdf.js` — update contact injection to text-only
- No changes to web page, data pipeline, or site styles

## Visual Failures to Correct

### 1. Iconography Problem

**Current state**: Contact information uses SVG icons (email, phone, location, web, GitHub, LinkedIn) that:
- Create visual clutter and unnecessary page weight
- Distract from document-first presentation
- Add web-derived decorative treatment inappropriate for formal CVs
- May cause ATS parsing issues if not rendered correctly

**Target state**: Text-only contact presentation that:
- Lists contact information in clean, readable format
- Uses minimal separators (pipes `|` or commas)
- Prioritizes readability over visual decoration
- Follows formal document conventions

### 2. Typography and Hierarchy Problem

**Current state**: Typography may not follow Harvard GSAS/FAS standards:
- Inconsistent font sizes across elements
- Line heights that may not optimize readability
- Letter-spacing that may not match formal document conventions
- Hierarchy that may not clearly distinguish sections

**Target state**: Harvard-standard typography:
- Consistent font sizes: 9-12pt body, larger for headers
- Proper line heights: 1.15-1.25 for readability
- Appropriate letter-spacing for formal presentation
- Clear hierarchy: name > section headers > entry headers > body text

### 3. Spacing and Section Rhythm Problem

**Current state**: Spacing may be compressed for page density at expense of formality:
- Insufficient spacing between sections
- Inconsistent spacing within entries
- Poor visual rhythm that undermines document structure

**Target state**: Proper spacing and rhythm:
- Consistent section spacing (8-12pt between sections)
- Appropriate entry spacing (4-6pt between entries)
- Visual rhythm that guides the reader through the document
- Balance between density and readability

### 4. Document-First Presentation Problem

**Current state**: Residual web-derived styling decisions persist:
- Layout patterns that may feel web-like rather than document-like
- Visual treatments that prioritize screen reading over print formalism
- Inconsistencies with academic document conventions

**Target state**: True document-first presentation:
- Layout that feels like a formal academic or professional document
- Print-optimized spacing and typography
- Consistency with Harvard GSAS/FAS and industry formal standards

## Formal Harvard Principles to Restore

### 1. Document-First Philosophy

- **Principle**: The CV is a formal document, not a web page or visual design artifact
- **Application**: Remove all decorative elements, prioritize content over styling, ensure print-first design decisions

### 2. Typography Standards

- **Principle**: Harvard GSAS/FAS documents use specific typography conventions for readability and formality
- **Application**:
  - Font: Times New Roman (serif) or Arial (sans-serif) for ATS compatibility
  - Body text: 10-12pt
  - Section headers: 10-12pt, bold, uppercase
  - Entry headers: 10-11pt, bold
  - Line height: 1.15-1.25
  - Letter-spacing: minimal (0-0.5pt)

### 3. Spacing and Rhythm

- **Principle**: Formal documents use consistent spacing to create visual rhythm and guide the reader
- **Application**:
  - Section spacing: 8-12pt before each section header
  - Entry spacing: 4-6pt between entries
  - Element spacing: 2-4pt between related elements
  - Consistent rhythm throughout the document

### 4. Hierarchy and Structure

- **Principle**: Clear visual hierarchy helps readers scan and understand the document quickly
- **Application**:
  - Name: Largest, centered, bold (14-16pt)
  - Section headers: Bold, uppercase, with bottom border (10-12pt)
  - Entry headers: Bold, company/institution name (10-11pt)
  - Period: Right-aligned, regular weight
  - Body text: Regular weight, appropriate size (10-12pt)

### 5. Minimalism and Functionality

- **Principle**: Formal documents prioritize functionality over decoration
- **Application**:
  - No icons or decorative elements
  - No shading, text boxes, or visual effects
  - Minimal color (black text on white background)
  - Clean, readable presentation

### 6. ATS Compatibility

- **Principle**: Formal CVs must be parseable by ATS systems without visual noise
- **Application**:
  - Text-only contact information
  - Single-column layout
  - Standard bullets or simple dashes
  - No content as images
  - Consistent formatting for parsing

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
        (web page, unchanged)                                 (formal PDF, no icons)
```

### Icon Removal Implementation

**Before (current state)**:
```javascript
// generate-pdf.js creates contact items with SVG icons
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

**After (target state)**:
```javascript
// generate-pdf.js creates text-only contact items
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

### Typography and Spacing Restoration

**Key changes in `cv-pdf.astro`**:
- Remove icon-related CSS: `.contact-item svg` rules
- Restore formal typography: font sizes, line heights, letter-spacing
- Implement proper spacing: section, entry, and element spacing
- Ensure document-first presentation in `@media print` rules

## Acceptance Criteria

- **Icon removal**: PDF contains no icons; contact information is text-only
- **Typography**: Font sizes, line heights, and letter-spacing match Harvard standards
- **Spacing**: Section and entry spacing follows academic document conventions
- **Hierarchy**: Clear visual hierarchy guides reader through document
- **Document-first**: PDF feels like a formal document, not a web page
- **ATS-friendly**: Clean, parseable output without visual noise
- **Build**: `cd /src/cv-pipeline/web && pnpm build` passes without errors
- **PDF generation**: `cd /src/cv-pipeline/web && pnpm export-pdf` generates valid PDF
- **Format**: PDF is 1-2 pages, A4, white background, professional layout
- **Private data**: Contact info and recommenders injected correctly from `.env`

## Risks

- **Page density**: Restoring formal spacing may expand page count beyond 2 pages; accept this in first pass and revisit density later if needed
- **User preference**: Some users may prefer icons for visual identification; defer to formal Harvard standards in this change
- **ATS parsing**: While text-only contact info is more ATS-friendly, test to ensure parsers correctly identify email, phone, and URLs
- **Preview inconsistency**: Screen preview may look different from PDF; acceptable per user decision that only final PDF matters

## Rollback Plan

- Revert `web/scripts/generate-pdf.js` to restore SVG icon injection
- Revert `web/src/pages/cv-pdf.astro` to restore icon-related CSS
- No data migration or schema changes needed
- Web page and existing functionality remain untouched
- Rollback is straightforward: both files are version-controlled with clear before/after states

## Success Criteria

- PDF contains zero icons; all contact information is text-only
- Typography and spacing match Harvard GSAS/FAS formal standards
- Document feels formal and document-first, not web-derived
- Page count is acceptable (1-2 pages preferred, but formal quality takes priority)
- ATS systems can reliably parse all content sections
- Implementation remains within the 400-line review budget
- No regression in web page functionality or appearance

## First-Slice Recommendation

Implement the first slice as a focused formalization effort:

1. Remove all icons from `generate-pdf.js` contact injection
2. Update `cv-pdf.astro` to remove icon-related CSS
3. Restore Harvard typography standards (font sizes, line heights, letter-spacing)
4. Implement proper spacing and section rhythm
5. Test PDF output for formality and ATS compatibility
6. Accept page density tradeoffs in this first pass; defer density optimization if needed

Avoid:
- Web page changes
- Site style modifications
- Data schema changes
- Preview fidelity optimization
- Page density compromises that undermine formality

## Review Workload Forecast

Estimated implementation diff: approximately 200–300 changed lines
- `web/src/pages/cv-pdf.astro`: ~150–200 lines (icon removal, typography, spacing)
- `web/scripts/generate-pdf.js`: ~50–100 lines (icon removal, contact injection changes)
- Total: within the configured 400-line review budget

## Open Questions

No blocking open questions remain. The user has provided clear decisions:

- **Icon strategy**: Remove all icons from the formal PDF ✅
- **Preview vs PDF**: Only final PDF matters; preview fidelity not a requirement ✅
- **Priority**: Formal Harvard formality first, even if revisiting density later ✅

Optional refinements for later review:
- Should page density be revisited if the formal PDF exceeds 2 pages?
- Should contact separators use pipes (`|`) or commas for better readability?
- Should specific font sizes be adjusted based on final page count?

## Dependencies

- `cv-processed.json` must be generated before PDF export (existing pipeline)
- `.env` file with `PRIVATE_*` variables must exist for contact data injection
- Puppeteer must be installed (already in `web/package.json`)

## Related Artifacts

- Parent change: `cv-pdf-harvard` — initial Harvard-style PDF implementation
- Original proposal: `openspec/proposals/cv-pdf-harvard.md`
- Verification report: `openspec/verify-reports/cv-pdf-harvard.md`
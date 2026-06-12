# Proposal: CV PDF Reference Alignment

## Change ID

`cv-pdf-reference-alignment`

## Title

Align PDF output visually with reference document while improving typography and implementing technology label normalization

## Problem Statement

The current PDF implementation from `cv-pdf-harvard-formalization` produces a document that is on the right track but falls short of visual parity with the reference PDF. While it successfully removed icons and restored Harvard formality, the PDF exhibits typography and presentation problems that reduce its effectiveness:

1. **Type is too small and dense**: Font sizes and spacing are compressed, making the document feel cramped and harder to read
2. **Document rhythm is unclear**: Margins, spacing between sections, and visual hierarchy lack the clean separation present in the reference
3. **Technology labels are raw IDs**: The source uses internal tags like `csharp`, `typescript`, `python`, but these render as-is instead of human-friendly labels like `C#`, `TypeScript`, `Python`
4. **Whitespace is insufficient**: Margins and spacing between elements do not provide the breathing room that characterizes the reference's professional appearance
5. **Typography lacks refinement**: Font scale, line heights, and character spacing do not achieve the legibility and formality of the reference

This creates a product problem: the PDF does not achieve **paridad visual** with the reference, undermining its credibility in formal hiring contexts and making it less readable for recruiters and hiring committees.

## Intent

Improve the PDF to achieve visual parity with the reference document by refining typography, spacing, and document rhythm, while implementing a technology label normalization layer. This is a visual alignment change that prioritizes legibility and formality over exact metric cloning or web-preview fidelity.

## Goals

- Improve typography scale to feel closer to the reference PDF: larger, more readable fonts with appropriate line heights
- Make the PDF more legible and formal while keeping it compact enough for practical use
- Create a technology-label normalization layer in the PDF presentation path so source IDs like `csharp` render as user-facing labels like `C#`
- Establish clearer document rhythm with better margins, spacing between sections, and visual separation
- Preserve ATS-friendly single-column document structure
- Keep all changes isolated to PDF creation and export (no web changes)

## Non-Goals

- Do NOT modify `web/src/pages/index.astro` (public web page)
- Do NOT redesign the public site styles or branding
- Do NOT change the source data schema or Go data pipeline
- Do NOT attempt exact metric cloning of the reference (e.g., pixel-perfect measurements)
- Do NOT alter the reference PDF file
- Do NOT prioritize preview fidelity over final PDF output

## User and Product Impact

- Job applicants receive a PDF that achieves visual parity with the reference document
- Recruiters and hiring committees receive a more readable, professional CV with human-friendly technology labels
- ATS systems continue to parse the CV reliably with clean, structured content
- The CV owner maintains confidence that the PDF represents formal, reference-aligned standards
- Visual improvements enhance credibility without requiring data source changes

## Scope

### In Scope

- **Modified template**: `web/src/pages/cv-pdf.astro`
  - Refine typography scale: increase base font sizes, adjust line heights for readability
  - Improve document rhythm: better margins, section spacing, and visual hierarchy
  - Refine contact section layout to match reference proportions
  - Optimize section headers and separators for clearer visual separation
  - Ensure body text and bullet points match reference legibility

- **Modified script**: `web/scripts/generate-pdf.js`
  - Implement technology label mapper: source IDs → human-friendly labels
  - Normalize technology tags before rendering in the PDF
  - Maintain private data security and injection mechanism
  - Preserve existing data flow without schema changes

- **Print CSS refinement**: Inline in `cv-pdf.astro`
  - Adjust margin and spacing values for better whitespace
  - Refine typography rules (font-size, line-height, letter-spacing)
  - Ensure page-break control respects document structure

### Out of Scope

- Changes to `web/src/pages/index.astro` (public web page)
- Changes to web site styles, branding, or visual design
- Modifications to the Go data pipeline or `cv-processed.json`
- Changes to source data schema or YAML structure
- Exact metric cloning of reference measurements
- Preview fidelity optimization (screen styles in `cv-pdf.astro`)

## Affected Areas

- `web/src/pages/cv-pdf.astro` — refine typography, spacing, and document rhythm
- `web/scripts/generate-pdf.js` — implement technology label mapper
- No changes to web page, data pipeline, or site styles

## Visual Issues to Correct

### 1. Typography Scale Problem

**Current state**: Font sizes are too small and dense:
- Body text and bullets are cramped, reducing readability
- Line heights may be compressed for page density at expense of legibility
- Font scale does not match the reference's proportions

**Target state**: Improved typography scale:
- Larger base font sizes for body text and bullets
- Appropriate line heights (1.2–1.4) for comfortable reading
- Font scale that feels proportional to the reference
- Balance between compactness and readability

### 2. Document Rhythm Problem

**Current state**: Margins and spacing are insufficient:
- Section spacing may be too tight, reducing visual separation
- Margins do not provide the breathing room of the reference
- Visual hierarchy lacks the clear rhythm present in the reference

**Target state**: Clearer document rhythm:
- Consistent section spacing (10–14pt) for better separation
- Appropriate margins that frame content without wasting space
- Visual rhythm that guides the reader through the document
- Whitespace that feels intentional and professional

### 3. Technology Label Problem

**Current state**: Technology tags render as raw source IDs:
- `csharp` displays as "csharp" instead of "C#"
- `typescript` displays as "typescript" instead of "TypeScript"
- Other tags appear as internal IDs rather than user-facing labels

**Target state**: Human-friendly technology labels:
- Implement mapper: `csharp` → "C#", `typescript` → "TypeScript", etc.
- Render normalized labels in the PDF
- Keep source data unchanged; transformation happens in presentation layer

### 4. Contact Section Layout Problem

**Current state**: Contact section may not match reference proportions:
- Contact items may be too cramped or poorly spaced
- Separator treatment may not align with reference visual style

**Target state**: Reference-aligned contact layout:
- Contact items with appropriate spacing and separators
- Layout proportions that match the reference's rhythm
- Clean, readable presentation without visual clutter

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
        (web page, unchanged)                                 (PDF with mapper)
```

### Technology Label Mapper Implementation

**Before (current state)**:
```javascript
// generate-pdf.js renders technology tags as-is
function renderSkills(skills) {
  return skills.map(skill => `<span class="skill-tag">${skill}</span>`).join('');
}
// Output: "csharp", "typescript", "python"
```

**After (target state)**:
```javascript
// generate-pdf.js applies normalization mapper
const TECH_LABEL_MAP = {
  csharp: 'C#',
  typescript: 'TypeScript',
  javascript: 'JavaScript',
  python: 'Python',
  go: 'Go',
  rust: 'Rust',
  // ... additional mappings
};

function normalizeTechLabel(tag) {
  return TECH_LABEL_MAP[tag] || tag.charAt(0).toUpperCase() + tag.slice(1);
}

function renderSkills(skills) {
  return skills.map(skill => {
    const label = normalizeTechLabel(skill);
    return `<span class="skill-tag">${label}</span>`;
  }).join('');
}
// Output: "C#", "TypeScript", "Python"
```

### Typography and Spacing Refinement

**Key changes in `cv-pdf.astro`**:
- Increase base font sizes: body from ~9pt to ~10–11pt
- Adjust line heights: from ~1.1 to ~1.25–1.4
- Refine section spacing: from ~6–8pt to ~10–14pt
- Adjust margins: increase side margins for better framing
- Optimize contact section spacing and separators

## Reference Alignment Principles

### 1. Visual Similitude, Not Metric Cloning

- **Principle**: The goal is visual parity with the reference, not exact pixel-perfect measurements
- **Application**: Adjust typography and spacing to match the reference's feel and proportions, but do not attempt metric cloning

### 2. Typography Priority: More Legible

- **Principle**: Typography should be larger, more readable, and more formal than the current implementation
- **Application**:
  - Increase base font sizes for body text and bullets
  - Use appropriate line heights for comfortable reading
  - Ensure font scale matches reference proportions

### 3. Technology Label Normalization

- **Principle**: Source data uses internal IDs, but presentation should show human-friendly labels
- **Application**:
  - Implement mapper in `generate-pdf.js` for normalization
  - Keep source data unchanged; transform in presentation layer only
  - Provide fallback capitalization for unmapped tags

### 4. Document Rhythm and Whitespace

- **Principle**: Margins and spacing should create clear visual rhythm and guide the reader
- **Application**:
  - Consistent section spacing for better separation
  - Appropriate margins that frame content professionally
  - Whitespace that feels intentional, not accidental

### 5. Preserve ATS Compatibility

- **Principle**: Visual improvements should not compromise ATS parsing
- **Application**:
  - Maintain single-column layout
  - Use standard bullets or simple dashes
  - No content as images
  - Consistent formatting for parsing

## Acceptance Criteria

- **Typography**: Font sizes and line heights are larger and more legible than current implementation
- **Document rhythm**: Margins and section spacing provide clear visual separation
- **Technology labels**: Source IDs (e.g., `csharp`) render as human-friendly labels (e.g., `C#`)
- **Visual parity**: PDF achieves visual similitude with the reference document
- **Contact section**: Layout matches reference proportions with appropriate spacing
- **ATS-friendly**: Clean, parseable output maintained throughout
- **Build**: `cd /src/cv-pipeline/web && pnpm build` passes without errors
- **PDF generation**: `cd /src/cv-pipeline/web && pnpm export-pdf` generates valid PDF
- **Format**: PDF is 1-2 pages, A4, white background, professional layout
- **Private data**: Contact info and recommenders injected correctly from `.env`

## Risks

- **Page density**: Larger typography and spacing may expand page count beyond 2 pages; accept visual parity as higher priority
- **Unmapped technology tags**: Some tags may lack explicit mappings; fallback capitalization should handle gracefully
- **Visual interpretation**: "Visual parity" may involve subjective judgment; focus on readability and formal rhythm
- **Preview inconsistency**: Screen preview may look different from PDF; acceptable as only final PDF matters
- **Reference interpretation**: Exact reference proportions may not be achievable without metric cloning; prioritize visual similitude

## Rollback Plan

- Revert `web/scripts/generate-pdf.js` to remove technology label mapper
- Revert `web/src/pages/cv-pdf.astro` to previous typography and spacing values
- No data migration or schema changes needed
- Web page and existing functionality remain untouched
- Rollback is straightforward: both files are version-controlled with clear before/after states

## Success Criteria

- PDF typography is larger and more legible than current implementation
- Technology labels render as human-friendly names (e.g., `C#`, not `csharp`)
- Document rhythm feels similar to the reference with clear margins and section spacing
- Visual parity is achieved without exact metric cloning
- ATS systems can reliably parse all content sections
- Implementation remains within the 400-line review budget
- No regression in web page functionality or appearance

## First-Slice Recommendation

Implement the first slice as a focused visual alignment effort:

1. Implement technology label mapper in `generate-pdf.js`
2. Update `cv-pdf.astro` to refine typography scale (font sizes, line heights)
3. Adjust margins and section spacing for better document rhythm
4. Refine contact section layout to match reference proportions
5. Test PDF output for visual parity and ATS compatibility
6. Accept page density tradeoffs in this first pass; prioritize visual alignment

Avoid:
- Web page changes
- Site style modifications
- Data schema changes
- Exact metric cloning
- Preview fidelity optimization

## Review Workload Forecast

Estimated implementation diff: approximately 250–350 changed lines
- `web/src/pages/cv-pdf.astro`: ~150–200 lines (typography, spacing, rhythm)
- `web/scripts/generate-pdf.js`: ~100–150 lines (technology label mapper)
- Total: within the configured 400-line review budget

## Open Questions

No blocking open questions remain. The user has provided clear decisions:

- **Alignment goal**: Similitud visual with reference, not exact metric cloning ✅
- **Typography priority**: More legible ✅
- **Technology normalization**: Implement mapper (csharp → C#) ✅
- **Success definition**: Paridad visual ✅
- **Scope**: PDF-only, no web changes ✅

Optional refinements for later review:
- Should additional technology mappings be added beyond common languages?
- Should section spacing values be fine-tuned based on final page count?
- Should fallback capitalization use title case or sentence case for unmapped tags?

## Dependencies

- `cv-processed.json` must be generated before PDF export (existing pipeline)
- `.env` file with `PRIVATE_*` variables must exist for contact data injection
- Puppeteer must be installed (already in `web/package.json`)
- Reference PDF must be available for visual comparison

## Related Artifacts

- Parent change: `cv-pdf-harvard-formalization` — Harvard formality restoration
- Parent proposal: `openspec/proposals/cv-pdf-harvard-formalization.md`
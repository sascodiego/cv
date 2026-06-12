## Summary

**Spec written to:** `/src/cv-pipeline/openspec/changes/cv-pdf-reference-alignment/specs/cv-pdf/spec.md`

This is a full new domain spec (no canonical spec exists for cv-pdf) that defines visual and structural requirements for PDF output to achieve visual parity with the reference document.

### Core Technical Rules

**Typography Rules:**
- Name: 14–15pt, bold, uppercase, letter-spacing 1pt
- Title: 10–11pt, italic, line-height ≥1.15
- Section headers: 10.5–11pt, uppercase, bold, letter-spacing 0.35–0.5pt, bottom border 1pt
- Body text: 9.5–10.5pt, line-height 1.2–1.4
- Bullets: 9–10pt, line-height 1.15–1.25

**Spacing Rules:**
- Page margins: 16mm (top/bottom), 15mm (left/right)
- Section spacing: ≥10pt above header, ≥4pt below header, ≥10pt between sections
- Entry spacing: ≥8pt between entries, ≥4pt header-to-role, ≥2pt role-to-bullets, ≥2pt between bullets
- Contact items: 8.75–9pt, separated by " · " with 4pt padding

**Technology Normalization:**
- Mapper in `generate-pdf.js`: `csharp`→`C#`, `typescript`→`TypeScript`, etc.
- Fallback: capitalize first character for unmapped tags
- Isolated to presentation layer; source data unchanged

**PDF-Only Constraints:**
- Only `cv-pdf.astro` and `generate-pdf.js` modified
- Public web (`index.astro`) untouched
- ATS-friendly: single-column, selectable text, standard bullets
- A4 format, white background, 1-2 pages

The spec focuses on WHAT must be true (visual parity, legibility, normalization) without prescribing HOW to achieve it.
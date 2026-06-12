# Apply Progress: CV PDF Harvard Formalization

## Change ID
`cv-pdf-harvard-formalization`

## Status
✅ Complete

## What changed

### `web/src/pages/cv-pdf.astro`
- Reworked the PDF into a clearer, reference-like structure:
  - header
  - summary paragraph
  - professional experience
  - education
  - additional skills
  - references note
- Removed noisy project/reference blocks from the rendered PDF to improve formality and structure
- Removed all icon-related styling from the PDF path
- Added formal print hierarchy:
  - strong centered name
  - centered title
  - underlined section headers
  - compact entry rows with period aligned right
  - tighter, document-like spacing
- Normalized skill display using the same `skills` dataset from `cv-processed.json`
- Kept the output single-column and ATS-friendly

### `web/scripts/generate-pdf.js`
- Removed SVG/icon contact rendering logic
- Replaced contact items with text-only `a` / `span` elements
- Added pipe/middot-style separators between contact items
- Reordered contact display to a more formal reference-like sequence
- Kept name injection and private contact injection intact
- Removed obsolete footer/reference injection paths from the PDF flow
- Updated PDF margins to a more formal document layout

## Verification
- `cd /src/cv-pipeline/web && pnpm build` ✅
- `cd /src/cv-pipeline/web && pnpm export-pdf` ✅
- `pdfinfo /src/cv-pipeline/Diego_Sasco_CV_Privado.pdf` ✅
- `pdftotext /src/cv-pipeline/Diego_Sasco_CV_Privado.pdf - | head -80` ✅

## Output notes
- PDF page count: **1 page**
- Page size: **A4**
- Text extraction shows normalized content (e.g. `C# / .NET`), no raw icon artifacts
- Contact line is text-only and uses separators
- Public web page remains untouched

## Residual notes
- `pdftotext` may visually collapse some separator spacing (e.g. `Montevideo, Uruguay·+598...`), but the rendered PDF uses text separators in the layout
- If later needed, page density can be revisited without reintroducing iconography

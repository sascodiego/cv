# Apply Progress: CV PDF Reference Alignment

## Change ID
`cv-pdf-reference-alignment`

## Status
✅ Complete

## What changed

### `web/src/pages/cv-pdf.astro`
- Increased typographic scale to feel closer to the right-side reference and improve legibility
- Increased the name/header prominence and raised the body/title/section scales
- Tightened vertical rhythm while keeping the layout single-column and ATS-friendly
- Switched the experience/education header row to a grid so dates sit more clearly on the right
- Enlarged the date column and strengthened company/institution/role hierarchy for better scanability
- Reworked the skills section into a compact, Spanish-labeled grid with metadata-driven truncation instead of long inline lists
- Reduced page margins to preserve a one-page PDF after the typography increase
- Preserved the same underlying CV data source used by the web

### `web/scripts/generate-pdf.js`
- Added a technology label mapper so raw tags such as `csharp` render as `C#`
- Added normalization helpers with fallback capitalization for unmapped tags
- Preserved private data injection and PDF-only behavior
- Kept the public web untouched

### `openspec/apply-progress/cv-pdf-reference-alignment.md`
- Created apply-progress documentation for this change

## Verification
- `cd /src/cv-pipeline/web && pnpm build` ✅
- `cd /src/cv-pipeline/web && pnpm export-pdf` ✅
- `pdfinfo /src/cv-pipeline/Diego_Sasco_CV_Privado.pdf` ✅
- `pdftotext /src/cv-pipeline/Diego_Sasco_CV_Privado.pdf - | head -80` ✅

## Output notes
- PDF page count: **1 page**
- Page size: **A4**
- Text extraction confirms `C# / .NET` and other normalized labels instead of raw tech IDs
- The references footer was removed
- The public web remains untouched

## Residual notes
- The typography is now noticeably larger and more legible while still fitting on one page
- The extra room came from tighter page margins, and the date column/hierarchy/skills block are now more structured, Spanish-localized, and easier to scan

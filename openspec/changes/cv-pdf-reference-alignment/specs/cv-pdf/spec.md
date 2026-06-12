# CV PDF Reference Alignment Specification

## Purpose

Define the visual and structural requirements for the PDF CV output to achieve visual parity with a reference document while improving typography and presenting human-readable technology labels. The spec prioritizes visual similitude over exact metric cloning and legibility over page density.

## Requirements

### Requirement: Typography Scale and Hierarchy

The PDF MUST present a clear, readable typographic hierarchy that prioritizes legibility and formality.

#### Scenario: Header name and title appear at appropriate sizes

- GIVEN the PDF CV is generated
- WHEN the header section is rendered
- THEN the name MUST be displayed at a size of approximately 14–15pt with bold weight and uppercase letter-spacing
- AND the professional title MUST be displayed at approximately 10–11pt with italic styling
- AND the line height for the title MUST be at least 1.15 for readability
- AND the name and title MUST be centered and separated by appropriate spacing (approximately 2–4pt)

#### Scenario: Section headers are clearly distinguished

- GIVEN any section header is rendered (e.g., "EXPERIENCIA PROFESIONAL", "EDUCACIÓN", "HABILIDADES ADICIONALES")
- WHEN the header is displayed
- THEN the section header MUST be uppercase, bold, and approximately 10.5–11pt
- AND the header MUST have letter-spacing of approximately 0.35–0.5pt for formality
- AND the header MUST include a bottom border (1pt solid)
- AND the spacing above and below the header MUST be at least 8–10pt to create visual separation

#### Scenario: Body text and bullets are legible

- GIVEN any body text or bullet point is rendered
- WHEN the text is displayed
- THEN the font size MUST be at least 9.5–10.5pt for body text
- AND the font size for bullet points MUST be at least 9–10pt
- AND the line height MUST be at least 1.2 for body text and 1.15 for bullets
- AND the line height MUST NOT exceed 1.4 to maintain formal compactness
- AND the text MUST use Times New Roman or similar serif font family

#### Scenario: Entry headers show clear hierarchy

- GIVEN a work experience or education entry is rendered
- WHEN the entry header (company/institution name, period) is displayed
- THEN the company or institution name MUST be bold and approximately 10–10.5pt
- AND the period MUST be italic and approximately 9–9.5pt
- AND the role title or degree MUST be approximately 9.5–10pt with normal weight
- AND these elements MUST be separated by appropriate spacing (2–4pt)

### Requirement: Spacing and Margins

The PDF MUST maintain consistent margins and spacing that create clear document rhythm and visual separation.

#### Scenario: Page margins frame content appropriately

- GIVEN the PDF is generated
- WHEN the document is viewed
- THEN the top and bottom margins MUST be approximately 16mm
- AND the left and right margins MUST be approximately 15mm
- AND the margins MUST create a frame that matches the reference's proportions
- AND the content MUST fit within A4 page dimensions (210mm × 297mm)

#### Scenario: Section spacing provides visual rhythm

- GIVEN two consecutive sections are rendered
- WHEN the sections are displayed
- THEN the spacing above a section header MUST be at least 10pt
- AND the spacing below a section header MUST be at least 4pt
- AND the spacing between the last entry of one section and the header of the next MUST be at least 10pt
- AND the spacing MUST be consistent across all sections

#### Scenario: Entry spacing maintains readability

- GIVEN multiple entries are rendered within a section
- WHEN the entries are displayed
- THEN the spacing between entries MUST be at least 8pt
- AND the spacing between the entry header and the role/degree text MUST be at least 4pt
- AND the spacing between the role/degree text and the bullets MUST be at least 2pt
- AND the spacing between bullet items MUST be at least 2pt

#### Scenario: Contact section spacing matches reference proportions

- GIVEN the contact information section is rendered
- WHEN contact items are displayed
- THEN contact items MUST be separated by " · " characters with 4pt padding on each side
- AND the contact items MUST be approximately 8.75–9pt in size
- AND the spacing between the professional title and the contact items MUST be at least 4pt
- AND the contact items MUST be centered and wrapped appropriately

### Requirement: Reference-Like Section Structure

The PDF MUST present sections in a structure that mirrors the reference document's layout and alignment.

#### Scenario: Document follows reference section order

- GIVEN the PDF CV is generated
- WHEN the document is rendered
- THEN sections MUST appear in this order: header, summary, EXPERIENCIA PROFESIONAL, EDUCACIÓN, HABILIDADES ADICIONALES, references note
- AND each section MUST be separated by appropriate spacing
- AND the structure MUST match the reference's formal layout

#### Scenario: Contact section aligns with reference layout

- GIVEN the header section is rendered
- WHEN contact information is displayed
- THEN contact items MUST be centered horizontally
- AND contact items MUST be wrapped across lines if needed
- AND each contact item MUST use a separator (" · ") without icons
- AND the layout proportions MUST match the reference's visual rhythm

#### Scenario: Skills section maintains reference grouping

- GIVEN the HABILIDADES ADICIONALES section is rendered
- WHEN skills are displayed
- THEN skills MUST be grouped by category (Languages & Frameworks, Databases & Messaging, Tools & Platforms, IoT & Embedded)
- AND each category MUST display as a list item with a bold label
- AND skills within a category MUST be separated by commas
- AND each skill MUST include experience text in parentheses (e.g., "5 años")

#### Scenario: Work experience entries follow reference format

- GIVEN a work experience entry is rendered
- WHEN the entry is displayed
- THEN the company name MUST be bold and on the left
- AND the period MUST be italic and aligned to the right
- AND the role title MUST appear below the company/period line
- AND bullet points MUST appear below the role title with appropriate indentation (approximately 16pt)
- AND bullets MUST use standard bullets or simple dashes for ATS compatibility

### Requirement: Technology Normalization Layer

The PDF MUST present technology labels as human-readable names instead of raw source IDs.

#### Scenario: Common technology IDs map to readable labels

- GIVEN a technology tag is rendered from the source data
- WHEN the tag corresponds to a known technology (e.g., `csharp`, `typescript`, `javascript`, `python`, `go`, `rust`)
- THEN the PDF MUST display the human-readable label (e.g., `C#`, `TypeScript`, `JavaScript`, `Python`, `Go`, `Rust`)
- AND the source data MUST remain unchanged (transformation happens in presentation layer only)

#### Scenario: Fallback capitalization handles unmapped tags

- GIVEN a technology tag is rendered that does not have an explicit mapping
- WHEN the tag is displayed
- THEN the first character MUST be capitalized
- AND the rest of the characters MUST remain in lowercase (sentence case)
- AND the fallback MUST prevent raw IDs like `csharp` from displaying as-is

#### Scenario: Technology normalization occurs before PDF rendering

- GIVEN the PDF generation process begins
- WHEN technology tags are rendered
- THEN the normalization mapper MUST be applied before the HTML is captured by Puppeteer
- AND the normalized labels MUST be present in the final PDF output
- AND the mapping logic MUST be isolated to `web/scripts/generate-pdf.js`

#### Scenario: Technology labels remain compatible with existing data source

- GIVEN the CV data source includes technology tags
- WHEN the normalization mapper is applied
- THEN the mapper MUST accept all existing tag formats from the data source
- AND the mapper MUST NOT require changes to the data schema or Go pipeline
- AND the mapper MUST preserve the original tag structure if normalization fails

### Requirement: PDF-Only Constraints

The PDF generation MUST remain isolated from the public web site and preserve ATS compatibility.

#### Scenario: Public web site remains untouched

- GIVEN the change is applied
- WHEN the public web site is viewed
- THEN `web/src/pages/index.astro` MUST NOT be modified
- AND public site styles and branding MUST remain unchanged
- AND the web preview functionality MUST continue to work as before

#### Scenario: PDF generation path is isolated

- GIVEN the PDF is generated via `pnpm export-pdf`
- WHEN the generation script runs
- THEN `web/src/pages/cv-pdf.astro` MUST be the only modified template
- AND `web/scripts/generate-pdf.js` MUST include the normalization mapper
- AND no other files in the project MUST be affected

#### Scenario: ATS compatibility is preserved

- GIVEN the PDF is generated
- WHEN an ATS system parses the document
- THEN the PDF MUST use a single-column layout
- AND all content MUST be selectable text (not images)
- AND standard bullets or simple dashes MUST be used for lists
- AND the structure MUST remain consistent with formal CV conventions
- AND the document MUST NOT use tables, images, or other ATS-hostile elements

#### Scenario: Private data injection mechanism is preserved

- GIVEN the PDF is generated
- WHEN private data is injected
- THEN the Puppeteer-based injection mechanism MUST continue to work
- AND contact information MUST be populated from `.env` variables
- AND the injection MUST occur before PDF capture
- AND the recommendation system MUST remain functional

### Requirement: Page Layout and Format

The PDF MUST maintain a professional, ATS-friendly format within A4 dimensions.

#### Scenario: PDF format meets basic requirements

- GIVEN the PDF is generated
- WHEN the document is viewed
- THEN the PDF MUST be A4 format (210mm × 297mm)
- AND the PDF MUST have a white background
- AND the PDF MUST be 1-2 pages in length
- AND the PDF MUST use Times New Roman or similar serif font for all text

#### Scenario: Page breaks respect document structure

- GIVEN the PDF is generated
- WHEN a page break occurs
- THEN page breaks MUST NOT occur within a single work experience or education entry
- AND page breaks MUST NOT occur within a section header
- AND page breaks SHOULD NOT occur between a section header and its first entry
- AND the PDF MUST use `page-break-inside: avoid` for entries and sections

#### Scenario: Visual parity with reference is achieved

- GIVEN the PDF is generated and compared to the reference document
- WHEN the documents are viewed side-by-side
- THEN the PDF MUST achieve visual similitude with the reference (not exact metric cloning)
- AND the typography scale MUST feel similar to the reference
- AND the spacing and margins MUST match the reference's proportions
- AND the document rhythm MUST feel similar to the reference

## Acceptance Criteria

- Typography is larger and more legible than the current implementation (body text at least 9.5–10.5pt, line height at least 1.2)
- Section headers are clearly distinguished with appropriate uppercase, bold, and letter-spacing
- Margins and spacing create clear document rhythm (section spacing at least 10pt, entry spacing at least 8pt)
- Contact section layout matches reference proportions with appropriate spacing and separators
- Technology labels render as human-readable names (e.g., `C#`, not `csharp`)
- Unmapped technology tags use fallback capitalization (first character capitalized)
- PDF achieves visual similitude with the reference document
- Public web site remains unchanged (no modifications to `web/src/pages/index.astro`)
- PDF generation succeeds with `pnpm export-pdf`
- PDF is ATS-friendly (single-column, selectable text, standard bullets)
- PDF is 1-2 pages, A4, white background, professional layout
- Private data injection mechanism continues to work correctly
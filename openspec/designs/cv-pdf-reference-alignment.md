# Design: CV PDF Reference Alignment

**Change ID:** `cv-pdf-reference-alignment`

**Design written to:** `/src/cv-pipeline/openspec/designs/cv-pdf-reference-alignment.md`

## Overview

This design specifies concrete typography, spacing, and structural changes to achieve visual parity with the reference PDF. The design prioritizes legibility and formality over exact metric cloning, implementing a technology label normalization layer while preserving ATS compatibility.

### Design Goals

1. Typography: Larger, more readable fonts with appropriate line heights
2. Document rhythm: Clear margins and spacing for visual separation
3. Technology normalization: Source IDs (`csharp`) → human labels (`C#`)
4. Visual parity: Match reference proportions through similitude, not cloning
5. ATS compatibility: Maintain single-column, parseable output

### Design Principles

- **Visual similitude**: Match the feel and proportions of the reference, not exact pixels
- **Typography priority**: Larger and more legible than current implementation
- **Data isolation**: Transform technology labels in presentation layer only
- **Rhythm over density**: Accept slight page count increases for better readability
- **PDF-only scope**: No changes to public web page or site styles

---

## 1. Typography Tuning

### 1.1 Font Scale (Print CSS)

| Element | Current Value | Target Value | Justification |
|---------|---------------|--------------|---------------|
| Name (h1) | 14pt | 14–15pt | Reference shows substantial, uppercase name with breathing room |
| Title | 9.5pt | 10–11pt | Title needs more prominence; reference shows clearer hierarchy |
| Section headers (h2) | 10pt | 10.5–11pt | Headers should command attention without overwhelming content |
| Body text | 9.25pt | 9.5–10.5pt | Reference legibility suggests larger body text |
| Bullets (li) | 8.95pt | 9–10pt | Bullets need to be readable, not subordinate to body |
| Contact items | 8.75pt | 8.75–9pt | Already close to reference; minimal adjustment |
| Period dates | 8.75pt | 8.75–9pt | Consistent with contact treatment |

### 1.2 Line Height Guidelines

| Element | Current | Target | Rationale |
|---------|---------|--------|-----------|
| Name (h1) | 1.05 | 1.05–1.08 | Tighter line height for stacked uppercase letters |
| Title | 1.18 | ≥1.15 | Maintain clear separation from name |
| Section headers | Default | 1.1–1.2 | Headers need breathability but not excessive spacing |
| Body text | 1.16 | 1.2–1.4 | Reference shows more generous line height for readability |
| Bullets | 1.15 | 1.15–1.25 | Bullets need vertical room without orphaning |
| Summary text | 1.18 | 1.2–1.35 | Summary paragraph requires excellent legibility |

### 1.3 Letter Spacing Guidelines

| Element | Current | Target | Rationale |
|---------|---------|--------|-----------|
| Name (h1) | 1pt | 1pt | Uppercase tracking maintains reference feel |
| Section headers (h2) | 0.35pt | 0.35–0.5pt | Reference shows moderate tracking for uppercase headers |
| All other elements | Default | Normal | Avoid over-spacing; maintain professional rhythm |

### 1.4 Font Weight and Style

| Element | Current | Target | Rationale |
|---------|---------|--------|-----------|
| Name | Bold | Bold | Substantial presence for identification |
| Title | Italic, normal | Italic, normal | Distinguished from name, not competing |
| Section headers | Bold, uppercase | Bold, uppercase | Clear visual hierarchy for sections |
| Company/institution | Bold | Bold | Primary entry identifier |
| Role/degree | Normal | Normal | Subordinate to company/institution |
| Period dates | Italic | Italic | Distinct but not prominent |
| Summary | Italic | Italic | Introductory character without competing |
| Skill labels | Bold | Bold | Identify categories quickly |

### 1.5 Typography Implementation Details

```css
/* Print CSS adjustments in cv-pdf.astro */

/* Name: more prominence */
.cv-header h1 {
  font-size: 14.5pt;      /* increased from 14pt */
  letter-spacing: 1pt;
  line-height: 1.06;      /* slight increase for breathability */
}

/* Title: more legible */
.professional-title {
  font-size: 10.5pt;      /* increased from 9.5pt */
  margin: 3pt 0 5pt;      /* increased spacing around title */
  line-height: 1.2;       /* increased for readability */
}

/* Section headers: clearer separation */
h2.section-header {
  font-size: 10.75pt;     /* increased from 10pt */
  letter-spacing: 0.4pt;  /* increased from 0.35pt */
  margin: 11pt 0 5pt;     /* increased spacing */
  line-height: 1.15;
}

/* Body text: more legible */
body {
  font-size: 9.75pt;      /* increased from 9.25pt */
  line-height: 1.3;       /* increased from 1.16 */
}

/* Bullets: more readable */
.entry-bullets li {
  font-size: 9.5pt;       /* increased from 8.95pt */
  line-height: 1.2;       /* increased from 1.15 */
}

/* Summary: better paragraph rhythm */
.summary-text {
  font-size: 9.5pt;       /* increased from 9pt */
  line-height: 1.3;       /* increased from 1.18 */
  margin: 0 0 10pt;       /* increased bottom margin */
}
```

---

## 2. Whitespace and Rhythm Tuning

### 2.1 Page Margins

| Side | Current | Target | Rationale |
|------|---------|--------|-----------|
| Top | 16mm | 16mm | Matches reference; adequate header space |
| Bottom | 16mm | 16mm | Matches reference; balanced with top |
| Left | 15mm | 15mm | Matches reference; sufficient reading comfort |
| Right | 15mm | 15mm | Matches reference; balanced with left |

### 2.2 Section Spacing

| Location | Current | Target | Rationale |
|----------|---------|--------|-----------|
| Header to summary | 6pt | 8–10pt | Create breathing room after name block |
| Summary to first section | 8pt | 10–12pt | Clear transition from intro to content |
| Section header spacing | 10pt 0 4pt | 11–12pt 0 5–6pt | More separation above, controlled below |
| Between sections | 0 (implicitly) | 10–12pt | Clear visual rhythm between content blocks |

### 2.3 Entry Spacing

| Location | Current | Target | Rationale |
|----------|---------|--------|-----------|
| Company header to role | 2pt | 4pt | Clear hierarchy between org and position |
| Role to bullets | 4pt (via margin) | 4–5pt | Sufficient separation for readability |
| Between bullets | 2pt | 2–3pt | Maintain list rhythm without excessive spacing |
| Between entries | 8pt | 10pt | Clear separation between jobs/degrees |
| Education status below degree | N/A | 3–4pt | Clear but not prominent |

### 2.4 Contact Block Spacing

| Location | Current | Target | Rationale |
|----------|---------|--------|-----------|
| Title to contact | 4pt | 5–6pt | Clear separation from name/title block |
| Contact items | Inline with · | Inline with · + 4pt padding | Matches reference treatment |
| Contact to summary | N/A | 8–10pt | Clear transition from header to content |

### 2.5 Whitespace Implementation Details

```css
/* Print CSS adjustments in cv-pdf.astro */

/* Header rhythm */
.cv-header {
  padding-bottom: 8pt;     /* increased from 6pt */
  margin-bottom: 10pt;     /* increased from 8pt */
}

/* Contact block treatment */
.contact-info {
  margin: 5pt 0 0;         /* clear spacing from title */
}

.contact-separator {
  padding: 0 4pt;          /* maintain current */
  font-size: 8.75pt;
}

/* Section header rhythm */
h2.section-header {
  margin: 12pt 0 6pt;      /* increased from 10pt 0 4pt */
  padding: 0 0 3pt;        /* increased from 0 0 2pt */
}

/* Entry rhythm */
.entry {
  margin: 0 0 10pt;        /* increased from 0 0 8pt */
}

.entry-header {
  margin: 0 0 4pt;         /* increased from 0 0 2pt */
}

.role-title,
.degree-title {
  margin: 0 0 5pt;         /* increased from 0 0 4pt */
}

/* Bullet rhythm */
.entry-bullets {
  margin: 4pt 0 0 16pt;    /* increased from 2pt 0 0 16pt */
}

.entry-bullets li {
  margin: 0 0 3pt;         /* increased from 0 0 2pt */
}

/* Education status */
.education-status {
  margin: 3pt 0 0;         /* clear spacing from degree */
}

/* Summary rhythm */
.summary-text {
  margin: 0 0 10pt;        /* increased from 0 0 8pt */
}

/* References note */
.references-note {
  margin: 12pt 0 0;        /* increased from 10pt 0 0 */
}
```

---

## 3. Section Architecture

### 3.1 Document Structure

```
┌─────────────────────────────────────────┐
│  Name (uppercase, 14.5pt)              │
│  Title (italic, 10.5pt)                │
│  Contact line (8.75–9pt, · separated)  │
├─────────────────────────────────────────┤
│  Summary paragraph (italic, 9.5pt)     │
├─────────────────────────────────────────┤
│  EXPERIENCIA PROFESIONAL (header)      │
│    • Company name (bold, 9.8pt)        │
│      Period (italic, 8.75–9pt)         │
│    • Role title (normal, 9pt)          │
│      • Bullet 1                        │
│      • Bullet 2                        │
│    [repeat entries]                    │
├─────────────────────────────────────────┤
│  EDUCACIÓN (header)                    │
│    • Institution (bold, 9.8pt)         │
│      Period (italic, 8.75–9pt)         │
│    • Degree (normal, 9pt)              │
│      Status (italic, 8.75pt)           │
│    [repeat entries]                    │
├─────────────────────────────────────────┤
│  HABILIDADES ADICIONALES (header)      │
│    • Category label (bold): item1,     │
│      item2 (item3 years)               │
│    [repeat categories]                 │
├─────────────────────────────────────────┤
│  Referencias disponibles a solicitud.  │
└─────────────────────────────────────────┘
```

### 3.2 Section Decisions

| Section | Decision | Rationale |
|---------|----------|-----------|
| Header block | Stay single, name/title/contact | Reference shows unified header, no separation |
| Summary | Keep as paragraph, italic | Reference maintains intro paragraph style |
| Experience | Keep current structure | Clear org/role/bullet hierarchy works |
| Education | Keep current structure | Consistent with experience format |
| Skills | Keep grouped categories | Organizes technology logically |
| References note | Keep centered at bottom | Reference shows this placement |
| Language | Keep Spanish headers | Matches reference and target audience |

### 3.3 Section Separation

- **Visual separation**: Use consistent 10–12pt spacing between major sections
- **Header treatment**: Uppercase, bold, bottom border for clear section identification
- **No structural changes**: Keep existing HTML structure; refine CSS only
- **Page breaks**: `page-break-after: avoid` on headers, `page-break-inside: avoid` on entries

### 3.4 Content Hierarchy

1. **Primary**: Name, company names, institution names
2. **Secondary**: Title, section headers, role titles, degrees
3. **Tertiary**: Period dates, contact items, status
4. **Quaternary**: Bullets, skill items, references note

### 3.5 Section-Specific Refinements

**Experience Section:**
- Company names: 9.8pt, bold, left-aligned
- Period dates: 8.75–9pt, italic, right-aligned
- Role titles: 9pt, normal, clear hierarchy
- Bullets: 2 per entry max, 9–10pt, 1.2 line-height

**Education Section:**
- Institution names: 9.8pt, bold, left-aligned
- Period dates: 8.75–9pt, italic, right-aligned
- Degrees: 9pt, normal
- Status: 8.75pt, italic, clear from degree

**Skills Section:**
- Category labels: Bold, inline with items
- Item format: `Technology Name (experience years)`
- Spacing: 3pt between skill items, grouped by category

---

## 4. Technology Mapper Design

### 4.1 Mapper Structure

**Location:** `web/scripts/generate-pdf.js`

**Implementation:**
```javascript
// Technology label normalization mapper
const TECH_LABEL_MAP = {
  // Programming languages
  'csharp': 'C#',
  'typescript': 'TypeScript',
  'javascript': 'JavaScript',
  'python': 'Python',
  'go': 'Go',
  'rust': 'Rust',
  'java': 'Java',
  'kotlin': 'Kotlin',
  'swift': 'Swift',
  'c': 'C',
  'cpp': 'C++',
  'ruby': 'Ruby',
  'php': 'PHP',
  'scala': 'Scala',
  'haskell': 'Haskell',
  'elixir': 'Elixir',
  'erlang': 'Erlang',

  // Frameworks
  'react': 'React',
  'vue': 'Vue.js',
  'angular': 'Angular',
  'svelte': 'Svelte',
  'nextjs': 'Next.js',
  'nuxt': 'Nuxt.js',
  'express': 'Express.js',
  'fastapi': 'FastAPI',
  'django': 'Django',
  'flask': 'Flask',
  'spring': 'Spring',
  'springboot': 'Spring Boot',
  'aspnet': 'ASP.NET',
  'laravel': 'Laravel',

  // Databases
  'postgresql': 'PostgreSQL',
  'mysql': 'MySQL',
  'sqlite': 'SQLite',
  'mongodb': 'MongoDB',
  'redis': 'Redis',
  'elasticsearch': 'Elasticsearch',

  // Cloud and infrastructure
  'aws': 'AWS',
  'azure': 'Azure',
  'gcp': 'GCP',
  'kubernetes': 'Kubernetes',
  'docker': 'Docker',
  'terraform': 'Terraform',
  'ansible': 'Ansible',

  // DevOps and tools
  'git': 'Git',
  'jenkins': 'Jenkins',
  'githubactions': 'GitHub Actions',
  'circleci': 'CircleCI',
  'prometheus': 'Prometheus',
  'grafana': 'Grafana',

  // Messaging
  'kafka': 'Apache Kafka',
  'rabbitmq': 'RabbitMQ',
  'activemq': 'ActiveMQ',

  // IoT and embedded
  'arduino': 'Arduino',
  'esp32': 'ESP32',
  'raspberrypi': 'Raspberry Pi',
  'micropython': 'MicroPython',

  // Testing
  'jest': 'Jest',
  'pytest': 'PyTest',
  'selenium': 'Selenium',
  'cypress': 'Cypress',

  // Build tools
  'webpack': 'Webpack',
  'vite': 'Vite',
  'babel': 'Babel',
  'maven': 'Maven',
  'gradle': 'Gradle',
  'npm': 'npm',
  'yarn': 'Yarn',
  'pnpm': 'pnpm',

  // Other
  'graphql': 'GraphQL',
  'grpc': 'gRPC',
  'rest': 'REST',
  'soap': 'SOAP',
  'oauth': 'OAuth',
  'jwt': 'JWT',
  'json': 'JSON',
  'xml': 'XML',
  'yaml': 'YAML',
  'markdown': 'Markdown',
  'latex': 'LaTeX',
  'vim': 'Vim',
  'emacs': 'Emacs',
  'vscode': 'VS Code',
  'linux': 'Linux',
  'windows': 'Windows',
  'macos': 'macOS',
  'bash': 'Bash',
  'powershell': 'PowerShell',
  'zsh': 'Zsh',
};

/**
 * Normalizes a technology tag to a human-friendly label.
 * Uses the explicit mapping if available; otherwise falls back to
 * capitalizing the first character.
 *
 * @param {string} tag - The source technology tag (e.g., 'csharp')
 * @returns {string} - The normalized label (e.g., 'C#')
 */
function normalizeTechLabel(tag) {
  if (!tag || typeof tag !== 'string') {
    return '';
  }

  const normalizedTag = tag.toLowerCase().trim();

  // Check explicit mapping first
  if (TECH_LABEL_MAP[normalizedTag]) {
    return TECH_LABEL_MAP[normalizedTag];
  }

  // Fallback: capitalize first character, keep rest as-is
  return normalizedTag.charAt(0).toUpperCase() + normalizedTag.slice(1);
}
```

### 4.2 Integration with Data Flow

**Current flow:**
```
cv-processed.json → cv-pdf.astro → HTML → Puppeteer → PDF
```

**New flow (unchanged schema):**
```
cv-processed.json → cv-pdf.astro → HTML → Puppeteer → PDF
                                           (inject mapper in DOM)
```

**Key decision:** The mapper runs in `generate-pdf.js` during DOM injection, **not** in `cv-pdf.astro`. This keeps the mapper isolated to PDF generation and doesn't affect the public web page.

### 4.3 Mapper Invocation

```javascript
// In injectPrivateData() function within generate-pdf.js
async function injectPrivateData(page) {
  await page.evaluate((data, techMap) => {
    // ... existing contact injection code ...

    // Normalize technology labels in skills section
    const skillItems = document.querySelectorAll('.skill-list li');
    skillItems.forEach(item => {
      const textContent = item.textContent;
      // Technology names appear before parentheses
      const updatedText = textContent.replace(
        /([a-zA-Z0-9_\-]+)\s*\(/g,
        (match, techName) => {
          const normalized = normalizeTechLabel(techName);
          return `${normalized} (`;
        }
      );
      item.textContent = updatedText;
    });

    // Helper function available in the evaluate context
    function normalizeTechLabel(tag) {
      if (!tag || typeof tag !== 'string') return '';
      const normalizedTag = tag.toLowerCase().trim();
      return techMap[normalizedTag] ||
             normalizedTag.charAt(0).toUpperCase() + normalizedTag.slice(1);
    }
  }, privateData, TECH_LABEL_MAP);
}
```

### 4.4 Fallback Behavior

1. **Explicit mapping:** Use predefined map (e.g., `csharp` → `C#`)
2. **Implicit fallback:** Capitalize first character (e.g., `rust` → `Rust`)
3. **Empty/invalid:** Return empty string, skip rendering

### 4.5 Data Invariants

- Source data (`cv-processed.json`) remains unchanged
- Mapper is only applied during PDF generation
- Public web page renders raw tags (no normalization)
- No schema changes required in Go pipeline or YAML source

### 4.6 Technology Label Map Maintenance

**Extension guideline:**
- Add new mappings to `TECH_LABEL_MAP` as needed
- Use canonical capitalization (e.g., `TypeScript`, not `Typescript`)
- Preserve existing mappings to avoid regressions
- Keep map sorted alphabetically for maintainability

---

## 5. Contact Block Presentation

### 5.1 Contact Block Structure

```
[Name]
[Title]
[Address] · [LinkedIn] · [Phone] · [Email] · [Portafolio] · [GitHub]
```

### 5.2 Visual Treatment

| Element | Font Size | Style | Separator | Spacing |
|---------|-----------|-------|-----------|---------|
| Contact items | 8.75–9pt | Normal | ` · ` (with 4pt padding) | Inline |
| All items | No links (text-only) | Normal | Between each item | No break |

### 5.3 Ordering Rule

```
1. Address (if present)
2. LinkedIn (if present)
3. Phone (if present)
4. Email (if present)
5. Website/Portafolio (if present)
6. GitHub (if present)
```

**Rationale:** Matches current implementation and provides logical progression from physical location to professional profiles to contact methods.

### 5.4 Separator Treatment

- **Character:** Middle dot `·` (U+00B7)
- **Spacing:** 4pt padding on each side
- **Color:** Black (#000), same as text
- **Font:** Times New Roman, same as contact items

```css
.contact-separator {
  padding: 0 4pt;
  font-family: 'Times New Roman', Times, serif;
  font-size: 8.75pt;
  color: #000;
}
```

### 5.5 No Icons

- **Decision:** Remove all icons from contact block
- **Implementation:** Text-only presentation via Puppeteer injection
- **Rationale:** Matches reference formal style and ATS compatibility

### 5.6 Link Treatment

- **Decision:** Remove hyperlinks from contact items in PDF
- **Implementation:** Text-only in PDF, even if source has URLs
- **Rationale:** PDF is static document; links add visual clutter and don't function in print

### 5.7 Contact Block Spacing

```css
/* Title to contact */
.professional-title {
  margin: 3pt 0 5pt;
}

/* Contact block */
.contact-info {
  margin: 5pt 0 0;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  align-items: baseline;
  gap: 0;
}

/* Contact to summary transition */
.cv-header {
  padding-bottom: 8pt;
  margin-bottom: 10pt;
  border-bottom: 1pt solid #000;
}
```

---

## 6. Constraints

### 6.1 PDF-Only Scope

| File | Modified? | Reason |
|------|-----------|--------|
| `web/src/pages/cv-pdf.astro` | ✅ Yes | Typography, spacing, rhythm refinements |
| `web/scripts/generate-pdf.js` | ✅ Yes | Technology label mapper |
| `web/src/pages/index.astro` | ❌ No | Public web page unchanged |
| `web/src/styles/*.css` | ❌ No | No site style changes |
| `cv-processed.json` | ❌ No | No data schema changes |
| `cv-source.yaml` | ❌ No | No source data changes |

### 6.2 ATS Compatibility

| Requirement | Implementation |
|-------------|----------------|
| Single-column layout | Maintain current structure |
| Selectable text | No content as images; use standard HTML/CSS |
| Standard bullets | Use `ul`/`li` with default bullets or simple dashes |
| No visual tricks | Avoid icons, graphics, or complex formatting |
| Clear headings | Use semantic `h1`, `h2`, `h3` elements |
| Consistent formatting | Uniform typography and spacing across sections |

### 6.3 Document Format Constraints

- **Paper size:** A4 (210mm × 297mm)
- **Margins:** 16mm (top/bottom), 15mm (left/right)
- **Orientation:** Portrait
- **Background:** White (#fff)
- **Font family:** Times New Roman, Times, serif
- **Page count:** 1–2 pages (accept slight expansion for readability)

### 6.4 Language Constraints

- **Headers:** Spanish (`EXPERIENCIA PROFESIONAL`, `EDUCACIÓN`, etc.)
- **Content:** Spanish (names, titles, descriptions)
- **Technology labels:** English (`C#`, `TypeScript`, etc.)
- **Rationale:** Matches reference and target hiring market

### 6.5 Visual Alignment Constraints

- **No metric cloning:** Use similitude, not exact pixel measurements
- **No screen optimization:** Preview styles (`@media screen`) are secondary
- **No visual tricks:** Avoid CSS hacks for layout; use standard properties
- **No experimental features:** Use stable, cross-browser compatible CSS

---

## 7. Implementation Sequence

### 7.1 Phase Order

**Rationale:** Minimize risk and enable incremental testing. Implement mapper first (no visual impact), then typography and spacing in progressive refinement.

```
Phase 1: Technology Label Mapper (low risk)
  ↓
Phase 2: Typography Scale (medium risk, visible change)
  ↓
Phase 3: Whitespace and Rhythm (medium risk, visible change)
  ↓
Phase 4: Contact Block Refinement (low risk, localized change)
  ↓
Phase 5: Integration Testing (validation phase)
```

### 7.2 Detailed Implementation Steps

#### Phase 1: Technology Label Mapper

1. **Add mapper object to `generate-pdf.js`**
   - Insert `TECH_LABEL_MAP` constant at file top
   - Populate with common technology mappings
   - Sort alphabetically for maintainability

2. **Implement normalization function**
   - Add `normalizeTechLabel(tag)` function
   - Apply explicit mapping lookup
   - Implement fallback capitalization

3. **Integrate mapper with DOM injection**
   - Modify `injectPrivateData()` to accept mapper
   - Add DOM manipulation to normalize skill labels
   - Use regex to find and replace technology names

4. **Test mapper in isolation**
   - Run `pnpm export-pdf` and inspect generated PDF
   - Verify technology labels are normalized
   - Check fallback behavior for unmapped tags

**Files modified:** `web/scripts/generate-pdf.js`
**Lines changed:** ~80–100 lines
**Risk:** Low (isolated to PDF generation)

---

#### Phase 2: Typography Scale

1. **Update font sizes in print CSS**
   - Increase body text: 9.25pt → 9.75pt
   - Increase section headers: 10pt → 10.75pt
   - Increase title: 9.5pt → 10.5pt
   - Increase bullets: 8.95pt → 9.5pt

2. **Update line heights in print CSS**
   - Body text: 1.16 → 1.3
   - Bullets: 1.15 → 1.2
   - Section headers: Add explicit 1.15

3. **Update letter spacing in print CSS**
   - Section headers: 0.35pt → 0.4pt
   - Keep name tracking at 1pt

4. **Test typography changes**
   - Run `pnpm export-pdf` and review legibility
   - Compare with reference for proportional similitude
   - Verify page count hasn't expanded beyond 2 pages

**Files modified:** `web/src/pages/cv-pdf.astro`
**Lines changed:** ~50–70 lines
**Risk:** Medium (visible change, affects all content)

---

#### Phase 3: Whitespace and Rhythm

1. **Update page margins**
   - No change needed (already 16mm/15mm)
   - Verify values match `@page` and Puppeteer margin config

2. **Update section spacing**
   - Section headers: 10pt 0 4pt → 12pt 0 6pt
   - Summary to first section: 8pt → 11pt
   - Between sections: Implicit → 11pt

3. **Update entry spacing**
   - Company to role: 2pt → 4pt
   - Role to bullets: 4pt → 5pt
   - Between bullets: 2pt → 3pt
   - Between entries: 8pt → 10pt

4. **Update header block spacing**
   - Name to title: 2pt → 3pt
   - Title to contact: 4pt → 5pt
   - Contact to summary: 8pt → 10pt

5. **Test rhythm changes**
   - Run `pnpm export-pdf` and review visual flow
   - Check for orphaned content or awkward breaks
   - Verify clear separation between sections

**Files modified:** `web/src/pages/cv-pdf.astro`
**Lines changed:** ~60–80 lines
**Risk:** Medium (visible change, affects spacing globally)

---

#### Phase 4: Contact Block Refinement

1. **Verify contact block treatment**
   - Confirm icons are removed (already done in parent change)
   - Verify text-only presentation
   - Check separator treatment (` · ` with 4pt padding)

2. **Refine contact block spacing**
   - Title to contact: 5pt
   - Contact to summary: 10pt
   - Verify inline layout with centering

3. **Test contact block**
   - Run `pnpm export-pdf` and review contact line
   - Verify ordering: address, LinkedIn, phone, email, portfolio, GitHub
   - Check for line wrapping issues

**Files modified:** `web/src/pages/cv-pdf.astro`
**Lines changed:** ~10–20 lines
**Risk:** Low (localized to header block)

---

#### Phase 5: Integration Testing

1. **Full PDF generation test**
   ```bash
   cd /src/cv-pipeline/web
   pnpm build
   pnpm export-pdf
   ```

2. **Visual validation checklist**
   - [ ] Typography is larger and more legible than before
   - [ ] Section spacing provides clear visual separation
   - [ ] Technology labels are normalized (e.g., `C#`, not `csharp`)
   - [ ] Document rhythm feels similar to reference
   - [ ] Contact block is clean and readable
   - [ ] No visual clutter or awkward breaks
   - [ ] Page count is 1–2 pages

3. **ATS compatibility test**
   - [ ] Selectable text works (no images)
   - [ ] Single-column layout maintained
   - [ ] Semantic HTML structure preserved
   - [ ] No formatting that blocks parsers

4. **Regression check**
   - [ ] Public web page (`index.astro`) unchanged
   - [ ] No changes to data pipeline
   - [ ] Build passes without errors
   - [ ] PDF generation succeeds

**Files modified:** None (testing phase)
**Risk:** None (validation only)

---

### 7.3 Rollback Plan

If any phase introduces issues:

1. **Phase 5:** No rollback needed (testing only)
2. **Phase 4:** Revert contact block spacing changes (~15 lines)
3. **Phase 3:** Revert margin and spacing changes (~70 lines)
4. **Phase 2:** Revert font size and line-height changes (~60 lines)
5. **Phase 1:** Remove `TECH_LABEL_MAP` and `normalizeTechLabel` (~90 lines)

**Rollback procedure:**
```bash
git checkout web/src/pages/cv-pdf.astro
git checkout web/scripts/generate-pdf.js
pnpm build
pnpm export-pdf
```

---

## 8. Acceptance Validation

### 8.1 Typography Validation

| Criterion | How to Validate |
|-----------|-----------------|
| Font sizes larger | Compare generated PDF with previous version; measure text |
| Line heights improved | Check readability; compare with reference |
| Letter spacing appropriate | Review uppercase headers for balanced tracking |
| Font hierarchy clear | Name > section headers > body > bullets > dates |

### 8.2 Document Rhythm Validation

| Criterion | How to Validate |
|-----------|-----------------|
| Margins adequate | Verify 16mm/15mm on all sides |
| Section spacing clear | Count space between major sections |
| Entry separation visible | Check space between jobs/degrees |
| Whitespace intentional | No cramped or orphaned content |

### 8.3 Technology Label Validation

| Criterion | How to Validate |
|-----------|-----------------|
| Labels normalized | Inspect skills section for `C#`, `TypeScript`, `Python` |
| No raw IDs | Verify no lowercase IDs like `csharp`, `typescript` |
| Fallback works | Check unmapped tags are capitalized |
| Source unchanged | Confirm `cv-processed.json` still has raw IDs |

### 8.4 Contact Block Validation

| Criterion | How to Validate |
|-----------|-----------------|
| Icons removed | Visual inspection; no symbols or graphics |
| Text-only presentation | No hyperlinks in PDF |
| Separator correct | Middle dot `· ` with 4pt padding |
| Ordering maintained | Address → LinkedIn → Phone → Email → Portfolio → GitHub |

### 8.5 Visual Parity Validation

| Criterion | How to Validate |
|-----------|-----------------|
| Reference similitude | Side-by-side visual comparison |
| Proportions feel right | Layout matches reference rhythm |
| Not metric cloning | Accept differences if overall feel matches |
| Legibility improved | Text is easier to read than before |

### 8.6 ATS Compatibility Validation

| Criterion | How to Validate |
|-----------|-----------------|
| Selectable text | Highlight text in PDF viewer |
| Single-column | No multi-column layout detected |
| Semantic structure | HTML uses `h1`, `h2`, `h3`, `ul`, `li` |
| Parseable formatting | No images as text, no complex layouts |

---

## 9. Success Criteria

### 9.1 Functional Success

- [ ] PDF generates successfully with `pnpm export-pdf`
- [ ] Technology labels are normalized (`C#`, `TypeScript`, `Python`)
- [ ] Typography is larger and more legible than before
- [ ] Document rhythm provides clear visual separation
- [ ] Contact block is clean, text-only, with correct separators

### 9.2 Visual Success

- [ ] PDF achieves visual similitude with reference
- [ ] Typography feels balanced and professional
- [ ] Whitespace creates intentional rhythm
- [ ] No cramped or orphaned content
- [ ] Page count remains 1–2 pages (accept slight expansion)

### 9.3 Technical Success

- [ ] No changes to public web page or site styles
- [ ] No schema changes in data pipeline
- [ ] Build passes without errors
- [ ] Changes fit within 400-line review budget
- [ ] Rollback is straightforward if needed

### 9.4 Product Success

- [ ] ATS systems can reliably parse the CV
- [ ] Recruiters receive a more readable, professional CV
- [ ] Technology labels are human-friendly
- [ ] Visual improvements enhance credibility
- [ ] Private data remains secure in `.env`

---

## 10. Risk Mitigation

### 10.1 Page Density Risk

**Risk:** Larger typography and spacing may expand page count beyond 2 pages.

**Mitigation:**
- Accept visual parity as higher priority than page count
- If page count exceeds 2 pages significantly, consider minor spacing reductions (not typography)
- Reference itself may use more than 2 pages with similar rhythm

### 10.2 Unmapped Technology Tags

**Risk:** Some tags may lack explicit mappings in `TECH_LABEL_MAP`.

**Mitigation:**
- Fallback capitalization handles unknown tags gracefully
- Map can be extended in future iterations without code changes
- Review common tags in `cv-processed.json` to ensure coverage

### 10.3 Visual Interpretation

**Risk:** "Visual parity" involves subjective judgment.

**Mitigation:**
- Focus on objective criteria: legibility, rhythm, hierarchy
- Use reference as guide, not strict template
- Prioritize readability over exact similitude
- Document design decisions for future reference

### 10.4 Preview Inconsistency

**Risk:** Screen preview may look different from final PDF.

**Mitigation:**
- Accept preview differences; prioritize final PDF output
- Print CSS is authoritative; screen styles are secondary
- Test final PDF, not browser preview

### 10.5 Regression in Web Page

**Risk:** Changes might accidentally affect public web page.

**Mitigation:**
- Only modify `cv-pdf.astro`, never `index.astro`
- Verify `index.astro` unchanged after implementation
- Test public site deployment to ensure no regressions

---

## 11. Design Tradeoffs

### 11.1 Typography vs. Page Count

**Tradeoff:** Larger typography improves legibility but may increase page count.

**Decision:** Prioritize legibility; accept 1–2 pages with possible slight expansion. Reference visual quality is more important than minimal page count.

### 11.2 Exact Mapping vs. Fallback

**Tradeoff:** Exact mappings for all tags vs. simple fallback capitalization.

**Decision:** Use explicit mappings for common technologies, fallback for others. Extensible design allows future additions without refactoring.

### 11.3 Visual Similitude vs. Metric Cloning

**Tradeoff:** Exact pixel measurements vs. proportional similitude.

**Decision:** Similitude over cloning. Reference proportions guide decisions, but exact metrics are not required.

### 11.4 Contact Icons vs. Text-Only

**Tradeoff:** Icons add visual interest vs. ATS compatibility and formal style.

**Decision:** Text-only presentation (already decided in parent change). Icons compromise ATS parsing and formality.

### 11.5 Screen Preview vs. PDF Output

**Tradeoff:** Optimize for browser preview vs. optimize for final PDF.

**Decision:** Prioritize PDF output. Screen styles are secondary; the artifact is the PDF, not the preview.

---

## 12. Open Questions

No blocking questions. The design specifies all implementation details.

**Optional refinements for future consideration:**
- Should additional technology mappings be added beyond the initial set?
- Should fallback capitalization use title case for multi-word tags (e.g., `aspnet` → `Aspnet` vs `ASP.NET`)?
- Should section spacing be fine-tuned based on final page count?

---

## 13. Related Artifacts

### 13.1 Parent Changes

- `cv-pdf-harvard-formalization` — Harvard formality restoration and icon removal
- This change builds on that foundation without revisiting previous decisions

### 13.2 Proposal

- `/src/cv-pipeline/openspec/proposals/cv-pdf-reference-alignment.md`
- Defines problem statement, goals, scope, and acceptance criteria

### 13.3 Spec

- `/src/cv-pipeline/openspec/specs/cv-pdf-reference-alignment.md`
- Defines WHAT must be true (visual parity, legibility, normalization)

### 13.4 Reference PDF

- Original reference document (not modified)
- Used as visual guide for proportions and rhythm

---

## Appendix A: CSS Reference

### A.1 Print CSS Structure

```css
@page {
  size: A4;
  margin: 16mm 15mm 16mm 15mm;
}

@media print {
  body {
    font-family: 'Times New Roman', Times, serif;
    font-size: 9.75pt;          /* increased from 9.25pt */
    line-height: 1.3;           /* increased from 1.16 */
    color: #000;
    background: #fff;
    margin: 0;
    padding: 0;
    -webkit-print-color-adjust: exact;
    print-color-adjust: exact;
  }

  .cv-header {
    text-align: center;
    padding-bottom: 8pt;        /* increased from 6pt */
    margin-bottom: 10pt;        /* increased from 8pt */
    border-bottom: 1pt solid #000;
  }

  .cv-header h1 {
    margin: 0;
    font-family: 'Times New Roman', Times, serif;
    font-size: 14.5pt;          /* increased from 14pt */
    font-weight: bold;
    text-transform: uppercase;
    letter-spacing: 1pt;
    line-height: 1.06;          /* increased from 1.05 */
  }

  .professional-title {
    margin: 3pt 0 5pt;          /* increased from 2pt 0 4pt */
    font-family: 'Times New Roman', Times, serif;
    font-size: 10.5pt;          /* increased from 9.5pt */
    font-style: italic;
    font-weight: normal;
    color: #333;
    line-height: 1.2;
  }

  .contact-info {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    align-items: baseline;
    gap: 0;
    margin: 5pt 0 0;           /* spacing from title */
    padding: 0;
  }

  .contact-item,
  .contact-separator {
    display: inline-flex;
    align-items: baseline;
    font-family: 'Times New Roman', Times, serif;
    font-size: 8.75pt;
    line-height: 1;
    color: #000;
    white-space: nowrap;
  }

  .contact-separator {
    padding: 0 4pt;
  }

  .summary-text {
    margin: 0 0 10pt;          /* increased from 0 0 8pt */
    font-family: 'Times New Roman', Times, serif;
    font-size: 9.5pt;          /* increased from 9pt */
    font-style: italic;
    line-height: 1.3;          /* increased from 1.18 */
    text-align: justify;
    color: #111;
  }

  .section {
    margin-top: 0;
  }

  h2.section-header {
    margin: 12pt 0 6pt;        /* increased from 10pt 0 4pt */
    padding: 0 0 3pt;          /* increased from 0 0 2pt */
    font-family: 'Times New Roman', Times, serif;
    font-size: 10.75pt;        /* increased from 10pt */
    font-weight: bold;
    text-transform: uppercase;
    letter-spacing: 0.4pt;     /* increased from 0.35pt */
    border-bottom: 1pt solid #000;
    page-break-after: avoid;
    page-break-inside: avoid;
    line-height: 1.15;
  }

  .entry {
    margin: 0 0 10pt;          /* increased from 0 0 8pt */
    page-break-inside: avoid;
  }

  .entry-header {
    display: flex;
    justify-content: space-between;
    align-items: baseline;
    gap: 10pt;
    margin: 0 0 4pt;           /* increased from 0 0 2pt */
  }

  h3.company-name,
  h3.institution-name,
  h3.project-name {
    margin: 0;
    font-family: 'Times New Roman', Times, serif;
    font-size: 9.8pt;
    font-weight: bold;
    line-height: 1.1;
  }

  .period {
    margin: 0;
    font-family: 'Times New Roman', Times, serif;
    font-size: 8.75pt;
    font-style: italic;
    white-space: nowrap;
    color: #111;
  }

  .role-title,
  .degree-title,
  .project-title {
    margin: 0 0 5pt;           /* increased from 0 0 4pt */
    font-family: 'Times New Roman', Times, serif;
    font-size: 9pt;
    font-weight: normal;
    line-height: 1.16;
  }

  .entry-bullets {
    margin: 4pt 0 0 16pt;     /* increased from 2pt 0 0 16pt */
    padding: 0;
  }

  .entry-bullets li {
    margin: 0 0 3pt;          /* increased from 0 0 2pt */
    font-family: 'Times New Roman', Times, serif;
    font-size: 9.5pt;         /* increased from 8.95pt */
    line-height: 1.2;         /* increased from 1.15 */
  }

  .skill-list {
    margin: 0;
    padding: 0 0 0 12pt;
  }

  .skill-list li {
    margin: 0 0 3pt;
    font-family: 'Times New Roman', Times, serif;
    font-size: 9.5pt;         /* increased from 8.95pt */
    line-height: 1.2;         /* increased from 1.15 */
  }

  .skill-label {
    font-weight: bold;
  }

  .education-status {
    margin: 3pt 0 0;
    font-family: 'Times New Roman', Times, serif;
    font-size: 8.75pt;
    font-style: italic;
    line-height: 1.12;
  }

  .references-note {
    margin: 12pt 0 0;         /* increased from 10pt 0 0 */
    font-family: 'Times New Roman', Times, serif;
    font-size: 8.5pt;
    font-style: italic;
    text-align: center;
  }

  .section,
  .entry,
  .skill-list,
  .references-note {
    page-break-inside: avoid;
  }
}
```

---

## Appendix B: Technology Mapper Reference

### B.1 Mapper Object (Initial Set)

```javascript
const TECH_LABEL_MAP = {
  // Programming languages
  'csharp': 'C#',
  'typescript': 'TypeScript',
  'javascript': 'JavaScript',
  'python': 'Python',
  'go': 'Go',
  'rust': 'Rust',
  'java': 'Java',
  'kotlin': 'Kotlin',
  'swift': 'Swift',
  'c': 'C',
  'cpp': 'C++',
  'ruby': 'Ruby',
  'php': 'PHP',
  'scala': 'Scala',
  'haskell': 'Haskell',
  'elixir': 'Elixir',
  'erlang': 'Erlang',

  // Frameworks
  'react': 'React',
  'vue': 'Vue.js',
  'angular': 'Angular',
  'svelte': 'Svelte',
  'nextjs': 'Next.js',
  'nuxt': 'Nuxt.js',
  'express': 'Express.js',
  'fastapi': 'FastAPI',
  'django': 'Django',
  'flask': 'Flask',
  'spring': 'Spring',
  'springboot': 'Spring Boot',
  'aspnet': 'ASP.NET',
  'laravel': 'Laravel',

  // Databases
  'postgresql': 'PostgreSQL',
  'mysql': 'MySQL',
  'sqlite': 'SQLite',
  'mongodb': 'MongoDB',
  'redis': 'Redis',
  'elasticsearch': 'Elasticsearch',

  // Cloud and infrastructure
  'aws': 'AWS',
  'azure': 'Azure',
  'gcp': 'GCP',
  'kubernetes': 'Kubernetes',
  'docker': 'Docker',
  'terraform': 'Terraform',
  'ansible': 'Ansible',

  // DevOps and tools
  'git': 'Git',
  'jenkins': 'Jenkins',
  'githubactions': 'GitHub Actions',
  'circleci': 'CircleCI',
  'prometheus': 'Prometheus',
  'grafana': 'Grafana',

  // Messaging
  'kafka': 'Apache Kafka',
  'rabbitmq': 'RabbitMQ',
  'activemq': 'ActiveMQ',

  // IoT and embedded
  'arduino': 'Arduino',
  'esp32': 'ESP32',
  'raspberrypi': 'Raspberry Pi',
  'micropython': 'MicroPython',

  // Testing
  'jest': 'Jest',
  'pytest': 'PyTest',
  'selenium': 'Selenium',
  'cypress': 'Cypress',

  // Build tools
  'webpack': 'Webpack',
  'vite': 'Vite',
  'babel': 'Babel',
  'maven': 'Maven',
  'gradle': 'Gradle',
  'npm': 'npm',
  'yarn': 'Yarn',
  'pnpm': 'pnpm',

  // Other
  'graphql': 'GraphQL',
  'grpc': 'gRPC',
  'rest': 'REST',
  'soap': 'SOAP',
  'oauth': 'OAuth',
  'jwt': 'JWT',
  'json': 'JSON',
  'xml': 'XML',
  'yaml': 'YAML',
  'markdown': 'Markdown',
  'latex': 'LaTeX',
  'vim': 'Vim',
  'emacs': 'Emacs',
  'vscode': 'VS Code',
  'linux': 'Linux',
  'windows': 'Windows',
  'macos': 'macOS',
  'bash': 'Bash',
  'powershell': 'PowerShell',
  'zsh': 'Zsh',
};
```

### B.2 Normalization Function

```javascript
function normalizeTechLabel(tag) {
  if (!tag || typeof tag !== 'string') {
    return '';
  }

  const normalizedTag = tag.toLowerCase().trim();

  // Check explicit mapping first
  if (TECH_LABEL_MAP[normalizedTag]) {
    return TECH_LABEL_MAP[normalizedTag];
  }

  // Fallback: capitalize first character
  return normalizedTag.charAt(0).toUpperCase() + normalizedTag.slice(1);
}
```

---

## Appendix C: Validation Checklist

### C.1 Pre-Implementation Checklist

- [ ] Reference PDF available for comparison
- [ ] `.env` file contains `PRIVATE_*` variables
- [ ] `cv-processed.json` is up-to-date
- [ ] Astro build passes: `cd web && pnpm build`
- [ ] Current PDF generation works: `cd web && pnpm export-pdf`

### C.2 Post-Implementation Checklist

- [ ] PDF generates without errors
- [ ] Typography changes are visible and positive
- [ ] Spacing changes improve document rhythm
- [ ] Technology labels are normalized
- [ ] Contact block is clean and correct
- [ ] Visual parity with reference is achieved
- [ ] ATS compatibility is maintained
- [ ] Page count is 1–2 pages
- [ ] Public web page is unchanged
- [ ] Build passes without errors

### C.3 Rollback Validation (if needed)

- [ ] Reverted files are correct
- [ ] Build passes after rollback
- [ ] PDF generation works after rollback
- [ ] No residual changes in unrelated files

---

**Design document complete.**
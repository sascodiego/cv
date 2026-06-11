# SDD Tasks: cv-pdf-harvard

**Change ID:** `cv-pdf-harvard`  
**Phase:** SDD Tasks  
**Status:** Ready  
**Created:** 2026-06-10  

---

## Review Workload Forecast

| Field | Value |
|-------|-------|
| Estimated changed lines | 280-330 |
| 400-line budget risk | Low |
| Chained PRs recommended | No |
| Suggested split | Single PR |
| Delivery strategy | single-pr |
| Chain strategy | size-exception |

```
Decision needed before apply: No
Chained PRs recommended: No
Chain strategy: size-exception
400-line budget risk: Low
```

**Risk Assessment:**
- New file `cv-pdf.astro`: ~250-300 lines (HTML ~150, CSS ~80, frontmatter ~50)
- Modified file `generate-pdf.js`: 3 lines changed
- Total: ~280-330 lines, comfortably under 400-line budget
- Changes are isolated (new template + 3 script modifications)
- Low risk of merging conflicts with existing web page

---

## Phase 1: Template Setup (cv-pdf.astro)

- [ ] **1.1** Create `web/src/pages/cv-pdf.astro` with basic Astro page structure (frontmatter, HTML shell, `<style>` container)
- [ ] **1.2** Import `cv-processed.json` in frontmatter: `import cvData from '../data/cv-processed.json'`
- [ ] **1.3** Destructure top-level data: `const { personal_info, skills, work_experience, education, recommendations } = cvData`
- [ ] **1.4** Set HTML `<title>` to `{personal_info.name} - CV` and `<html lang="es">`

---

## Phase 2: Frontmatter Logic Implementation

- [ ] **2.1** Define `categoryMapping` constant in frontmatter with mappings:
  - `'Languages & Frameworks'` → `'Languages & Frameworks'`
  - `'Databases & Obs'` → `'Databases & Messaging'`
  - `'Architecture & Protocols'` → `'Databases & Messaging'`
  - `'Infrastructure & DevOps'` → `'Tools & Platforms'`
  - `'Automation & QA'` → `'Tools & Platforms'`
  - `'IoT & Embedded'` → `'IoT & Embedded'`
  - `'Metodologías & Prácticas'` → `null` (omit)
  - `'Habilidades Funcionales'` → `null` (omit)
  - `'Competencias Blandas'` → `null` (omit)
- [ ] **2.2** Implement `skillGroups` reduction function to group skills by mapped category (filter out `null` categories)
- [ ] **2.3** Implement `sortedCategories` array from `Object.keys(skillGroups).sort()`
- [ ] **2.4** Implement `allProjects` extraction: filter `work_experience` for entries with `projects` array, then `flatMap` to extract each project with `company` and `entryPeriod` context
- [ ] **2.5** Implement `sortedProjects` by `start_date` descending: `allProjects.sort((a, b) => new Date(b.start_date) - new Date(a.start_date))`
- [ ] **2.6** Implement `selectedProjects` as `sortedProjects.slice(0, 5)` (limit to 5 projects)
- [ ] **2.7** Implement `truncateSummary(text, maxLength = 200)` function:
  - Return text if `length <= maxLength`
  - Find last sentence boundary (`.`) within truncated substring
  - Return text up to last sentence end, trimmed
- [ ] **2.8** Assign `summaryText = truncateSummary(personal_info.about_me, 200)`
- [ ] **2.9** Implement `sortedWorkExperience`: sort by `end_date` descending (handle `'Presente'` → `'2999-12-31'` for sorting)
- [ ] **2.10** Implement `sortedEducation`: sort by `end_date` descending
- [ ] **2.11** Implement `splitDescriptionToBullets(description, maxBullets = 3)` function:
  - Match sentences with regex `/[^.!?]+[.!?]+/g` or return array with description
  - Return `slice(0, maxBullets).map(s => s.trim())`

---

## Phase 3: HTML Structure Implementation

- [ ] **3.1** Create root `<div class="cv-container">` containing all sections
- [ ] **3.2** Implement Section 1 (Header) `.cv-header`:
  - `<h1>{personal_info.name}</h1>`
  - `<p class="professional-title">{personal_info.title}</p>`
  - `<div class="contact-info">` (empty, Puppeteer injects content here)
- [ ] **3.3** Implement Section 2 (Summary) `.summary-section`:
  - `<h2 class="section-header">PROFESSIONAL SUMMARY</h2>`
  - `<p class="summary-text">{summaryText}</p>`
- [ ] **3.4** Implement Section 3 (Skills) `.skills-section`:
  - `<h2 class="section-header">TECHNICAL SKILLS</h2>`
  - Loop through `sortedCategories` to create `.skill-category` blocks
  - Each category: `<h3 class="skill-category-header">{category}</h3>` and `<p class="skill-list">` with skill names + `experience_text` in parentheses, comma-separated
- [ ] **3.5** Implement Section 4 (Experience) `.experience-section`:
  - `<h2 class="section-header">PROFESSIONAL EXPERIENCE</h2>`
  - Loop through `sortedWorkExperience` to create `.experience-entry` blocks
  - Each entry: `.entry-header` (flex row) with `<h3 class="company-name">` and `<span class="period">` (right-aligned)
  - Below header: `<p class="role-title">` with role
  - Below role: `<ul class="entry-bullets">` with bullets from `splitDescriptionToBullets(description)`
  - Below bullets: `<p class="technologies">Technologies: {technologies.join(', ')}</p>` (if technologies array exists)
- [ ] **3.6** Implement Section 5 (Projects) `.projects-section`:
  - Wrap entire section in `{selectedProjects.length > 0 && ( ... )}` conditional
  - `<h2 class="section-header">KEY PROJECTS</h2>`
  - Loop through `selectedProjects` to create `.project-entry` blocks
  - Each project: `.entry-header` (flex row) with `<h3 class="project-name">{name} ({company})</h3>` and `<span class="period">` (right-aligned)
  - Below header: `<p class="project-tech">{technologies.join(', ')}</p>` (smaller font)
  - Below tech: `<ul class="entry-bullets">` with bullets from `splitDescriptionToBullets(description)`
- [ ] **3.7** Implement Section 6 (Education) `.education-section`:
  - `<h2 class="section-header">EDUCATION</h2>`
  - Loop through `sortedEducation` to create `.education-entry` blocks
  - Each entry: `.entry-header` (flex row) with `<h3 class="institution-name">` and `<span class="period">` (right-aligned)
  - Below header: `<p class="degree-title">` with degree
  - Below degree: `<p class="education-status">` with status (if exists)
- [ ] **3.8** Implement Section 7 (References) `.references-section`:
  - `<h2 class="section-header">REFERENCES</h2>`
  - Loop through `recommendations` (up to 3) to create `.recommendation-card` blocks
  - Each card: `<blockquote class="recommendation-text">{text}</blockquote>`
  - Note: `<footer class="recommendation-author">` is created by Puppeteer, not in template
  - Add fallback: if no recommenders have data (or after Puppeteer injection check), show `<p class="references-fallback">References available on request.</p>`
- [ ] **3.9** Implement Optional Footer Note `.pdf-note` (empty div, Puppeteer injects content if `PRIVATE_CV_NOTE` exists)

---

## Phase 4: Inline CSS Implementation

- [ ] **4.1** Implement `@page` rule in `@media print`:
  - `size: A4`
  - `margin: 0mm` (Puppeteer adds 15mm margins via PDF options)
- [ ] **4.2** Implement body styles in `@media print`:
  - `font-family: 'Times New Roman', Times, serif`
  - `font-size: 11pt`
  - `line-height: 1.3`
  - `color: #000000`
  - `background-color: #ffffff`
  - `margin: 0; padding: 0`
- [ ] **4.3** Implement `.cv-container` styles in `@media print`:
  - `width: 100%; max-width: 100%`
  - `margin: 0; padding: 0`
  - `background: #ffffff`
- [ ] **4.4** Implement section header styles (`h2.section-header`) in `@media print`:
  - `font-size: 12pt; font-weight: bold`
  - `text-transform: uppercase`
  - `margin: 16pt 0 8pt 0`
  - `border-bottom: none; letter-spacing: 0.5pt`
- [ ] **4.5** Implement entry header styles (`.entry-header`, `h3.company-name`, `h3.project-name`, `h3.institution-name`, `.period`) in `@media print`:
  - `.entry-header`: `display: flex; justify-content: space-between; align-items: baseline; margin-bottom: 4pt`
  - `h3.*`: `font-size: 11pt; font-weight: bold; margin: 0; padding: 0`
  - `.period`: `font-size: 10pt; font-weight: normal; margin: 0; padding: 0`
- [ ] **4.6** Implement role and degree title styles (`.role-title`, `.degree-title`) in `@media print`:
  - `font-size: 11pt; font-weight: normal`
  - `margin: 0 0 6pt 0; padding: 0`
- [ ] **4.7** Implement entry spacing styles (`.experience-entry`, `.project-entry`, `.education-entry`) in `@media print`:
  - `margin-bottom: 12pt`
  - `page-break-inside: avoid`
- [ ] **4.8** Implement bullet styles (`.entry-bullets`, `.entry-bullets li`) in `@media print`:
  - `.entry-bullets`: `margin: 6pt 0 6pt 0; padding-left: 18pt; list-style-type: disc`
  - `.entry-bullets li`: `font-size: 11pt; line-height: 1.3; margin-bottom: 4pt; padding-left: 2pt`
- [ ] **4.9** Implement technologies list styles (`.technologies`, `.project-tech`) in `@media print`:
  - `font-size: 10pt; font-style: italic`
  - `margin: 6pt 0 0 0; padding: 0`
- [ ] **4.10** Implement skills section styles (`.skill-category`, `h3.skill-category-header`, `.skill-list`) in `@media print`:
  - `.skill-category`: `margin-bottom: 10pt; page-break-inside: avoid`
  - `h3.skill-category-header`: `font-size: 11pt; font-weight: bold; margin: 0 0 4pt 0; padding: 0`
  - `.skill-list`: `font-size: 11pt; line-height: 1.3; margin: 0; padding: 0`
- [ ] **4.11** Implement summary text styles (`.summary-text`) in `@media print`:
  - `font-size: 11pt; line-height: 1.4`
  - `margin: 0; padding: 0`
  - `text-align: justify`
- [ ] **4.12** Implement references section styles (`.recommendation-card`, `.recommendation-text`, `.recommendation-author`, `.author-name`, `.author-role`, `.author-relation`, `.references-fallback`) in `@media print`:
  - `.recommendation-card`: `margin-bottom: 12pt; padding: 0; border: none; background: #ffffff; page-break-inside: avoid`
  - `.recommendation-text`: `font-size: 11pt; line-height: 1.4; margin: 0 0 8pt 0; padding: 0 0 0 18pt; border-left: 1px solid #000000; quotes: "“" "”" "‘" "’"`
  - `.recommendation-text::before`: `content: open-quote; font-size: 12pt; margin-left: -4pt`
  - `.recommendation-text::after`: `content: close-quote`
  - `.recommendation-author`: `font-size: 10pt; margin: 0; padding: 0`
  - `.author-name`: `font-style: italic; font-weight: normal`
  - `.author-role`, `.author-relation`: `font-style: normal`
  - `.references-fallback`: `font-size: 11pt; font-style: italic; margin: 0; padding: 0`
- [ ] **4.13** Implement page break control styles in `@media print`:
  - `section`: `page-break-after: auto`
  - `section:last-child`: `page-break-after: avoid`
  - `.section-header`: `page-break-after: avoid; page-break-inside: avoid`
- [ ] **4.14** Implement PDF note styles (`.pdf-note`) in `@media print`:
  - `font-size: 9pt; font-style: italic`
  - `margin: 20pt 0 0 0; padding: 0`
  - `text-align: center`
- [ ] **4.15** Implement `@media screen` rules for preview:
  - `body`: `font-family: 'Arial', sans-serif; background-color: #f5f5f5; margin: 20px; padding: 20px`
  - `.cv-container`: `background: #ffffff; max-width: 210mm; margin: 0 auto; padding: 15mm; box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1)`
  - Override `h2.section-header` to use Arial font
- [ ] **4.16** Ensure no dark theme variables leak into print output:
  - Add `*, *::before, *::after { box-sizing: border-box; }` in `@media print`

---

## Phase 5: Script Modifications (generate-pdf.js)

- [ ] **5.1** In `web/scripts/generate-pdf.js`, line ~117: change `const baseUrl = "http://localhost:4321/cv-pipeline/";` to `const baseUrl = "http://localhost:4321/cv-pipeline/cv-pdf";`
- [ ] **5.2** In `web/scripts/generate-pdf.js`, line ~49: change selector from `const h1 = document.querySelector(".hero h1");` to `const h1 = document.querySelector(".cv-header h1");`
- [ ] **5.3** In `web/scripts/generate-pdf.js`, line ~158: change `printBackground: true` to `printBackground: false` (ensures white background in PDF)

---

## Phase 6: Build Verification

- [ ] **6.1** Run `cd /src/cv-pipeline/web && pnpm build` and verify build passes without errors
- [ ] **6.2** Verify no TypeScript errors in Astro frontmatter (if applicable)
- [ ] **6.3** Verify CSS is valid and no syntax errors in `<style>` block
- [ ] **6.4** Verify Astro routes are correctly registered (check dist output if needed)

---

## Phase 7: PDF Generation and Manual Testing

- [ ] **7.1** Run `pnpm export-pdf` and verify PDF is generated at `/src/cv-pipeline/Diego_Sasco_CV_Privado.pdf`
- [ ] **7.2** Open PDF and verify it is 1-2 pages total
- [ ] **7.3** Verify PDF page size is A4
- [ ] **7.4** Verify background is pure white (#ffffff) - no dark theme remnants
- [ ] **7.5** Verify all text is selectable (not rendered as images)
- [ ] **7.6** Verify all 7 sections are present and in correct order:
  - Header (name, title)
  - Professional Summary
  - Technical Skills (4 categories)
  - Professional Experience (reverse chronological)
  - Key Projects (if data exists, up to 5)
  - Education (reverse chronological)
  - References
- [ ] **7.7** Verify font appears as Times New Roman or Arial fallback
- [ ] **7.8** Verify font sizes are within 10-12pt range for body text
- [ ] **7.9** Verify margins are approximately 15mm on all sides
- [ ] **7.10** Verify skills are grouped into 4 Harvard-standard categories (not 9 original)
- [ ] **7.11** Verify summary is truncated to 2-3 sentences (approx. 150-200 chars)
- [ ] **7.12** Verify work experience entries are sorted most recent first
- [ ] **7.13** Verify projects are extracted from `work_experience[].projects` (not top-level)
- [ ] **7.14** Verify page breaks do not orphan section headers or split entries awkwardly
- [ ] **7.15** Verify no decorative elements (borders, underlines, shading, text boxes)

---

## Phase 8: Private Data Injection Verification

- [ ] **8.1** Verify `.cv-header h1` contains private name from `PRIVATE_NAME` (not placeholder)
- [ ] **8.2** Verify `.contact-info` contains all injected contact items (email, phone, address, website, GitHub, LinkedIn) with SVG icons
- [ ] **8.3** Verify each `.recommendation-card` has `<footer class="recommendation-author">` with author name, role, and relation
- [ ] **8.4** If `PRIVATE_CV_NOTE` is set, verify `.pdf-note` is appended to main with the note text
- [ ] **8.5** Verify no `.env` variable values are visible in HTML source (only injected after Puppeteer renders)

---

## Phase 9: ATS Compatibility Verification

- [ ] **9.1** Verify all content uses semantic HTML elements (`<h1>`, `<h2>`, `<h3>`, `<p>`, `<ul>`, `<li>`, `<blockquote>`, `<footer>`, `<cite>`)
- [ ] **9.2** Verify no tables or complex layouts used
- [ ] **9.3** Verify no text boxes or shapes (only standard HTML elements)
- [ ] **9.4** Verify no images for text content (all text is selectable)
- [ ] **9.5** Verify no custom Unicode bullets (use standard disc bullets)
- [ ] **9.6** (Optional) Test with online ATS parser if available

---

## Phase 10: Rollback Verification (if needed)

- [ ] **10.1** Verify original web page (`index.astro`) remains unchanged and functional
- [ ] **10.2** Verify `generate-pdf.js` can be reverted to navigate to `/cv-pipeline/` if needed
- [ ] **10.3** Verify deleting `cv-pdf.astro` restores original state without breaking build

---

## Task Count Summary

- **Phase 1 (Template Setup):** 4 tasks
- **Phase 2 (Frontmatter Logic):** 11 tasks
- **Phase 3 (HTML Structure):** 9 tasks
- **Phase 4 (Inline CSS):** 16 tasks
- **Phase 5 (Script Modifications):** 3 tasks
- **Phase 6 (Build Verification):** 4 tasks
- **Phase 7 (PDF Testing):** 15 tasks
- **Phase 8 (Private Data Verification):** 5 tasks
- **Phase 9 (ATS Verification):** 6 tasks
- **Phase 10 (Rollback Verification):** 3 tasks

**Total:** 76 tasks across 10 phases

---

**Next Phase:** SDD Apply — Execute this task list and track progress in `openspec/apply-progress/cv-pdf-harvard.md`
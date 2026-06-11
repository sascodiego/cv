# Harvard-Style CV PDF Template — Detailed Design Document

## Change ID

`cv-pdf-harvard`

## Document Status

**Phase:** SDD Design  
**Status:** Complete  
**Created:** 2025-06-10  

---

## 1. Implementation Overview

This document provides the complete implementation blueprint for creating a dedicated Harvard-style CV PDF template (`web/src/pages/cv-pdf.astro`) that produces a professional, ATS-optimized document from the existing `cv-processed.json` data source.

### 1.1 Key Design Principles

- **ATS-first**: Text-only, selectable content, no images for text
- **Harvard standards**: Serif fonts, 10-12pt body, 14-16pt name, equal margins
- **Single-column**: No grids, tables, or multi-section layouts
- **Zero decorative elements**: No shading, borders, text boxes, or underlining
- **Print-optimized**: Page-break control, white background, proper spacing
- **Puppeteer-injectable**: DOM selectors match private data injection contract

---

## 2. File Structure and Route

### 2.1 New File

**Location:** `web/src/pages/cv-pdf.astro`  
**Route:** `/cv-pipeline/cv-pdf`  
**Type:** Astro page template with inline print CSS

### 2.2 Modified File

**Location:** `web/scripts/generate-pdf.js`  
**Change:** Single URL update (line ~53)  
**Before (line ~138):** `await page.goto(baseUrl, { waitUntil: 'networkidle0' });`  
**After (line ~138):** `await page.goto('http://localhost:4321/cv-pipeline/cv-pdf', { waitUntil: 'networkidle0' });`  

**Note:** Change `const baseUrl = "http://localhost:4321/cv-pipeline/";` to the new URL at line ~117.

---

## 3. Complete HTML Structure

### 3.1 Root Element

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8" />
  <title>CV - Diego Sasco</title>
  <style>
    <!-- Inline CSS goes here -->
  </style>
</head>
<body>
  <div class="cv-container">
    <!-- All sections here -->
  </div>
</body>
</html>
```

### 3.2 Section 1: Header (`.cv-header`)

```html
<header class="cv-header">
  <h1>Diego Sasco</h1>
  <p class="professional-title">Desarrollador de Software Senior</p>
  <div class="contact-info">
    <!-- Puppeteer injects contact items here -->
  </div>
</header>
```

**Selectors for Puppeteer:**
- `.cv-header h1` — Target for name injection
- `.contact-info` — Container cleared and rebuilt by Puppeteer

### 3.3 Section 2: Professional Summary (`.summary-section`)

```html
<section class="summary-section">
  <h2 class="section-header">PROFESSIONAL SUMMARY</h2>
  <p class="summary-text">
    Soy desarrollador de software con más de 7 años de experiencia,
    especializado en construir soluciones backend (C#/.NET, Go) y herramientas
    de automatización para los sectores retail y servicios. Mi experiencia
    fuerte está en la integración de sistemas, la integración de pasarelas de
    pago y la automatización de procesos administrativos críticos mediante
    software de escritorio y arquitecturas distribuidas. De forma complementaria,
    desarrollo soluciones IoT y software embebido en Linux (Raspberry Pi,
    ESP32) para el control de dispositivos en puntos de venta y entornos
    comerciales.
  </p>
</section>
```

**Note:** Truncate to first 2-3 sentences (approx. 150-200 characters) for brevity.

### 3.4 Section 3: Technical Skills (`.skills-section`)

```html
<section class="skills-section">
  <h2 class="section-header">TECHNICAL SKILLS</h2>
  
  <div class="skill-category">
    <h3 class="skill-category-header">Languages & Frameworks</h3>
    <p class="skill-list">
      C# / .NET (7 años y 1 mes), Go (2 años), Python (3 años),
      JavaScript / TypeScript (1 año)
    </p>
  </div>
  
  <div class="skill-category">
    <h3 class="skill-category-header">Databases & Messaging</h3>
    <p class="skill-list">
      SQL (PostgreSQL / SQL Server) (6 años y 10 meses), RabbitMQ (4 años y 2 meses)
    </p>
  </div>
  
  <div class="skill-category">
    <h3 class="skill-category-header">Tools & Platforms</h3>
    <p class="skill-list">
      Docker (2 años), Linux (Ubuntu/Debian) (8 años y 10 meses),
      Playwright / UI Automation (5 años y 2 meses), Astro (1 mes)
    </p>
  </div>
  
  <div class="skill-category">
    <h3 class="skill-category-header">IoT & Embedded</h3>
    <p class="skill-list">
      Fyne UI (Go Desktop) (3 meses), gRPC (2 años)
    </p>
  </div>
</section>
```

**Logic:**
- Group skills by `category` field
- Display `experience_text` next to each skill in parentheses
- Map category names to Harvard-standard headers:
  - "Languages & Frameworks" → "Languages & Frameworks"
  - "Databases & Obs" → "Databases & Messaging"
  - "Architecture & Protocols" → "Databases & Messaging"
  - "Infrastructure & DevOps" → "Tools & Platforms"
  - "Automation & QA" → "Tools & Platforms"
  - "IoT & Embedded" → "IoT & Embedded"
  - "Metodologías & Prácticas" → omit from PDF (competencies, not technical)
  - "Habilidades Funcionales" → omit from PDF (functional skills, not technical)
  - "Competencias Blandas" → omit from PDF (soft skills)

### 3.5 Section 4: Professional Experience (`.experience-section`)

```html
<section class="experience-section">
  <h2 class="section-header">PROFESSIONAL EXPERIENCE</h2>
  
  <article class="experience-entry">
    <div class="entry-header">
      <h3 class="company-name">Desarrollo Independiente & Consultoría</h3>
      <span class="period">Jul 2024 - Presente</span>
    </div>
    <p class="role-title">Desarrollador de Software & IoT (Freelance)</p>
    <ul class="entry-bullets">
      <li>Diseñé e implementé sistemas distribuidos en Go y C#/.NET para
        automatización IoT e integración legacy.</li>
      <li>Colaboré con especialistas en hardware electrónico en el
        prototipado de soluciones embebidas.</li>
      <li>Aplicación sistemática de metodologías de desarrollo asistido por IA
        para maximizar la calidad y velocidad de entrega.</li>
    </ul>
    <p class="technologies">Technologies: golang, csharp, grpc, soap, fyne, linux</p>
  </article>
  
  <!-- Additional work entries in reverse chronological order -->
</section>
```

**Logic:**
- Render `work_experience` in reverse chronological order
- For each entry, display:
  - Company name (bold)
  - Period (right-aligned or on same line as company)
  - Role title (below company)
  - Description split into 1-3 bullet points (action verbs, CAR framework)
  - Technologies as comma-separated list below bullets

### 3.6 Section 5: Key Projects (`.projects-section`)

**Note:** This section is only rendered if projects exist in `work_experience[].projects`. If no projects exist, the entire section is omitted.

```html
<section class="projects-section">
  <h2 class="section-header">KEY PROJECTS</h2>
  
  <article class="project-entry">
    <div class="entry-header">
      <h3 class="project-name">RpPOS - Arquitectura IoT (Golden Wash)</h3>
      <span class="period">Abr 2026 - Presente</span>
    </div>
    <p class="project-tech">golang, fyne, linux</p>
    <ul class="entry-bullets">
      <li>Diseñé e implementé la interfaz de usuario y el control de
        hardware (GPIO) en Go con Fyne UI para un Punto de Venta embebido
        en Raspberry Pi.</li>
    </ul>
  </article>
  
  <!-- Additional projects (up to 5 total) -->
</section>
```

**Logic:**
- Extract projects from `work_experience[].projects` (not top-level `projects` array)
- Select up to 5 projects (prioritize most recent by `start_date`)
- Display:
  - Project name (bold)
  - Period (right-aligned)
  - Tech stack (comma-separated, smaller font)
  - Description as 1-2 bullets
- Omit section if no projects exist

### 3.7 Section 6: Education (`.education-section`)

```html
<section class="education-section">
  <h2 class="section-header">EDUCATION</h2>
  
  <article class="education-entry">
    <div class="entry-header">
      <h3 class="institution-name">Universidad ORT Uruguay</h3>
      <span class="period">Mar 2019 - Dic 2020</span>
    </div>
    <p class="degree-title">Analista en Infraestructura Informática</p>
    <p class="education-status">Incompleto</p>
  </article>
  
  <!-- Additional education entries in reverse chronological order -->
</section>
```

**Logic:**
- Render `education` in reverse chronological order
- Display:
  - Institution name (bold)
  - Period (right-aligned)
  - Degree title (below institution)
  - Status (below degree)

### 3.8 Section 7: References (`.references-section`)

```html
<section class="references-section">
  <h2 class="section-header">REFERENCES</h2>
  
  <div class="recommendation-card">
    <blockquote class="recommendation-text">
      He tenido el gusto de trabajar con Diego, es un excelente compañero
      como profesional.
    </blockquote>
    <footer class="recommendation-author">
      <!-- Puppeteer injects recommender details here -->
    </footer>
  </div>
  
  <!-- Additional recommendation cards (up to 3) -->
</section>
```

**Selectors for Puppeteer:**
- `.recommendation-card` — Target container for recommender attribution injection (function creates `<footer>` inside)

**Puppeteer injection structure:**
```html
<footer class="recommendation-author">
  <cite class="author-name">Juan Pérez</cite>
  <span class="author-role">CTO at Tech Corp</span>
  <span class="author-relation">Former Manager</span>
</footer>
```

**Fallback:**
- If any `.env` recommender variable is missing, display:
  ```html
  <p class="references-fallback">References available on request.</p>
  ```
  And omit the `<footer>` injection.

### 3.9 Optional: PDF Note (`.pdf-note`)

```html
<div class="pdf-note">
  <!-- Optional footer note for versioning or additional info -->
</div>
```

**Selector for Puppeteer:**
- `.pdf-note` — Optional injection point for footer note

---

## 4. Complete Inline CSS

### 4.1 Base Print CSS Rules

```css
<style>
  @media print {
    @page {
      size: A4;
      margin: 0mm; /* Puppeteer handles 15mm margins via PDF options */
    }

    body {
      font-family: 'Times New Roman', Times, serif;
      font-size: 11pt;
      line-height: 1.3;
      color: #000000;
      background-color: #ffffff;
      margin: 0;
      padding: 0;
      -webkit-print-color-adjust: exact;
      print-color-adjust: exact;
    }

    .cv-container {
      width: 100%;
      max-width: 100%;
      margin: 0;
      padding: 0;
      background: #ffffff;
    }

    /* Section headers - bold, uppercase, no underline */
    h2.section-header {
      font-family: 'Times New Roman', Times, serif;
      font-size: 12pt;
      font-weight: bold;
      text-transform: uppercase;
      color: #000000;
      margin: 16pt 0 8pt 0;
      padding: 0;
      border-bottom: none;
      letter-spacing: 0.5pt;
    }

    /* Entry headers */
    .entry-header {
      display: flex;
      justify-content: space-between;
      align-items: baseline;
      margin-bottom: 4pt;
    }

    h3.company-name,
    h3.project-name,
    h3.institution-name {
      font-family: 'Times New Roman', Times, serif;
      font-size: 11pt;
      font-weight: bold;
      color: #000000;
      margin: 0;
      padding: 0;
    }

    .period {
      font-family: 'Times New Roman', Times, serif;
      font-size: 10pt;
      font-weight: normal;
      color: #000000;
      margin: 0;
      padding: 0;
    }

    /* Role and degree titles */
    .role-title,
    .degree-title {
      font-family: 'Times New Roman', Times, serif;
      font-size: 11pt;
      font-weight: normal;
      color: #000000;
      margin: 0 0 6pt 0;
      padding: 0;
    }

    /* Entry spacing */
    .experience-entry,
    .project-entry,
    .education-entry {
      margin-bottom: 12pt;
      page-break-inside: avoid;
    }

    /* Bullets */
    .entry-bullets {
      margin: 6pt 0 6pt 0;
      padding-left: 18pt;
      list-style-type: disc;
      color: #000000;
    }

    .entry-bullets li {
      font-family: 'Times New Roman', Times, serif;
      font-size: 11pt;
      line-height: 1.3;
      color: #000000;
      margin-bottom: 4pt;
      padding-left: 2pt;
    }

    /* Technologies list */
    .technologies,
    .project-tech {
      font-family: 'Times New Roman', Times, serif;
      font-size: 10pt;
      font-style: italic;
      color: #000000;
      margin: 6pt 0 0 0;
      padding: 0;
    }

    /* Skills section */
    .skill-category {
      margin-bottom: 10pt;
      page-break-inside: avoid;
    }

    h3.skill-category-header {
      font-family: 'Times New Roman', Times, serif;
      font-size: 11pt;
      font-weight: bold;
      color: #000000;
      margin: 0 0 4pt 0;
      padding: 0;
    }

    .skill-list {
      font-family: 'Times New Roman', Times, serif;
      font-size: 11pt;
      line-height: 1.3;
      color: #000000;
      margin: 0;
      padding: 0;
    }

    /* Summary text */
    .summary-text {
      font-family: 'Times New Roman', Times, serif;
      font-size: 11pt;
      line-height: 1.4;
      color: #000000;
      margin: 0 0 0 0;
      padding: 0;
      text-align: justify;
    }

    /* References section */
    .recommendation-card {
      margin-bottom: 12pt;
      padding: 0;
      border: none;
      background: #ffffff;
      page-break-inside: avoid;
    }

    .recommendation-text {
      font-family: 'Times New Roman', Times, serif;
      font-size: 11pt;
      line-height: 1.4;
      color: #000000;
      margin: 0 0 8pt 0;
      padding: 0 0 0 18pt;
      border-left: 1px solid #000000;
      quotes: "“" "”" "‘" "’";
    }

    .recommendation-text::before {
      content: open-quote;
      font-size: 12pt;
      margin-left: -4pt;
    }

    .recommendation-text::after {
      content: close-quote;
    }

    .recommendation-author {
      font-family: 'Times New Roman', Times, serif;
      font-size: 10pt;
      color: #000000;
      margin: 0;
      padding: 0;
    }

    .author-name {
      font-style: italic;
      font-weight: normal;
    }

    .author-role {
      font-style: normal;
    }

    .author-relation {
      font-style: normal;
    }

    .references-fallback {
      font-family: 'Times New Roman', Times, serif;
      font-size: 11pt;
      color: #000000;
      margin: 0;
      padding: 0;
      font-style: italic;
    }

    /* Page break control */
    section {
      page-break-after: auto;
    }

    section:last-child {
      page-break-after: avoid;
    }

    .section-header {
      page-break-after: avoid;
      page-break-inside: avoid;
    }

    /* PDF note (optional) */
    .pdf-note {
      font-family: 'Times New Roman', Times, serif;
      font-size: 9pt;
      color: #000000;
      margin: 20pt 0 0 0;
      padding: 0;
      text-align: center;
      font-style: italic;
    }

    /* Ensure no dark theme remnants */
    *,
    *::before,
    *::after {
      box-sizing: border-box;
    }
  }

  /* Screen styles (for preview, not PDF) */
  @media screen {
    body {
      font-family: 'Arial', sans-serif;
      font-size: 11pt;
      line-height: 1.3;
      color: #000000;
      background-color: #f5f5f5;
      margin: 20px;
      padding: 20px;
    }

    .cv-container {
      background: #ffffff;
      max-width: 210mm;
      margin: 0 auto;
      padding: 15mm;
      box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
    }

    h2.section-header {
      font-family: 'Arial', sans-serif;
      font-size: 12pt;
      font-weight: bold;
      text-transform: uppercase;
      color: #000000;
      margin: 16pt 0 8pt 0;
      padding: 0;
      border-bottom: none;
    }
  }
</style>
```

### 4.2 CSS Design Decisions

1. **Font Stack**: Times New Roman (serif) for ATS compatibility, Arial for screen preview
2. **Font Sizes**: 11pt body, 10-12pt headers, 9pt footer notes
3. **Colors**: Pure black (#000000) on pure white (#ffffff)
4. **Spacing**: 15mm margins, 12-16pt section spacing, 8-12pt entry spacing
5. **Page Breaks**: `page-break-inside: avoid` on entries, `page-break-after: avoid` on headers
6. **Bullets**: Standard disc bullets, no custom Unicode
7. **No Decorations**: No borders, underlines, shading, or text boxes
8. **Print-Only Rules**: `@media print` enforces white background and black text
9. **Screen Preview**: Separate `@media screen` rules for development/preview

---

## 5. Astro Frontmatter Implementation

### 5.1 Data Import

```astro
---
import cvData from '../data/cv-processed.json';

// Access top-level data
const { personal_info, skills, work_experience, education, recommendations } = cvData;
---
```

### 5.2 Skill Grouping Logic

```astro
---
// Define category mapping to Harvard-standard headers
const categoryMapping = {
  'Languages & Frameworks': 'Languages & Frameworks',
  'Databases & Obs': 'Databases & Messaging',
  'Architecture & Protocols': 'Databases & Messaging',
  'Infrastructure & DevOps': 'Tools & Platforms',
  'Automation & QA': 'Tools & Platforms',
  'IoT & Embedded': 'IoT & Embedded',
  'Metodologías & Prácticas': null, // omit from PDF
  'Habilidades Funcionales': null,   // omit from PDF
  'Competencias Blandas': null       // omit from PDF
};

// Group skills by mapped category
const skillGroups = skills.reduce((groups, skill) => {
  const mappedCategory = categoryMapping[skill.category];
  if (mappedCategory) {
    if (!groups[mappedCategory]) {
      groups[mappedCategory] = [];
    }
    groups[mappedCategory].push(skill);
  }
  return groups;
}, {});

// Get sorted category keys
const sortedCategories = Object.keys(skillGroups).sort();
---
```

### 5.3 Project Extraction Logic

```astro
---
// Extract all projects from work_experience entries
const allProjects = work_experience
  .filter(entry => entry.projects && entry.projects.length > 0)
  .flatMap(entry => 
    entry.projects.map(project => ({
      ...project,
      company: entry.company, // Track source company for context
      entryPeriod: entry.period_text
    }))
  );

// Sort projects by start_date (most recent first)
const sortedProjects = allProjects.sort((a, b) => 
  new Date(b.start_date) - new Date(a.start_date)
);

// Select up to 5 projects
const selectedProjects = sortedProjects.slice(0, 5);
---
```

### 5.4 Summary Truncation Logic

```astro
---
// Truncate summary to first 2-3 sentences (approx. 200 chars)
const truncateSummary = (text, maxLength = 200) => {
  if (text.length <= maxLength) return text;
  
  // Find last sentence boundary within limit
  const truncated = text.substring(0, maxLength);
  const lastPeriod = truncated.lastIndexOf('.');
  const lastSentenceEnd = lastPeriod > 0 ? lastPeriod + 1 : maxLength;
  
  return text.substring(0, lastSentenceEnd).trim();
};

const summaryText = truncateSummary(personal_info.about_me, 200);
---
```

### 5.5 Work Experience Reverse Chronological Sort

```astro
---
// Sort work_experience by end_date (most recent first)
const sortedWorkExperience = [...work_experience].sort((a, b) => {
  // Handle "Presente" or future dates
  const aEnd = a.end_date === '2026-06-10' ? '2999-12-31' : a.end_date;
  const bEnd = b.end_date === '2026-06-10' ? '2999-12-31' : b.end_date;
  
  return new Date(bEnd) - new Date(aEnd);
});
---
```

### 5.6 Education Reverse Chronological Sort

```astro
---
// Sort education by end_date (most recent first)
const sortedEducation = [...education].sort((a, b) => 
  new Date(b.end_date) - new Date(a.end_date)
);
---
```

### 5.7 Bullet Point Splitting Logic

```astro
---
// Split description into 1-3 bullet points
const splitDescriptionToBullets = (description, maxBullets = 3) => {
  // Split by sentence boundaries (period + space)
  const sentences = description.match(/[^.!?]+[.!?]+/g) || [description];
  
  // Take up to maxBullets sentences
  return sentences.slice(0, maxBullets).map(s => s.trim());
};
---
```

---

## 6. Complete Template Wireframe

### 6.1 Pseudo-HTML Hierarchy

```
<html>
  <head>
    <style>
      @media print { /* 15mm margins, Times New Roman, 11pt */ }
      @media screen { /* Arial, preview styles */ }
    </style>
  </head>
  <body>
    <div class="cv-container">
      <!-- Section 1: Header -->
      <header class="cv-header">
        <h1>Diego Sasco</h1>
        <p class="professional-title">Desarrollador de Software Senior</p>
        <div class="contact-info">
          <!-- Puppeteer injects: Email, Phone, Address, Website, GitHub, LinkedIn -->
        </div>
      </header>

      <!-- Section 2: Summary -->
      <section class="summary-section">
        <h2 class="section-header">PROFESSIONAL SUMMARY</h2>
        <p class="summary-text">
          Soy desarrollador de software con más de 7 años de experiencia...
        </p>
      </section>

      <!-- Section 3: Skills -->
      <section class="skills-section">
        <h2 class="section-header">TECHNICAL SKILLS</h2>
        <div class="skill-category">
          <h3 class="skill-category-header">Languages & Frameworks</h3>
          <p class="skill-list">C# / .NET (7 años), Go (2 años), Python (3 años)...</p>
        </div>
        <div class="skill-category">
          <h3 class="skill-category-header">Databases & Messaging</h3>
          <p class="skill-list">SQL (6 años), RabbitMQ (4 años)...</p>
        </div>
        <div class="skill-category">
          <h3 class="skill-category-header">Tools & Platforms</h3>
          <p class="skill-list">Docker (2 años), Linux (8 años)...</p>
        </div>
        <div class="skill-category">
          <h3 class="skill-category-header">IoT & Embedded</h3>
          <p class="skill-list">gRPC (2 años), Fyne UI (3 meses)...</p>
        </div>
      </section>

      <!-- Section 4: Experience -->
      <section class="experience-section">
        <h2 class="section-header">PROFESSIONAL EXPERIENCE</h2>
        <article class="experience-entry">
          <div class="entry-header">
            <h3 class="company-name">Desarrollo Independiente & Consultoría</h3>
            <span class="period">Jul 2024 - Presente</span>
          </div>
          <p class="role-title">Desarrollador de Software & IoT (Freelance)</p>
          <ul class="entry-bullets">
            <li>Diseñé e implementé sistemas distribuidos en Go...</li>
            <li>Colaboré con especialistas en hardware...</li>
            <li>Aplicación sistemática de metodologías...</li>
          </ul>
          <p class="technologies">Technologies: golang, csharp, grpc, soap, fyne, linux</p>
        </article>
        <!-- More experience entries -->
      </section>

      <!-- Section 5: Projects -->
      <section class="projects-section">
        <h2 class="section-header">KEY PROJECTS</h2>
        <article class="project-entry">
          <div class="entry-header">
            <h3 class="project-name">RpPOS - Arquitectura IoT (Golden Wash)</h3>
            <span class="period">Abr 2026 - Presente</span>
          </div>
          <p class="project-tech">golang, fyne, linux</p>
          <ul class="entry-bullets">
            <li>Diseñé e implementé la interfaz de usuario...</li>
          </ul>
        </article>
        <!-- More projects (up to 5) -->
      </section>

      <!-- Section 6: Education -->
      <section class="education-section">
        <h2 class="section-header">EDUCATION</h2>
        <article class="education-entry">
          <div class="entry-header">
            <h3 class="institution-name">Universidad ORT Uruguay</h3>
            <span class="period">Mar 2019 - Dic 2020</span>
          </div>
          <p class="degree-title">Analista en Infraestructura Informática</p>
          <p class="education-status">Incompleto</p>
        </article>
        <!-- More education entries -->
      </section>

      <!-- Section 7: References -->
      <section class="references-section">
        <h2 class="section-header">REFERENCES</h2>
        <div class="recommendation-card">
          <blockquote class="recommendation-text">
            He tenido el gusto de trabajar con Diego...
          </blockquote>
          <footer class="recommendation-author">
            <!-- Puppeteer injects: Juan Pérez, CTO, Former Manager -->
          </footer>
        </div>
        <!-- More recommendation cards (up to 3) -->
      </section>

      <!-- Optional: PDF Note -->
      <div class="pdf-note">
        <!-- Optional footer note -->
      </div>
    </div>
  </body>
</html>
```

---

## 7. Puppeteer Script Modifications

### 7.1 File: `web/scripts/generate-pdf.js`

**Location:** `web/scripts/generate-pdf.js`  
**Line to change:** ~53 (navigation URL)

#### Change 1: URL Update

```javascript
// Before:
await page.goto('http://localhost:4321/cv-pipeline/');

// After:
await page.goto('http://localhost:4321/cv-pipeline/cv-pdf');
```

#### Change 2: Print Background (Verify)

```javascript
// Ensure this option is set to false (line ~158):
await page.pdf({
  path: pdfPath,
  format: 'A4',
  margin: {
    top: '15mm',
    bottom: '15mm',
    left: '15mm',
    right: '15mm',
  },
  printBackground: false, // ← Must be false for white background (currently true)
});
```

**Note:** The margins in PDF options are `15mm` on all sides (lines ~152-156). Since Puppeteer adds its margins **in addition** to CSS margins, the CSS `@page` rule should use `0mm` margins to avoid double-margins.

### 7.2 Selector Updates in `injectPrivateData()` Function

The existing `injectPrivateData()` function in `generate-pdf.js` needs to work with new selectors:

#### Current Selectors (Web Page)
- Name injection: `.hero h1` (line ~49)
- Contact info: `.contact-info` (line ~54)
- Recommender: `.recommendation-card` (line ~158) - function creates `<footer>` inside

#### New Selectors (PDF Template)
- Name injection: `.cv-header h1` ← Updated (change line ~49)
- Contact info: `.contact-info` ← Same selector, different HTML structure (no changes needed)
- Recommender: `.recommendation-card` ← Same selector, different HTML structure (no changes needed)

#### Injection Logic (Unchanged)

The existing `injectPrivateData()` function logic remains the same with selector update:

**Line ~49** (Name injection - change selector):
```javascript
const h1 = document.querySelector(".cv-header h1"); // Was: ".hero h1"
if (h1) h1.textContent = data.name;
```

**Lines ~54-56** (Contact injection - no change needed):
```javascript
const contactInfo = document.querySelector(".contact-info");
if (!contactInfo) return;
// Clears innerHTML and rebuilds with contact items (email, phone, address, website, github, linkedin)
```

**Lines ~158-185** (Recommender injection - no change needed):
```javascript
const quotes = document.querySelectorAll(".recommendation-card");
quotes.forEach((quote, i) => {
  const rec = data.recommenders[i];
  if (!rec || !rec.name) return;
  const footer = document.createElement("footer");
  footer.className = "recommendation-author";
  // Creates <cite class="author-name"> with rec.name
  // Creates <span class="author-role"> with rec.role
  // Creates <span class="author-relation"> with rec.relation
  quote.appendChild(footer);
});
```

---

## 8. DOM Selector Reference

### 8.1 Required Selectors for Puppeteer

| Selector | Purpose | Element | Content |
|----------|---------|---------|---------|
| `.cv-header h1` | Name injection | `<h1>` | Full name from `PRIVATE_NAME` |
| `.contact-info` | Contact injection | `<div>` | Container for contact items (cleared and rebuilt) |
| `.recommendation-card` | Recommender injection | `<div>` | Function creates `<footer>` inside with author name, role, relation |
| `.pdf-note` | Optional footer note | `<div>` | Versioning or additional info |

### 8.2 Selector Existence Requirements

All selectors must exist in the DOM even if initially empty or containing placeholder content. Puppeteer will:
- Update `.cv-header h1` textContent with private name
- Clear `.contact-info` innerHTML and rebuild with contact items (each with SVG icon + text)
- Create `<footer class="recommendation-author">` inside each `.recommendation-card` with author name, role, relation
- Create `<footer class="pdf-note">` if `PRIVATE_CV_NOTE` is present and append to `<main>`

---

## 9. Data Flow Diagram

```
cv-source.yaml
    ↓ (Go pipeline)
cv-processed.json (web/src/data/)
    ↓ (shared data source)
    ├─────────────────┬─────────────────┐
    ↓                 ↓                 ↓
index.astro     cv-pdf.astro     generate-pdf.js
(web page)      (PDF template)   (Puppeteer script)
    │                 │                 │
    │                 │                 ├─ Navigate to /cv-pipeline/cv-pdf
    │                 │                 ├─ Inject PRIVATE_NAME → .cv-header h1
    │                 │                 ├─ Inject contact data → .contact-info
    │                 │                 ├─ Inject recommenders → .recommendation-card (creates footer inside)
    │                 │                 └─ Generate PDF with print CSS
    │                 │
    │                 └─ Render with @media print rules
    │                   (white background, Times New Roman, 15mm margins)
    │
    └─ Render with dark zinc theme (unchanged)
```

---

## 10. Implementation Checklist

### 10.1 New File Creation

- [ ] Create `web/src/pages/cv-pdf.astro`
- [ ] Implement HTML structure with all 7 sections
- [ ] Add inline `<style>` with `@media print` rules
- [ ] Add `@media screen` rules for preview
- [ ] Ensure all DOM selectors exist: `.cv-header h1`, `.contact-info`, `.recommendation-card`, `.pdf-note`

### 10.2 Astro Frontmatter Logic

- [ ] Import `cv-processed.json`
- [ ] Implement skill category mapping and grouping
- [ ] Implement project extraction from `work_experience[].projects`
- [ ] Implement summary truncation (2-3 sentences)
- [ ] Implement reverse chronological sorting for work experience
- [ ] Implement reverse chronological sorting for education
- [ ] Implement bullet point splitting for descriptions

### 10.3 CSS Implementation

- [ ] Set `@page { size: A4; margin: 15mm; }`
- [ ] Enforce white background: `body { background-color: #ffffff; color: #000000; }`
- [ ] Set font stack: `'Times New Roman', Times, serif`
- [ ] Set font sizes: 11pt body, 12pt headers, 10pt contact
- [ ] Add page-break control: `page-break-inside: avoid` on entries
- [ ] Remove all decorative elements: no borders, underlines, shading
- [ ] Ensure no dark theme variables or colors

### 10.4 Script Modifications

- [ ] Update navigation URL in `generate-pdf.js` to `/cv-pipeline/cv-pdf` (line ~117: baseUrl, line ~138: page.goto)
- [ ] Update name selector from `.hero h1` to `.cv-header h1` (line ~49)
- [ ] Change `printBackground: true` to `printBackground: false` (line ~158)
- [ ] Keep Puppeteer margins at `15mm` (lines ~152-156) - CSS `@page` uses `0mm`
- [ ] Test `injectPrivateData()` with new selectors

### 10.5 Testing

- [ ] Run `pnpm build` — should pass without errors
- [ ] Run `pnpm export-pdf` — should generate PDF at `/src/cv-pipeline/Diego_Sasco_CV_Privado.pdf`
- [ ] Verify PDF is 1-2 pages, A4, white background
- [ ] Verify no dark theme remnants in PDF
- [ ] Verify all 7 sections present and correctly ordered
- [ ] Verify private data (contact, recommenders) injected correctly
- [ ] Verify text is selectable (not images)
- [ ] Verify page breaks don't orphan sections
- [ ] Test ATS parsing with sample tools

---

## 11. Design Decisions and Rationale

### 11.1 Font Stack Decision

**Decision:** Prioritize Times New Roman (serif) over Arial (sans-serif)

**Rationale:**
- Harvard GSAS/FAS standards traditionally use serif fonts
- ATS systems historically parse serif fonts more reliably
- Times New Roman is widely available across all platforms
- Fallback to Arial ensures compatibility if Times New Roman is unavailable

### 11.2 Margin Configuration

**Decision:** 0mm CSS margins + 15mm Puppeteer margins

**Rationale:**
- 15mm = 0.59 inch, exceeding the minimum 0.5 inch requirement
- Puppeteer adds its margins in addition to CSS margins
- Keeping Puppeteer margins at 15mm and CSS at 0mm prevents double-margins
- Matches existing implementation pattern in `generate-pdf.js`

### 11.3 Skill Category Mapping

**Decision:** Map original categories to Harvard-standard headers, omit non-technical categories

**Rationale:**
- "Databases & Obs" → "Databases & Messaging" (standard naming)
- "Architecture & Protocols" → "Databases & Messaging" (consolidated category)
- "Infrastructure & DevOps" + "Automation & QA" → "Tools & Platforms" (consolidated)
- "Metodologías & Prácticas", "Habilidades Funcionales", "Competencias Blandas" → omit (not technical skills)
- Produces 4 categories instead of 9, better for ATS parsing

### 11.4 Project Selection Strategy

**Decision:** Extract from `work_experience[].projects`, not top-level `projects` array

**Rationale:**
- Top-level `projects` array is empty in current data
- Nested projects within work entries are more contextually relevant
- Sorting by `start_date` ensures most recent projects first
- Limiting to 5 projects prevents page overflow

### 11.5 Summary Truncation

**Decision:** Truncate to first 2-3 sentences (~200 characters)

**Rationale:**
- Harvard-style CVs use concise summaries (2-3 lines)
- Prevents summary from dominating the page
- Sentence-boundary truncation maintains readability
- Action verbs and key accomplishments remain visible

### 11.6 Private Data Injection Strategy

**Decision:** Continue using Puppeteer DOM manipulation (not server-side rendering)

**Rationale:**
- Private data remains in `.env`, never exposed to Git or web build
- No changes to build process or deployment pipeline
- Existing `injectPrivateData()` function works with minimal selector updates
- Maintains security posture: data only exists during PDF generation

### 11.7 Inline CSS Strategy

**Decision:** Embed all CSS in `<style>` tag (no external stylesheet)

**Rationale:**
- Ensures portability: template is self-contained
- Avoids conflicts with global web styles
- `@media print` rules only apply to PDF rendering
- Preview (`@media screen`) uses separate rules

### 11.8 Page-Break Control

**Decision:** `page-break-inside: avoid` on entries, `page-break-after: avoid` on headers

**Rationale:**
- Prevents orphaned bullets or incomplete entries
- Keeps section headers with their content
- Logical break points between major sections
- Improves ATS parsing (entries not split across pages)

---

## 12. Risks and Mitigations

### 12.1 Risk: Font Rendering Variability

**Description:** System fonts may render differently across Windows, macOS, and Linux

**Mitigation:**
- Prioritize Times New Roman (widely available)
- Provide Arial fallback (universally available)
- Test PDF generation on all three platforms
- Consider font embedding if variability is severe

### 12.2 Risk: Page Overflow

**Description:** Dense content (21 skills, 6 work entries, 4 education, 3+ projects) may exceed 2 pages

**Mitigation:**
- Compress skill categories from 9 to 4
- Truncate summary to 2-3 sentences
- Limit bullets to 1-3 per role
- Limit projects to 5
- If still overflowing, reduce font size to 10pt

### 12.3 Risk: Private Data Missing

**Description:** Incomplete `.env` variables may show placeholders instead of contact info

**Mitigation:**
- `injectPrivateData()` checks for missing data before injection
- References section shows "References available on request" if any data missing
- Document `.env` variable requirements in README

### 12.4 Risk: ATS Parsing Issues

**Description:** ATS systems may not parse content correctly despite ATS-friendly design

**Mitigation:**
- Test with common ATS parsers (e.g., Jobvite, Taleo)
- Ensure text is selectable (not images)
- Use standard HTML elements (`<ul>`, `<li>`, `<h3>`)
- Avoid tables, text boxes, or decorative elements

### 12.5 Risk: Print CSS Conflicts

**Description:** Global CSS from `global.css` may interfere with print-specific rules

**Mitigation:**
- Inline CSS in `cv-pdf.astro` (not external stylesheet)
- Use specific class names (`.cv-header`, `.experience-entry`) not used in web styles
- Set `!important` on critical print rules if needed
- Verify no dark theme variables leak into print output

---

## 13. Rollout Plan

### 13.1 Implementation Order

1. Create `web/src/pages/cv-pdf.astro` with basic HTML structure
2. Implement inline CSS with `@media print` rules
3. Add Astro frontmatter with data import and logic
4. Update `generate-pdf.js` URL to `/cv-pipeline/cv-pdf`
5. Test build: `pnpm build`
6. Test PDF generation: `pnpm export-pdf`
7. Verify PDF meets all formatting standards
8. Test ATS compatibility with sample parsing tools

### 13.2 Rollback Plan

If issues arise:

1. Revert `generate-pdf.js` to navigate back to `/cv-pipeline/`
2. Delete `web/src/pages/cv-pdf.astro`
3. Web page and existing functionality remain untouched
4. No data migration or schema changes needed

### 13.3 Success Criteria

- [ ] Build passes: `pnpm build` succeeds
- [ ] PDF generates: `pnpm export-pdf` produces valid PDF
- [ ] PDF is 1-2 pages, A4, white background
- [ ] No dark theme remnants in PDF
- [ ] All 7 sections present and correctly ordered
- [ ] Times New Roman or Arial, 10-12pt body text
- [ ] 15mm margins on all sides
- [ ] Private data injected correctly from `.env`
- [ ] Text is selectable, not images
- [ ] Page breaks don't orphan sections
- [ ] ATS systems can parse content reliably

---

## 14. Appendix: DOM Selector Validation

### 14.1 Selector Existence Test

Run this in browser console on `/cv-pipeline/cv-pdf` to validate selectors:

```javascript
const selectors = [
  '.cv-header h1',
  '.contact-info',
  '.recommendation-card',
  '.pdf-note'
];

selectors.forEach(selector => {
  const element = document.querySelector(selector);
  console.log(selector, element ? '✓ EXISTS' : '✗ MISSING');
});
```

Expected output:
```
.cv-header h1 ✓ EXISTS
.contact-info ✓ EXISTS
.recommendation-card ✓ EXISTS
.pdf-note ✓ EXISTS (optional - function creates if missing)
```

### 14.2 Puppeteer Injection Test

Run this in browser console to validate injection targets:

```javascript
// Test name injection
document.querySelector('.cv-header h1').textContent = 'Test Name';

// Test contact info clearing
document.querySelector('.contact-info').innerHTML = '';

// Test recommender footer creation (function creates new footer element)
const card = document.querySelector('.recommendation-card');
const footer = document.createElement('footer');
footer.className = 'recommendation-author';
const cite = document.createElement('cite');
cite.className = 'author-name';
cite.textContent = 'Juan Pérez';
footer.appendChild(cite);
card.appendChild(footer);

// Test pdf-note creation (function appends to main)
const main = document.querySelector('main');
const pdfNote = document.createElement('footer');
pdfNote.className = 'pdf-note';
pdfNote.innerHTML = '<p>Version 1.0 - 2025-06-10</p>';
main.appendChild(pdfNote);
```

---

## 15. References

- Harvard GSAS CV Guidelines: https://gsas.harvard.edu/academic-careers/cv-and-resume-guide
- Harvard FAS CV Guidelines: https://ocs.fas.harvard.edu/cvs-and-resumes
- ATS Compatibility Best Practices: https://www.jobscan.co/blog/ats-resume-format/
- Puppeteer PDF Generation: https://pptr.dev/api/puppeteer.page.pdf
- Astro Data Import: https://docs.astro.build/en/guides/data-import/

---

**Document End**
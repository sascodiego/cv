# Delta for CV Web Optimization

## ADDED Requirements

### Requirement: REQ-C1

The system MUST include an "About Me" narrative section rendered between the hero and work experience sections, displaying up to 80 words without first-person pronouns.

#### Scenario: About Me Rendering
- GIVEN a user visits the CV web page
- WHEN the page loads the processed CV data
- THEN an "About Me" section MUST be displayed after the hero section and before the work experience section
- AND the content MUST be sourced from the `about_me` field in the YAML source
- AND the text MUST be ≤80 words
- AND the text MUST NOT contain first-person pronouns ("I", "my", "me")
- AND the text MUST use professional identity verbs (e.g., "Desarrollador de Software Senior")

#### Scenario: About Me Content Structure
- GIVEN the About Me section is rendered
- THEN the content MUST reflect Diego Sasco's professional profile
- AND MUST mention sectors: "financiero, retail e industrial"
- AND MUST highlight technical specialties: "integración de sistemas legacy, arquitectura IoT y soluciones RPA"
- AND MUST include key technologies: "C#/.NET, Go, gRPC"
- AND MUST reference security background: "formación en seguridad informática"
- AND MUST end with professional differentiation: "comunicación asertiva y enfoque en la confiabilidad de sistemas críticos"

### Requirement: REQ-C2

The Real2B work experience entry MUST be reframed to emphasize technical support and infrastructure roles, removing low-signal terms and elevating technical competencies.

#### Scenario: Real2B Role Reframing
- GIVEN the CV processes work experience data
- THEN the Real2B role title MUST be changed from "Analista Operacional & Control de Flujo" to "Soporte Técnico y Operaciones de TI"
- AND the description MUST follow the CAR qualitative framework: Context → Action → Result
- AND MUST mention "soporte técnico (Help Desk) a usuarios en múltiples locales"
- AND MUST include "diagnosticando y resolviendo incidencias de infraestructura de TI"
- AND MUST reference "sistemas transaccionales" instead of "transacciones de caja"
- AND MUST include "Cultura de Seguridad" in technologies instead of "Atención al Público"

#### Scenario: Real2B Technology Skills Update
- GIVEN the Real2B entry is processed
- THEN the technologies list MUST include: "Help Desk & Soporte Técnico", "Resolución de Problemas", "Trabajo en Equipo", "Comunicación Asertiva"
- AND MUST exclude "Atención al Público" from the technologies list
- AND MUST include "Cultura de Seguridad" under "Competencias Blandas"

### Requirement: REQ-C3

The PIXI Supermercados work experience entry MUST be reframed to emphasize leadership and operational management, removing low-signal terms and highlighting transferable management skills.

#### Scenario: PIXI Role Reframing
- GIVEN the CV processes work experience data
- THEN the PIXI role title MUST be changed from "Encargado de Operaciones" to "Supervisor de Operaciones y Logística"
- AND the description MUST follow the CAR qualitative framework: Context → Action → Result
- AND MUST mention "Coordiné equipos operativos y supervisé la logística de inventarios"
- AND MUST include "implementando controles sistemáticos de stock"
- AND MUST reference "optimizando los procesos de recepción de mercadería"
- AND MUST mention "Gestión de proveedores y resolución de conflictos operativos"

#### Scenario: PIXI Technology Skills Update
- GIVEN the PIXI entry is processed
- THEN the technologies list MUST include: "Liderazgo de Equipos", "Trabajo en Equipo", "Comunicación Asertiva", "Resolución de Problemas"
- AND MUST exclude "Atención al Público" from the technologies list

### Requirement: REQ-C4

All work experience descriptions and project descriptions MUST be rewritten using the CAR/CCCR qualitative framework with action verbs and impact statements, without numerical metrics.

#### Scenario: CAR Qualitative Framework Application
- GIVEN all experience and project descriptions are processed
- THEN each description MUST follow the structure: Context → Action → Result
- AND MUST use past-tense action verbs: "diseñó", "implementó", "desplegó", "dirigió", "coordinó", "automatizó"
- AND MUST include specific technologies mentioned in the action
- AND MUST conclude with impact statements using qualitative result verbs: "consolidó", "optimizó", "redujo", "estandarizó", "habilitó"
- AND MUST NOT contain numerical metrics percentages or quantities
- AND MUST NOT use passive voice with abstract nouns ("Diseño y desarrollo de...")

#### Scenario: Action Verb Taxonomy
- GIVEN the CAR framework is applied
- THEN development tasks MUST use verbs: "diseñó", "programó", "construyó", "optimizó", "integró", "desplegó"
- THEN leadership tasks MUST use verbs: "dirigió", "coordinó", "encabezó", "administró"
- THEN management tasks MUST use verbs: "automatizó", "implementó", "simplificó", "estandarizó"
- THEN research tasks MUST use verbs: "analizó", "diagnosticó", "validó"

#### Scenario: Impact Statement Qualitative Language
- GIVEN result statements are formulated
- THEN MUST use qualitative impact language: "optimizó significativamente", "simplificó procesos", "garantizó operación continua", "desplegó en producción estable"
- AND MUST connect impact to business context: "para automatización IoT", "para procesos contables", "para entornos de producción"

### Requirement: REQ-C5

The project list MUST be curated to display only 5 flagship projects that maximize technical differentiation signals, with remaining projects hidden or removed from display.

#### Scenario: Project Curation Display
- GIVEN the CV processes project data
- THEN exactly 5 projects MUST be displayed in the rendered output
- AND the visible projects MUST be: RpPOS IoT, Proxy Inverso Financiero, Plataforma de Monitoreo Remoto de Activos IoT, Framework RPA Financiero, Motor de Extracción PDF
- AND projects MUST be ordered by technical differentiation signal: IoT, Legacy Integration, IoT Platform, RPA, PDF Processing
- AND hidden projects MUST NOT appear in any rendered sections

#### Scenario: Project Technical Signal Maximization
- GIVEN the 5 flagship projects are displayed
- THEN RpPOS IoT MUST emphasize "Go + Raspberry Pi + GPIO + Fyne UI + hardware limitado"
- THEN Proxy Inverso Financiero MUST highlight "C# + SOAP→gRPC + pasarela financiera"
- THEN Plataforma de Monitoreo Remoto MUST feature "Go + gRPC + protocolos seriales + IoT"
- THEN Framework RPA MUST showcase "C# + RabbitMQ + UI Automation + observabilidad"
- THEN Motor de Extracción MUST highlight "C# + PdfPig + procesamiento masivo"

### Requirement: REQ-C6

The education entries MUST be contextualized to complement the technical narrative, with the Contador Público entry reformatted to emphasize its complementary nature.

#### Scenario: Education Contextualization
- GIVEN the CV processes education data
- THEN the Contador Público entry MUST be reformatted as: "Contador Público — 1er año cursado (formación complementaria en gestión financiera y contable)"
- AND MUST appear at the end of the education section
- AND the Analista en Infraestructura Informática entry MUST maintain its position
- AND the Hacking Ético certification MUST remain visible under "Graduado" status

#### Scenario: Education Section Ordering
- GIVEN the education section is rendered
- THEN MUST be ordered: Analista en Informática (Incompleto), Ingeniería Sistemas (1er año), Hacking Ético (Graduado), Contador Público (1er año)
- AND the Contador Público entry MUST use the contextualized text format

## MODIFIED Requirements

### Requirement: REQ-C7

The skills grid layout MUST be modified to display in a single column for print output while maintaining responsive grid layout for web display.

#### Scenario: Print Layout Optimization
- GIVEN the CSS media query for `@media print`
- THEN the `.skills-grid` container MUST use `grid-template-columns: 1fr` instead of `grid-template-columns: repeat(3, 1fr)`
- AND the grid layout MUST remain unchanged for screen display
- AND the skills MUST stack vertically in single column for PDF generation
- AND the category titles MUST remain aligned and visible in print

#### Scenario: Web Layout Preservation
- GIVEN the CSS for screen display
- THEN the `.skills-grid` MUST continue using `grid-template-columns: repeat(auto-fill, minmax(250px, 1fr))`
- AND the responsive behavior MUST be preserved for different screen sizes
- AND the visual hierarchy MUST remain consistent with the current design

### Requirement: REQ-C8

The recommendations grid layout MUST be modified to display in a single column for print output while maintaining responsive grid layout for web display.

#### Scenario: Print Recommendations Layout
- GIVEN the CSS media query for `@media print`
- THEN the `.recommendations-grid` container MUST use `grid-template-columns: 1fr` instead of `grid-template-columns: repeat(2, 1fr)`
- AND each recommendation MUST display as a full-width block in print
- AND the recommendation text MUST maintain readability in single column format
- AND the attribution placeholders MUST remain but not be emphasized

#### Scenario: Web Recommendations Layout
- GIVEN the CSS for screen display
- THEN the `.recommendations-grid` MUST continue using `grid-template-columns: repeat(auto-fill, minmax(280px, 1fr))`
- AND the responsive grid behavior MUST be preserved
- AND the visual card layout MUST remain unchanged for web display

### Requirement: REQ-C9

The print stylesheet MUST include 0.75 inch margins for the body and appropriate padding adjustments to ensure proper document formatting.

#### Scenario: Print Margin Application
- GIVEN the CSS media query for `@media print`
- THEN the `body` element MUST have `padding: 0.75in`
- AND main content areas MUST adjust padding to compensate for the increased margins
- AND text content MUST remain readable and properly spaced within the new margins
- AND the overall document MUST conform to standard ATS formatting requirements

#### Scenario: Margin Consistency
- GIVEN the print margins are applied
- THEN all section content MUST be properly contained within the margin boundaries
- AND no content should extend beyond the 0.75 inch margins
- AND the spacing between elements MUST be consistent and readable

### Requirement: REQ-C10

The print stylesheet MUST use ATS-compatible fonts (Georgia for titles, Arial for body) to ensure compatibility with Applicant Tracking Systems.

#### Scenario: ATS Font Application
- GIVEN the CSS media query for `@media print`
- THEN the `:root` variables MUST be overridden for print
- AND `--font-title` MUST be set to "Georgia", serif
- AND `--font-body` MUST be set to "Arial", sans-serif
- AND all title elements MUST use Georgia font in print
- AND all body text MUST use Arial font in print
- AND the font changes MUST only apply to print media, not screen display

#### Scenario: Font Compatibility
- GIVEN the ATS fonts are applied
- THEN the fonts MUST be standard web fonts available in all printing environments
- AND the typography MUST maintain proper hierarchy and readability
- AND the font sizes MUST be optimized for print readability

### Requirement: REQ-C11

The HTML head MUST include meta description, Open Graph tags, and viewport fix to improve SEO and discoverability.

#### Scenario: Meta Tags Implementation
- GIVEN the `<head>` section of `index.astro`
- THEN MUST include `<meta name="description" content="Diego Sasco — Arquitecto de Software y Especialista en Automatización e IA. Más de 7 años diseñando sistemas distribuidos, soluciones IoT e integración legacy para el sector financiero, retail e industrial.">`
- AND MUST include complete viewport tag: `<meta name="viewport" content="width=device-width, initial-scale=1">`
- AND MUST include canonical URL link with appropriate base URL
- AND MUST NOT include personal contact information in description

#### Scenario: Open Graph Tags Implementation
- GIVEN the `<head>` section
- THEN MUST include: `<meta property="og:title" content="Diego Sasco | Arquitecto de Software & Especialista en Automatización / IA">`
- AND MUST include: `<meta property="og:description" content="... (same as meta description)">`
- AND MUST include: `<meta property="og:type" content="website">`
- AND MUST include: `<meta property="og:url" content="...">`
- AND MUST include: `<meta property="og:locale" content="es_ES">`

### Requirement: REQ-C12

The HTML MUST include Schema.org JSON-LD markup with Person type to enhance search engine understanding of the CV content.

#### Scenario: Schema.org JSON-LD Implementation
- GIVEN the `<head>` section of `index.astro`
- THEN MUST include `<script type="application/ld+json">` with Person schema
- AND the JSON MUST include: `@type: "Person"`, `name: "Diego Sasco"`, `jobTitle: "Arquitecto de Software y Especialista en Automatización / IA"`
- AND MUST include `url` with the canonical URL
- AND MUST include `knowsAbout` array with key skills: "C#/.NET", "Go", "gRPC", "IoT", "integración de sistemas legacy"
- AND the JSON MUST be properly formatted and valid

#### Scenario: Schema Validation
- GIVEN the JSON-LD is included
- THEN the structure MUST conform to Schema.org Person specification
- AND all required fields MUST be present
- AND the JSON MUST be parseable and valid according to Schema.org standards

### Requirement: REQ-C13

The HTML head MUST include preconnect links to Google Fonts to optimize loading performance.

#### Scenario: Font Preconnect Implementation
- GIVEN the `<head>` section of `index.astro`
- THEN MUST include: `<link rel="preconnect" href="https://fonts.googleapis.com">`
- AND MUST include: `<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>`
- AND these links MUST be placed before the font import statement
- AND the preconnect MUST use appropriate crossorigin attribute for fonts.gstatic.com

### Requirement: REQ-C14

The CSS MUST include visible focus styles for keyboard navigation to improve accessibility.

#### Scenario: Focus Visible Styles
- GIVEN the CSS styles for interactive elements
- THEN MUST include `:focus-visible` pseudo-classes for links and buttons
- AND focus styles MUST include visible outline or background color changes
- AND focus styles MUST maintain sufficient color contrast
- AND the styles MUST work consistently across different browsers
- AND the focus styles MUST be visually distinct but not distracting

#### Scenario: Keyboard Navigation
- GIVEN a user navigates via keyboard
- THEN each interactive element MUST have a clear focus indicator
- AND the focus indicator MUST appear on tab navigation
- AND the focus styles MUST not interfere with the overall design

### Requirement: REQ-C15

The HTML MUST include appropriate ARIA roles and alt text to improve accessibility for screen readers.

#### Scenario: ARIA Roles Implementation
- GIVEN the HTML structure
- THEN the `<header>` element MUST have `role="banner"`
- AND the `<main>` element MUST have `role="main"`
- AND the recommendations section MUST have `role="complementary"`
- AND the contact SVG MUST have `aria-label="Email"`
- AND all ARIA roles MUST be semantically appropriate

#### Scenario: Accessibility Validation
- GIVEN the accessibility attributes are added
- THEN all ARIA roles MUST be valid and appropriate for their context
- AND alt text MUST be provided for meaningful images
- AND the structure MUST be navigable via screen readers
- AND the CV MUST meet WCAG 2.1 AA accessibility standards

## Data Contracts

### YAML Schema Changes

The `cv-source.yaml` file MUST be updated to include:

```yaml
# New field in personal_info
about_me: "Desarrollador de Software Senior con más de 7 años de experiencia en el diseño y automatización de sistemas para el sector financiero, retail e industrial..."

# Updated work_experience entries
- company: "Real2B"
  role: "Soporte Técnico y Operaciones de TI"
  description: "Proporcioné soporte técnico (Help Desk) a usuarios en múltiples locales, diagnosticando y resolviendo incidencias de infraestructura de TI y sistemas de comunicación. Administré la continuidad operativa de sistemas transaccionales y coordiné la resolución de problemas técnicos con equipos internos."
  technologies: ["help_desk", "problem_solving", "teamwork", "assertive_communication", "security_mindset"]

- company: "PIXI Supermercados"
  role: "Supervisor de Operaciones y Logística"
  description: "Coordiné equipos operativos y supervisé la logística de inventarios en un entorno de alta rotación, implementando controles sistemáticos de stock y optimizando los procesos de recepción de mercadería. Gestión de proveedores y resolución de conflictos operativos."
  technologies: ["leadership", "teamwork", "assertive_communication", "problem_solving"]

# Updated education entry
- institution: "Universidad de la República"
  degree: "Contador Público"
  start_date: "2015-03-01"
  end_date: "2016-12-15"
  status: "Contador Público — 1er año cursado (formación complementaria en gestión financiera y contable)"
```

### Go Model Changes

The `models.go` file MUST be updated to:

1. Add `AboutMe` field to `PersonalInfo` struct:
```go
type PersonalInfo struct {
    Name     string `yaml:"name" json:"name"`
    Title    string `yaml:"title" json:"title"`
    AboutMe  string `yaml:"about_me" json:"about_me"`  // NEW FIELD
    // ... other fields
}
```

2. Ensure the field passes through to `CVProcessed` struct without modification

### Processing Pipeline Changes

The `calculator.go` file MUST be updated to:

1. Pass through the `AboutMe` field from raw to processed data:
```go
PersonalInfo: PersonalInfo{
    Name:     raw.PersonalInfo.Name,
    Title:    raw.PersonalInfo.Title,
    AboutMe:  raw.PersonalInfo.AboutMe,  // NEW PASS-THROUGH
    // ... other fields
}
```

### Template Changes

The `index.astro` file MUST be updated to:

1. Add About Me section after hero:
```astro
{personalInfo.about_me && (
    <section class="about-me">
        <h2>Sobre Mí</h2>
        <p>{personalInfo.about_me}</p>
    </section>
)}
```

2. Add SEO meta tags, Schema.org JSON-LD, and ARIA roles as specified in requirements

### CSS Changes

The `global.css` file MUST be updated to:

1. Modify print media queries for single-column layouts and ATS fonts
2. Add `:focus-visible` styles for accessibility
3. Update margin and padding for print layout

## Non-Functional Requirements

### Performance Requirements
- The page MUST load within 3 seconds on average broadband connection
- Font loading MUST be optimized using preconnect links
- SEO optimization MUST not impact initial page load performance

### Compatibility Requirements
- The website MUST be compatible with all modern browsers (Chrome, Firefox, Safari, Edge)
- The PDF export MUST work reliably in headless Chrome environments
- Print styles MUST be compatible with standard printers and PDF viewers

### Print Behavior
- The PDF MUST export with proper margins (0.75 inches) and ATS-compatible fonts
- Contact information MUST be hidden in PDF output (already implemented)
- The layout MUST be optimized for readability on paper

## Out of Scope

The following elements are explicitly excluded from this change:

1. **Additional contact channels** - Email, LinkedIn, and GitHub contact information remain locked to Google Form only per user decision
2. **Quantitative metrics** - No numerical percentages or hard metrics will be added as per user decision
3. **Recommendation attribution** - Author placeholders will remain as `<RECOMMENDER_X>` per user decision
4. **Light/dark mode toggle** - Deferred to future change with P3 priority
5. **Project URL integration** - Projects with empty URLs remain unchanged until real content is available
6. **Low-skill removal** - Only reframing of existing skills, not complete removal
7. **Font self-hosting** - Preconnect optimization is sufficient, no self-hosting planned
8. **Advanced SEO** - robots.txt and sitemap.xml deferred to future changes

## Implementation Dependencies

The changes MUST be implemented in the following order:

1. **Go Pipeline Updates** - Add AboutMe field to models and ensure pass-through to processed data
2. **YAML Content Updates** - Reframe jobs, apply CAR framework, curate projects, contextualize education
3. **Pipeline Re-execution** - Regenerate cv-processed.json and verify skill duration calculations
4. **Template Updates** - Add About Me section, SEO tags, Schema.org, ARIA roles
5. **CSS Updates** - Print layout, focus styles, ATS fonts
6. **Verification** - Run tests, verify build, validate PDF output and web display
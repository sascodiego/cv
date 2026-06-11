# SDD Implementation Tasks — cv-web-optimization

**Change ID**: `cv-web-optimization`
**Phase**: Tasks
**Date**: 2026-06-10
**Status**: Ready for implementation

---

## Phase 1: Go Pipeline Changes

> Files: `pipeline/models.go`, `pipeline/calculator.go`, `pipeline/calculator_test.go`

- [x] **1a. models.go — Add `AboutMe` field to `PersonalInfo` struct (line ~94)**
  Insert `AboutMe  string \`yaml:"about_me" json:"about_me,omitempty"\`` after the `Title` field.
  The `omitempty` tag ensures backward compatibility when the field is absent.

- [x] **1b. models.go — Add `Hidden` field to `Project` struct (line ~131)**
  Append `Hidden bool \`yaml:"hidden" json:"-"\`` after the `URL` field.
  The `json:"-"` tag prevents the flag from appearing in processed JSON output.

- [x] **1c. calculator.go — Pass through `AboutMe` in `ProcessCV()` (line ~81)**
  Add `AboutMe: raw.PersonalInfo.AboutMe,` to the `PersonalInfo` literal inside the `processed := &CVProcessed{...}` block, after the `Title` assignment.

- [x] **1d. calculator.go — Filter hidden projects from associated work output (line ~119)**
  Inside the `for _, projID := range work.Projects` loop, wrap the `processedProjects = append(...)` call in `if !proj.Hidden { ... }`. Tech interval registration (`techIntervals[tech] = append(...)`) stays OUTSIDE the guard so hidden projects still contribute to skill durations.

- [x] **1e. calculator.go — Skip hidden independent projects (line ~156)**
  In the independent projects loop (`for _, proj := range raw.Projects`), add `if proj.Hidden { continue }` after the `associatedProjects` check. Add a preceding loop to register tech intervals for hidden independent projects before the rendering loop.

- [x] **1f. calculator_test.go — Add `TestAboutMePassThrough`**
  Test that a non-empty `AboutMe` string flows from `CVRaw.PersonalInfo` to `CVProcessed.PersonalInfo` unchanged.

- [x] **1g. calculator_test.go — Add `TestAboutMeEmptyString`**
  Test that an empty/missing `AboutMe` produces an empty string in processed output without errors.

- [x] **1h. calculator_test.go — Add `TestHiddenProjectFiltered`**
  Test with a work entry containing one visible and one hidden project. Assert: (a) only 1 project in rendered output, (b) skill duration includes both projects' intervals (12 months total).

- [x] **1i. calculator_test.go — Add `TestHiddenIndependentProject`**
  Test with a hidden independent project. Assert: (a) 0 independent projects in rendered output, (b) skill duration still includes the hidden project's interval (6 months).

- [x] **1j. Verify — Run `go test ./...` in `/src/cv-pipeline/pipeline`**
  All existing tests + 4 new tests must pass with zero failures.

---

## Phase 2: YAML Content Changes

> File: `data/cv-source.yaml`

- [x] **2a. `personal_info.title` — Change to "Desarrollador de Software Senior" (line ~4)**
  Replace `title: "Arquitecto de Software & Especialista en Automatización / IA"` with `title: "Desarrollador de Software Senior"`.

- [x] **2b. `personal_info.about_me` — Add new field (after title, line ~5)**
  Insert `about_me: "Desarrollador de Software Senior con más de 7 años de experiencia en el diseño y automatización de sistemas para el sector financiero, retail e industrial. Especializado en integración de sistemas legacy, arquitectura IoT y soluciones RPA, con competencias en backend distribuido (C#/.NET, Go, gRPC, RabbitMQ) y formación en seguridad informática. Profesional orientado a resultados, con comunicación asertiva y enfoque en la confiabilidad de sistemas críticos."`

- [x] **2c. Skills — Remove `customer_service` entry entirely**
  Delete the skill block with `id: customer_service` / `name: "Atención al Público & Gestión de Clientes"`.

- [x] **2d. Real2B — Reframe role, description, technologies (REQ-C2)**
  - `role`: → `"Soporte Técnico y Operaciones de TI"`
  - `description`: → `"Proporcioné soporte técnico (Help Desk) a usuarios en múltiples locales, diagnosticando y resolviendo incidencias de infraestructura de TI y sistemas de comunicación. Administré la continuidad operativa de sistemas transaccionales y coordiné la resolución de problemas técnicos con equipos internos."`
  - `technologies`: Replace `customer_service` with `security_mindset`. Final list: `["help_desk", "problem_solving", "teamwork", "assertive_communication", "security_mindset"]`

- [x] **2e. PIXI Supermercados — Reframe role, description, technologies (REQ-C3)**
  - `role`: → `"Supervisor de Operaciones y Logística"`
  - `description`: → `"Coordiné equipos operativos y supervisé la logística de inventarios en un entorno de alta rotación, implementando controles sistemáticos de stock y optimizando los procesos de recepción de mercadería. Gestión de proveedores y resolución de conflictos operativos."`
  - `technologies`: Remove `customer_service`, add `problem_solving`. Final list: `["leadership", "teamwork", "assertive_communication", "problem_solving"]`

- [x] **2f. All work descriptions — Apply CAR qualitative framework (REQ-C4)**
  Rewrite each description with past-tense action verbs:
  - **Desarrollo Independiente**: → `"Diseñé e implementé sistemas distribuidos en Go y C#/.NET para automatización IoT e integración legacy, colaborando con especialistas en hardware electrónico en el prototipado de soluciones embebidas. Aplicación sistemática de metodologías de desarrollo asistido por IA para maximizar la calidad y velocidad de entrega."`
  - **Sluckis Hermanos SA**: → `"Lideré un proyecto integral de normalización de inventarios (mayoristas y minoristas), diagnosticando y depurando discrepancias de stock histórico en bases de datos SQL, y definiendo métodos de trabajo estandarizados que simplificaron los procesos de control."`
  - **Bionico / Figital Tech**: → `"Dirigí el desarrollo de soluciones de automatización de procesos financieros (RPA), microservicios y procesamiento masivo de datos, diseñando una arquitectura observable con RabbitMQ que garantizó la trazabilidad de operaciones críticas."`
  - **TopBrands Int.**: → `"Automatizé procesos de retail mediante herramientas personalizadas en C# y SQL, administrando la infraestructura de TI y sistemas de comunicación corporativos, y proporcionando soporte técnico especializado a usuarios finales en un entorno multi-sucursal."`

- [x] **2g. Project curation — Add `hidden: true` to 7 projects (REQ-C5)**
  Add `hidden: true` to each of these project blocks:
  - `ai_engineering_workflow` — Flujo de Ingeniería Asistida por IA (SDD & TDD)
  - `cv_pipeline` — CV Pipeline CLI
  - `normalizacion_inventario` — Sluckis - Normalización de Inventario
  - `sasco_agency_ai` — Soluciones de IA Generativa
  - `retail_automation` — Automatización de Inventario y Retail
  - `opt_administrativa` — Optimización de Flujos Administrativos

- [x] **2h. Project 3 rename — Update `iot_hardware_automation` name and description**
  - `name`: → `"Plataforma de Monitoreo Remoto de Activos IoT"`
  - `description`: → `"Diseñé y programé la plataforma de monitoreo remoto de activos IoT, implementando la capa de comunicación (GPIO, protocolos seriales) y control de firmware para enlazar microcontroladores con APIs distribuidas en Go y gRPC, habilitando el control remoto de dispositivos industriales en entornos de producción."`

- [x] **2i. Visible project descriptions — CAR rewrite for 4 remaining projects**
  - `rppos_iot`: → `"Diseñé e implementé la interfaz de usuario y el control de hardware (GPIO) en Go con Fyne UI para un Punto de Venta embebido en Raspberry Pi, optimizando el rendimiento en un entorno de recursos limitados para control de maquinaria industrial."`
  - `proxy_getnet`: → `"Implementé un proxy inverso en C# que traduce protocolos legacy SOAP a gRPC, habilitando la comunicación eficiente con pasarelas financieras de GetNet/Santander y consolidando la integración de sistemas transaccionales."`
  - `framework_rpa`: → `"Desarrollé un framework de bots de escritorio (Windows UI Automation) en C# con observabilidad vía RabbitMQ, automatizando la inserción de datos en sistemas contables y reduciendo significativamente la intervención manual en procesos financieros críticos."`
  - `extraccion_pdf`: → `"Construí un motor inteligente de extracción basado en coordenadas X/Y (PdfPig) sobre PDFs semi-estructurados en C#, habilitando el procesamiento masivo automatizado de facturas y estandarizando la ingesta de documentos financieros."`

- [x] **2j. Education — Contextualize Contador Público status (REQ-C6)**
  Change the Contador Público `status` field to: `"1er año cursado (formación complementaria en gestión financiera y contable)"`. Confirm it remains at the end of the education list.

---

## Phase 3: Pipeline Re-run

- [x] **3a. Run Go pipeline to regenerate `cv-processed.json`**
  Execute: `cd /src/cv-pipeline/pipeline && go run main.go`

- [x] **3b. Verify `about_me` present in JSON output**
  Check `web/src/data/cv-processed.json` → `personal_info.about_me` exists and matches the YAML text.

- [x] **3c. Verify exactly 5 visible projects in output**
  Count rendered projects in `cv-processed.json`: 3 under Desarrollo Independiente work entry (rppos_iot, proxy_getnet, iot_hardware_automation) + 2 under Bionico (framework_rpa, extraccion_pdf).

- [x] **3d. Verify hidden projects' skills retain correct durations**
  Check these skills have non-zero `months_experience`:
  - `python`: ~36 months (from ai_engineering_workflow + sasco_agency_ai)
  - `docker`: ~24 months (from ai_engineering_workflow)
  - `javascript`: ~12 months (from sasco_agency_ai)
  - `generative_ai`: ~36 months
  - `ai_assisted_dev`: ~24 months

- [x] **3e. Verify `customer_service` absent from output**
  Confirm no skill with `id: "customer_service"` exists in `cv-processed.json` skills array.

- [x] **3f. Verify Real2B and PIXI roles/descriptions**
  Confirm Real2B shows role `"Soporte Técnico y Operaciones de TI"` and PIXI shows `"Supervisor de Operaciones y Logística"`.

---

## Phase 4: Template Changes

> File: `web/src/pages/index.astro`

- [x] **4a. Fix viewport meta tag (line 12)**
  Replace `<meta name="viewport" content="width=device-width" />` with `<meta name="viewport" content="width=device-width, initial-scale=1" />`.

- [x] **4b. Add SEO meta tags — description, canonical, Open Graph (after line 13)**
  Insert after `<title>`:
  - `<meta name="description" content={...} />` with text referencing the title + 7+ years + sectors
  - `<link rel="canonical" href={...} />` pointing to `https://diegosasco.github.io`
  - OG tags: `og:title`, `og:description`, `og:type`, `og:url`, `og:locale`

- [x] **4c. Add Schema.org JSON-LD (REQ-C12)**
  Insert `<script type="application/ld+json" set:html={JSON.stringify({...})} />` with `@type: "Person"`, `name`, `jobTitle`, `url`, `knowsAbout` array.

- [x] **4d. Add Google Fonts preconnect links (REQ-C13)**
  Insert `<link rel="preconnect" href="https://fonts.googleapis.com" />` and `<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />` before the font import in the CSS or after the OG tags.

- [x] **4e. Add About Me section (REQ-C1)**
  Insert between `</header>` (line 40) and the `{/* Experiencia Laboral */}` comment (line 42):
  ```astro
  {personalInfo.about_me && (
    <section class="about-me">
      <h2>Sobre Mí</h2>
      <p class="about-me-text">{personalInfo.about_me}</p>
    </section>
  )}
  ```

- [x] **4f. Add ARIA roles (REQ-C15)**
  - `<header class="hero">` → `<header class="hero" role="banner">` (line 20)
  - `<main>` → `<main role="main">` (line 18)
  - Recommendations `<section class="section-card">` → add `role="complementary"` (line ~103)
  - Contact SVG → add `aria-label="Email"` (line 23)

- [x] **4g. Verify — Run `pnpm build` in `/src/cv-pipeline/web`**
  Build must complete with zero errors.

---

## Phase 5: CSS Changes

> File: `web/src/styles/global.css`

- [x] **5a. Add About Me section styles (after `.contact-form-btn svg` block, ~line 122)**
  Insert `.about-me` styles: centered text, padding, border-bottom, `max-inline-size: 720px`, muted color, 0.95rem font, 1.7 line-height.

- [x] **5b. Add focus-visible styles (REQ-C14, after About Me styles)**
  Insert `a:focus-visible, button:focus-visible, .contact-form-btn:focus-visible` with `outline: 2px solid var(--primary-color)`, `outline-offset: 2px`, `border-radius: 4px`.

- [x] **5c. Print — Skills grid single column (REQ-C7)**
  Inside `@media print`, change `.skills-grid { grid-template-columns: repeat(3, 1fr); }` → `grid-template-columns: 1fr;` (line ~334).

- [x] **5d. Print — Recommendations grid single column (REQ-C8)**
  Inside `@media print`, change `.recommendations-grid { grid-template-columns: repeat(2, 1fr); }` → `grid-template-columns: 1fr;` (line ~419).

- [x] **5e. Print — 0.75" page margins via `@page` (REQ-C9)**
  Inside `@media print`, after the `body` rule block, add:
  ```css
  @page {
    margin: 0.75in;
  }
  ```

- [x] **5f. Print — ATS-compatible font overrides (REQ-C10)**
  Inside the existing print `:root` block, add:
  ```css
  --font-title: "Georgia", serif;
  --font-body: "Arial", sans-serif;
  ```

- [x] **5g. Print — About Me compact styles**
  Inside `@media print`, after the `.hero` print styles:
  ```css
  .about-me {
    padding: 0.5rem 0;
    border-block-end: 1px solid #e2e8f0;
  }
  .about-me h2 {
    font-size: 12pt;
    color: #0f172a;
  }
  .about-me-text {
    font-size: 9.5pt;
    color: #475569;
  }
  ```

---

## Phase 6: Build & Verify

- [x] **6a. Go tests — `cd /src/cv-pipeline/pipeline && go test ./...`**
  All tests pass (existing + 4 new).

- [x] **6b. Astro build — `cd /src/cv-pipeline/web && pnpm build`**
  Build completes with zero errors.

- [x] **6c. Visual verification — `pnpm dev` checklist**
  - [x] About Me section visible between hero and work experience
  - [x] About Me text matches the approved 75-word text exactly
  - [x] Hero subtitle shows "Desarrollador de Software Senior"
  - [x] Real2B role shows "Soporte Técnico y Operaciones de TI"
  - [x] Real2B description starts with "Proporcioné soporte técnico..."
  - [x] PIXI role shows "Supervisor de Operaciones y Logística"
  - [x] Exactly 3 projects under Desarrollo Independiente, 2 under Bionico
  - [x] Project 3 named "Plataforma de Monitoreo Remoto de Activos IoT" (no "Confidencial")
  - [x] Hidden projects (ai_engineering_workflow, cv_pipeline, etc.) NOT visible
  - [x] `customer_service` / "Atención al Público" absent from Skills section
  - [x] Contador Público shows contextualized status text
  - [x] All descriptions use past-tense action verbs

- [x] **6d. Print/PDF verification — browser Print Preview checklist**
  - [x] Skills in single column
  - [x] Recommendations in single column
  - [x] ~0.75" page margins
  - [x] Georgia for titles, Arial for body text
  - [x] About Me visible but compact
  - [x] Contact button hidden
  - [x] White background, dark text

- [x] **6e. SEO/Accessibility verification**
  - [x] `<meta name="description">` present in page source
  - [x] OG tags present (`og:title`, `og:description`, `og:type`, `og:url`, `og:locale`)
  - [x] `<link rel="canonical">` present
  - [x] Schema.org JSON-LD valid (inspect `<script type="application/ld+json">`)
  - [x] Preconnect links present before font import
  - [x] `<header>` has `role="banner"`, `<main>` has `role="main"`
  - [x] Recommendations section has `role="complementary"`
  - [x] Contact SVG has `aria-label="Email"`
  - [x] Tab through page → focus-visible outline on interactive elements

---

## Review Workload Forecast

| Field | Value |
|-------|-------|
| Estimated changed lines | ~205–233 |
| 400-line budget risk | Low |
| Chained PRs recommended | No |
| Suggested split | Single PR — all phases |
| Delivery strategy | auto-chain (single PR within budget) |
| Chain strategy | single-pr |

Decision needed before apply: No
Chained PRs recommended: No
Chain strategy: single-pr
400-line budget risk: Low

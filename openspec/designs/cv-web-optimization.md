# SDD Technical Design — cv-web-optimization

**Change ID**: `cv-web-optimization`
**Phase**: Design
**Date**: 2026-06-10
**Status**: Ready for implementation
**Estimated Size**: ~150–200 changed lines (within 400-line PR budget)

---

## 1. Architecture Overview

The CV system follows a strict single-direction pipeline:

```
cv-source.yaml ──► Go Pipeline ──► cv-processed.json ──► Astro Template ──► HTML
   (YAML)        (models.go +         (JSON)          (index.astro)      (web)
                  calculator.go)
```

### Pipeline Stages

| Stage | File | Responsibility |
|-------|------|----------------|
| **Source** | `data/cv-source.yaml` | Single source of truth. All content, skills, projects, education live here. |
| **Models** | `pipeline/models.go` | Go structs mirroring YAML schema (`CVRaw`, `CVProcessed`, etc.). |
| **Processing** | `pipeline/calculator.go` | `ProcessCV()` transforms raw data: merges overlapping skill intervals, formats durations, resolves project associations, filters. |
| **Output** | `web/src/data/cv-processed.json` | Computed JSON consumed by the Astro template. |
| **Template** | `web/src/pages/index.astro` | Renders sections from processed JSON. Single-page CV. |
| **Styles** | `web/src/styles/global.css` | Dark premium theme for web + `@media print` overrides for ATS-compatible PDF. |

### Data Ownership Rule

- **YAML owns content**: role names, descriptions, skill definitions, project metadata.
- **Go owns computation**: date formatting, duration calculation, interval merging, project filtering.
- **Astro owns rendering**: section layout, SEO tags, accessibility attributes.
- **CSS owns presentation**: typography, spacing, print layout, focus states.

---

## 2. Component Changes

### 2.1 `data/cv-source.yaml` — Content Changes

This file receives the bulk of changes. All edits are pure content — no structural changes to the YAML schema beyond adding two new fields.

#### 2.1.1 `personal_info` — New `about_me` + Title Change

**Current** (lines 3–4):
```yaml
personal_info:
  name: "Diego Sasco"
  title: "Arquitecto de Software & Especialista en Automatización / IA"
```

**Target**:
```yaml
personal_info:
  name: "Diego Sasco"
  title: "Desarrollador de Software Senior"
  about_me: "Desarrollador de Software Senior con más de 7 años de experiencia en el diseño y automatización de sistemas para el sector financiero, retail e industrial. Especializado en integración de sistemas legacy, arquitectura IoT y soluciones RPA, con competencias en backend distribuido (C#/.NET, Go, gRPC, RabbitMQ) y formación en seguridad informática. Profesional orientado a resultados, con comunicación asertiva y enfoque en la confiabilidad de sistemas críticos."
```

**Locked decisions applied**:
- Title changed to "Desarrollador de Software Senior" per locked decision.
- `about_me` text is ≤80 words (75 words), no first-person pronouns, opens with professional identity.

#### 2.1.2 Skills — Remove `customer_service`

**Remove** this skill entry entirely (lines ~50–52):
```yaml
  - id: customer_service
    name: "Atención al Público & Gestión de Clientes"
    category: "Habilidades Funcionales"
```

**Rationale**: After reframing Real2B and PIXI, no remaining work entry or project references `customer_service`. Keeping it would produce 0 months experience ("Sin experiencia").

#### 2.1.3 Real2B — Reframe (REQ-C2)

**Current**:
```yaml
  - company: "Real2B"
    role: "Analista Operacional & Control de Flujo"
    description: "Gestión integral de operaciones de venta de servicios, transacciones de caja, soporte técnico (Help Desk) a usuarios y atención directa al público en múltiples locales."
    technologies: ["help_desk", "customer_service", "teamwork", "assertive_communication", "problem_solving"]
```

**Target**:
```yaml
  - company: "Real2B"
    role: "Soporte Técnico y Operaciones de TI"
    description: "Proporcioné soporte técnico (Help Desk) a usuarios en múltiples locales, diagnosticando y resolviendo incidencias de infraestructura de TI y sistemas de comunicación. Administré la continuidad operativa de sistemas transaccionales y coordiné la resolución de problemas técnicos con equipos internos."
    technologies: ["help_desk", "problem_solving", "teamwork", "assertive_communication", "security_mindset"]
```

**Changes**:
- `role`: "Analista Operacional & Control de Flujo" → "Soporte Técnico y Operaciones de TI"
- `description`: Replaced with CAR-qualitative text (action verbs + no low-signal terms).
- `technologies`: Removed `customer_service`, added `security_mindset`.

**Skill duration impact**: Real2B has no projects, so its work-level interval (Jan 2024 – Jun 2024, 6 months) applies to all its technologies. After the change, `help_desk` keeps this interval; `security_mindset` gains 6 months (added to its existing TopBrands interval of 82 months).

#### 2.1.4 PIXI Supermercados — Reframe (REQ-C3)

**Current**:
```yaml
  - company: "PIXI Supermercados"
    role: "Encargado de Operaciones"
    description: "Supervisión de equipos, administración de personal, control de inventario, arqueos de caja y atención al público (resolución de conflictos y gestión de proveedores)."
    technologies: ["customer_service", "leadership", "teamwork", "assertive_communication"]
```

**Target**:
```yaml
  - company: "PIXI Supermercados"
    role: "Supervisor de Operaciones y Logística"
    description: "Coordiné equipos operativos y supervisé la logística de inventarios en un entorno de alta rotación, implementando controles sistemáticos de stock y optimizando los procesos de recepción de mercadería. Gestión de proveedores y resolución de conflictos operativos."
    technologies: ["leadership", "teamwork", "assertive_communication", "problem_solving"]
```

**Changes**:
- `role`: "Encargado de Operaciones" → "Supervisor de Operaciones y Logística"
- `description`: Replaced with CAR-qualitative text.
- `technologies`: Removed `customer_service`, added `problem_solving`.

**Skill duration impact**: No projects, so work-level interval (Jul 2016 – Mar 2017, 9 months) applies to technologies. `problem_solving` gains 9 months; `customer_service` loses 9 months (but is removed from skills list entirely).

#### 2.1.5 All Work Experience Descriptions — CAR Qualitative (REQ-C4)

Each description must be rewritten using past-tense action verbs + specific technology + impact statement. Pattern: `[Action verb] + [solution] + [tech] + para + [context], lo que + [qualitative impact]`.

| Work Entry | New Description |
|------------|----------------|
| **Desarrollo Independiente** | "Diseñé e implementé sistemas distribuidos en Go y C#/.NET para automatización IoT e integración legacy, colaborando con especialistas en hardware electrónico en el prototipado de soluciones embebidas. Aplicación sistemática de metodologías de desarrollo asistido por IA para maximizar la calidad y velocidad de entrega." |
| **Sluckis Hermanos SA** | "Lideré un proyecto integral de normalización de inventarios (mayoristas y minoristas), diagnosticando y depurando discrepancias de stock histórico en bases de datos SQL, y definiendo métodos de trabajo estandarizados que simplificaron los procesos de control." |
| **Bionico / Figital Tech** | "Dirigí el desarrollo de soluciones de automatización de procesos financieros (RPA), microservicios y procesamiento masivo de datos, diseñando una arquitectura observable con RabbitMQ que garantizó la trazabilidad de operaciones críticas." |
| **TopBrands Int.** | "Automatizé procesos de retail mediante herramientas personalizadas en C# y SQL, administrando la infraestructura de TI y sistemas de comunicación corporativos, y proporcionando soporte técnico especializado a usuarios finales en un entorno multi-sucursal." |

**Note**: Real2B and PIXI descriptions are already specified in sections 2.1.3 and 2.1.4 above.

#### 2.1.6 Project Curation — 12 → 5 Flagship (REQ-C5)

Add `hidden: true` to 7 projects. Keep them in YAML for data integrity and skill interval calculation. The Go pipeline will filter them from rendered output.

**Projects to mark `hidden: true`**:

| Project ID | Name | Reason for Hiding |
|------------|------|-------------------|
| `ai_engineering_workflow` | Flujo de Ingeniería Asistida por IA (SDD & TDD) | Methodology, not deliverable product |
| `cv_pipeline` | CV Pipeline CLI | Meta-project, low signal for external recruiter |
| `normalizacion_inventario` | Sluckis - Normalización de Inventario | Low technical differentiation |
| `sasco_agency_ai` | Soluciones de IA Generativa | Too generic without demos |
| `retail_automation` | Automatización de Inventario y Retail | Moderate signal, less differentiated than flagship 5 |
| `opt_administrativa` | Optimización de Flujos Administrativos | Low technical signal |
| *(none additional)* | | |

**5 remaining visible projects**:

| # | Project ID | Name (updated) | Company | Signal |
|---|-----------|----------------|---------|--------|
| 1 | `rppos_iot` | RpPOS - Arquitectura IoT (Golden Wash) | Desarrollo Independiente | Go + Raspberry Pi + GPIO + Fyne |
| 2 | `proxy_getnet` | Proxy Inverso Financiero | Desarrollo Independiente | C# + SOAP→gRPC + financial gateway |
| 3 | `iot_hardware_automation` | Plataforma de Monitoreo Remoto de Activos IoT | Desarrollo Independiente | Go + gRPC + serial protocols + IoT |
| 4 | `framework_rpa` | Framework RPA Financiero | Bionico | C# + RabbitMQ + UI Automation |
| 5 | `extraccion_pdf` | Motor de Extracción PDF | Bionico | C# + PdfPig + massive processing |

**Rename project 3** (locked decision):
```yaml
  - id: iot_hardware_automation
    name: "Plataforma de Monitoreo Remoto de Activos IoT"
```
The description should also be updated to remove "(Confidencial)" and emphasize the monitoring platform angle:
```yaml
    description: "Diseñé y programé la plataforma de monitoreo remoto de activos IoT, implementando la capa de comunicación (GPIO, protocolos seriales) y control de firmware para enlazar microcontroladores con APIs distribuidas en Go y gRPC, habilitando el control remoto de dispositivos industriales en entornos de producción."
```

**YAML change pattern** — add `hidden: true` to each hidden project:
```yaml
  - id: ai_engineering_workflow
    name: "Flujo de Ingeniería Asistida por IA (SDD & TDD)"
    hidden: true
    # ... rest unchanged
```

**Visible project descriptions — CAR rewrite**:

| Project ID | New Description |
|------------|----------------|
| `rppos_iot` | "Diseñé e implementé la interfaz de usuario y el control de hardware (GPIO) en Go con Fyne UI para un Punto de Venta embebido en Raspberry Pi, optimizando el rendimiento en un entorno de recursos limitados para control de maquinaria industrial." |
| `proxy_getnet` | "Implementé un proxy inverso en C# que traduce protocolos legacy SOAP a gRPC, habilitando la comunicación eficiente con pasarelas financieras de GetNet/Santander y consolidando la integración de sistemas transaccionales." |
| `framework_rpa` | "Desarrollé un framework de bots de escritorio (Windows UI Automation) en C# con observabilidad vía RabbitMQ, automatizando la inserción de datos en sistemas contables y reduciendo significativamente la intervención manual en procesos financieros críticos." |
| `extraccion_pdf` | "Construí un motor inteligente de extracción basado en coordenadas X/Y (PdfPig) sobre PDFs semi-estructurados en C#, habilitando el procesamiento masivo automatizado de facturas y estandarizando la ingesta de documentos financieros." |

#### 2.1.7 Education Contextualization (REQ-C6)

**Change Contador Público status** (locked at end of education list — already in correct position):

```yaml
  - institution: "Universidad de la República"
    degree: "Contador Público"
    start_date: "2015-03-01"
    end_date: "2016-12-15"
    status: "1er año cursado (formación complementaria en gestión financiera y contable)"
```

**No reordering needed**: The current YAML order (Analista → Ingeniería → Hacking Ético → Contador Público) already matches the spec requirement.

---

### 2.2 `pipeline/models.go` — Struct Changes

#### 2.2.1 Add `AboutMe` to `PersonalInfo`

**Location**: `PersonalInfo` struct (line ~93)

**Current**:
```go
type PersonalInfo struct {
	Name     string `yaml:"name" json:"name"`
	Title    string `yaml:"title" json:"title"`
	Email    string `yaml:"email" json:"email,omitempty"`
	Website  string `yaml:"website" json:"website,omitempty"`
	GitHub   string `yaml:"github" json:"github,omitempty"`
	LinkedIn string `yaml:"linkedin" json:"linkedin,omitempty"`
}
```

**Target** — add field after `Title`:
```go
type PersonalInfo struct {
	Name     string `yaml:"name" json:"name"`
	Title    string `yaml:"title" json:"title"`
	AboutMe  string `yaml:"about_me" json:"about_me,omitempty"`
	Email    string `yaml:"email" json:"email,omitempty"`
	Website  string `yaml:"website" json:"website,omitempty"`
	GitHub   string `yaml:"github" json:"github,omitempty"`
	LinkedIn string `yaml:"linkedin" json:"linkedin,omitempty"`
}
```

- `omitempty` ensures backward compatibility — if `about_me` is absent, JSON output omits the key.
- `CVProcessed.PersonalInfo` reuses the same `PersonalInfo` struct, so the field propagates automatically.

#### 2.2.2 Add `Hidden` to `Project`

**Location**: `Project` struct (line ~122)

**Current**:
```go
type Project struct {
	ID           string   `yaml:"id" json:"id"`
	Name         string   `yaml:"name" json:"name"`
	Description  string   `yaml:"description" json:"description"`
	StartDate    CVDate   `yaml:"start_date" json:"start_date"`
	EndDate      CVDate   `yaml:"end_date" json:"end_date"`
	Technologies []string `yaml:"technologies" json:"technologies"`
	URL          string   `yaml:"url" json:"url"`
}
```

**Target** — add field after `URL`:
```go
type Project struct {
	ID           string   `yaml:"id" json:"id"`
	Name         string   `yaml:"name" json:"name"`
	Description  string   `yaml:"description" json:"description"`
	StartDate    CVDate   `yaml:"start_date" json:"start_date"`
	EndDate      CVDate   `yaml:"end_date" json:"end_date"`
	Technologies []string `yaml:"technologies" json:"technologies"`
	URL          string   `yaml:"url" json:"url"`
	Hidden       bool     `yaml:"hidden" json:"-"`
}
```

- `json:"-"` ensures hidden flag never appears in processed JSON output.
- Default Go zero-value (`false`) means existing projects without the field remain visible — no migration needed.

---

### 2.3 `pipeline/calculator.go` — Logic Changes

#### 2.3.1 Pass Through `AboutMe` in `ProcessCV()`

**Location**: `ProcessCV()` function, `processed := &CVProcessed{...}` block (line ~80)

**Current**:
```go
processed := &CVProcessed{
	PersonalInfo: PersonalInfo{
		Name:  raw.PersonalInfo.Name,
		Title: raw.PersonalInfo.Title,
	},
	// ...
}
```

**Target**:
```go
processed := &CVProcessed{
	PersonalInfo: PersonalInfo{
		Name:    raw.PersonalInfo.Name,
		Title:   raw.PersonalInfo.Title,
		AboutMe: raw.PersonalInfo.AboutMe,
	},
	// ...
}
```

**Lines changed**: 1 line added.

#### 2.3.2 Filter Hidden Projects from Rendered Output

**Location**: Inside the `for _, work := range raw.WorkExperience` loop, in the inner `for _, projID := range work.Projects` block (line ~100).

**Current behavior**: All associated projects are added to `processedProjects`.

**Target**: Register tech intervals for ALL projects (including hidden), but only add non-hidden projects to `processedProjects`.

**Current code** (approximately line 105–130):
```go
for _, projID := range work.Projects {
	proj, found := projectsMap[projID]
	if !found {
		fmt.Printf("⚠️ Advertencia: ...")
		continue
	}
	associatedProjects[projID] = true
	pStart := proj.StartDate.Time
	pEnd := proj.EndDate.Time
	pMonths := CalculateMonths(pStart, pEnd)

	// ... period text formatting ...

	// Register tech intervals
	for _, tech := range proj.Technologies {
		techIntervals[tech] = append(techIntervals[tech], Interval{Start: pStart, End: pEnd})
	}

	processedProjects = append(processedProjects, ProjectProcessed{
		// ...
	})
}
```

**Target** — wrap the `processedProjects` append in a hidden check:
```go
	// Register tech intervals for ALL projects (hidden ones still contribute to skill durations)
	for _, tech := range proj.Technologies {
		techIntervals[tech] = append(techIntervals[tech], Interval{Start: pStart, End: pEnd})
	}

	// Only include visible projects in rendered output
	if !proj.Hidden {
		processedProjects = append(processedProjects, ProjectProcessed{
			Name:         proj.Name,
			Description:  proj.Description,
			// ... rest unchanged ...
		})
	}
```

**Location**: Also in the "Proyectos Independientes" loop (line ~155):

```go
for _, proj := range raw.Projects {
	if associatedProjects[proj.ID] {
		continue
	}
	if proj.Hidden {
		continue // Hidden independent projects are not rendered
	}
	// ... rest unchanged ...
}
```

**Important**: Hidden independent projects' tech intervals must still be registered. Add a separate loop BEFORE the independent rendering loop:

```go
// Register tech intervals for hidden independent projects (they contribute to skill durations)
for _, proj := range raw.Projects {
	if associatedProjects[proj.ID] || !proj.Hidden {
		continue
	}
	pStart := proj.StartDate.Time
	pEnd := proj.EndDate.Time
	for _, tech := range proj.Technologies {
		techIntervals[tech] = append(techIntervals[tech], Interval{Start: pStart, End: pEnd})
	}
}
```

**Note**: In the current data, all hidden projects are associated with work entries, so this hidden-independent loop is a safety measure. It costs ~8 lines and prevents future bugs.

**Total lines changed in calculator.go**: ~12–15 lines.

---

### 2.4 `web/src/pages/index.astro` — Template Changes

#### 2.4.1 Add SEO Meta Tags (REQ-C11, REQ-C13)

**Location**: Inside `<head>`, after the existing `<title>` tag (line 9).

**Insert** (after `<title>{personalInfo.name} | {personalInfo.title}</title>`):

```astro
		<meta name="description" content={`Diego Sasco — ${personalInfo.title}. Más de 7 años diseñando sistemas distribuidos, soluciones IoT e integración legacy para el sector financiero, retail e industrial.`} />
		<meta name="viewport" content="width=device-width, initial-scale=1" />
		<link rel="canonical" href={`https://diegosasco.github.io${base}`} />
		<!-- Open Graph -->
		<meta property="og:title" content={`${personalInfo.name} | ${personalInfo.title}`} />
		<meta property="og:description" content={`Diego Sasco — ${personalInfo.title}. Más de 7 años diseñando sistemas distribuidos, soluciones IoT e integración legacy para el sector financiero, retail e industrial.`} />
		<meta property="og:type" content="website" />
		<meta property="og:url" content={`https://diegosasco.github.io${base}`} />
		<meta property="og:locale" content="es_ES" />
		<!-- Preconnect Google Fonts -->
		<link rel="preconnect" href="https://fonts.googleapis.com" />
		<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
```

**Note**: The existing `<meta name="viewport" content="width=device-width" />` (line 7) is replaced with the complete viewport tag that includes `initial-scale=1`.

**Canonical URL**: Assumes deployment to `https://diegosasco.github.io`. If the actual URL differs, update accordingly.

#### 2.4.2 Add Schema.org JSON-LD (REQ-C12)

**Location**: Inside `<head>`, after the OG tags, before `</head>`.

```astro
		<script type="application/ld+json" set:html={JSON.stringify({
			"@context": "https://schema.org",
			"@type": "Person",
			"name": personalInfo.name,
			"jobTitle": personalInfo.title,
			"url": `https://diegosasco.github.io${base}`,
			"knowsAbout": [
				"C#/.NET", "Go", "gRPC", "IoT", "RabbitMQ",
				"integración de sistemas legacy", "automatización RPA",
				"seguridad informática", "desarrollo de software"
			]
		})} />
```

#### 2.4.3 Add About Me Section (REQ-C1)

**Location**: Between the `</header>` closing tag and the `{/* Experiencia Laboral */}` comment (after line 44, before line 46).

```astro
			{/* About Me */}
			{personalInfo.about_me && (
				<section class="about-me">
					<h2>Sobre Mí</h2>
					<p class="about-me-text">{personalInfo.about_me}</p>
				</section>
			)}
```

**Conditional rendering**: The `personalInfo.about_me &&` guard ensures backward compatibility — if the field is missing or empty, the section is omitted entirely.

#### 2.4.4 Add ARIA Roles (REQ-C15)

**Location**: Various elements in the template.

| Element | Current | Target |
|---------|---------|--------|
| `<header class="hero">` (line 13) | `<header class="hero">` | `<header class="hero" role="banner">` |
| `<main>` (line 12) | `<main>` | `<main role="main">` |
| Recommendations `<section>` (line ~103) | `<section class="section-card">` | `<section class="section-card" role="complementary">` |
| Contact SVG (line ~22) | `<svg xmlns="..." viewBox="...">` | `<svg xmlns="..." viewBox="..." aria-label="Email">` |

---

### 2.5 `web/src/styles/global.css` — Style Changes

#### 2.5.1 About Me Section Styles

**Location**: After `.contact-form-btn svg` block (after line ~121), before `/* Secciones genéricas */`.

```css
/* About Me Section */
.about-me {
	text-align: center;
	padding: 1.5rem 0;
	border-block-end: 1px solid var(--card-border);
}

.about-me h2 {
	font-family: var(--font-title);
	font-size: 1.4rem;
	font-weight: 700;
	color: var(--text-main);
	margin-block-end: 0.75rem;
	border: none;
	padding: 0;
}

.about-me-text {
	color: var(--text-muted);
	font-size: 0.95rem;
	line-height: 1.7;
	max-inline-size: 720px;
	margin-inline: auto;
}
```

#### 2.5.2 Focus Visible Styles (REQ-C14)

**Location**: After the About Me styles, before `/* Secciones genéricas */`.

```css
/* Accessibility: Focus Visible */
a:focus-visible,
button:focus-visible,
.contact-form-btn:focus-visible {
	outline: 2px solid var(--primary-color);
	outline-offset: 2px;
	border-radius: 4px;
}
```

#### 2.5.3 Print Layout — Single Column (REQ-C7, REQ-C8)

**Location**: Inside `@media print` block.

**Skills grid** — find the existing print rule (approximately line ~296):
```css
	.skills-grid {
		grid-template-columns: repeat(3, 1fr);
		gap: 1rem;
	}
```

**Replace with**:
```css
	.skills-grid {
		grid-template-columns: 1fr;
		gap: 1rem;
	}
```

**Recommendations grid** — find (approximately line ~340):
```css
	.recommendations-grid {
		grid-template-columns: repeat(2, 1fr);
		gap: 1rem;
	}
```

**Replace with**:
```css
	.recommendations-grid {
		grid-template-columns: 1fr;
		gap: 1rem;
	}
```

#### 2.5.4 Print Layout — 0.75" Margins (REQ-C9)

**Location**: Inside `@media print`, find the `body` rule.

**Current**:
```css
	body {
		background-color: #ffffff;
		color: #0f172a;
		padding-block: 0;
		padding-inline: 0;
		font-size: 10pt;
		line-height: 1.4;
	}
```

**Target**:
```css
	body {
		background-color: #ffffff;
		color: #0f172a;
		padding-block: 0;
		padding-inline: 0;
		font-size: 10pt;
		line-height: 1.4;
	}

	@page {
		margin: 0.75in;
	}
```

**Note**: The `@page { margin }` rule is the correct mechanism for PDF page margins. The `body` padding remains 0 since `@page` handles the margin. If Puppeteer doesn't respect `@page`, add `padding: 0.75in` to `body` as fallback.

#### 2.5.5 Print Layout — ATS Fonts (REQ-C10)

**Location**: Inside `@media print`, at the top of the block (after `:root` overrides).

**Insert** after the existing `:root` print overrides:
```css
	:root {
		/* ... existing print overrides ... */
		--font-title: "Georgia", serif;
		--font-body: "Arial", sans-serif;
	}
```

**Current print `:root` already overrides colors. Add the two font variables to the same block.**

#### 2.5.6 Print — About Me Section

**Location**: Inside `@media print`, add after the `.hero` print styles:

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

## 3. Data Flow — `about_me` Field Lifecycle

This section traces the new `about_me` field through every stage of the pipeline.

```
┌─────────────────────────────────────────────────────────────────────┐
│ Stage 1: YAML Source                                               │
│ File: data/cv-source.yaml                                         │
│ Key: personal_info.about_me                                        │
│ Value: "Desarrollador de Software Senior con más de 7 años..."    │
│ YAML tag: `yaml:"about_me"`                                       │
└──────────────────────────┬──────────────────────────────────────────┘
                           │ Go yaml.Unmarshal
                           ▼
┌─────────────────────────────────────────────────────────────────────┐
│ Stage 2: Go Struct                                                 │
│ File: pipeline/models.go                                          │
│ Struct: PersonalInfo.AboutMe string                               │
│ Populated by YAML unmarshaling via `yaml:"about_me"` tag          │
└──────────────────────────┬──────────────────────────────────────────┘
                           │ ProcessCV() pass-through
                           ▼
┌─────────────────────────────────────────────────────────────────────┐
│ Stage 3: Processing                                                │
│ File: pipeline/calculator.go                                      │
│ Line: PersonalInfo: PersonalInfo{                                 │
│         Name: ..., Title: ..., AboutMe: raw.PersonalInfo.AboutMe  │
│       }                                                            │
│ No transformation — pure pass-through.                             │
└──────────────────────────┬──────────────────────────────────────────┘
                           │ json.Marshal
                           ▼
┌─────────────────────────────────────────────────────────────────────┐
│ Stage 4: JSON Output                                               │
│ File: web/src/data/cv-processed.json                              │
│ Key: personal_info.about_me                                       │
│ JSON tag: `json:"about_me,omitempty"`                             │
│ Value: Same string as YAML source                                 │
└──────────────────────────┬──────────────────────────────────────────┘
                           │ Astro import
                           ▼
┌─────────────────────────────────────────────────────────────────────┐
│ Stage 5: Astro Template                                            │
│ File: web/src/pages/index.astro                                   │
│ Variable: cvData.personal_info.about_me                           │
│ Alias: personalInfo.about_me (via destructuring)                  │
│ Rendered: Conditional <section> after hero, before experience     │
│ Guard: personalInfo.about_me && (...)                             │
└──────────────────────────┬──────────────────────────────────────────┘
                           │ Astro build
                           ▼
┌─────────────────────────────────────────────────────────────────────┐
│ Stage 6: HTML Output                                               │
│ Element: <section class="about-me">                               │
│ Contains: <h2>Sobre Mí</h2> + <p class="about-me-text">...</p>   │
│ Position: Between hero and work experience sections                │
└─────────────────────────────────────────────────────────────────────┘
```

### Parallel Data Flow: `hidden` Field

```
YAML: project.hidden: true
  │
  ├─► models.go: Project.Hidden bool (yaml:"hidden", json:"-")
  │
  ├─► calculator.go: Tech intervals registered REGARDLESS of hidden
  │   └─► filteredProjects: only !Hidden projects added to output
  │
  └─► JSON output: Hidden flag is json:"-", never appears in output
      └─► Template: Simply never receives hidden projects
```

---

## 4. Implementation Order

Dependencies flow top-to-bottom. Each step assumes the previous is complete.

```
Step 1: Go Pipeline — Struct & Logic Changes
├── 1a. models.go: Add AboutMe to PersonalInfo
├── 1b. models.go: Add Hidden to Project
├── 1c. calculator.go: Pass through AboutMe
├── 1d. calculator.go: Filter hidden projects from rendered output
├── 1e. calculator.go: Register tech intervals for hidden projects
└── 1f. Verify: go test ./... passes (existing tests unchanged)

Step 2: YAML Content — All Content Changes
├── 2a. personal_info: Add about_me, change title
├── 2b. Skills: Remove customer_service entry
├── 2c. Real2B: Reframe role, description, technologies
├── 2d. PIXI: Reframe role, description, technologies
├── 2e. All work descriptions: Apply CAR qualitative
├── 2f. Projects: Add hidden: true to 7 projects
├── 2g. Project 3: Rename to "Plataforma de Monitoreo Remoto de Activos IoT"
├── 2h. Visible project descriptions: Apply CAR qualitative
└── 2i. Education: Contextualize Contador Público status

Step 3: Pipeline Re-execution
├── 3a. Run Go pipeline to regenerate cv-processed.json
├── 3b. Verify about_me appears in personal_info
├── 3c. Verify exactly 5 projects visible (3 in Desarrollo Indep., 2 in Bionico)
├── 3d. Verify hidden projects' skills still have correct durations
├── 3e. Verify customer_service is absent from output
└── 3f. Verify Real2B and PIXI show new roles/descriptions

Step 4: Template — index.astro
├── 4a. Fix viewport meta tag
├── 4b. Add SEO meta tags (description, canonical, OG)
├── 4c. Add Schema.org JSON-LD
├── 4d. Add preconnect links
├── 4e. Add About Me section (conditional)
├── 4f. Add ARIA roles (banner, main, complementary, aria-label)
└── 4g. Verify: pnpm build succeeds

Step 5: Styles — global.css
├── 5a. Add About Me section styles
├── 5b. Add focus-visible styles
├── 5c. Print: skills grid → 1fr
├── 5d. Print: recommendations grid → 1fr
├── 5e. Print: Add @page margin 0.75in
├── 5f. Print: Override font variables for ATS
└── 5g. Print: Add About Me print styles

Step 6: Verification
├── 6a. go test ./... (Go unit tests pass)
├── 6b. pnpm build (Astro build succeeds)
├── 6c. pnpm dev → Manual visual verification
│   ├── About Me section visible between hero and experience
│   ├── 5 projects visible (3 under Desarrollo Indep., 2 under Bionico)
│   ├── No "Confidencial" in project names
│   ├── Real2B shows "Soporte Técnico y Operaciones de TI"
│   ├── PIXI shows "Supervisor de Operaciones y Logística"
│   ├── Title shows "Desarrollador de Software Senior"
│   ├── No customer_service skill in Knowledge section
│   └── ARIA roles present in DOM
├── 6d. Print preview verification
│   ├── Skills in single column
│   ├── Recommendations in single column
│   ├── 0.75" margins
│   ├── Georgia/Georgia fonts applied
│   ├── About Me visible but compact
│   └── Contact button hidden
└── 6e. Schema.org validation (Google Rich Results Test)
```

---

## 5. Edge Cases

### 5.1 Skill Duration Recalculation After Project Curation

**Risk**: Hiding projects removes their rendered output but MUST NOT remove their contribution to skill duration calculations.

**Current intervals at risk**:

| Skill ID | Projects Contributing Intervals | Hidden Projects Affecting It | Risk |
|----------|--------------------------------|------------------------------|------|
| `python` | `ai_engineering_workflow` (24mo), `sasco_agency_ai` (12mo) | Both hidden | **HIGH** — Python would drop to 0 months if hidden projects' intervals aren't registered. Both are associated projects (in work entries), so intervals ARE registered in the work experience loop. ✅ Safe. |
| `docker` | `ai_engineering_workflow` (24mo) | Hidden | **HIGH** — Docker would drop to 0. Same as above — associated project, interval registered. ✅ Safe. |
| `javascript` | `sasco_agency_ai` (12mo) | Hidden | **HIGH** — JavaScript would drop to 0. Same as above. ✅ Safe. |
| `generative_ai` | `ai_engineering_workflow` (24mo), `sasco_agency_ai` (12mo) | Both hidden | Same pattern. ✅ Safe. |
| `ai_assisted_dev` | `ai_engineering_workflow` (24mo) | Hidden | Same pattern. ✅ Safe. |
| `astro` | `cv_pipeline` (1mo) | Hidden | Drops to 0. ✅ Safe (correct behavior — this is a meta-project). |
| `sql` | `normalizacion_inventario` (12mo) + work-level intervals | Hidden | SQL has multiple work-level intervals (Sluckis work-level, Bionico work-level, TopBrands work-level). Losing project interval is fine since work-level covers it. ✅ Safe. |

**Verification step (3d)**: After pipeline re-execution, manually check that these skills still have non-zero durations in `cv-processed.json`:
- `python`: Should remain ~36 months (2 overlapping intervals merged)
- `docker`: Should remain ~24 months
- `javascript`: Should remain ~12 months
- `generative_ai`: Should remain ~36 months
- `ai_assisted_dev`: Should remain ~24 months

### 5.2 Work Entries With All Projects Hidden

**Affected entries**:
- **Sluckis Hermanos SA**: Only project `normalizacion_inventario` is hidden. After hiding, this entry renders with NO sub-projects, only technologies at the work level. The template already handles this case (`{(!work.projects \|\| work.projects.length === 0) && ...}` renders tech tags at the work level).
- **TopBrands**: Both `retail_automation` and `opt_administrativa` are hidden. Same behavior as Sluckis — shows work-level technologies without sub-projects.

**Template behavior**: The Astro template has two rendering branches for each work item:
1. If `work.projects.length > 0` → renders project sub-items
2. If `work.projects.length === 0` AND `work.technologies.length > 0` → renders flat tech tags

After hiding all projects for Sluckis/TopBrands, `processedProjects` will be an empty slice. The `len(work.Projects) == 0` check in `calculator.go` determines whether tech intervals use work-level or project-level. Since `work.Projects` (the YAML string array) still contains the project IDs (they're just hidden), the pipeline logic will still process them as "having projects" for interval purposes.

**Wait — important subtlety**: In `calculator.go`, the check `if len(work.Projects) == 0` uses the YAML `work.Projects` array (project ID strings), NOT the `processedProjects` array. Even after hiding, the YAML array still has IDs. So:
- Hidden projects' tech intervals are still registered ✅
- Work-level techs that are ALSO in projects are still excluded from work-level intervals ✅
- But the rendered `processedProjects` will be empty, and the template will show no projects ✅

Actually, there's a problem. The template checks `work.projects && work.projects.length > 0`. After hiding, `processedProjects` could be empty `[]ProjectProcessed{}`. In JSON, this serializes as `[]` not `null`. So `work.projects.length > 0` would be `false` for empty arrays — the template would render the fallback tech tags branch. This is correct behavior.

But wait — if ALL projects in a work entry are hidden, the `processedProjects` variable will be `nil` (since `append` to a nil slice with no appends remains nil). In JSON, nil serializes as `null`. The template check `work.projects && work.projects.length > 0` would be `null && ...` which is falsy → falls through to tech tags. ✅ Correct.

### 5.3 Empty `processedProjects` vs `null`

**Current behavior**: In Go, `var processedProjects []ProjectProcessed` starts as `nil`. If no projects are appended (all hidden), it remains `nil`, which serializes to JSON `null`.

**Template guard**: `{work.projects && work.projects.length > 0}` handles both `null` and `[]` correctly — both are falsy for the second condition.

**No code change needed** for this edge case.

### 5.4 `customer_service` Skill Orphaning

After removing `customer_service` from both Real2B and PIXI technologies and removing the skill definition from YAML:

- The `techIntervals` map will have no entries for `customer_service`.
- The skill won't be in `raw.Skills`, so no `SkillProcessed` entry is created.
- **Result**: `customer_service` disappears entirely from output. ✅ Correct.

### 5.5 About Me Word Count Validation

**Current**: 75 words. Requirement: ≤80 words.

**Validation**: No automated validation is implemented. The word count is enforced at content-authoring time. The pipeline does not validate word count — it simply passes the string through.

**Rationale**: Adding a word count validator to the Go pipeline would be over-engineering for a single static field. The content is author-controlled and locked.

### 5.6 Title Change Propagation

The `personal_info.title` change from "Arquitecto de Software & Especialista en Automatización / IA" to "Desarrollador de Software Senior" propagates to:

| Consumer | Field | Current | Target |
|----------|-------|---------|--------|
| Hero subtitle | `<p class="subtitle">` | Arquitecto de Software... | Desarrollador de Software Senior |
| Page title | `<title>` | Diego Sasco \| Arquitecto... | Diego Sasco \| Desarrollador de Software Senior |
| OG title | `og:title` | (new) | Diego Sasco \| Desarrollador de Software Senior |
| Schema.org | `jobTitle` | (new) | Desarrollador de Software Senior |
| Meta description | N/A — hardcoded in template | (new) | References the title inline |

---

## 6. Test Strategy

### 6.1 Go Unit Tests (`pipeline/calculator_test.go`)

#### 6.1.1 Existing Tests — Must Pass Unchanged

The existing `TestProcessCV`, `TestCalculateMonths`, and `TestMergeIntervals` must continue passing after all code changes. **No modifications** to existing test functions.

#### 6.1.2 New Test: `TestAboutMePassThrough`

**Purpose**: Verify that `about_me` flows from raw to processed without modification.

```go
func TestAboutMePassThrough(t *testing.T) {
	raw := &CVRaw{
		PersonalInfo: PersonalInfo{
			Name:    "Test",
			Title:   "Dev",
			AboutMe: "Short bio for testing purposes.",
		},
	}
	processed := ProcessCV(raw)
	if processed.PersonalInfo.AboutMe != "Short bio for testing purposes." {
		t.Errorf("AboutMe not passed through. Got %q", processed.PersonalInfo.AboutMe)
	}
}
```

#### 6.1.3 New Test: `TestAboutMeEmptyString`

**Purpose**: Verify empty `about_me` doesn't break the pipeline.

```go
func TestAboutMeEmptyString(t *testing.T) {
	raw := &CVRaw{
		PersonalInfo: PersonalInfo{Name: "Test", Title: "Dev"},
	}
	processed := ProcessCV(raw)
	if processed.PersonalInfo.AboutMe != "" {
		t.Errorf("Expected empty AboutMe, got %q", processed.PersonalInfo.AboutMe)
	}
}
```

#### 6.1.4 New Test: `TestHiddenProjectFiltered`

**Purpose**: Verify hidden projects are excluded from rendered output but still contribute to skill durations.

```go
func TestHiddenProjectFiltered(t *testing.T) {
	raw := &CVRaw{
		PersonalInfo: PersonalInfo{Name: "Test"},
		Skills:       []SkillRaw{{ID: "go", Name: "Go", Category: "Lang"}},
		WorkExperience: []WorkExperience{
			{
				Company:      "Co",
				Role:         "Dev",
				StartDate:    CVDate{Time: parseDate("2020-01-01")},
				EndDate:      CVDate{Time: parseDate("2020-12-31")},
				Technologies: []string{"go"},
				Projects:     []string{"proj-visible", "proj-hidden"},
			},
		},
		Projects: []Project{
			{
				ID: "proj-visible", Name: "Visible",
				StartDate: CVDate{Time: parseDate("2020-01-01")},
				EndDate:   CVDate{Time: parseDate("2020-06-30")},
				Technologies: []string{"go"},
			},
			{
				ID: "proj-hidden", Name: "Hidden", Hidden: true,
				StartDate: CVDate{Time: parseDate("2020-07-01")},
				EndDate:   CVDate{Time: parseDate("2020-12-31")},
				Technologies: []string{"go"},
			},
		},
	}

	processed := ProcessCV(raw)

	// Hidden project should NOT appear in rendered projects
	if len(processed.WorkExperience[0].Projects) != 1 {
		t.Fatalf("Expected 1 visible project, got %d", len(processed.WorkExperience[0].Projects))
	}
	if processed.WorkExperience[0].Projects[0].Name != "Visible" {
		t.Errorf("Expected 'Visible', got %q", processed.WorkExperience[0].Projects[0].Name)
	}

	// BUT skill duration should include both projects (6 + 6 = 12 months)
	if processed.Skills[0].MonthsExperience != 12 {
		t.Errorf("Expected 12 months for 'go', got %d", processed.Skills[0].MonthsExperience)
	}
}
```

#### 6.1.5 New Test: `TestHiddenIndependentProject`

**Purpose**: Verify hidden independent projects are not rendered.

```go
func TestHiddenIndependentProject(t *testing.T) {
	raw := &CVRaw{
		PersonalInfo: PersonalInfo{Name: "Test"},
		Skills:       []SkillRaw{{ID: "go", Name: "Go", Category: "Lang"}},
		WorkExperience: []WorkExperience{
			{
				Company:   "Co", Role: "Dev",
				StartDate: CVDate{Time: parseDate("2020-01-01")},
				EndDate:   CVDate{Time: parseDate("2020-12-31")},
			},
		},
		Projects: []Project{
			{
				ID: "proj-indie", Name: "Indie", Hidden: true,
				StartDate:    CVDate{Time: parseDate("2020-01-01")},
				EndDate:      CVDate{Time: parseDate("2020-06-30")},
				Technologies: []string{"go"},
			},
		},
	}

	processed := ProcessCV(raw)

	// Hidden independent project should NOT appear in processed.Projects
	if len(processed.Projects) != 0 {
		t.Errorf("Expected 0 independent projects, got %d", len(processed.Projects))
	}

	// But tech intervals should still be registered (6 months)
	if processed.Skills[0].MonthsExperience != 6 {
		t.Errorf("Expected 6 months for 'go', got %d", processed.Skills[0].MonthsExperience)
	}
}
```

### 6.2 Web Layer — Manual Verification

No automated web tests exist (per config.yaml: `web.command: null`). All web verification is manual.

#### 6.2.1 Visual Verification Checklist (`pnpm dev`)

| # | Check | Pass Criteria |
|---|-------|---------------|
| V1 | About Me section visible | Section appears between hero and work experience |
| V2 | About Me text content | Matches the approved 75-word text exactly |
| V3 | About Me title | "Desarrollador de Software Senior" in hero subtitle |
| V4 | Real2B role | Shows "Soporte Técnico y Operaciones de TI" |
| V5 | Real2B description | Starts with "Proporcioné soporte técnico..." |
| V6 | PIXI role | Shows "Supervisor de Operaciones y Logística" |
| V7 | Project count | Exactly 3 projects under Desarrollo Independiente, 2 under Bionico |
| V8 | Project 3 name | "Plataforma de Monitoreo Remoto de Activos IoT" (not "Confidencial") |
| V9 | No hidden projects | ai_engineering_workflow, cv_pipeline, etc. NOT visible |
| V10 | No customer_service skill | "Atención al Público" absent from Skills section |
| V11 | Contador Público status | Shows "(formación complementaria en gestión financiera y contable)" |
| V12 | CAR verbs | All descriptions use past-tense action verbs (not passive nouns) |

#### 6.2.2 Print/PDF Verification Checklist (`pnpm dev` → Print Preview)

| # | Check | Pass Criteria |
|---|-------|---------------|
| P1 | Skills layout | Single column (not 3-column grid) |
| P2 | Recommendations layout | Single column (not 2-column grid) |
| P3 | Page margins | Content starts ~0.75" from edges |
| P4 | Fonts | Georgia for titles, Arial for body text |
| P5 | Contact button hidden | Google Form button not visible |
| P6 | About Me visible | Compact but readable in print |
| P7 | Dark theme removed | White background, black text |

#### 6.2.3 SEO/Accessibility Verification

| # | Check | Verification Method |
|---|-------|---------------------|
| S1 | Meta description present | View page source → find `<meta name="description">` |
| S2 | OG tags present | View source → find `og:title`, `og:description`, `og:type`, `og:url`, `og:locale` |
| S3 | Canonical URL | View source → find `<link rel="canonical">` |
| S4 | Schema.org valid | Paste URL into Google Rich Results Test |
| S5 | Preconnect | View source → find preconnect links BEFORE font import |
| S6 | ARIA roles | Inspect DOM → header has `role="banner"`, main has `role="main"` |
| S7 | SVG aria-label | Contact SVG has `aria-label="Email"` |
| S8 | Focus visible | Tab through page → outline visible on interactive elements |

### 6.3 Build Verification

```bash
# Go pipeline tests
cd /src/cv-pipeline/pipeline && go test ./...

# Go pipeline coverage (informational)
cd /src/cv-pipeline/pipeline && go test -cover ./...

# Rebuild processed data
cd /src/cv-pipeline/pipeline && go run main.go

# Astro build
cd /src/cv-pipeline/web && pnpm build
```

All three commands must succeed with zero errors.

---

## 7. Change Summary by File

| File | Lines Changed (est.) | Type | Risk |
|------|---------------------|------|------|
| `data/cv-source.yaml` | ~65–80 | Content rewrite | Low — YAML only |
| `pipeline/models.go` | ~3 lines added | Struct fields | Low — additive |
| `pipeline/calculator.go` | ~12–15 lines modified | Logic + passthrough | Medium — filter logic |
| `pipeline/calculator_test.go` | ~65 lines added | New test cases | Low — additive |
| `web/src/pages/index.astro` | ~35–40 lines added | Template + SEO + ARIA | Low — additive |
| `web/src/styles/global.css` | ~25–30 lines added/modified | Styles + print | Low — CSS only |
| **Total** | **~205–233** | | **Within 400-line budget** |

---

## 8. Design Decisions Log

| Decision | Rationale | Alternative Rejected |
|----------|-----------|---------------------|
| `hidden: true` on projects instead of deletion | Preserves skill duration intervals and data completeness | Deleting projects from YAML would lose interval data for Python, Docker, JS |
| `json:"-"` on Hidden field | Hidden is a processing directive, not output data | Including it in JSON would leak implementation details to template |
| `omitempty` on AboutMe JSON tag | Backward compatible — missing field produces clean JSON | Required field would break if YAML is temporarily missing about_me |
| `@page { margin }` for print margins | Correct CSS mechanism for PDF page margins | `body { padding }` as primary is hacky; used as fallback only |
| Remove `customer_service` skill entirely | No work entry references it after reframe → would show "Sin experiencia" | Keeping with 0 experience looks broken |
| Conditional About Me rendering (`&&` guard) | Graceful degradation if field is missing | Hard render would show empty section on missing field |
| Keep project IDs in `work.Projects` arrays even for hidden projects | Pipeline needs to process them for tech intervals | Moving hidden projects to a separate list would duplicate data |

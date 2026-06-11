# Propuesta SDD — cv-web-optimization

**Change ID**: `cv-web-optimization`
**Fase**: Propuesta
**Fecha**: 2026-06-10
**Estado**: Pendiente de aprobación

---

## 1. Declaración del Problema

El CV web de Diego Sasco presenta una base técnica sólida —Astro como generador estático, pipeline de datos en Go con cálculo de experiencia sin solapamientos, soporte dual web/PDF desde un único YAML— pero exhibe deficiencias significativas cuando se evalúa contra los criterios de optimización ATS, escaneo ocular de 7.4 segundos y señalización de credibilidad técnica que establece el informe de referencia (`docs/informe.md`).

El diagnóstico completo (`docs/diagnostico-cv-web.md`) identifica **18 recomendaciones** agrupadas en 4 dimensiones. Los problemas críticos son:

1. **Ausencia de canales de contacto directo** — Solo un Google Form anónimo. Sin email, LinkedIn ni GitHub. El usuario ha decidido mantener Google Form como único canal (decisión LOCKED).

2. **Empleos puente sin reformulación** — Real2B (Help Desk) y PIXI Supermercados (Operaciones) son roles de baja señal técnica que penalizan la tasa de callback según evidencia del NBER. El usuario ha decidido mantener ambos con reframing técnico (decisión LOCKED).

3. **Layout multi-columna en PDF** — Skills en grid de 3 columnas y recomendaciones en 2 columnas violan la regla de columna única para ATS.

4. **Cero métricas cuantificables** — Ningún proyecto ni experiencia incluye resultados medibles. El usuario no dispone de métricas duras (decisión LOCKED), por lo que se usará CAR cualitativo con verbos de acción.

5. **Sección "About Me" inexistente** — Violación directa de la estructura de e-portfolio prescrita (máximo 80 palabras).

6. **Descripciones en lenguaje pasivo** — Todas usan sustantivos abstractos ("Diseño y desarrollo de...", "Gestión integral de...") en lugar de verbos de acción en pasado.

7. **Sin meta tags SEO, Open Graph ni Schema.org** — El sitio es prácticamente invisible para motores de búsqueda y no genera preview cards.

8. **Fuentes no ATS-compatibles en PDF** — Inter y Outfit no están entre las fuentes estándar prescritas por Harvard/MIT/Stanford.

9. **12 proyectos sin curaduría** — Excede el límite de 5 flagship del informe.

10. **Recomendaciones sin atribución** — Placeholders `<RECOMMENDER_X>`. El usuario ha decidido mantener quotes sin nombres de autor (decisión LOCKED).

---

## 2. Solución Propuesta

### 2.1 Área de Trabajo: Contenido y Narrativa

#### 2.1.1 Sección "About Me" (Nueva)

**Qué**: Agregar un campo `about_me` al YAML fuente, procesarlo en Go, y renderizarlo como sección narrativa entre el hero y la experiencia laboral.

**Archivos**: `cv-source.yaml` + `models.go` (campo nuevo en `CVRaw`) + `calculator.go` (passthrough a `CVProcessed`) + `index.astro` (nueva sección HTML) + `global.css` (estilos).

**Borrador propuesto** (~75 palabras):

> Desarrollador de Software Senior con más de 7 años de experiencia en el diseño y automatización de sistemas robustos para el sector financiero, retail e industrial. Especializado en integración de sistemas legacy, arquitectura IoT y soluciones RPA, combinando competencias en backend (C#/.NET, Go, gRPC) con una formación en seguridad informática (Hacking Ético). Profesional orientado a resultados, con comunicación asertiva y enfoque en la confiabilidad de sistemas críticos.

**Fundamentos del texto**:
- "Desarrollador de Software Senior con más de 7 años" — del perfil técnico (`CV_Diego_Sasco_Perfil_Tecnico.txt`), ajustado según decisión del usuario.
- "sector financiero, retail e industrial" — mapeo directo de proyectos: GetNet/Santander (financiero), TopBrands/Neutral (retail), Golden Wash/IoT (industrial).
- "integración de sistemas legacy, arquitectura IoT y soluciones RPA" — señalización de las 3 áreas técnicas diferenciadoras verificables en proyectos.
- "comunicación asertiva" y "seguridad informática" — de las recomendaciones y la certificación en Hacking Ético.
- "confiabilidad de sistemas críticos" — refuerza la señal de infraestructura/seguridad sin inventar datos.

**Criterio**: ≤80 palabras, sin pronombres de primera persona (cumple restricción del informe), verbo inicial de identidad profesional, cierra con valor diferencial.

#### 2.1.2 Reframing de Empleos Puente

**Real2B** — Reframe de "Analista Operacional & Control de Flujo" → Señal de soporte técnico + infraestructura.

Estrategia:
- Eliminar menciones a "transacciones de caja" y "atención al público" (señales de baja cualificación).
- Reenfocar hacia: soporte técnico (Help Desk) a usuarios, resolución de incidencias de infraestructura en múltiples locales, administración de sistemas de comunicación.
- Renombrar rol de `cv-source.yaml` a algo como "Soporte Técnico y Operaciones de TI".
- Reescribir descripción con verbos de acción en pasado y marco CAR cualitativo.
- Reemplazar tecnologías blandas (Help Desk, Atención al Público) por señales técnicas: Linux, networking, administración de sistemas, o eliminar skills de baja señal y mantener solo las transferibles.

**PIXI Supermercados** — Reframe de "Encargado de Operaciones" → Señal de liderazgo de equipos + optimización de procesos.

Estrategia:
- Eliminar "arqueos de caja" y "atención al público" (señales de baja cualificación).
- Reenfocar hacia: supervisión de equipos operativos, coordinación logística, gestión de inventarios, resolución de conflictos.
- Mantener skills transferibles: Liderazgo, Trabajo en Equipo, Comunicación Asertiva.
- Agregar señal de gestión de procesos: administración de stock, control de proveedores.
- Reescribir descripción enfatizando la dimensión de gestión y liderazgo operativo.

#### 2.1.3 Descripciones con Marco CAR/CCCR Cualitativo

Dado que el usuario no dispone de métricas cuantificables (decisión LOCKED), se aplica una adaptación cualitativa del marco CAR:

**Estructura adaptada (CAR Cualitativo)**:
- **Contexto (C)**: Situación operativa o problema de negocio.
- **Acción (A)**: Verbo de acción en pasado + tecnología específica + decisión de diseño.
- **Resultado (R)**: Impacto cualitativo descrito con verbos de resultado (consolidó, optimizó, redujo, estandarizó) + beneficio de negocio (sin porcentajes inventados).

**Ejemplo de reframing**:
- *Actual*: "Diseño y desarrollo de sistemas distribuidos, software a medida y soluciones de automatización IoT..."
- *Propuesto*: "Diseñé e implementé sistemas distribuidos en Go y C#/.NET para automatización IoT, integrando protocolos seriales (GPIO, gRPC) con hardware embebido, lo que permitió el control remoto de dispositivos industriales en entornos de producción."

**Taxonomía de verbos de acción** (del informe, categorías aplicables):
- Desarrollo: diseñó, programó, construyó, optimizó, integró, desplegó
- Liderazgo: dirigió, coordinó, encabezó, administró
- Gestión: automatizó, implementó, simplificó, estandarizó
- Investigación: analizó, diagnosticó, validó

**Alcance**: Reformular las 6 descripciones de experiencia + descripciones de proyectos que se mantengan tras la curaduría.

#### 2.1.4 Curaduría de Proyectos (12 → 5 Flagship)

Reducir de 12 proyectos a 5 flagship que maximicen la señal técnica diferenciadora:

| # | Proyecto | Empresa | Señal Técnica | Justificación |
|---|----------|---------|---------------|---------------|
| 1 | RpPOS IoT | Freelance | Go + Raspberry Pi + GPIO + Fyne UI | Sistemas embebidos, hardware limitado |
| 2 | Proxy Inverso Financiero | Freelance | C# + SOAP→gRPC + pasarela financiera | Integración legacy, protocolos |
| 3 | Plataforma de Monitoreo Remoto de Activos IoT | Freelance | Go + gRPC + protocolos seriales + IoT | Hardware embebido + APIs distribuidas |
| 4 | Framework RPA Financiero | Bionico | C# + RabbitMQ + UI Automation | Automatización financiera, observabilidad |
| 5 | Motor de Extracción PDF | Bionico | C# + PdfPig + procesamiento masivo | Procesamiento documentos semi-estructurados |

**Proyectos eliminados del CV**:
- Flujo IA SDD/TDD (metodología, no producto entregable)
- CV Pipeline CLI (meta-proyecto, baja señal para reclutador externo)
- Normalización Inventario Sluckis (baja señal técnica)
- Soluciones IA Generativa (demasiado genérico sin demos)
- Automatización Retail TopBrands (señal moderada, pero menos diferenciadora)
- Optimización Flujos Administrativos TopBrands (baja señal técnica)
- Proyectos eliminados permanecen en YAML con `hidden: true` o se eliminan si el pipeline no soporta flag.

**Nota técnica**: Los proyectos eliminados afectan los meses de experiencia calculados para ciertas skills (el pipeline calcula desde intervalos de proyecto). Se debe re-ejecutar el pipeline y verificar que las duraciones de skills permanezcan correctas. Si un proyecto aporta intervalos únicos para una skill, considerar mover esa skill a technologies del work_experience en lugar del proyecto.

#### 2.1.5 Contextualización de Educación

- Reformular "Contador Público (1er año cursado)" como complemento a la narrativa de perfil que cruza negocio y tecnología: "Contador Público — 1er año cursado (formación complementaria en gestión financiera y contable)" o similar.
- Considerar si mover esta entrada al final de la sección educación para reducir riesgo de clasificación errónea.

### 2.2 Área de Trabajo: ATS / Layout

#### 2.2.1 Skills Grid → Columna Única en Print

**Qué**: Cambiar `grid-template-columns: repeat(3, 1fr)` a `grid-template-columns: 1fr` en `@media print` para `.skills-grid`.

**Archivo**: `global.css` (~1 línea).

**Web**: Se mantiene el grid auto-fill actual (aceptable para web).

#### 2.2.2 Recomendaciones → Columna Única en Print

**Qué**: Cambiar `grid-template-columns: repeat(2, 1fr)` a `grid-template-columns: 1fr` en `@media print` para `.recommendations-grid`.

**Archivo**: `global.css` (~1 línea).

#### 2.2.3 Márgenes de 0.75" en PDF

**Qué**: Agregar `padding: 0.75in` al body en `@media print` y ajustar los elementos internos para compensar.

**Archivo**: `global.css` (~2-3 líneas).

#### 2.2.4 Fuentes ATS-compatibles en Print

**Qué**: En `@media print`, sobrescribir `--font-body` y `--font-title` con fuentes ATS-safe: Georgia para títulos, Arial para body (o Calibri).

**Archivo**: `global.css` (~4 líneas).

### 2.3 Área de Trabajo: SEO y Descubribilidad

#### 2.3.1 Meta Tags y Open Graph

**Qué**: Agregar al `<head>` de `index.astro`:
- `<meta name="description" content="...">`
- `<meta name="viewport" content="width=device-width, initial-scale=1">` (fix incompleto actual)
- Open Graph: `og:title`, `og:description`, `og:type`, `og:url`, `og:locale`
- `<link rel="canonical">`

**Archivo**: `index.astro` (~12-15 líneas).

**Description propuesta**: "Diego Sasco — Arquitecto de Software y Especialista en Automatización e IA. Más de 7 años diseñando sistemas distribuidos, soluciones IoT e integración legacy para el sector financiero, retail e industrial."

#### 2.3.2 Schema.org JSON-LD

**Qué**: Agregar un `<script type="application/ld+json">` con tipo `Person` de Schema.org.

**Archivo**: `index.astro` (~15 líneas).

**Campos**: name, jobTitle, url, sameAs (GitHub/LinkedIn si se habilitan en futuro), knowsAbout (skills clave).

#### 2.3.3 Preconnect a Google Fonts

**Qué**: Agregar `<link rel="preconnect" href="https://fonts.googleapis.com">` y `<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>`.

**Archivo**: `index.astro` (~2 líneas).

### 2.4 Área de Trabajo: Accesibilidad

#### 2.4.1 Estilos `:focus-visible`

**Qué**: Agregar estilos de foco visible para enlaces y botones (navegación por teclado).

**Archivo**: `global.css` (~10-15 líneas).

#### 2.4.2 ARIA y Alt Text

**Qué**:
- Agregar `role="banner"` al `<header>`.
- Agregar `role="main"` al `<main>`.
- Agregar `aria-label` al SVG del ícono de contacto (`aria-label="Email"`).
- Agregar `role="complementary"` a la sección de recomendaciones.

**Archivo**: `index.astro` (~5-6 atributos).

---

## 3. Alcance

### 3.1 Dentro del Alcance (IN Scope)

| ID | Cambio | Prioridad | Archivos |
|----|--------|-----------|----------|
| C1 | Sección "About Me" (campo nuevo + renderizado) | P1 | `cv-source.yaml`, `models.go`, `calculator.go`, `index.astro`, `global.css` |
| C2 | Reframe Real2B con señal técnica | P0 | `cv-source.yaml` |
| C3 | Reframe PIXI con señal de liderazgo | P0 | `cv-source.yaml` |
| C4 | Descripciones CAR/CCCR cualitativo (6 experiencias + 5 proyectos) | P0 | `cv-source.yaml` |
| C5 | Curaduría de proyectos (12→5 flagship) | P1 | `cv-source.yaml` |
| C6 | Contextualización educación | P2 | `cv-source.yaml` |
| C7 | Skills columna única en print | P0 | `global.css` |
| C8 | Recomendaciones columna única en print | P0 | `global.css` |
| C9 | Márgenes 0.75" en print | P0 | `global.css` |
| C10 | Fuentes ATS-compatibles en print | P1 | `global.css` |
| C11 | Meta description + OG tags + viewport fix | P1 | `index.astro` |
| C12 | Schema.org JSON-LD (Person) | P1 | `index.astro` |
| C13 | Preconnect Google Fonts | P2 | `index.astro` |
| C14 | Estilos :focus-visible | P2 | `global.css` |
| C15 | ARIA roles + alt text SVG | P2 | `index.astro` |

### 3.2 Fuera del Alcance (OUT of Scope)

| Elemento | Razón |
|----------|-------|
| Agregar email/LinkedIn/GitHub al hero | Decisión LOCKED del usuario: contacto solo vía Google Form |
| Métricas cuantificables inventadas | Decisión LOCKED del usuario: no dispone de datos duros |
| Atribución de autoría en recomendaciones | Decisión LOCKED del usuario: mantener quotes sin nombres |
| Toggle light/dark mode | Prioridad P3, diferido a cambio futuro |
| URLs de demos/repositorios en proyectos | Requiere crear contenido externo (READMEs, deploys); diferido |
| Renombrar URLs de proyectos en template | El template ya soporta `proj.url` pero las URLs están vacías; diferido hasta que haya URLs reales |
| Eliminar skills de baja señal del YAML | Se aborda parcialmente via reframing de empleos puente, pero las skills funcionales (Help Desk, Atención al Público) se reevalúan caso a caso |
| Self-hosting de fuentes | Complejidad alta vs. beneficio marginal; preconnect suficiente |
| `robots.txt` / `sitemap.xml` | Diferido a cambio futuro de SEO avanzado |

---

## 4. Borrador "About Me" — Propuesta

### Texto Propuesto (~75 palabras)

> Desarrollador de Software Senior con más de 7 años de experiencia en el diseño y automatización de sistemas para el sector financiero, retail e industrial. Especializado en integración de sistemas legacy, arquitectura IoT y soluciones RPA, con competencias en backend distribuido (C#/.NET, Go, gRPC, RabbitMQ) y formación en seguridad informática. Profesional orientado a resultados, con comunicación asertiva y enfoque en la confiabilidad de sistemas críticos.

### Fundamentos del Texto

| Fragmento | Fuente |
|-----------|--------|
| "Desarrollador de Software Senior con más de 7 años" | Perfil técnico: "Arquitecto de Software y Experto en Automatización con más de 7 años de experiencia", ajustado según decisión del usuario |
| "sector financiero, retail e industrial" | Proyectos GetNet (financiero), TopBrands (retail), Golden Wash/IoT (industrial) |
| "integración de sistemas legacy, arquitectura IoT y soluciones RPA" | 3 áreas técnicas diferenciadoras verificables en proyectos |
| "C#/.NET, Go, gRPC, RabbitMQ" | Stack técnico dominante y verificable |
| "formación en seguridad informática" | Certificación Hacking Ético (Universidad ORT) |
| "comunicación asertiva" | Recomendación: "asertivo en la comunicación" |
| "confiabilidad de sistemas críticos" | Señal de infraestructura + seguridad, reforzada por recomendación sobre seguridad |

### Criterios Cumplidos

- ✅ ≤80 palabras (75 palabras)
- ✅ Sin pronombres de primera persona ("yo", "mi")
- ✅ Verbo inicial de identidad profesional ("Arquitecto de Software")
- ✅ Cierra con valor diferencial (confiabilidad + comunicación)
- ✅ Menciona tecnologías clave verificables
- ✅ No inventa datos ni métricas

---

## 5. Estrategia de Reescritura CAR/CCCR (Cualitativa)

### Adaptación sin Métricas Duras

Dado que el usuario no dispone de porcentajes, cantidades ni KPIs verificables, se adapta el marco CAR prescrito por el informe:

**Estructura CAR Cualitativo**:

1. **Contexto (C)**: Situación operativa, sistema afectado o problemática de negocio.
2. **Acción (A)**: Verbo de acción en pasado + tecnología específica + decisión de diseño o arquitectura.
3. **Resultado (R)**: Beneficio de negocio o impacto operativo descrito con verbos de resultado, sin inventar cifras.

**Verbos de resultado cualitativos** (adaptación de la taxonomía del informe):
- En lugar de "redujo un X%": "redujo significativamente", "optimizó", "simplificó"
- En lugar de "procesó Y facturas/día": "permitió el procesamiento masivo", "habilitó la automatización a escala"
- En lugar de "soportó Z usuarios": "garantizó la operación continua", "desplegó en producción estable"

**Patrón de viñeta resultante**:
> "[Verbo de acción] + [solución técnica] + [tecnología] + para + [contexto de negocio], lo que + [resultado cualitativo con verbo de impacto]."

### Ejemplos de Aplicación

| Experiencia | Antes (Pasivo) | Después (CAR Cualitativo) |
|-------------|----------------|---------------------------|
| Desarrollo Independiente | "Diseño y desarrollo de sistemas distribuidos..." | "Diseñé e implementé sistemas distribuidos en Go y C#/.NET para automatización IoT, integrando protocolos seriales con hardware embebido, lo que habilitó el control remoto de dispositivos industriales en entornos de producción." |
| Bionico | "Liderazgo técnico en el diseño de soluciones de automatización..." | "Dirigí el desarrollo de un framework RPA financiero en C#/.NET con observabilidad vía RabbitMQ, automatizando la inserción de datos en sistemas contables y reduciendo la intervención manual en procesos críticos." |
| Real2B (reframed) | "Gestión integral de operaciones de venta de servicios..." | "Proporcioné soporte técnico (Help Desk) a usuarios en múltiples locales, diagnosticando y resolviendo incidencias de infraestructura de TI y sistemas de comunicación, manteniendo la continuidad operativa del servicio." |

---

## 6. Reframing de Empleos Puente

### 6.1 Real2B — De "Analista Operacional" a Señal de Soporte Técnico + Infraestructura

**Diagnóstico del problema**: La descripción actual menciona "transacciones de caja" y "atención directa al público", que son señales de baja cualificación técnica según la taxonomía del NBER citada en el informe.

**Estrategia de reframing**:

| Elemento | Antes | Después |
|----------|-------|---------|
| **Rol** | "Analista Operacional & Control de Flujo" | "Soporte Técnico y Operaciones de TI" |
| **Descripción** | "Gestión integral de operaciones de venta de servicios, transacciones de caja, soporte técnico (Help Desk) a usuarios y atención directa al público en múltiples locales." | "Proporcioné soporte técnico (Help Desk) a usuarios en múltiples locales, diagnosticando y resolviendo incidencias de infraestructura de TI y sistemas de comunicación. Administré la continuidad operativa de sistemas transaccionales y coordiné la resolución de problemas técnicos con equipos internos." |
| **Tecnologías** | Help Desk, Atención al Público, Trabajo en Equipo, Comunicación Asertiva, Resolución de Problemas | Help Desk & Soporte Técnico, Resolución de Problemas, Trabajo en Equipo, Comunicación Asertiva, Cultura de Seguridad |

**Señal técnica recuperada**:
- "infraestructura de TI" → conexión con el perfil de infraestructura/backend.
- "sistemas transaccionales" → eufemismo técnico para lo que antes era "transacciones de caja".
- "diagnosticando y resolviendo incidencias" → verbo de acción + señal de troubleshooting.
- "Cultura de Seguridad" → skill transferible al perfil de seguridad (Hacking Ético).

**Lo que se elimina**: "Atención al Público" como skill (se reemplaza por señal más técnica), "transacciones de caja" como texto.

### 6.2 PIXI Supermercados — De "Encargado de Operaciones" a Señal de Liderazgo + Procesos

**Diagnóstico del problema**: La descripción actual menciona "arqueos de caja" y "atención al público", señales de baja cualificación.

**Estrategia de reframing**:

| Elemento | Antes | Después |
|----------|-------|---------|
| **Rol** | "Encargado de Operaciones" | "Supervisor de Operaciones y Logística" |
| **Descripción** | "Supervisión de equipos, administración de personal, control de inventario, arqueos de caja y atención al público (resolución de conflictos y gestión de proveedores)." | "Coordiné equipos operativos y supervisé la logística de inventarios en un entorno de alta rotación, implementando controles sistemáticos de stock y optimizando los procesos de recepción de mercadería. Gestión de proveedores y resolución de conflictos operativos." |
| **Tecnologías** | Atención al Público, Liderazgo, Trabajo en Equipo, Comunicación Asertiva | Liderazgo de Equipos, Trabajo en Equipo, Comunicación Asertiva, Resolución de Problemas |

**Señal técnica recuperada**:
- "Coordiné equipos operativos" → verbo de liderazgo + gestión.
- "logística de inventarios" → señal de gestión de procesos (mejor que "control de inventario").
- "controles sistemáticos de stock" → disciplina de procesos transferible.
- "optimizando los procesos de recepción" → verbo de mejora continua.

**Lo que se elimina**: "arqueos de caja" (reemplazado por "controles sistemáticos"), "Atención al Público" como skill (reemplazado por Resolución de Problemas).

---

## 7. Evaluación de Riesgos

| # | Riesgo | Probabilidad | Impacto | Mitigación |
|---|--------|--------------|---------|------------|
| R1 | El reframing de empleos puente resulta poco creíble ante un reclutador que verifique el rol real | Media | Alto | Los reframings se mantienen dentro de la verdad (no inventar responsabilidades). Se priorizan verbos precisos sobre exageración. El riesgo de inclusión sin reframing (callback penalizado) es mayor que el riesgo de reframing moderado. |
| R2 | Eliminar proyectos afecta las duraciones calculadas de skills por el pipeline | Alta | Medio | Re-ejecutar pipeline y verificar duraciones. Si una skill pierde meses significativos, mover la tecnología a `work_experience.technologies` del rol correspondiente para preservar el intervalo. |
| R3 | El "About Me" propuesto no refleja la voz del candidato | Media | Bajo | El texto es un borrador sujeto a aprobación. Se proporcionan los fundamentos de cada frase para que el usuario valide o ajuste. |
| R4 | Los cambios en Go (campo `about_me`) requieren recompilar el pipeline y puede haber regresiones | Baja | Medio | Los tests existentes (`calculator_test.go`) deben pasar. Agregar test para el campo nuevo antes de modificar código. |
| R5 | Las fuentes ATS en print (Arial/Georgia) se ven significativamente diferentes en web | Baja | Bajo | Solo se aplican en `@media print`. La web mantiene Inter/Outfit. |
| R6 | El budget de 400 líneas changed por PR podría excederse si todas las descripciones se reescriben en detalle | Media | Medio | Estimación actual: 138-205 líneas (exploración). Monitorear durante implementación. Si se excede, priorizar P0s y diferir P2s. |
| R7 | La curaduría de proyectos elimina señal de ciertas skills (ej: Python, Docker, JavaScript) si los proyectos eliminados eran la única fuente de intervalos | Media | Medio | Verificar cada skill eliminada contra los intervalos calculados. Preservar las skills clave en `work_experience.technologies` del rol correspondiente (ej: Python y Docker en Desarrollo Independiente). |

---

## 8. Criterios de Éxito

### 8.1 Criterios Técnicos

- [ ] `go test ./...` pasa sin regresiones en el pipeline.
- [ ] `pnpm build` completa sin errores en la capa web.
- [ ] El PDF generado respeta: columna única en skills y recomendaciones, márgenes de 0.75", fuentes ATS-compatibles.
- [ ] Schema.org JSON-LD válido (verificar con Google Rich Results Test).
- [ ] Meta description y OG tags presentes en el `<head>`.

### 8.2 Criterios de Contenido

- [ ] Sección "About Me" presente con ≤80 palabras, sin pronombres de primera persona.
- [ ] Todas las descripciones de experiencia usan verbos de acción en pasado.
- [ ] Las descripciones aplican la estructura CAR cualitativo (Contexto → Acción → Resultado).
- [ ] Real2B no menciona "transacciones de caja" ni "atención al público".
- [ ] PIXI no menciona "arqueos de caja" ni "atención al público".
- [ ] Máximo 5 proyectos visibles en el CV renderizado.
- [ ] Las recomendaciones se muestran sin placeholders `<RECOMMENDER_X>` (solo texto entre comillas).

### 8.3 Criterios de Accesibilidad

- [ ] Estilos `:focus-visible` presentes en enlaces y botones.
- [ ] ARIA roles en `<header>`, `<main>`, y secciones relevantes.
- [ ] SVG de contacto con `aria-label`.

### 8.4 Criterios de No-Regresión

- [ ] Orden cronológico inverso preservado en experiencia.
- [ ] Cálculo de años de experiencia sin solapamientos funciona correctamente (test Go).
- [ ] Soporte dual web/PDF mantenido (dark mode web, light mode print).
- [ ] Botón de contacto (Google Form) funcional en web, oculto en PDF.

---

## 9. Cadena de Dependencias de Implementación

```
1. Go Pipeline — Agregar campo about_me en models.go + calculator.go
2. YAML Contenido — Reframe Real2B, PIXI, descripciones CAR, About Me, curaduría proyectos, educación
3. Re-run Pipeline — Regenerar cv-processed.json + verificar duraciones de skills
4. Template — About Me section, ARIA, SEO tags, Schema.org, viewport fix
5. CSS — Print layout (columnas, márgenes, fuentes), focus-visible
6. Verify — go test + pnpm build + verificación visual web/PDF
```

---

## 10. Presupuesto Estimado

| Categoría | Líneas Estimadas |
|-----------|-----------------|
| Go Pipeline (models + calculator) | ~17-21 |
| YAML Contenido (descripciones, reframes, about me, proyectos) | ~55-75 |
| Template (SEO, Schema, ARIA, About Me) | ~35-50 |
| CSS (print, focus) | ~20-30 |
| **Total** | **~127-176** |

Dentro del budget de 400 líneas ✅. Entregable como **PR único** (decisión LOCKED del usuario).

---

## 11. Decisiones del Usuario (LOCKED)

Estas decisiones fueron tomadas por el usuario y no deben modificarse:

1. **Contacto solo vía Google Form** — No agregar email, LinkedIn ni GitHub al hero.
2. **Real2B era Help Desk** — Mantenerlo para cubrir el gap temporal, reframizar con señal técnica.
3. **PIXI Supermercados** — Mantenerlo para cubrir el gap temporal, reframizar con señal de liderazgo.
4. **Sin métricas cuantificables** — Usar CAR cualitativo con verbos de acción, sin números inventados.
5. **About Me** — Proponer borrador basado en docs de perfil y textos de recomendación.
6. **Recomendaciones sin atribución** — Mantener quotes, sin nombres de autor.
7. **Entrega en PR único**.

# Diagnóstico del CV Web — Diego Sasco

> Documento de análisis y diagnóstico generado contra los criterios del marco teórico establecido en `docs/informe.md`. Fecha: 2026-06-10.

---

## 1. Resumen Ejecutivo

El CV web de Diego Sasco presenta una base técnica sólida en cuanto a infraestructura —Astro como generador estático, pipeline de datos en Go, CSS con soporte de impresión— pero exhibe deficiencias significativas cuando se evalúa contra los criterios de optimización ATS, escaneo ocular de 7.4 segundos, y señalización de credibilidad técnica que establece el informe de referencia.

La estructura de contenido enfrenta un problema fundamental de identidad: el documento opera simultáneamente como CV, portfolio y ficha de contacto genérica, sin cumplir de forma completa ninguna de las tres funciones. El hero carece de datos de contacto directos y enlaces a GitHub/LinkedIn (criterio 1 del informe para la Hero Section). No existe sección "About Me" narrativa (criterio 2 del informe, máximo 80 palabras). Los proyectos no se presentan como casos de estudio con métricas cuantificables, sino como descripciones funcionales planas. Y el currículum incluye roles de empleo puente (Real2B como "Analista Operacional & Control de Flujo", PIXI Supermercados como "Encargado de Operaciones") que, según la evidencia del NBER citada en el informe, penalizan significativamente la tasa de callback.

En el plano técnico-ATS, el CV usa columnas múltiples en la sección de skills (`grid-template-columns: repeat(auto-fill, minmax(250px, 1fr))`) y la hoja de impresión fuerza un grid de 3 columnas para skills —ambos incompatibles con el parseo lineal de los ATS. Las fuentes web (Inter, Outfit) no están entre las recomendadas por Harvard/MIT/Stanford para PDF de envío. Las descripciones de experiencia no aplican el marco CAR/CCCR, carecen de verbos de acción al inicio, y no cuantifican resultados. La sección de recomendaciones presenta placeholders sin atribución de autoría (`<RECOMMENDER_1>`), lo que degrada la credibilidad.

Las fortalezas a preservar incluyen: columna única en experiencia principal, orden cronológico inverso correcto, cálculo automático de años de experiencia sin solapamientos, soporte dual web/PDF desde un único source, y un cuerpo de proyectos freelance con alta señal técnica (IoT, gRPC, RPA financiera).

---

## 2. Diagnóstico por Dimensión

### 2.1 Estructura de Contenido y Narrativa Profesional

#### Hero Section vs. Criterio de Escaneo de 7.4 Segundos

El informe establece que en los primeros 7.4 segundos el reclutador busca seis elementos estructurados: nombre, cargo actual, empresa actual, fechas del rol actual, cargo previo, y sección de educación. El hero actual del CV cumple parcialmente:

- ✅ **Nombre completo**: "Diego Sasco" en `<h1>`.
- ✅ **Título profesional**: "Arquitecto de Software & Especialista en Automatización / IA" como subtítulo.
- ❌ **Datos de contacto directos**: No hay email, teléfono, LinkedIn ni GitHub. Solo un botón "Contacto" que redirige a un Google Form anónimo. El informe exige "enlaces hipertextuales limpios a GitHub, LinkedIn y correo electrónico corporativo" en el hero.
- ❌ **Empresa/cargo actual visible en el hero**: El primer rol ("Desarrollo Independiente & Consultoría / Freelance") aparece más abajo en la sección de experiencia, no en el hero.

El hero está centrado y visualmente atractivo, pero funcionalmente incompleto. Un reclutador que dedique 1.8 segundos al encabezado (según el eye-tracking citado) encontrará el nombre y título, pero no tendrá forma de contactar al candidato ni de verificar su perfil profesional en un solo vistazo.

#### Sección "About Me" (Narrativa Profesional)

- ❌ **Ausente por completo**. El informe exige "Breve síntesis (máximo 80 palabras) que describe las competencias del candidato, su enfoque en la resolución de problemas técnicos y su valor diferencial dentro de un equipo de desarrollo." No existe ningún párrafo de resumen profesional entre el hero y la experiencia laboral.

Esta ausencia es significativa porque la narrativa profesional actúa como puente semántico entre el título del hero y la experiencia detallada. Sin ella, el reclutador debe inferir la propuesta de valor del candidato exclusivamente a partir de la lista de trabajos.

#### Cantidad y Selección de Proyectos

El informe establece un límite de 3 a 5 proyectos flagship. El CV lista **12 proyectos** distribuidos entre las experiencias laborales:

| Experiencia | Proyectos |
|---|---|
| Desarrollo Independiente & Consultoría | 5 (RpPOS IoT, Proxy Inverso Financiero, Telemetría Hardware, Flujo IA SDD/TDD, CV Pipeline) |
| Sluckis Hermanos SA | 1 (Normalización de Inventario) |
| Bionico / Figital Tech | 3 (Framework RPA, Motor Extracción PDF, Soluciones IA Generativa) |
| TopBrands Int. | 2 (Automatización Retail, Optimización Administrativa) |

El informe advierte explícitamente: "La inclusión de más de 5 proyectos diluye la atención del evaluador y suele denotar falta de criterio en la selección de trabajos." Con 12 proyectos, la señal se diluye. Además, no todos deberían estar en un CV orientado a un rol de ingeniería de software: la "Normalización de Inventario" de Sluckis y la "Optimización de Flujos Administrativos" de TopBrands tienen baja señal técnica.

#### Ordenamiento Cronológico

- ✅ Las experiencias están en orden cronológico inverso correcto (Jul 2024 → Ene 2024 → Ene 2023 → Nov 2018 → Mar 2017 → Jul 2016).
- ⚠️ Los proyectos dentro de cada experiencia no siempre respetan el orden cronológico inverso. En TopBrands, "Automatización de Inventario y Retail" (Mar 2017 - Dic 2018) aparece antes que "Optimización de Flujos Administrativos" (Ene 2020 - Dic 2023), lo cual es cronológicamente incorrecto dentro de la entrada.

#### Uso del Marco CAR/CCCR en Descripciones

El marco CAR (Contexto, Acción, Resultado) y CCCR (Contexto, Desafío, Contribución, Resultado) exige que cada logro contenga: (C) la situación o problema, (A) la acción técnica específica, y (R) el resultado cuantificable.

**Análisis de las descripciones actuales:**

| Descripción actual | Marco CAR/CCCR | Diagnóstico |
|---|---|---|
| "Diseño y desarrollo de sistemas distribuidos, software a medida y soluciones de automatización IoT..." (Desarrollo Independiente) | Sin Contexto, Acción genérica, Sin Resultado | Frase descriptiva genérica, no un logro |
| "Gestión integral de operaciones de venta de servicios, transacciones de caja, soporte técnico..." (Real2B) | Sin Contexto, Acción genérica, Sin Resultado | Enumeración de tareas, no logros |
| "Liderazgo técnico en el diseño de soluciones de automatización de procesos de negocio (RPA), microservicios y procesamiento masivo de datos." (Bionico) | Contexto implícito, Acción parcial, Sin Resultado | Carece de métricas |

Ninguna descripción aplica el marco CAR/CCCR. Ninguna cuantifica resultados. El informe muestra ejemplos concretos del contraste:

- **Formulación débil** (tipo actual): "Fui responsable de desarrollar la base de datos de un rastreador financiero en Node.js y MongoDB."
- **Formulación optimizada CAR/CCCR**: "Diseñé e implementé la arquitectura de persistencia de datos en MongoDB utilizando índices compuestos, lo que redujo la latencia de las consultas de transacciones en un 35% para una plataforma de gestión financiera activa en producción."

#### Uso de Verbos de Acción vs. Lenguaje Pasivo

El informe prescribe verbos de acción al inicio de cada viñeta (diseñó, programó, optimizó, integró, desplegó, automatizó, etc.). Las descripciones del CV usan principalmente:

- Sustantivos abstractos: "Diseño y desarrollo de...", "Gestión integral de...", "Liderazgo técnico en...", "Automatización de..."
- Infinitivos como sustantivo: "colaborando estrechamente", "Aplicación de metodologías"

Estas formulaciones son pasivas y abstractas. Deberían reformularse en pasado con verbos de acción: "Diseñé e implementé...", "Lideré el desarrollo de...", "Automatizé procesos de...".

---

### 2.2 Optimización ATS y Accesibilidad

#### Layout de Columna Única

- ✅ **Experiencia profesional**: Presentada en una sola columna, con flujo de lectura lineal vertical. Esto es compatible con el parseo ATS.
- ❌ **Sección de Skills (web)**: Usa `grid-template-columns: repeat(auto-fill, minmax(250px, 1fr))`, generando múltiples columnas en pantallas anchas. ATS parsers que lean el HTML renderizado encontrarán texto mezclado entre categorías.
- ❌ **Sección de Skills (impresión/PDF)**: Fuerza `grid-template-columns: repeat(3, 1fr)` — tres columnas fijas. El informe es explícito: "Los formatos de una sola columna, alineados a la izquierda, reducen la fatiga ocular del evaluador y garantizan que toda la información relevante sea procesada dentro del umbral crítico de 7.4 segundos." Y: "Los motores de análisis sintáctico (parsers) de los ATS leen el documento de manera lineal horizontal; la presencia de dos columnas provoca una lectura cruzada que fusiona líneas de texto paralelas e inconexas."
- ❌ **Recomendaciones (impresión)**: Fuerza `grid-template-columns: repeat(2, 1fr)` — dos columnas en PDF.

#### Tipografía y Compatibilidad ATS

- ❌ **Fuentes web (Inter, Outfit)**: No están entre las fuentes ATS-compatible prescritas por el informe (Arial, Calibri, Times New Roman, Georgia, Tahoma, Verdana). En la hoja de impresión, se mantiene Inter como fuente del body (`font-family: var(--font-body)` = Inter). Si el PDF se genera desde el navegador, Inter puede no estar embebida correctamente, resultando en sustitución de fuente por parte del ATS.
- ⚠️ **El informe prescribe**: "Se exige el uso exclusivo de fuentes limpias y convencionales (como Arial, Calibri, Times New Roman, Georgia, Tahoma o Verdana) con tamaños de fuente consistentes de 10 pt a 12 pt."

#### Elementos Estructurales Problemáticos para ATS

- ⚠️ **`backdrop-filter: blur(10px)`** en `.section-card`: Efecto visual que no afecta el texto pero indica un enfoque estético sobre funcional. En el modo impresión se desactiva correctamente.
- ⚠️ **`-webkit-background-clip: text` / `-webkit-text-fill-color: transparent`** en `.hero h1`: El nombre "Diego Sasco" se renderiza con texto transparente y gradiente de fondo. Los ATS que extraen texto del DOM encontrarán el texto, pero los que procesen PDFs rasterizados pueden perder la información visual si el gradiente afecta el contraste en el documento impreso. La hoja de impresión corrige esto (`-webkit-text-fill-color: initial`), lo cual es correcto.
- ✅ **No hay tablas HTML, cuadros de texto, ni elementos gráficos que bloqueen el parseo**. La estructura es semántica con `<article>`, `<section>`, `<header>`, `<blockquote>`.
- ✅ **La hoja de impresión oculta el botón de contacto** (`.contact-form-btn { display: none }`), eliminando el enlace al Google Form del PDF.

#### Estructura HTML Semántica

- ✅ Uso de `<header>`, `<main>`, `<section>`, `<article>`, `<blockquote>` — semántica correcta.
- ❌ No hay atributos de accesibilidad ARIA, `role`, ni `alt` en el SVG del ícono de contacto.
- ❌ No hay `<nav>`, `<footer>`, ni estructura de encabezados completa (h1 → h2 → h3 es correcta, pero no hay landmark regions).

#### Compatibilidad PDF para Envío ATS

- ⚠️ La hoja de impresión tiene margen cero (`padding-block: 0; padding-inline: 0`). El informe prescribe márgenes de 0.5 a 1.0 pulgadas (0.75" óptimo). Sin márgenes, el contenido impreso toca los bordes del papel, lo cual: (a) puede ser cortado por impresoras, (b) no cumple los estándares de Harvard/MIT/Stanford, (c) se ve poco profesional.
- ⚠️ El tamaño de fuente body en impresión es 10pt, dentro del rango prescrito (10-12pt). Pero los tamaños de sección varían: 26pt para h1, 14pt para h2, 11pt para company-role, 9pt para fechas, 9.5pt para descripciones, 8.5pt para recomendaciones. Esta variación es excesiva según el informe que prescribe consistencia.

---

### 2.3 Arquitectura Web y Performance

#### Cumplimiento Mobile-First

El informe establece que "el 68% de las visualizaciones iniciales de portafolios se realizan desde teléfonos inteligentes" y exige filosofía mobile-first.

- ⚠️ **Meta viewport**: `<meta name="viewport" content="width=device-width">` — correcto pero incompleto. Debería incluir `initial-scale=1` para evitar comportamiento de zoom por defecto en algunos dispositivos.
- ⚠️ **No hay media queries específicas para móvil**. El CSS usa `clamp()` para tipografía fluida y `auto-fill` en grids, lo cual es un enfoque intrínsecamente responsive. Pero:
  - La clase `.skills-grid` con `minmax(250px, 1fr)` colapsa a 1 columna en móvil correctamente.
  - El `.experience-header` usa `flex-wrap: wrap` que funciona en móvil.
  - No hay breakpoints explícitos, lo cual es aceptable con el enfoque fluido actual.
- ❌ **No hay `@media (max-width: ...)` para optimizar la experiencia móvil**. El padding de 4rem en body es generoso en desktop pero puede ser excesivo o insuficiente en móviles muy estrechos sin ajuste.

#### Tiempo de Carga (< 3 segundos)

- ✅ **Astro genera HTML estático** — sin JavaScript en runtime por defecto. El `dist/index.html` es un archivo HTML puro con CSS inline/bundled. Esto debería cargar en sub-segundos.
- ❌ **Carga de fuentes externas desde Google Fonts**: `@import url("https://fonts.googleapis.com/css2?family=Inter:...&family=Outfit:...")`. Esto genera:
  1. Una petición DNS + conexión + descarga al CDN de Google.
  2. Posibles FOUT (Flash of Unstyled Text) o FOIT (Flash of Invisible Text).
  3. Dependencia de un recurso de terceros que puede ser bloqueado por ad-blockers o políticas CSP.
- ⚠️ **No hay preconnect a Google Fonts** ni `<link rel="preload">` para las fuentes.

#### SEO y Descubribilidad

- ❌ **No hay meta description**: Falta `<meta name="description" content="...">` — crítico para snippets de búsqueda.
- ❌ **No hay Open Graph tags**: Sin `og:title`, `og:description`, `og:image`, `og:type` — cuando alguien comparte el enlace en LinkedIn, WhatsApp o Slack, no hay preview card.
- ❌ **No hay datos estructurados (JSON-LD / Schema.org)**: El informe menciona la importancia de la indexación. Sin Schema.org `Person` o `ProfilePage`, el sitio pierde descubribilidad en búsquedas de reclutadores.
- ❌ **No hay `<link rel="canonical">`**: Potencial contenido duplicado si el sitio se sirve desde múltiples URLs.
- ❌ **No hay `robots.txt` ni `sitemap.xml`** visibles en el proyecto.
- ✅ **`<html lang="es">`** — correcto, indica el idioma del contenido.
- ✅ **Title tag descriptivo**: "Diego Sasco | Arquitecto de Software & Especialista en Automatización / IA".

#### Accesibilidad (WCAG)

- ❌ **Sin atributos ARIA** en ningún elemento.
- ❌ **Sin `alt` text** en el SVG del botón de contacto.
- ⚠️ **Contraste de colores en modo web**: `--text-muted: #94a3b8` sobre `--bg-color: #0b0f19` tiene ratio de contraste aproximado de 6.5:1 (pasa AA para texto normal, pero las tech-tags usan `--text-muted` sobre fondo `rgba(255,255,255,0.05)` que puede reducir el contraste efectivo).
- ⚠️ **`:focus` styles ausentes**: No hay estilos de foco definidos para enlaces ni botones, lo cual es un problema de accesibilidad por teclado.

---

### 2.4 Diseño Visual y Carga Cognitiva

#### Compatibilidad con Patrón de Lectura en F y Z

El informe detalla que el escaneo ocular sigue un patrón en F: fijación horizontal en el tercio superior, segundo barrido más corto, y barrido vertical por el margen izquierdo.

- ✅ **Hero centrado**: El nombre y título están en el tercio superior — compatible con la primera fijación horizontal.
- ✅ **Columna única en experiencia**: El flujo lineal facilita el barrido vertical.
- ⚠️ **Experiencia-header con layout flex**: La empresa y rol están a la izquierda, las fechas a la derecha (`justify-content: space-between`). El patrón F prioriza el margen izquierdo, donde está la información más importante (empresa, rol). Esto es correcto.
- ❌ **Skills en grid multi-columna**: Rompe el barrido F. El lector debe saltar entre columnas, incrementando la carga cognitiva.
- ❌ **Recomendaciones en grid multi-columna**: Mismo problema. Las blockquotes están distribuidas en 2-3 columnas, obligando al lector a escanear en zigzag.

#### Densidad de Información

El CV presenta una densidad de información muy alta. Contabilizando el contenido renderizado:

- 6 experiencias laborales con descripciones
- 12 proyectos con descripciones y tags
- 8 categorías de skills con ~20 habilidades individuales
- 4 entradas de educación
- 3 recomendaciones

Esto genera un documento extenso. Para la versión web es tolerable (scroll infinito), pero para la versión PDF impresa resulta en un volumen de información que probablemente exceda las 2 páginas recomendadas, especialmente considerando que la hoja de impresión usa `page-break-inside: avoid` agresivamente.

#### Contraste y Legibilidad

- **Modo web (dark)**: Fondo `#0b0f19` con texto `#f1f5f9` — alto contraste, buena legibilidad. Pero los elementos `--text-muted` (`#94a3b8`) sobre el fondo oscuro tienen menor contraste. Las tech-tags (`#94a3b8` sobre `rgba(255,255,255,0.05)`) tienen contraste borderline.
- **Modo impresión (light)**: Fondo `#ffffff` con texto `#0f172a` — excelente contraste. Los textos muted usan `#475569` sobre blanco (~7:1), lo cual es correcto.

#### Dark Mode vs. Contexto del Reclutador

El informe no aborda explícitamente el dark mode, pero el contexto es relevante: un reclutador que abre el CV web desde un email o enlace probablemente lo ve en un navegador con tema claro por defecto. El sitio fuerza dark mode (`--bg-color: #0b0f19`) sin opción de cambio. Esto puede generar:
- Impacto visual fuerte (positivo o negativo según preferencia).
- Posibles problemas de impresión directa desde el navegador (el CSS de impresión lo mitiga).
- Incompatibilidad con la expectativa de un documento profesional "limpio y convencional" que prescribe el informe.

---

### 2.5 Perfil Profesional y Señalización

#### Enlaces a GitHub y LinkedIn

- ❌ **No hay enlace a GitHub** en todo el documento. El informe dedica una sección completa a la optimización del perfil de GitHub como "registro empírico más confiable del flujo de trabajo habitual de un desarrollador." Sin enlace a GitHub, el reclutador no puede verificar las credenciales técnicas declaradas. El informe indica que optimizar el perfil de GitHub "puede incrementar las visitas de reclutadores en un 300% y los retornos de llamada en un 40%."
- ❌ **No hay enlace a LinkedIn**. LinkedIn es la plataforma estándar de networking profesional. Su ausencia reduce la capacidad del reclutador de validar la identidad y trayectoria del candidato.
- ❌ **No hay email visible**. El único canal de contacto es un Google Form anónimo. Esto: (a) no permite al reclutador enviar un email directo, (b) no permite respuesta asíncrona, (c) genera fricción en el proceso de contacto.

#### Señales de Credibilidad Técnica

- ✅ **Proyectos con tecnologías relevantes**: Go, gRPC, C#/.NET, RabbitMQ, IoT, Docker — stack técnico coherente y demandado.
- ✅ **Experiencia como fundador/líder**: El rol "Director & Líder de Desarrollo de Software" en Bionico/Figital Tech (5 años) demuestra capacidad de liderazgo técnico.
- ✅ **Certificación en Hacking Ético**: Diferenciador específico.
- ❌ **Sin enlaces a repositorios o demos**: Los proyectos listan tecnologías pero no tienen URLs a código, demos o documentación. El informe exige "enlaces de acceso al despliegue activo" y "enlace visible al código en producción." El cv_pipeline tiene `url: ""` vacío.
- ❌ **Sin métricas de impacto**: Ningún proyecto incluye métricas cuantificables ("redujo latencia en X%", "procesó Y facturas/día", "soportó Z usuarios concurrentes").

#### Presentación de la Educación

- ⚠️ **2 de 4 entradas son incompletas**: "Analista en Infraestructura Informática" (Incompleto) e "Ingeniería en Sistemas Computacionales" (1er año aprobado). El informe recomienda escribir los nombres completos de las instituciones, lo cual se cumple ("Universidad ORT Uruguay", "Universidad de la República").
- ⚠️ **"Contador Público" (1er año cursado)**: Esta entrada puede confundir al reclutador sobre el foco profesional del candidato. Un reclutador que escanee rápidamente puede ver "Contador Público" y clasificar al candidato incorrectamente.
- ✅ **Certificación en Hacking Ético (Graduado)**: La única credencial completa y directamente relevante para roles técnicos.

#### Manejo de Empleos Puente / Roles de Baja Señal Técnica

El informe cita investigación del NBER que demuestra: "aceptar un puesto temporal de baja cualificación técnica reduce significativamente la probabilidad de recibir una llamada de retorno en comparación con mantenerse desempleado." Identifica explícitamente "cajero, barista" y roles de "subsistencia económica" como penalizantes.

El CV incluye dos roles que califican como empleos puente según esta taxonomía:

1. **Real2B — "Analista Operacional & Control de Flujo"** (Ene 2024 - Jun 2024, 6 meses): Descripción: "Gestión integral de operaciones de venta de servicios, transacciones de caja, soporte técnico (Help Desk) a usuarios y atención directa al público en múltiples locales." Las skills listadas son: Help Desk, Atención al Público, Trabajo en Equipo, Comunicación Asertiva, Resolución de Problemas — todas habilidades blandas/fucionales, ninguna técnica de software. Este rol es, por definición del informe, un "empleo puente de baja cualificación técnica."

2. **PIXI Supermercados — "Encargado de Operaciones"** (Jul 2016 - Mar 2017, 9 meses): Descripción: "Supervisión de equipos, administración de personal, control de inventario, arqueos de caja y atención al público." Skills: Atención al Público, Liderazgo, Trabajo en Equipo, Comunicación Asertiva. Sin tecnologías de software.

Según el informe, estos roles deberían: (a) eliminarse del CV para puestos de ingeniería de software, o (b) reformularse mínimamente si contribuyen a la narrativa de evolución profesional. Su inclusión actual penaliza la tasa de callback.

#### Efectividad de la Sección de Recomendaciones

- ⚠️ **Sin atribución de autoría**: Las tres recomendaciones tienen placeholders: `author: "<RECOMMENDER_1>"`, `author: "<RECOMMENDER_2>"`, `author: "<RECOMMENDER_3>"`. En el HTML renderizado, las blockquotes solo muestran el texto entre comillas, sin nombre, rol ni relación del autor. Esto reduce significativamente la credibilidad — una recomendación anónima vale menos que una atribuida.
- ⚠️ **Contenido genérico**: "He tenido el gusto de trabajar con Diego, es un excelente compañero como profesional." — no menciona ninguna competencia técnica específica, proyecto o resultado. El informe enfatiza la importancia de la señalización de credibilidad técnica.
- ✅ **Formato blockquote**: La elección de `<blockquote>` es semánticamente correcta para testimonios.

---

## 3. Hallazgos Críticos

Los siguientes problemas se clasifican por severidad, de mayor a menor impacto en la tasa de conversión del CV:

1. **Ausencia total de canales de contacto directo** — Sin email, LinkedIn ni GitHub. Solo un Google Form. Esto viola el criterio fundamental de la Hero Section del informe y genera fricción extrema para el reclutador. Severidad: **Crítica**.

2. **Empleos puente incluidos sin reformulación** — Real2B (Analista Operacional) y PIXI Supermercados (Encargado de Operaciones) son roles de baja cualificación técnica que, según la evidencia del NBER citada en el informe, penalizan la tasa de callback en puestos de ingeniería de software. Severidad: **Crítica**.

3. **Layout multi-columna en Skills y Recomendaciones (PDF)** — El grid de 3 columnas en skills y 2 columnas en recomendaciones en la hoja de impresión viola directamente la regla de columna única para ATS y escaneo ocular. Severidad: **Alta**.

4. **Cero métricas cuantificables en proyectos y experiencias** — Ningún proyecto ni experiencia incluye resultados medibles. El marco CAR/CCCR prescribe explícitamente "impacto cuantitativo medible del proyecto." Severidad: **Alta**.

5. **Fuentes no ATS-compatibles en PDF** — Inter y Outfit no están entre las fuentes estándar prescritas por Harvard/MIT/Stanford para CVs enviados a ATS. Severidad: **Alta**.

6. **Márgenes cero en PDF impreso** — La hoja de impresión elimina todo padding. El informe prescribe 0.75 pulgadas. Severidad: **Alta**.

7. **Sección "About Me" inexistente** — Violación directa de la estructura de e-portfolio prescrita (sección 2 obligatoria). Severidad: **Media-Alta**.

8. **12 proyectos sin curaduría** — Excede el límite de 5 flagship del informe. Diluye la señal técnica. Severidad: **Media**.

9. **Descripciones en lenguaje pasivo sin verbos de acción** — Todas las descripciones usan sustantivos abstractos en lugar de verbos de acción pasados. Severidad: **Media**.

10. **Recomendaciones sin atribución de autoría** — Placeholders `<RECOMMENDER_X>` visibles en el contenido. Severidad: **Media**.

11. **Sin meta tags SEO, Open Graph ni Schema.org** — El sitio es prácticamente invisible para motores de búsqueda y no genera preview cards al compartirse. Severidad: **Media**.

12. **Educación incompleta y carrera "Contador Público" sin contextualizar** — Potencial clasificación errónea por parte del reclutador. Severidad: **Baja-Media**.

---

## 4. Fortalezas a Preservar

1. **Pipeline de datos automatizado con cálculo de experiencia sin solapamientos** — El sistema Go calcula años de experiencia reales descontando períodos superpuestos. Esto es un diferenciador técnico real que demuestra capacidad de ingeniería.

2. **Arquitectura de columna única en la sección de experiencia** — La experiencia laboral, que es la sección más importante del CV, usa un layout lineal compatible con el escaneo en F y el parseo ATS.

3. **Soporte dual web/PDF desde un único source YAML** — La arquitectura de un solo archivo de datos (`cv-source.yaml`) que alimenta tanto la web como el PDF es eficiente y mantenible. La hoja de impresión está bien diseñada en términos de cambios de color y supresión de elementos decorativos.

4. **Orden cronológico inverso correcto** — Las experiencias están ordenadas de la más reciente a la más antigua, cumpliendo el estándar prescrito.

5. **HTML semántico** — Uso correcto de `<header>`, `<main>`, `<section>`, `<article>`, `<blockquote>`, `<h1>`-`<h3>` con jerarquía coherente.

6. **Generación estática con Astro** — HTML puro sin JavaScript en runtime, tiempos de carga óptimos, compatibilidad con hosting estático (GitHub Pages, Netlify, Cloudflare Pages).

7. **Proyectos freelance con alta señal técnica** — RpPOS IoT (Go + Raspberry Pi + GPIO), Proxy Inverso Financiero (SOAP→gRPC), Telemetría Hardware (protocolos seriales + gRPC) son proyectos diferenciadores que demuestran competencia en sistemas embebidos, protocolos de comunicación e integración con hardware.

8. **Stack técnico coherente y verificable** — C#/.NET (7 años), Go (2 años), Python (3 años), Linux (8+ años), SQL (6+ años), gRPC, RabbitMQ, Docker — un perfil de backend/infraestructura sólido.

9. **Experiencia como fundador/líder** — "Director & Líder de Desarrollo de Software" en Bionico/Figital Tech durante 5 años con proyectos de RPA financiera e IA generativa es una señal fuerte de capacidad técnica y de gestión.

10. **Certificación en Hacking Ético** — Un diferenciador específico que valida competencias de seguridad, complementando el perfil de infraestructura/backend.

---

## 5. Matriz de Recomendaciones Priorizadas

| # | Recomendación | Dimensión | Impacto | Esfuerzo | Prioridad |
|---|---------------|-----------|---------|----------|-----------|
| 1 | Agregar email directo, enlace a GitHub y enlace a LinkedIn en el hero section | Contenido / Señalización | Alto | Bajo | P0 |
| 2 | Eliminar o reformular completamente Real2B (empleo puente) — si se mantiene, reenfocar hacia automatización/soporte técnico eliminando "transacciones de caja" y "atención al público" | Contenido / Señalización | Alto | Bajo | P0 |
| 3 | Eliminar PIXI Supermercados del CV para postulaciones de ingeniería de software (empleo puente de baja cualificación técnica sin señalización de software) | Contenido / Señalización | Alto | Bajo | P0 |
| 4 | Convertir Skills grid a layout de columna única en la hoja de impresión para compatibilidad ATS total | ATS / Layout | Alto | Bajo | P0 |
| 5 | Agregar márgenes de 0.75 pulgadas (0.56cm aprox.) en la hoja de impresión | ATS / Diseño | Alto | Bajo | P0 |
| 6 | Reformular todas las descripciones de experiencia usando el marco CAR/CCCR con verbos de acción y métricas cuantificables | Contenido / Narrativa | Alto | Medio | P0 |
| 7 | Agregar sección "About Me" (máximo 80 palabras) entre el hero y la experiencia laboral | Contenido / Estructura | Alto | Bajo | P1 |
| 8 | Cambiar fuentes de impresión a Arial, Calibri o Georgia (ATS-compatibles) | ATS / Tipografía | Medio | Bajo | P1 |
| 9 | Curar proyectos a máximo 5 flagship: RpPOS IoT, Proxy Inverso Financiero, Telemetría Hardware, Framework RPA Financiero, Motor de Extracción PDF | Contenido / Proyectos | Alto | Medio | P1 |
| 10 | Completar autoría de las 3 recomendaciones (nombre, rol, relación) o eliminar la sección si no se puede verificar | Contenido / Credibilidad | Medio | Bajo | P1 |
| 11 | Agregar meta description, Open Graph tags y Schema.org JSON-LD (tipo Person) en el `<head>` | SEO / Performance | Medio | Bajo | P1 |
| 12 | Agregar preconnect a Google Fonts y considerar self-hosting de fuentes para eliminar dependencia externa | Performance | Medio | Medio | P2 |
| 13 | Agregar estilos `:focus-visible` para accesibilidad por teclado | Accesibilidad | Medio | Bajo | P2 |
| 14 | Reconvertir Skills en impresión a lista plana en lugar de grid multi-columna | ATS / Layout | Medio | Bajo | P2 |
| 15 | Agregar enlaces a repositorios GitHub y/o demos desplegadas en cada proyecto flagship | Señalización / Credibilidad | Alto | Alto | P2 |
| 16 | Contextualizar educación incompleta — eliminar "Contador Público (1er año cursado)" o reformular como complemento a la narrativa de "perfil que cruza negocio y tecnología" | Contenido / Educación | Bajo | Bajo | P2 |
| 17 | Agregar `<meta name="viewport" content="width=device-width, initial-scale=1">` completo | Mobile / Performance | Bajo | Bajo | P2 |
| 18 | Implementar opción de tema claro/oscuro para mejor recepción por reclutadores en contexto profesional estándar | Diseño Visual | Medio | Alto | P3 |

---

## 6. Red Flags Específicas del Informe

El informe define una taxonomía de red flags específicas para postulaciones tecnológicas. A continuación se mapea cada una al estado actual del CV:

| Red Flag del Informe | Estado en el CV | Detalle |
|---|---|---|
| **Proyectos basados exclusivamente en tutoriales públicos** | ⚠️ No verificable | Sin enlaces a repositorios ni demos, es imposible para el reclutador determinar si los proyectos son originales o derivados de tutoriales. El informe advierte que herramientas como GitMeter detectan "proyectos clonados sin aportaciones individuales reales." La ausencia de enlaces a GitHub genera la sospecha por omisión. |
| **Enlaces de demostración caídos u obsoletos** | ❌ Presente por ausencia total | Ningún proyecto tiene enlace de demostración. El Cv Pipeline tiene `url: ""`. Esto es peor que tener enlaces caídos: es no tener enlaces en absoluto. |
| **Ausencia de manejo de variables de entorno de producción** | ⚠️ No verificable | Las descripciones no mencionan prácticas de deployment, configuración ni manejo de entornos. |
| **Colecciones inconexas sin hipótesis de negocio** | ⚠️ Parcialmente presente | Los proyectos de Bionico (Framework RPA, Motor PDF, Soluciones IA) están agrupados bajo una empresa pero no se articulan como un ecosistema coherente. Las descripciones son funcionales pero no conectan el problema de negocio con la solución técnica. |
| **Sobreingeniería en proyectos** | ✅ Ausente | Los proyectos muestran un nivel de complejidad apropiado para su escala (e.g., un proxy inverso SOAP→gRPC para una pasarela financiera es razonable, no sobre-ingeniería). |
| **Empleos puente / trabajos de subsistencia** | ❌ Presente | Real2B ("transacciones de caja", "atención al público") y PIXI Supermercados ("arqueos de caja", "atención al público") son empleos puente según la taxonomía del informe. Su presencia penaliza la tasa de callback. |
| **Formato multicolumna incompatible con ATS** | ❌ Presente | Skills grid en 3 columnas (impresión) y auto-fill multi-columna (web). Recomendaciones en 2 columnas (impresión). |
| **Fuentes no estándar para ATS** | ❌ Presente | Inter y Outfit en lugar de Arial, Calibri, Times New Roman, Georgia, Tahoma o Verdana. |
| **Ausencia de verbos de acción** | ❌ Presente | Todas las descripciones usan sustantivos abstractos o infinitivos. Ninguna usa verbos de acción en pasado al inicio. |
| **Cero métricas cuantificables** | ❌ Presente | Ningún proyecto ni experiencia incluye porcentajes, cantidades, tiempos de mejora ni KPIs de negocio. |
| **Saturación artificial de palabras clave** | ✅ Ausente | Las tecnologías listadas corresponden a proyectos reales y verificables. No hay relleno de keywords sin contexto. |
| **Documentos generados por IA sin autenticidad** | ✅ Ausente | El contenido tiene voz auténtica y específica. No presenta los indicadores de IA que el informe identifica (bloques densos de texto, formato plano, saturación de keywords). |
| **Currículums en dos columnas** | ⚠️ Parcialmente presente | La experiencia principal es columna única, pero skills y recomendaciones son multi-columna. No es un CV de dos columnas en el sentido tradicional (sidebar + contenido), pero las secciones de grid violan el mismo principio de lectura lineal. |

---

*Fin del diagnóstico. Los hallazgos y recomendaciones aquí presentados deben priorizarse según la columna "Prioridad" de la Matriz de Recomendaciones (sección 5), comenzando por las P0 que representan cambios de alto impacto y bajo esfuerzo.*

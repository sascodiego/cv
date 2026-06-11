# SDD Explore — cv-web-optimization

**Change ID**: `cv-web-optimization`
**Phase**: Explore
**Date**: 2026-06-10
**Status**: ✅ Complete

---

## 1. Mapping de Recomendaciones → Archivos

18 recomendaciones analizadas → **15 accionables, 1 bloqueada, 1 fusionada, 1 diferida.**

| # | Rec | Prioridad | Tipo | Archivos | Líneas Est. | Input usuario? |
|---|-----|-----------|------|----------|-------------|----------------|
| 1 | Links de contacto en hero | P0 | **BLOQUEADA** | N/A | 0 | N/A — usuario decidió Google Form único |
| 2 | Reformular Real2B | P0 | YAML contenido | `cv-source.yaml` | ~8 | No (reframe con señal técnica) |
| 3 | Reformular PIXI | P0 | YAML contenido | `cv-source.yaml` | ~8 | No (reframe con señal técnica) |
| 4 | Skills columna única en print | P0 | CSS estilo | `global.css` | ~3 | No |
| 5 | Márgenes 0.75" en print | P0 | CSS estilo | `global.css` | ~2 | No |
| 6 | Descripciones CAR/CCCR | P0 | YAML contenido | `cv-source.yaml` | ~36-54 | **SÍ — métricas requeridas** |
| 7 | Sección "About Me" | P1 | Cross-boundary | `cv-source.yaml` + `models.go` + `calculator.go` + `index.astro` + `global.css` | ~17-21 | **SÍ — texto requerido** |
| 8 | Fuentes ATS-compatibles en print | P1 | CSS estilo | `global.css` | ~4 | No |
| 9 | Curar a 5 proyectos flagship | P1 | YAML contenido | `cv-source.yaml` | ~10-15 | Parcial (confirmar selección) |
| 10 | Autoría de recomendaciones | P1 | YAML + Template | `cv-source.yaml` + `index.astro` | ~15 | Parcial (nombres necesarios) |
| 11 | Meta/OG/Schema.org | P1 | Template | `index.astro` | ~15-25 | No |
| 12 | Preconnect Google Fonts | P2 | Template | `index.astro` | ~2 | No |
| 13 | Estilos :focus-visible | P2 | CSS estilo | `global.css` | ~10-15 | No |
| 14 | Skills flat list print | P2 | **FUSIONADA → Rec 4** | See Rec 4 | 0 | No |
| 15 | URLs de proyectos | P2 | YAML + Template | `cv-source.yaml` + `index.astro` | ~11 | Parcial (URLs necesarias) |
| 16 | Contextualizar educación | P2 | YAML contenido | `cv-source.yaml` | ~3-5 | No |
| 17 | Viewport meta fix | P2 | Template | `index.astro` | ~1 | No |
| 18 | Toggle light/dark | P3 | **DIFERIDA** | N/A | 0 | N/A |

## 2. Presupuesto de Review

**138–205 líneas estimadas** — dentro del budget de 400 líneas ✅

## 3. Inputs Bloqueantes del Usuario

| ID | Rec | Dato requerido | Cómo resolver |
|----|-----|----------------|---------------|
| U1 | Rec 6 | Métricas cuantificables para cada experiencia y proyecto (porcentajes, cantidades, tiempos) | Usuario provee datos o aprueba placeholders razonables |
| U2 | Rec 7 | Texto "About Me" (~80 palabras, español) | Usuario redacta o aprueba borrón propuesto |

## 4. Descubrimientos Clave del Código

1. **Project URLs existen en datos pero el template nunca las renderiza** — `proj.url` está en `ProjectProcessed` y `cv-processed.json` pero `index.astro` lo ignora. Los estilos CSS `.indie-project-title a` existen sin uso.

2. **Datos de autoría de recomendaciones están en YAML pero el HTML no los renderiza** — `index.astro` solo renderiza `rec.text`, ignorando `rec.author`/`rec.role`/`rec.relation`. Todas las clases CSS (`.recommendation-author`, `.author-name`, `.author-role`, `.author-relation`) están en `global.css` como código muerto.

3. **Los meses de experiencia de skills se calculan desde project intervals, no desde work_experience.technologies** — Rec 9 (curaduría de proyectos) puede alterar las duraciones calculadas.

4. **PersonalInfo en Go ya tiene campos GitHub/LinkedIn/Email** — excluidos del output público por diseño (línea 98-100 en calculator.go). Reversible en 1 línea si el usuario cambia de opinión.

## 5. Cadena de Dependencias

```
cv-source.yaml → Go pipeline → cv-processed.json → Astro build → dist/index.html
                                                     ↑
                                               index.astro
                                               global.css
```

- Rec 7 requiere cambios en Go model ANTES de re-ejecutar el pipeline
- Rec 9 requiere re-ejecutar pipeline + verificar duraciones de skills
- El resto de cambios YAML necesitan re-ejecutar pipeline sin cambios en Go
- Cambios solo CSS/template (Rec 4, 5, 8, 11, 12, 13, 17) necesitan solo `pnpm build`

## 6. Orden de Fases Propuesto

1. **Go Pipeline** — Rec 7 (campo `about_me` en models.go + calculator.go)
2. **YAML Contenido** — Rec 2, 3, 6, 7, 9, 10, 15, 16
3. **Re-run Pipeline** — regenerar cv-processed.json
4. **Template** — Rec 7, 10, 11, 12, 15, 17
5. **CSS** — Rec 4/14, 5, 8, 13
6. **Verify** — `go test ./...` + `pnpm build` + verificación visual

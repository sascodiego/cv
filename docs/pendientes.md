# Pendientes

## Normalización de idioma en categorías de Skills

**Estado**: Pendiente
**Prioridad**: Media

Las categorías de skills en `data/cv-source.yaml` mezclan inglés y español:

- `Languages & Frameworks` (inglés)
- `Databases & Obs` (inglés, truncado)
- `Architecture & Protocols` (inglés)
- `Metodologías & Prácticas` (español)
- `Habilidades Funcionales` (español)
- `Competencias Blandas` (español)

**Acción**: Unificar todas las categorías a español, dado que `<html lang="es">` y todo el contenido del CV está en español.

**Archivos afectados**: `data/cv-source.yaml` → requiere re-procesar pipeline (`cd pipeline && go run .`).

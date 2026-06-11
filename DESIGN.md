# Guía de Diseño y Estándares: Diego Sasso Portfolio

Este documento establece las directrices de diseño, experiencia de usuario (UX), dirección de arte y especificaciones técnicas de exportación para la plataforma web y hoja de vida de **Diego Sasso**. Está basado en el minimalismo funcional, sobriedad ingenieril y las mejores prácticas de accesibilidad y optimización para lectores automáticos (ATS).

---

## 1. Filosofía del Diseño

El diseño rechaza la sobrecarga visual, los degradados sintéticos de alta saturación ("neón") y las animaciones intrusivas. En su lugar, adopta una estética **editorial y arquitectónica**, similar a los entornos de trabajo de plataformas de ingeniería de alto rendimiento (como Vercel, Linear o Stripe).

### Principios Fundamentales:

- **Legibilidad Radical:** El texto es el elemento principal. Cada decisión estética apoya la fluidez de lectura.
- **Alineación de Columna Única:** Una sola columna centralizada garantiza un orden de lectura intuitivo y una traducción perfecta (1:1) al formato de impresión física.
- **Densidad de Información Equilibrada:** Uso riguroso de espacios vacíos (padding/margins) para evitar la fatiga visual.

---

## 2. Paleta de Colores (Sistema Neutro de Contraste)

La interfaz utiliza una paleta monocromática basada en tonos tierra suaves y metales oscuros, eliminando la distracción del color para resaltar el contenido de ingeniería.

### Modo Oscuro (Predeterminado)

- **Fondo Principal:** `#09090b` (Zinc-950) — Oscuridad absoluta y reposo visual.
- **Texto Primario:** `#f4f4f5` (Zinc-100) — Alta visibilidad sin brillo excesivo.
- **Texto Secundario:** `#a1a1aa` (Zinc-400) — Para descripciones y sumarios.
- **Fronteras y Líneas:** `#18181b` (Zinc-900) — Separación sutil sin segmentar ruidosamente el espacio.

### Modo Claro (Alabastro)

- **Fondo Principal:** `#faf9f6` (Alabastro/Marfil) — Un tono hueso cálido que reduce la fatiga ocular a diferencia del blanco puro (`#ffffff`).
- **Texto Primario:** `#18181b` (Zinc-900) — Contraste nítido.
- **Texto Secundario:** `#71717a` (Zinc-500) — Jerarquía tipográfica secundaria.
- **Fronteras y Líneas:** `#e4e4e7` (Zinc-200) — Delimitadores limpios.

---

## 3. Tipografía y Jerarquía Visual

El sistema tipográfico utiliza fuentes sans-serif del sistema operativo para optimizar la carga del sitio ($t < 1\text{ s}$) y garantizar que los motores de renderizado no sufran desfases visuales (CLS).

- **Familia Tipográfica Principal:** `Inter`, `system-ui`, `-apple-system`, `sans-serif`.
- **Familia Tipográfica Secundaria (Datos/Métricas):** `SFMono-Regular`, `Consolas`, `monospace`.

### Escala de Tamaños:

- **H1 (Nombre del Ingeniero):** `2.25rem` / `36px` (`font-extrabold`, `tracking-tight`).
- **H2 (Secciones):** `0.75rem` / `12px` (`font-mono`, `font-bold`, `tracking-widest`, `uppercase`).
- **H3 (Roles/Puestos):** `1rem` / `16px` (`font-bold`).
- **Cuerpo de texto:** `0.875rem` a `1rem` (`font-light`, `leading-relaxed`).

---

## 4. Estándar de Contenidos (Modelo CAR)

Toda viñeta técnica dentro de la sección de experiencia laboral debe formularse bajo el estándar científico **Contexto-Acción-Resultado (CAR)**:

1. **Verbo de Acción Inicial:** Utilizar palabras de alto impacto técnico (p. ej., _Diseñó_, _Lideró_, _Automatizó_, _Optimizó_).
2. **Acción/Tecnología:** Detallar la solución de ingeniería empleada (p. ej., _arquitectura de microfrontends con React y Go_).
3. **Métrica Cuantificable:** Todo logro debe cerrar con un resultado numérico verificable (p. ej., _reduciendo el tiempo de carga en un 42% y mejorando la conversión en un 18%_).

---

## 5. Motor de Optimización para Impresión (ATS-Ready)

La hoja de estilos CSS cuenta con reglas estrictas `@media print` para garantizar que la exportación física desde el navegador sea impecable:

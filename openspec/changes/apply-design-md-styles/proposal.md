# Proposal: Apply DESIGN.md Styles to the Web CV

## Change ID

`apply-design-md-styles`

## Title

Apply DESIGN.md styles to the web CV

## Problem Statement

The current web CV presents a dark premium visual system with saturated blue/indigo accents, gradients, translucent cards, backdrop blur, multi-column grids, Google Fonts, and hover/entrance motion. That direction conflicts with `DESIGN.md`, which defines the CV as a sober, editorial, architectural, ATS-preserving document experience: zinc-neutral palette, strict single-column reading, system typography, minimal section structure, restrained interaction, and print output that maps cleanly from web to physical/PDF form.

This mismatch creates a product problem: the web CV does not fully express the intended professional positioning. It feels more like a polished portfolio landing page than an engineering CV optimized for readability, recruiter scanning, and print/ATS continuity.

## Intent

Realign the web CV presentation with `DESIGN.md` while preserving its conversion purpose. The result should feel like a refined engineering document: minimal, dark, editorial, readable, accessible, and credible, with a deliberate but sober contact call to action.

## Proposal Question Round

Interactive SDD mode normally offers product-shaping questions before finalizing a proposal. The user already supplied the key decisions for this proposal round:

1. Theme: dark only; remove or avoid the existing theme toggle; no light mode deliverable.
2. Layout: strict single column for web and print.
3. Effects: restrained microinteractions are allowed; no intrusive animation, synthetic/neon gradients, or card hover movement.
4. CTA: keep a conversion-oriented call to action that sells the CV while fitting the sober editorial zinc design.

These answers are treated as proposal assumptions. No blocking product questions remain for the proposal phase.

## Goals

- Apply the `DESIGN.md` dark zinc visual system to the web CV.
- Replace blue/indigo gradient premium styling with neutral editorial styling.
- Remove reliance on Google Fonts in favor of system font stacks as specified by `DESIGN.md`.
- Enforce a single-column reading model across web and print.
- Preserve or improve accessibility: contrast, focus visibility, reduced-motion support, semantic readability, and keyboard access.
- Preserve ATS/print readiness and avoid introducing visual structures that degrade PDF export.
- Keep a conversion-oriented contact CTA, but restyle it so it feels like a sober editorial action rather than a marketing gradient button.

## Non-Goals

- No implementation in this proposal phase.
- No light theme or theme toggle deliverable for this change.
- No redesign of CV content, data structure, or CAR bullet wording.
- No new animation system, illustration system, synthetic gradients, neon effects, or large decorative backgrounds.
- No multi-column desktop layout.
- No changes to backend/data generation unless later implementation discovers a styling blocker.

## User and Product Impact

- Recruiters and hiring managers get a more legible, document-like CV that emphasizes engineering credibility over decorative portfolio effects.
- Print/PDF users receive a closer 1:1 relationship between the web reading order and exported physical document.
- The CV owner retains a conversion path through a deliberate contact CTA instead of losing the ability to drive inbound opportunities.
- Maintenance improves because the style system becomes simpler: fewer visual modes, fewer imported assets, fewer motion/card variants.

## Scope

### In Scope

- Major stylesheet refactor in `web/src/styles/global.css`:
  - Define zinc-based CSS variables from `DESIGN.md`.
  - Remove Google Fonts import and use system font stacks.
  - Replace gradient/accent/card-heavy styling with editorial lines, spacing, and typographic hierarchy.
  - Normalize section layout to a strict single column.
  - Restyle headings to match `DESIGN.md`: H1 2.25rem, H2 mono uppercase 0.75rem, H3 1rem.
  - Remove or neutralize intrusive entrance animations and hover movement.
  - Preserve reduced-motion handling.
  - Preserve and review print rules/page-break behavior.
- Minor template adjustments in `web/src/pages/index.astro` if needed:
  - Update `theme-color` to `#09090b`.
  - Remove Google Fonts preconnect tags if no longer needed.
  - Adjust CTA copy/markup if needed to support the sober conversion-oriented style.
  - Remove/avoid any theme-toggle markup if present.

### Out of Scope

- CV data changes in `web/src/data/cv-processed.json`.
- New pages, routing, or form integrations.
- Analytics or conversion tracking changes.
- Visual QA automation beyond the implementation phase’s normal verification.

## Affected Areas

- `DESIGN.md` as the source of truth.
- `web/src/styles/global.css` as the main implementation surface.
- `web/src/pages/index.astro` for meta/theme/font-link/CTA markup adjustments.
- Browser print/PDF output, especially page breaks, section continuity, and ATS-friendly text flow.

## Design Decisions

- Dark-only implementation using `#09090b` background, `#f4f4f5` primary text, `#a1a1aa` muted text, and `#18181b` borders.
- Typography should use system sans-serif and monospace stacks; no external Google Font dependency.
- Layout should be single-column even for sections currently rendered as grids.
- Section structure should rely on typography, spacing, and subtle borders rather than cards, blur, shadows, or colorful accents.
- CTA remains active and conversion-oriented, likely as a high-contrast editorial button/link with measured border/fill treatment, not a passive inline link.
- Microinteractions may include restrained color/border/focus transitions only; no hover lift, card movement, or entrance choreography.

## Acceptance Criteria

- The web CV visually matches `DESIGN.md`’s sober editorial zinc direction.
- No blue/indigo primary gradients remain in the main visual system.
- External Google Fonts import/preconnects are removed or no longer used.
- Web layout remains single-column at all viewport sizes.
- Print layout remains single-column and preserves existing ATS/page-break intent.
- CTA is visually prominent enough to invite contact while matching the restrained editorial style.
- Keyboard focus indicators remain visible and high contrast.
- Motion respects `prefers-reduced-motion`; no intrusive animation is introduced.
- Text contrast meets WCAG AA for normal text and UI components.
- No theme toggle/light-mode requirement remains in this change.

## Risks

- Over-reducing the CTA could harm conversion if it becomes visually passive.
- Muted zinc text may become too low-contrast if applied to dense CV details without review.
- Removing card/grid patterns may require careful spacing so content remains scannable.
- Print CSS may regress if web refactor accidentally changes page-break rules or display properties.
- Some existing CSS selectors may be tightly coupled to current card/grid naming and require careful consolidation.

## Rollback Plan

- Revert the stylesheet and any small template/meta changes from this change.
- Because the change is presentation-focused, rollback should not require data migration.
- If the CTA treatment underperforms visually, adjust only CTA styling/copy while keeping the rest of the zinc editorial system.

## Success Criteria

- The site reads as an engineering CV first and a portfolio page second.
- The visual system is consistent with `DESIGN.md` without relying on saturated accent gradients.
- Recruiter scanning is improved through clear hierarchy, one-column flow, and strong typographic rhythm.
- Print/PDF export remains clean, readable, and ATS-friendly.
- Implementation diff remains within the 800-line review budget.

## First-Slice Recommendation

Implement the first slice as a focused presentation refactor:

1. Replace global tokens: colors, fonts, spacing, border radius, shadow usage.
2. Remove font imports/preconnects and update theme color.
3. Convert all grid/card-heavy sections to single-column editorial blocks.
4. Restyle CTA and focus states.
5. Re-check print CSS and reduced-motion behavior.

Avoid content rewriting or structural data changes in the first slice unless strictly necessary for styling correctness.

## Review Workload Forecast

Estimated implementation diff: approximately 400–500 changed lines, concentrated in `web/src/styles/global.css` with a smaller update in `web/src/pages/index.astro`. This is within the configured 800-line review budget.

## Open Questions

No blocking open questions remain for this proposal. Optional product refinements for later review:

- Should CTA copy stay as `Contacto`, or become more explicitly conversion-oriented in Spanish while remaining sober?
- Should any small typographic distinction be preserved for technology tags, or should they become plain inline metadata?

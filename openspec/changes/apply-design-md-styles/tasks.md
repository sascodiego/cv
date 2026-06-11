# Implementation Tasks: Apply DESIGN.md Styles to Web CV

**Change ID:** `apply-design-md-styles`
**Status:** Draft
**Created:** 2026-06-10

---

## Review Workload Forecast

| Field | Value |
|-------|-------|
| Estimated changed lines | 400–500 lines |
| 400-line budget risk | Low |
| 800-line review budget risk | Low |
| Chained PRs recommended | No |
| Suggested split | Single PR |
| Delivery strategy | auto-chain (single PR within budget) |
| Chain strategy | stacked-to-main |

**Rationale:** The change is concentrated in `web/src/styles/global.css` (~560 lines total) with minor template updates in `web/src/pages/index.astro`. The refactor replaces a gradient/card system with a zinc editorial system, but does not add new features or pages. A single PR is appropriate and within the 800-line review budget.

```
Decision needed before apply: No
Chained PRs recommended: No
Chain strategy: stacked-to-main
400-line budget risk: Low
800-line budget risk: Low
```

---

## Task List

### Phase 1: Foundation — Color Tokens and Typography

- [ ] **T1.1: Replace CSS color variables with zinc palette**
  - File: `web/src/styles/global.css`
  - Actions:
    - Replace `--bg-color: #0b0f19` with `#09090b` (Zinc-950)
    - Replace `--text-main: #f1f5f9` with `#f4f4f5` (Zinc-100)
    - Replace `--text-muted: #94a3b8` with `#a1a1aa` (Zinc-400)
    - Replace `--card-border` and border colors with `#18181b` (Zinc-900)
    - Remove `--primary-color: #38bdf8` and `--primary-gradient` (blue/indigo)
    - Remove `--text-accent` (no longer needed)
    - Remove `--shadow-sm` and `--shadow-md` (no shadows in editorial system)
  - Acceptance: Computed CSS variables match DESIGN.md zinc palette; no blue/indigo references remain.

- [ ] **T1.2: Remove Google Fonts import and use system font stacks**
  - File: `web/src/styles/global.css`
  - Actions:
    - Remove `@import url("https://fonts.googleapis.com/...")` at top of file
    - Define `--font-body: 'Inter', system-ui, -apple-system, sans-serif`
    - Define `--font-mono: 'SFMono-Regular', Consolas, 'Liberation Mono', Menlo, monospace`
    - Remove `--font-title` variable (merge with `--font-body`)
    - Update all references from `var(--font-title)` to `var(--font-body)` or `var(--font-mono)`
  - Acceptance: No external font URLs in CSS; computed font-family shows system fonts.

- [ ] **T1.3: Remove Google Fonts preconnect links from template**
  - File: `web/src/pages/index.astro`
  - Actions:
    - Remove `<link rel="preconnect" href="https://fonts.googleapis.com">`
    - Remove `<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>`
  - Acceptance: No Google Fonts links in rendered HTML.

- [ ] **T1.4: Update theme-color meta tag to match zinc background**
  - File: `web/src/pages/index.astro`
  - Actions:
    - Change `<meta name="theme-color" content="#0b0f19">` to `content="#09090b"`
  - Acceptance: Browser theme color matches DESIGN.md background.

---

### Phase 2: Layout — Single-Column Enforcement

- [ ] **T2.1: Convert skills section from grid to single-column stack**
  - File: `web/src/styles/global.css`
  - Actions:
    - Change `.skills-grid` from `grid-template-columns: repeat(auto-fill, minmax(250px, 1fr))` to `display: flex; flex-direction: column; gap: 1.5rem`
    - Remove responsive breakpoint that sets `grid-template-columns: 1fr` (no longer needed)
  - Acceptance: Skills display vertically at all viewport widths; no horizontal grid.

- [ ] **T2.2: Convert independent projects from grid to single-column stack**
  - File: `web/src/styles/global.css`
  - Actions:
    - Change `.indie-projects-grid` from `grid-template-columns: repeat(auto-fill, minmax(280px, 1fr))` to `display: flex; flex-direction: column; gap: 1.5rem`
    - Remove responsive breakpoint that sets `grid-template-columns: 1fr`
  - Acceptance: Projects display vertically at all viewport widths.

- [ ] **T2.3: Convert recommendations from grid to single-column stack**
  - File: `web/src/styles/global.css`
  - Actions:
    - Change `.recommendations-grid` from `grid-template-columns: repeat(auto-fill, minmax(280px, 1fr))` to `display: flex; flex-direction: column; gap: 1.5rem`
    - Remove responsive breakpoint that sets `grid-template-columns: 1fr`
  - Acceptance: Recommendations display vertically at all viewport widths.

- [ ] **T2.4: Remove print multi-column layouts**
  - File: `web/src/styles/global.css` (inside `@media print`)
  - Actions:
    - Change `.indie-projects-grid` print override from `grid-template-columns: repeat(2, 1fr)` to `display: flex; flex-direction: column; gap: 1rem`
    - Verify all print sections use single-column flow
  - Acceptance: Print output is strictly single-column; no 2-column layouts in PDF export.

---

### Phase 3: Editorial Structure — Remove Cards and Decorative Effects

- [ ] **T3.1: Remove card backgrounds, backdrop blur, and translucent styling**
  - File: `web/src/styles/global.css`
  - Actions:
    - Remove `--card-bg: rgba(22, 30, 49, 0.6)` variable
    - Set `.section-card`, `.skill-category-card`, `.indie-project-card`, `.recommendation-card` background to `transparent` or remove background property
    - Remove `backdrop-filter: blur(10px)` from all selectors
    - Remove `.about-me` border and padding overrides (merge with standard section styling)
  - Acceptance: No translucent backgrounds; no backdrop blur; sections use main background color.

- [ ] **T3.2: Remove drop shadows from all elevated elements**
  - File: `web/src/styles/global.css`
  - Actions:
    - Remove all `box-shadow` properties (including `var(--shadow-sm)` and `var(--shadow-md)`)
    - Remove shadow-related CSS variables if no longer referenced
  - Acceptance: No `box-shadow` properties remain; depth achieved through spacing only.

- [ ] **T3.3: Add subtle zinc borders for section separation**
  - File: `web/src/styles/global.css`
  - Actions:
    - Ensure section separators use `border: 1px solid #18181b` (Zinc-900)
    - Update `.hero` bottom border to zinc color
    - Update `.site-footer` top border to zinc color
    - Update `.internal-projects` border to zinc color
    - Update `.about-me` bottom border to zinc color
  - Acceptance: All borders use zinc palette; no colored borders remain.

- [ ] **T3.4: Remove decorative quote mark from recommendations**
  - File: `web/src/styles/global.css`
  - Actions:
    - Remove `.recommendation-card::before` pseudo-element (decorative quote mark)
  - Acceptance: No decorative quote mark in recommendations; clean editorial presentation.

---

### Phase 4: Typography Hierarchy — Match DESIGN.md Spec

- [ ] **T4.1: Restyle H1 (name) to DESIGN.md specification**
  - File: `web/src/styles/global.css`
  - Actions:
    - Set `.hero h1` font-size to `2.25rem` (36px)
    - Set font-weight to `800`
    - Remove `background: var(--primary-gradient)` and `-webkit-background-clip: text` and `-webkit-text-fill-color: transparent`
    - Set color to `var(--text-main)` (Zinc-100)
    - Keep `line-height: 1.1` for readability
  - Acceptance: H1 is 36px, 800 weight, solid zinc color; no gradient text effect.

- [ ] **T4.2: Restyle H2 (section headings) to mono uppercase**
  - File: `web/src/styles/global.css`
  - Actions:
    - Set `section h2` font-size to `0.75rem` (12px)
    - Set font-family to `var(--font-mono)` (system monospace)
    - Set font-weight to `bold`
    - Set `text-transform: uppercase`
    - Set `letter-spacing: 0.1em` (tracking-widest)
    - Remove left border accent (`border-inline-start: 4px solid var(--primary-color)`)
    - Remove `padding-inline-start: 0.75rem`
    - Add bottom border: `border-block-end: 1px solid var(--card-border)` for editorial separation
    - Add `margin-block-end: 1.5rem` for breathing room
  - Acceptance: H2 is 12px, mono, uppercase, widened tracking; no colored left border.

- [ ] **T4.3: Restyle H3 (subcategory headings) to 1rem bold**
  - File: `web/src/styles/global.css`
  - Actions:
    - Ensure all H3 elements (`.skill-category-title`, `.indie-project-title h3`, `.soft-skills-section h3`) use `font-size: 1rem`
    - Set font-weight to `bold`
    - Use `var(--font-body)` (sans-serif)
  - Acceptance: H3 elements are 16px, bold; consistent across all sections.

- [ ] **T4.4: Update body text to relaxed line-height**
  - File: `web/src/styles/global.css`
  - Actions:
    - Set `body` line-height to `1.7` (relaxed/leading-relaxed equivalent)
    - Ensure body font-size is `1rem` (16px) for optimal readability
  - Acceptance: Body text is readable with comfortable line-height; not cramped.

---

### Phase 5: CTA Restyling — Sober Yet Conversion-Oriented

- [ ] **T5.1: Restyle contact button with zinc palette**
  - File: `web/src/styles/global.css`
  - Actions:
    - Remove `background: var(--primary-gradient)` from `.contact-form-btn`
    - Set background to `var(--text-main)` (Zinc-100) or high-contrast zinc variant
    - Set text color to `var(--bg-color)` (Zinc-950) for contrast
    - Keep `border-radius` modest (e.g., `4px` or `6px`) for editorial feel
    - Keep padding `0.75rem 2rem` for prominence
    - Remove box-shadow properties
  - Acceptance: CTA uses zinc colors only; high contrast between background and text.

- [ ] **T5.2: Remove hover lift from contact button**
  - File: `web/src/styles/global.css`
  - Actions:
    - Remove `transform: translateY(-2px)` from `.contact-form-btn:hover`
    - Remove `box-shadow` increase on hover
    - Keep restrained color transition (e.g., slight border opacity change or brightness shift)
    - Ensure transition duration is ≤ 0.2s
  - Acceptance: No vertical movement on hover; only subtle color/border change.

- [ ] **T5.3: Ensure CTA has accessible focus indicator**
  - File: `web/src/styles/global.css`
  - Actions:
    - Verify `.contact-form-btn:focus-visible` has `outline: 2px solid` with high contrast
    - Ensure `outline-offset: 2px` or more for visibility
    - Test that focus outline meets WCAG AA 3:1 contrast against background
  - Acceptance: Focus indicator is at least 2px wide, offset, and high contrast.

- [ ] **T5.4: Position CTA prominently within hero section**
  - File: `web/src/styles/global.css`
  - Actions:
    - Ensure `.contact-form-btn` has `margin-block-start: 1.5rem` for separation from subtitle
    - Verify button is centered and visually distinct from surrounding contact info
  - Acceptance: CTA is clearly visible and accessible in hero section.

---

### Phase 6: Motion and Interactions — Restrained Microinteractions

- [ ] **T6.1: Remove entrance animations (fadeInUp)**
  - File: `web/src/styles/global.css`
  - Actions:
    - Remove `@keyframes fadeInUp` definition
    - Remove `main > * { animation: fadeInUp 0.6s ease-out both; }`
    - Remove all `main > :nth-child(N) { animation-delay: ... }` staggered delays
  - Acceptance: No entrance animations; elements appear immediately on load.

- [ ] **T6.2: Remove hover lift from cards and interactive elements**
  - File: `web/src/styles/global.css`
  - Actions:
    - Remove `transform: translateY(-2px)` from `.skill-category-card:hover`
    - Remove `transform: translateY(-2px)` from `.indie-project-card:hover`
    - Remove `transform` from `.experience-item:hover`
    - Remove any other `translateY` hover effects
  - Acceptance: No elements move vertically on hover; only color/border changes allowed.

- [ ] **T6.3: Add reduced-motion media query**
  - File: `web/src/styles/global.css`
  - Actions:
    - Add `@media (prefers-reduced-motion: reduce)` block
    - Inside: set `animation-duration: 0.01ms !important; animation-iteration-count: 1 !important; transition-duration: 0.01ms !important; scroll-behavior: auto !important;`
    - Apply to `*`, `*::before`, `*::after` selectors
  - Acceptance: All animations and transitions disabled when user prefers reduced motion.

- [ ] **T6.4: Restrain remaining microinteractions**
  - File: `web/src/styles/global.css`
  - Actions:
    - Ensure all hover/active transitions have `transition: color 0.2s ease, border-color 0.2s ease, background-color 0.2s ease`
    - Remove any transition durations > 0.2s
    - Remove any `transition: transform` (no movement transitions)
  - Acceptance: All transitions are ≤ 0.2s and limited to color/border/background; no movement transitions.

---

### Phase 7: Print Stylesheet — Preserve ATS Readiness

- [ ] **T7.1: Verify print colors remain paper-friendly**
  - File: `web/src/styles/global.css` (inside `@media print`)
  - Actions:
    - Ensure print background is white/off-white (not dark zinc)
    - Ensure print text is dark enough for readability (contrast ≥ 4.5:1)
    - Do NOT change existing print color logic unless broken
  - Acceptance: Print output uses light-on-dark colors; readable on paper.

- [ ] **T7.2: Verify print removes interactive elements**
  - File: `web/src/styles/global.css` (inside `@media print`)
  - Actions:
    - Confirm `.contact-form-btn { display: none; }` exists
    - Confirm `.site-footer { display: none; }` exists
    - Confirm `.recommendation-card::before { display: none; }` exists
  - Acceptance: Contact button and footer hidden in print; only CV content visible.

- [ ] **T7.3: Verify print maintains single-column layout**
  - File: `web/src/styles/global.css` (inside `@media print`)
  - Actions:
    - After Phase 2 changes, ensure all sections remain single-column in print
    - Verify `.indie-projects-grid` does not revert to 2-column layout
    - Verify page-break rules (`page-break-inside: avoid`) still apply
  - Acceptance: Print output is single-column; no multi-column grids; sections not orphaned.

- [ ] **T7.4: Verify print removes animations and transitions**
  - File: `web/src/styles/global.css` (inside `@media print`)
  - Actions:
    - Confirm `* { animation: none !important; transition: none !important; }` exists
  - Acceptance: No animations or transitions in print output.

---

### Phase 8: Accessibility Compliance

- [ ] **T8.1: Verify color contrast meets WCAG AA**
  - File: Visual verification (browser dev tools)
  - Actions:
    - Test primary text (Zinc-100 on Zinc-950) — expect contrast ≥ 7:1 (AAA)
    - Test muted text (Zinc-400 on Zinc-950) — expect contrast ≥ 4.5:1 (AA)
    - Test CTA button (zinc background/text) — expect contrast ≥ 3:1 (AA for UI components)
    - Test focus outline against background — expect contrast ≥ 3:1
  - Acceptance: All color combinations meet or exceed WCAG AA; ideally AAA for body text.

- [ ] **T8.2: Verify keyboard focus indicators are visible**
  - File: Visual verification (keyboard navigation)
  - Actions:
    - Tab through all interactive elements (contact button, project links)
    - Verify each element shows a visible focus indicator (outline)
    - Ensure focus outline is at least 2px wide and offset by 2px
  - Acceptance: All interactive elements show clear focus indicators when tabbed.

- [ ] **T8.3: Verify semantic structure is preserved**
  - File: Visual verification (HTML inspection)
  - Actions:
    - Confirm heading hierarchy: H1 (name) → H2 (sections) → H3 (subsections)
    - Confirm no heading levels are skipped
    - Confirm each major section has an H2 heading
  - Acceptance: Semantic heading structure is logical and complete.

- [ ] **T8.4: Verify reduced-motion preference is respected**
  - File: Visual verification (browser dev tools)
  - Actions:
    - Enable `prefers-reduced-motion: reduce` in browser dev tools
    - Reload page and observe no animations
    - Observe no hover lift effects
    - Observe scroll behavior is `auto` (not smooth)
  - Acceptance: All motion disabled when user prefers reduced motion.

---

### Phase 9: Build Verification

- [ ] **T9.1: Build web CV without errors**
  - Command: `cd /src/cv-pipeline/web && pnpm build`
  - Actions:
    - Run build command
    - Verify build completes successfully
    - Verify no CSS syntax errors or missing references
  - Acceptance: Build succeeds with no errors or warnings.

- [ ] **T9.2: Inspect built output for correct styles**
  - File: Built HTML/CSS (inspect in browser or dev tools)
  - Actions:
    - Open built site in browser (or `pnpm dev`)
    - Inspect computed CSS variables — should match zinc palette
    - Inspect font-family — should show system fonts
    - Inspect layout — should be single-column
    - Inspect H1/H2/H3 sizes — should match DESIGN.md
  - Acceptance: Visual inspection confirms all style requirements are applied.

- [ ] **T9.3: Test print/PDF export**
  - File: Browser print preview or PDF export
  - Actions:
    - Open built site in browser
    - Open print preview (Cmd+P / Ctrl+P)
    - Verify single-column layout
    - Verify no interactive elements (button, footer)
    - Verify page breaks do not cut sections
  - Acceptance: Print output is clean, single-column, ATS-friendly.

---

### Phase 10: Rollback Verification (Optional but Recommended)

- [ ] **T10.1: Document rollback plan**
  - File: N/A (documentation task)
  - Actions:
    - Note that rollback involves reverting `web/src/styles/global.css` and `web/src/pages/index.astro`
    - Confirm no data migration required (presentation-only change)
    - If CTA underperforms, note that only CTA styling can be adjusted independently
  - Acceptance: Rollback plan is documented and clear.

---

## Verification Steps

### Manual Verification Checklist

After implementation, verify the following:

- [ ] **Visual Inspection**
  - [ ] Background is `#09090b` (Zinc-950)
  - [ ] Primary text is `#f4f4f5` (Zinc-100)
  - [ ] Muted text is `#a1a1aa` (Zinc-400)
  - [ ] Borders are `#18181b` (Zinc-900)
  - [ ] No blue/indigo gradients anywhere
  - [ ] No card backgrounds or shadows
  - [ ] No backdrop blur effects
  - [ ] Layout is single-column at all viewport sizes

- [ ] **Typography**
  - [ ] H1 (name) is 36px / 2.25rem, 800 weight
  - [ ] H2 (sections) is 12px / 0.75rem, mono, uppercase, tracking-widest
  - [ ] H3 (subsections) is 16px / 1rem, bold
  - [ ] Body text is 16px / 1rem with relaxed line-height
  - [ ] System fonts are used (no Google Fonts)

- [ ] **CTA**
  - [ ] Contact button uses zinc palette only
  - [ ] Button is visually prominent but editorial
  - [ ] No hover lift movement
  - [ ] Focus indicator is visible and high contrast

- [ ] **Motion**
  - [ ] No entrance animations on page load
  - [ ] No hover lift effects on cards
  - [ ] Transitions are ≤ 0.2s
  - [ ] Reduced-motion preference is respected

- [ ] **Accessibility**
  - [ ] Color contrast meets WCAG AA (ideally AAA)
  - [ ] Keyboard focus indicators visible
  - [ ] Semantic heading structure preserved
  - [ ] Reduced-motion support functional

- [ ] **Print/PDF**
  - [ ] Single-column layout preserved
  - [ ] No interactive elements (button, footer)
  - [ ] Page breaks avoid cutting sections
  - [ ] Light-on-dark colors for readability

### Build Verification

- [ ] Build succeeds: `cd /src/cv-pipeline/web && pnpm build`
- [ ] No errors or warnings in build output
- [ ] Built CSS contains zinc palette variables
- [ ] Built HTML has no Google Fonts links
- [ ] Built HTML has correct theme-color meta tag

---

## Acceptance Evidence

After implementation, provide the following evidence:

1. **Screenshot of rendered web CV** showing zinc palette, single-column layout, editorial structure
2. **Screenshot of browser dev tools** showing computed CSS variables match DESIGN.md
3. **Screenshot of print preview** showing clean single-column output
4. **Build output** showing successful `pnpm build` with no errors
5. **Accessibility audit report** (Lighthouse or manual) confirming WCAG AA compliance
6. **Reduced-motion test screenshot** showing animations disabled

---

## Dependencies and Ordering

Tasks are ordered to minimize risk and enable incremental verification:

1. **Phase 1 (Foundation)** — Must complete first; establishes color/typography tokens
2. **Phase 2 (Layout)** — Can proceed after Phase 1; removes multi-column structure
3. **Phase 3 (Editorial Structure)** — Can proceed after Phase 2; removes cards/effects
4. **Phase 4 (Typography)** — Can proceed after Phase 3; applies heading hierarchy
5. **Phase 5 (CTA)** — Can proceed after Phase 4; restyles contact button
6. **Phase 6 (Motion)** — Can proceed after Phase 5; removes animations
7. **Phase 7 (Print)** — Can proceed after Phase 6; verifies print compatibility
8. **Phase 8 (Accessibility)** — Can proceed after Phase 7; verifies compliance
9. **Phase 9 (Build)** — Must complete after all implementation phases
10. **Phase 10 (Rollback)** — Optional documentation task

---

## Risks and Mitigations

| Risk | Impact | Mitigation |
|------|--------|------------|
| CTA becomes too passive and harms conversion | Medium | Ensure CTA uses high-contrast zinc background/weight; verify prominence in hero section |
| Muted zinc text becomes too low-contrast | Medium | Test contrast ratios; adjust zinc-400 if needed; ensure ≥ 4.5:1 for body text |
| Removing cards/grids harms scannability | Low | Use spacing and typography to create visual rhythm; add subtle borders for separation |
| Print CSS regresses due to web refactor | Medium | Carefully review print rules after web changes; test print preview before commit |
| Existing CSS selectors tightly coupled to card naming | Low | Consolidate selectors during refactor; remove unused classes |

---

## Notes

- This is a presentation-only change; no data or content modifications required
- The CTA should remain conversion-oriented but fit the sober editorial aesthetic
- Accessibility is preserved or improved; WCAG AA is the minimum target
- Print/ATS output must remain clean and readable; no regression allowed
- The change is within the 800-line review budget; single PR is appropriate
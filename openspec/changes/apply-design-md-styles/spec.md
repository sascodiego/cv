# Web CV Visual System Specification

## Purpose

Define the sober, editorial, and ATS-preserving visual system for the Diego Sasso web CV, based on `DESIGN.md`. This specification replaces the current "Dark Premium" gradient-heavy styling with a zinc-neutral, system-font, single-column reading experience optimized for recruiter scanning, accessibility, and print continuity.

## Requirements

### Requirement: Zinc-Neutral Color Tokens

The system MUST use the zinc-neutral palette defined in `DESIGN.md` as the exclusive color foundation.

#### Scenario: Dark mode color tokens

- GIVEN the web CV is loaded in a browser
- WHEN the CSS variables are inspected
- THEN the primary background MUST be `#09090b` (Zinc-950)
- AND the primary text color MUST be `#f4f4f5` (Zinc-100)
- AND the muted text color MUST be `#a1a1aa` (Zinc-400)
- AND the border/line color MUST be `#18181b` (Zinc-900)

#### Scenario: No blue/indigo gradient accents

- GIVEN the web CV is rendered
- WHEN visual inspection of all sections is performed
- THEN no blue or indigo gradient colors MUST be present
- AND no synthetic or neon gradient effects MUST be present

#### Scenario: Meta theme color matches background

- GIVEN the web CV is loaded
- WHEN the `<meta name="theme-color">` element is inspected
- THEN the content MUST be `#09090b`

### Requirement: System Typography Stack

The system MUST use system font stacks exclusively and remove all external font dependencies.

#### Scenario: System sans-serif fonts

- GIVEN the web CV is loaded
- WHEN the computed font-family of body text is inspected
- THEN it MUST prioritize `Inter`, `system-ui`, `-apple-system`, or a native OS sans-serif
- AND MUST NOT reference Google Fonts or external font URLs

#### Scenario: System monospace fonts for labels

- GIVEN the web CV is loaded
- WHEN the computed font-family of section headings (H2) is inspected
- THEN it MUST prioritize `SFMono-Regular`, `Consolas`, or a native OS monospace
- AND MUST NOT reference external font URLs

#### Scenario: Remove Google Fonts imports

- GIVEN the web CV markup is inspected
- WHEN the `<head>` section is examined
- THEN NO `<link rel="preconnect" href="https://fonts.googleapis.com">` MUST be present
- AND NO `<link rel="preconnect" href="https://fonts.gstatic.com">` MUST be present
- AND NO `@import url("https://fonts.googleapis.com/...")` MUST be present

#### Scenario: Typographic hierarchy matches DESIGN.md

- GIVEN the web CV is rendered
- WHEN the H1 element (name) is measured
- THEN the font-size MUST be `2.25rem` (36px)
- AND the font-weight MUST be `800` (extrabold)

- WHEN the H2 element (section headings) is measured
- THEN the font-size MUST be `0.75rem` (12px)
- AND the font-style MUST be monospace
- AND the font-weight MUST be `bold`
- AND the text MUST be uppercase
- AND letter-spacing MUST be widened (tracking-widest)

- WHEN the H3 element (roles/puestos) is measured
- THEN the font-size MUST be `1rem` (16px)
- AND the font-weight MUST be `bold`

- WHEN body text is measured
- THEN the font-size MUST be between `0.875rem` and `1rem`
- AND the line-height MUST be relaxed (`leading-relaxed`)

### Requirement: Single-Column Layout

The system MUST enforce a strict single-column reading model for all viewport sizes, with no multi-column grids.

#### Scenario: Skills section is single-column

- GIVEN the web CV is rendered at any viewport width
- WHEN the skills grid container is inspected
- THEN `grid-template-columns` MUST NOT contain multiple columns
- AND the layout MUST display as a vertical stack

#### Scenario: Independent projects is single-column

- GIVEN the web CV is rendered at any viewport width
- WHEN the independent projects container is inspected
- THEN `grid-template-columns` MUST NOT contain multiple columns
- AND projects MUST display as a vertical stack

#### Scenario: Recommendations is single-column

- GIVEN the web CV is rendered at any viewport width
- WHEN the recommendations container is inspected
- THEN `grid-template-columns` MUST NOT contain multiple columns
- AND recommendations MUST display as a vertical stack

#### Scenario: No horizontal grid breaking

- GIVEN the web CV is rendered on desktop viewport (> 768px)
- WHEN visual inspection of all section containers is performed
- THEN NO section MUST use a multi-column CSS grid layout
- AND all content MUST flow in a single vertical column

### Requirement: Editorial Section Structure

The system MUST use typography, spacing, and subtle borders to structure sections, replacing cards, backdrop blur, shadows, and decorative backgrounds.

#### Scenario: Remove card backgrounds

- GIVEN the web CV is rendered
- WHEN section cards are inspected
- THEN the background MUST NOT be translucent or use backdrop blur
- AND the background MUST be the same as the main background color `#09090b`

#### Scenario: Remove drop shadows

- GIVEN the web CV is rendered
- WHEN any elevated element is inspected
- THEN NO `box-shadow` property MUST be present for visual elevation
- AND depth MUST be achieved through spacing and typography only

#### Scenario: Use subtle borders for separation

- GIVEN the web CV is rendered
- WHEN section boundaries are inspected
- THEN borders MUST use the zinc border color `#18181b`
- AND border widths MUST be `1px` or smaller
- AND borders MUST NOT use colors outside the zinc palette

### Requirement: Conversion-Oriented CTA with Sober Styling

The system MUST retain a prominent, conversion-oriented contact CTA but restyle it to match the editorial zinc aesthetic.

#### Scenario: CTA is visually prominent

- GIVEN the web CV is rendered
- WHEN the contact button is viewed
- THEN the button MUST have a high-contrast background or border treatment
- AND the button MUST stand out from surrounding content
- AND the button MUST invite user action

#### Scenario: CTA uses zinc palette

- GIVEN the contact button is inspected
- WHEN its computed styles are examined
- THEN the button colors MUST use only zinc palette values
- AND NO blue, indigo, or gradient colors MUST be present

#### Scenario: CTA has accessible focus state

- GIVEN the contact button is focused via keyboard
- WHEN the focus indicator is visible
- THEN the outline MUST be at least `2px` wide
- AND the outline color MUST meet WCAG AA contrast against the background
- AND the outline offset MUST be at least `2px`

#### Scenario: CTA hover transition is restrained

- GIVEN the contact button is hovered
- WHEN the style change is observed
- THEN the transition duration MUST be `0.2s` or less
- AND the transform MUST NOT include vertical movement (`translateY(-2px)` is prohibited)
- AND NO box-shadow increase MUST occur

### Requirement: Restrained Motion and Microinteractions

The system MUST allow restrained microinteractions but prohibit intrusive animation and hover movement.

#### Scenario: Remove entrance animations

- GIVEN the web CV is loaded
- WHEN page elements are observed during load
- THEN NO `fadeInUp` or similar entrance animations MUST be present
- AND elements MUST appear immediately without staggered delays

#### Scenario: No hover lift movement

- GIVEN the web CV is rendered
- WHEN interactive elements are hovered
- THEN NO element MUST move vertically using `translateY`
- AND NO element MUST change position on hover

#### Scenario: Respect reduced motion preference

- GIVEN the user has `prefers-reduced-motion: reduce` enabled
- WHEN the web CV is rendered
- THEN all animations MUST have a duration of `0.01ms` or be disabled
- AND all transitions MUST have a duration of `0.01ms` or be disabled
- AND scroll-behavior MUST be `auto` (not smooth)

#### Scenario: Allowed microinteractions

- GIVEN the web CV is rendered
- WHEN focus, hover, or active states occur
- THEN color transitions of `0.2s` or less are permitted
- AND border transitions of `0.2s` or less are permitted
- AND focus outline transitions of `0.2s` or less are permitted

### Requirement: Accessibility Compliance

The system MUST maintain or improve accessibility to WCAG AA level, with particular attention to keyboard navigation, focus visibility, and color contrast.

#### Scenario: Keyboard focus indicators

- GIVEN the web CV is rendered
- WHEN tabbing through all interactive elements
- THEN every focusable element MUST display a visible focus indicator
- AND the focus indicator MUST be at least `2px` wide
- AND the focus indicator color MUST have at least `3:1` contrast ratio against its background

#### Scenario: Color contrast for normal text

- GIVEN the web CV is rendered
- WHEN normal body text is compared to its background
- THEN the contrast ratio MUST be at least `4.5:1` (WCAG AA)

#### Scenario: Color contrast for large text

- GIVEN the web CV is rendered
- WHEN headings or text larger than `18px` is compared to its background
- THEN the contrast ratio MUST be at least `3:1` (WCAG AA)

#### Scenario: Color contrast for UI components

- GIVEN the web CV is rendered
- WHEN the contact button or other UI components are compared to their background
- THEN the contrast ratio MUST be at least `3:1` (WCAG AA)

#### Scenario: Keyboard navigation works

- GIVEN the web CV is rendered
- WHEN using only the keyboard (Tab, Enter, Space)
- THEN all interactive elements MUST be reachable
- AND all interactive elements MUST be activable
- AND NO keyboard traps MUST exist

#### Scenario: Semantic structure preserved

- GIVEN the web CV markup is inspected
- WHEN the heading hierarchy is examined
- THEN headings MUST follow a logical order (H1 → H2 → H3)
- AND NO heading levels MUST be skipped
- AND each section MUST have an H2 heading

### Requirement: Print and ATS Preservation

The system MUST preserve print-friendly and ATS-compatible output, maintaining 1:1 mapping between web reading order and physical/PDF export.

#### Scenario: Print uses paper-friendly colors

- GIVEN the web CV is printed or exported to PDF
- WHEN the printed output is examined
- THEN the background MUST be white or off-white
- AND the text MUST be dark enough for print readability
- AND NO dark background with light text MUST be used

#### Scenario: Print removes interactive elements

- GIVEN the web CV is printed or exported to PDF
- WHEN the printed output is examined
- THEN the contact form button MUST NOT appear
- AND NO interactive-only elements MUST be present

#### Scenario: Print maintains single-column layout

- GIVEN the web CV is printed or exported to PDF
- WHEN the printed layout is examined
- THEN the reading order MUST be a single vertical column
- AND NO multi-column grids MUST appear in print

#### Scenario: Print preserves page breaks

- GIVEN the web CV is printed or exported to PDF
- WHEN the printed pages are examined
- THEN page breaks MUST avoid cutting experience items in half
- AND section content MUST not be orphaned at page bottom

#### Scenario: Print removes animations

- GIVEN the web CV is printed or exported to PDF
- WHEN the printed output is examined
- THEN all animations MUST be disabled
- AND all transitions MUST be disabled

#### Scenario: Print removes decorative elements

- GIVEN the web CV is printed or exported to PDF
- WHEN the printed output is examined
- THEN the site footer MUST NOT appear
- AND NO decorative quote marks or background patterns MUST appear

### Requirement: Removal of Obsolete Behavior

The system MUST remove theme toggle, light mode, and gradient-centric styling that no longer serves the editorial zinc direction.

#### Scenario: No theme toggle present

- GIVEN the web CV is rendered
- WHEN the page is inspected for theme controls
- THEN NO theme toggle button or switch MUST be present
- AND NO JavaScript for theme switching MUST be present

#### Scenario: No light mode tokens defined

- GIVEN the web CSS is inspected
- WHEN CSS variables are examined
- THEN NO separate light-mode color variables MUST be defined
- AND the system MUST support only dark mode

#### Scenario: No gradient text effects

- GIVEN the web CV is rendered
- WHEN the hero H1 element is inspected
- THEN NO `background-clip: text` gradient effect MUST be applied
- AND NO `-webkit-text-fill-color: transparent` MUST be present
- AND the text MUST use a solid zinc color

#### Scenario: No gradient backgrounds

- GIVEN the web CV is rendered
- WHEN all elements with background properties are inspected
- THEN NO `linear-gradient` MUST be used in background properties
- AND NO gradient fills MUST be present

### Requirement: Container and Spacing

The system MUST maintain appropriate container width and spacing for readability across all viewports.

#### Scenario: Container width respects readability

- GIVEN the web CV is rendered on desktop viewport
- WHEN the main container is measured
- THEN the max-width MUST be `900px` or less
- AND the container MUST be centered horizontally

#### Scenario: Responsive spacing on mobile

- GIVEN the web CV is rendered on mobile viewport (< 768px)
- WHEN the body padding is measured
- THEN the inline padding MUST be `1rem` or less
- AND the block padding MUST be `2rem` or less

#### Scenario: Section spacing supports scanning

- GIVEN the web CV is rendered
- WHEN the gap between major sections is measured
- THEN the vertical gap MUST be at least `2rem`
- AND the gap MUST be consistent throughout the document

### Requirement: Dark-Only Implementation

The system MUST implement dark mode exclusively, with no light mode deliverable or theme switching capability.

#### Scenario: Dark background is default

- GIVEN the web CV is loaded
- WHEN the body background color is inspected
- THEN the background MUST be `#09090b` (Zinc-950)

#### Scenario: No theme class toggling

- GIVEN the web CV markup is inspected
- WHEN the `<html>` or `<body>` elements are examined
- THEN NO class-based theme attributes MUST be present
- AND NO `data-theme` attributes MUST be present

#### Scenario: Dark colors hardcoded in CSS

- GIVEN the web CSS is inspected
- WHEN CSS custom properties are examined
- THEN the zinc dark palette MUST be the only palette defined
- AND NO alternate palette MUST be present

## Acceptance Criteria

All requirements above MUST be satisfied. In summary:

- [ ] Zinc color tokens are applied correctly
- [ ] System fonts are used exclusively; Google Fonts removed
- [ ] Single-column layout enforced at all viewport sizes
- [ ] Editorial structure replaces cards/gradients/shadows
- [ ] CTA is prominent but uses zinc palette
- [ ] Motion is restrained; no entrance animations or hover lift
- [ ] Accessibility meets WCAG AA (keyboard, focus, contrast)
- [ ] Print output preserves single-column and ATS readiness
- [ ] Theme toggle and light mode removed
- [ ] Gradient effects removed from all elements
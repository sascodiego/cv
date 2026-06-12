# Design: CV PDF Harvard Formalization

## Change ID

`cv-pdf-harvard-formalization`

## Purpose

This design document specifies the exact implementation details for removing all iconography from the Harvard-style CV PDF and restoring formal typography, spacing, and document structure. The design prioritizes formal Harvard GSAS/FAS document standards over preview fidelity and page density.

## Scope

- **Modified files**:
  - `web/src/pages/cv-pdf.astro` — print CSS and contact structure
  - `web/scripts/generate-pdf.js` — private data injection logic

- **Out of scope**:
  - `web/src/pages/index.astro` (public web page) — unchanged
  - Data pipeline or schema changes
  - Screen preview styles in `@media screen` block

---

## 1. Exact Header Structure

### 1.1 Name Placement

**Element**: `.cv-header h1`

**CSS specification**:
```css
.cv-header h1 {
  font-family: 'Times New Roman', Times, serif;
  font-size: 14pt;
  font-weight: bold;
  margin: 0;
  padding: 0;
  text-transform: uppercase;
  letter-spacing: 1pt;
  text-align: center;
}
```

**Behavior**:
- Name is uppercase, bold, 14pt
- Centered across the full document width
- 1pt letter-spacing for formal Harvard treatment
- No decorative elements or borders

---

### 1.2 Title Placement

**Element**: `.professional-title`

**CSS specification**:
```css
.professional-title {
  font-family: 'Times New Roman', Times, serif;
  font-size: 9.5pt;
  font-weight: normal;
  font-style: italic;
  margin: 2pt 0 4pt 0;
  padding: 0;
  text-align: center;
  color: #333333;
}
```

**Behavior**:
- Title is italic, normal weight, 9.5pt
- Centered below the name
- 2pt margin above, 4pt margin below
- Slightly darker gray (#333) for subtle hierarchy without decoration

---

### 1.3 Text-Only Contact Row

**Element**: `.contact-info` (container), `.contact-item` (individual items)

**CSS specification**:
```css
.contact-info {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 0; /* No gap — separators will handle spacing */
  margin-top: 4pt;
}

.contact-item {
  display: inline-flex;
  align-items: center;
  font-size: 8.5pt;
  color: #000000;
}

.contact-item a {
  color: #000000;
  text-decoration: none;
}
```

**Behavior**:
- Container is a flexbox with center alignment
- No gap between items — separators handle spacing
- Each contact item is an inline-flex element
- Links are black, no underlining (ATS-friendly)

**Separator strategy**:
- Pipe characters (` | `) injected between contact items
- Spaces before and after the pipe for readability
- Separator appended as a text-only span element

**DOM structure**:
```html
<div class="contact-info">
  <a class="contact-item" href="mailto:email@example.com">email@example.com</a>
  <span class="contact-separator"> | </span>
  <span class="contact-item">+1 555-123-4567</span>
  <span class="contact-separator"> | </span>
  <span class="contact-item">City, Country</span>
  <span class="contact-separator"> | </span>
  <a class="contact-item" href="https://portfolio.com">portfolio.com</a>
  <span class="contact-separator"> | </span>
  <a class="contact-item" href="https://github.com/username">github.com/username</a>
  <span class="contact-separator"> | </span>
  <a class="contact-item" href="https://linkedin.com/in/username">linkedin.com/in/username</a>
</div>
```

---

### 1.4 Separator Choice and Line-Break Behavior

**Separator character**: ` | ` (space, pipe, space)

**CSS specification**:
```css
.contact-separator {
  display: inline-flex;
  font-size: 8.5pt;
  color: #000000;
}
```

**Line-break behavior**:
- `flex-wrap: wrap` on `.contact-info` allows line wrapping on narrow pages
- Contact items flow naturally without forced breaks
- No explicit line heights or min-heights to avoid gaps
- Wrapping occurs at the container width, not between arbitrary items

---

## 2. Exact Section Rhythm

### 2.1 Section Headers

**Element**: `h2.section-header`

**CSS specification**:
```css
h2.section-header {
  font-family: 'Times New Roman', Times, serif;
  font-size: 10pt;
  font-weight: bold;
  text-transform: uppercase;
  color: #000000;
  margin: 8pt 0 3pt 0;
  padding: 0;
  border-bottom: 1pt solid #000000;
  padding-bottom: 2pt;
  letter-spacing: 0.3pt;
  page-break-after: avoid;
  page-break-inside: avoid;
}
```

**Behavior**:
- Bold, uppercase, 10pt
- 8pt margin above (clear separation from previous section)
- 3pt margin below, 2pt padding below border
- 1pt solid black bottom border
- 0.3pt letter-spacing for formal treatment
- Prevents page breaks within the header

---

### 2.2 Entry Stacking

**Element**: `.experience-entry`, `.project-entry`, `.education-entry`

**CSS specification**:
```css
.experience-entry,
.project-entry,
.education-entry {
  margin-bottom: 6pt;
  page-break-inside: avoid;
}
```

**Behavior**:
- Each entry has 6pt margin below
- Prevents page breaks within an entry
- Consistent spacing across all entry types

---

### 2.3 Period Alignment

**Element**: `.period`

**CSS specification**:
```css
.entry-header {
  display: flex;
  justify-content: space-between;
  align-items: baseline;
  margin-bottom: 4pt;
}

.period {
  font-family: 'Times New Roman', Times, serif;
  font-size: 9pt;
  font-weight: normal;
  color: #000000;
  margin: 0;
  padding: 0;
  white-space: nowrap;
}
```

**Behavior**:
- Flexbox with `space-between` pushes period to the right
- `align-items: baseline` aligns text baselines
- Period text is regular weight, 9pt
- `white-space: nowrap` prevents wrapping of date ranges
- 4pt margin below the entry header

---

### 2.4 Bullets and Technologies Lines

**Element**: `.entry-bullets`, `.technologies`

**CSS specification**:
```css
.entry-bullets {
  margin: 2pt 0 2pt 0;
  padding-left: 14pt;
  list-style-type: disc;
  color: #000000;
}

.entry-bullets li {
  font-family: 'Times New Roman', Times, serif;
  font-size: 9.5pt;
  line-height: 1.2;
  color: #000000;
  margin-bottom: 1pt;
  padding-left: 2pt;
}

.technologies {
  font-family: 'Times New Roman', Times, serif;
  font-size: 8.5pt;
  font-style: italic;
  color: #333333;
  margin: 2pt 0 0 0;
  padding: 0;
}
```

**Behavior**:
- Bullet list has 2pt margin above and below
- Standard disc markers (no custom characters)
- 14pt left padding for proper indentation
- Each bullet item has 1pt margin below
- Technologies line is italic, 8.5pt, slightly darker gray
- Technologies line has 2pt margin above the entry's end

---

## 3. Print CSS Blueprint

### 3.1 Rules That Stay Unchanged

The following CSS rules remain as-is (already Harvard-compliant):

```css
/* Base document styles */
body {
  font-family: 'Times New Roman', Times, serif;
  font-size: 9pt;
  line-height: 1.15;
  color: #000000;
  background-color: #ffffff;
  margin: 0;
  padding: 0;
  -webkit-print-color-adjust: exact;
  print-color-adjust: exact;
}

/* Header styles */
.cv-header {
  text-align: center;
  margin-bottom: 6pt;
}

.cv-header h1 {
  font-family: 'Times New Roman', Times, serif;
  font-size: 14pt;
  font-weight: bold;
  margin: 0;
  padding: 0;
  text-transform: uppercase;
  letter-spacing: 1pt;
}

/* Section header styles */
h2.section-header {
  font-family: 'Times New Roman', Times, serif;
  font-size: 10pt;
  font-weight: bold;
  text-transform: uppercase;
  color: #000000;
  margin: 8pt 0 3pt 0;
  padding: 0;
  border-bottom: 1pt solid #000000;
  padding-bottom: 2pt;
  letter-spacing: 0.3pt;
}

/* Entry header styles */
.entry-header {
  display: flex;
  justify-content: space-between;
  align-items: baseline;
  margin-bottom: 4pt;
}

h3.company-name,
h3.project-name,
h3.institution-name {
  font-family: 'Times New Roman', Times, serif;
  font-size: 9.5pt;
  font-weight: bold;
  color: #000000;
  margin: 0;
  padding: 0;
}

.period {
  font-family: 'Times New Roman', Times, serif;
  font-size: 9pt;
  font-weight: normal;
  color: #000000;
  margin: 0;
  padding: 0;
}

.role-title,
.degree-title {
  font-family: 'Times New Roman', Times, serif;
  font-size: 9.5pt;
  font-weight: normal;
  font-style: italic;
  color: #000000;
  margin: 0 0 2pt 0;
  padding: 0;
}

/* Entry spacing */
.experience-entry,
.project-entry,
.education-entry {
  margin-bottom: 6pt;
  page-break-inside: avoid;
}

/* Bullets */
.entry-bullets {
  margin: 2pt 0 2pt 0;
  padding-left: 14pt;
  list-style-type: disc;
  color: #000000;
}

.entry-bullets li {
  font-family: 'Times New Roman', Times, serif;
  font-size: 9.5pt;
  line-height: 1.2;
  color: #000000;
  margin-bottom: 1pt;
  padding-left: 2pt;
}

/* Technologies */
.technologies {
  font-family: 'Times New Roman', Times, serif;
  font-size: 8.5pt;
  font-style: italic;
  color: #333333;
  margin: 2pt 0 0 0;
  padding: 0;
}

/* Skills */
.skill-category {
  margin-bottom: 4pt;
  page-break-inside: avoid;
}

h3.skill-category-header {
  font-family: 'Times New Roman', Times, serif;
  font-size: 9.5pt;
  font-weight: bold;
  color: #000000;
  margin: 0 0 1pt 0;
  padding: 0;
}

.skill-list {
  font-family: 'Times New Roman', Times, serif;
  font-size: 9pt;
  line-height: 1.2;
  color: #000000;
  margin: 0;
  padding: 0;
}

/* Summary */
.summary-text {
  font-family: 'Times New Roman', Times, serif;
  font-size: 9.5pt;
  line-height: 1.25;
  color: #000000;
  margin: 0;
  padding: 0;
  text-align: justify;
}

/* Recommendations */
.recommendation-card {
  margin-bottom: 4pt;
  padding: 0;
  border: none;
  background: #ffffff;
  page-break-inside: avoid;
}

.recommendation-text {
  font-family: 'Times New Roman', Times, serif;
  font-size: 9pt;
  line-height: 1.2;
  color: #333333;
  margin: 0 0 2pt 0;
  padding: 0 0 0 12pt;
  border-left: 1pt solid #999999;
  font-style: italic;
}

.recommendation-text::before,
.recommendation-text::after {
  content: none;
}

.recommendation-author {
  font-family: 'Times New Roman', Times, serif;
  font-size: 8.5pt;
  color: #000000;
  margin: 0;
  padding: 0;
}

/* Page break control */
section {
  page-break-after: auto;
}

section:last-child {
  page-break-after: avoid;
}

.section-header {
  page-break-after: avoid;
  page-break-inside: avoid;
}

/* Footer note */
.pdf-note {
  font-family: 'Times New Roman', Times, serif;
  font-size: 8pt;
  color: #666666;
  margin: 8pt 0 0 0;
  padding: 0;
  text-align: center;
  font-style: italic;
}

/* Ensure no dark theme remnants */
*,
*::before,
*::after {
  box-sizing: border-box;
}
```

---

### 3.2 Rules That Are Removed

The following CSS rules are **completely removed** from `@media print`:

```css
/* REMOVE THIS - Icon-specific styling */
.contact-item svg {
  width: 9pt;
  height: 9pt;
  fill: #000000;
}

/* REMOVE THIS - Icon-related constraints on contact items */
.contact-item {
  display: inline-flex;
  align-items: center;
  gap: 3pt;  /* This gap was for icon spacing - remove it */
  font-size: 8.5pt;
  color: #000000;
}
```

---

### 3.3 Rules That Are Replaced

The following CSS rules are **modified or replaced**:

**Before (current)**:
```css
.contact-info {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 8pt;
  margin-top: 4pt;
}

.contact-item {
  display: inline-flex;
  align-items: center;
  gap: 3pt;
  font-size: 8.5pt;
  color: #000000;
}

.contact-item svg {
  width: 9pt;
  height: 9pt;
  fill: #000000;
}
```

**After (formalized)**:
```css
.contact-info {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 0;
  margin-top: 4pt;
}

.contact-item {
  display: inline-flex;
  align-items: center;
  font-size: 8.5pt;
  color: #000000;
}

.contact-item a {
  color: #000000;
  text-decoration: none;
}

.contact-separator {
  display: inline-flex;
  font-size: 8.5pt;
  color: #000000;
}
```

---

### 3.4 Explicit Note: Eliminate Icon-Specific Styling and Web-Like Preview Contamination

**Critical requirement**: The `@media print` CSS block must contain:

1. **No icon-related selectors**: No `.contact-item svg`, `.contact-item svg path`, or any SVG-related rules.

2. **No icon-specific widths/heights**: No width or height constraints designed to accommodate icon dimensions.

3. **No decorative gaps for icons**: The `gap` property on `.contact-info` is removed. Separators provide spacing instead.

4. **No web-derived decorative treatments**: No shadows, gradients, borders (except the 1pt bottom border on section headers), background colors, or text boxes.

5. **No ATS-hostile styling**: No content as images, no content in SVG elements, no decorative Unicode characters as bullets.

6. **Clean separation from screen styles**: The `@media screen` block remains unchanged, but the `@media print` block must not inherit or reference screen-specific styles.

**Verification**: After implementation, grep the `@media print` block for:
- `svg` — should have zero matches
- `gap: 3pt` or `gap: 8pt` — should have zero matches
- `width: 9pt` or `height: 9pt` in contact context — should have zero matches

---

## 4. Script Blueprint

### 4.1 Rewriting `createContactItem()` Without SVGs

**Before (current implementation)**:
```javascript
function createContactItem(svgPath, label, href) {
  const el = document.createElement(href ? "a" : "div");
  el.className = "contact-item";
  if (href) {
    el.href = href;
    el.target = "_blank";
    el.rel = "noopener noreferrer";
  }
  // SVG icon
  const template = document.createElement("template");
  template.innerHTML = svgPath;
  el.appendChild(template.content.firstChild);
  // Label
  const span = document.createElement("span");
  span.textContent = label;
  el.appendChild(span);
  contactInfo.appendChild(el);
}
```

**After (formalized)**:
```javascript
function createContactItem(label, href) {
  const el = document.createElement(href ? "a" : "span");
  el.className = "contact-item";
  if (href) {
    el.href = href;
    el.target = "_blank";
    el.rel = "noopener noreferrer";
  }
  el.textContent = label;
  contactInfo.appendChild(el);
}
```

**Key changes**:
- Removed `svgPath` parameter from function signature
- Changed element type from `div` (non-link) to `span` (inline)
- Removed SVG injection logic entirely
- Removed nested `span` element for label
- Set `textContent` directly on the element
- Removed `contactInfo.appendChild(el)` from the function (caller handles appending)

---

### 4.2 Appending Separators Cleanly

**Separator strategy**: Create a dedicated function for separators

```javascript
function createSeparator() {
  const separator = document.createElement("span");
  separator.className = "contact-separator";
  separator.textContent = " | ";
  contactInfo.appendChild(separator);
}
```

**Calling order**: After each contact item (except the last), append a separator

```javascript
// Email
if (data.email) {
  createContactItem(data.email, `mailto:${data.email}`);
  if (hasNextContactItem("email", data)) {
    createSeparator();
  }
}
// Phone
if (data.phone) {
  createContactItem(data.phone, null);
  if (hasNextContactItem("phone", data)) {
    createSeparator();
  }
}
// ... and so on for each contact field
```

**Alternative implementation (simpler)**:
```javascript
// Collect all present contact items into an array
const items = [];
if (data.email) items.push({ label: data.email, href: `mailto:${data.email}` });
if (data.phone) items.push({ label: data.phone, href: null });
if (data.address) items.push({ label: data.address, href: null });
if (data.website) items.push({ label: "Portafolio", href: data.website });
if (data.github) items.push({ label: "GitHub", href: data.github });
if (data.linkedin) items.push({ label: "LinkedIn", href: data.linkedin });

// Append with separators
items.forEach((item, index) => {
  createContactItem(item.label, item.href);
  if (index < items.length - 1) {
    createSeparator();
  }
});
```

---

### 4.3 What Stays Unchanged in Private Data Injection

The following logic in `injectPrivateData()` remains unchanged:

```javascript
// Update the PDF header name if provided
if (data.name) {
  const h1 = document.querySelector(".cv-header h1");
  if (h1) h1.textContent = data.name;
}

// Build the contact info block
const contactInfo = document.querySelector(".contact-info");
if (!contactInfo) return;

// Clear any existing content (e.g., the Google Form button)
while (contactInfo.firstChild) {
  contactInfo.removeChild(contactInfo.firstChild);
}

// Footer note
if (data.note) {
  const footer = document.createElement("footer");
  footer.className = "pdf-note";
  const p = document.createElement("p");
  p.textContent = data.note;
  footer.appendChild(p);
  document.querySelector(".cv-container").appendChild(footer);
}

// Inject recommender attribution into each blockquote
if (data.recommenders && data.recommenders.length > 0) {
  const quotes = document.querySelectorAll(".recommendation-card");
  quotes.forEach((quote, i) => {
    const rec = data.recommenders[i];
    if (!rec || !rec.name) return;
    const footer = document.createElement("footer");
    footer.className = "recommendation-author";
    const cite = document.createElement("cite");
    cite.className = "author-name";
    cite.textContent = rec.name;
    footer.appendChild(cite);
    if (rec.role) {
      const role = document.createElement("span");
      role.className = "author-role";
      role.textContent = rec.role;
      footer.appendChild(role);
    }
    if (rec.relation) {
      const relation = document.createElement("span");
      relation.className = "author-relation";
      relation.textContent = rec.relation;
      footer.appendChild(relation);
    }
    quote.appendChild(footer);
  });
}
```

**Removed SVG constants**:
```javascript
// REMOVE THESE - SVG icon constants
const svgEmail = '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">...</svg>';
const svgPhone = '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">...</svg>';
const svgLocation = '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 384 512">...</svg>';
const svgWeb = '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">...</svg>';
const svgGithub = '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 496 512">...</svg>';
const svgLinkedin = '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512">...</svg>';
```

---

## 5. Document-First Layout Rules

### 5.1 What Makes the PDF Feel Formal Instead of Web-Derived

**Typography hierarchy**:
- Clear size distinction: 14pt name > 10pt headers > 9.5pt entry headers > 9pt body
- Consistent font family: Times New Roman throughout
- Proper use of uppercase for headers only (not body text)
- Italic treatment for secondary information (title, role, degree, technologies)

**Spacing rhythm**:
- Consistent section spacing: 8pt above each header
- Consistent entry spacing: 6pt below each entry
- Consistent bullet spacing: 1pt below each bullet
- No arbitrary gaps or "web-like" whitespace

**Minimalist treatment**:
- No icons or graphical elements
- No decorative borders except the 1pt bottom border on section headers
- No background colors or shading
- No text boxes or visual effects
- Black text on white background only

**Content-first presentation**:
- Information density prioritizes readability over visual design
- Single-column layout for scannability
- Standard bullets (disc) not custom characters
- No content hidden in images or SVGs

**ATS-friendly output**:
- Text-only content is fully parseable
- No CSS-based content generation
- No content in pseudo-elements
- No decorative Unicode characters that confuse parsers

---

### 5.2 What Stylistic Treatments Are Forbidden

**Iconography**:
- ❌ SVG icons in contact information
- ❌ Any graphical elements in the document body
- ❌ Icon-based separators or bullets

**Decorative web elements**:
- ❌ Shadows or drop shadows
- ❌ Gradients or background colors
- ❌ Text boxes or containers with borders (except section header bottom border)
- ❌ Rounded corners or decorative padding

**Typography violations**:
- ❌ Inconsistent font families within the document
- ❌ Excessive letter-spacing (beyond 0.3-1pt for headers)
- ❌ Line heights outside 1.15-1.25 range
- ❌ Uppercase text in body or entry headers (reserved for section headers only)

**Spacing violations**:
- ❌ Arbitrary gaps between contact items (use separators instead)
- ❌ Inconsistent section or entry spacing
- ❌ Negative margins or overlapping elements
- ❌ Excessive whitespace to fill pages

**Color violations**:
- ❌ Colors other than black (#000), dark gray (#333), or medium gray (#666, #999)
- ❌ Colored text for emphasis
- ❌ Colored backgrounds or borders

**Layout violations**:
- ❌ Multi-column layouts
- ❌ Grid-based positioning
- ❌ Absolute positioning that breaks document flow
- ❌ Content as images (screenshots, charts, graphics)

---

## 6. Implementation Notes

### 6.1 Suggested Order of Code Edits

**Phase 1: Prepare the script logic** (in `generate-pdf.js`):

1. **Update `createContactItem()` function signature**:
   - Remove `svgPath` parameter
   - Simplify to accept only `label` and `href`
   - Remove SVG injection logic

2. **Create `createSeparator()` function**:
   - Add function to create pipe separator elements
   - Set class name to `contact-separator`

3. **Update contact item collection logic**:
   - Change from individual `if` blocks to array-based collection
   - Iterate through array and append with separators
   - Ensure no trailing separator after the last item

4. **Remove SVG icon constants**:
   - Delete `svgEmail`, `svgPhone`, `svgLocation`, `svgWeb`, `svgGithub`, `svgLinkedin`

5. **Test with local `.env`**:
   - Run `pnpm export-pdf` to verify contact injection works
   - Check that separators appear correctly

**Phase 2: Update CSS for print** (in `cv-pdf.astro`):

1. **Remove icon-related CSS**:
   - Delete `.contact-item svg` block
   - Remove `gap: 3pt` and `gap: 8pt` from contact-related rules

2. **Add `.contact-separator` rule**:
   - Define styling for pipe separators
   - Ensure consistent font size and color

3. **Update `.contact-item` rule**:
   - Remove icon-related width/height constraints
   - Ensure proper inline-flex behavior

4. **Update `.contact-info` rule**:
   - Set `gap: 0` (separators handle spacing)
   - Verify center alignment works correctly

5. **Verify no icon remnants**:
   - Grep `@media print` block for `svg`
   - Grep `@media print` block for icon-specific sizes

**Phase 3: Verify and test**:

1. **Build and generate PDF**:
   - Run `pnpm build` to ensure no syntax errors
   - Run `pnpm export-pdf` to generate final PDF

2. **Manual inspection**:
   - Open PDF and verify no icons appear
   - Check that contact info reads as a clean text-only line
   - Verify separators are visible and consistent

3. **ATS compatibility check**:
   - Copy text from PDF and verify it's selectable
   - Ensure contact info is parseable as plain text
   - No hidden content or image-based text

---

### 6.2 Edge Cases

**Missing phone**:
- **Behavior**: Contact block continues without phone entry
- **Implementation**: Array-based collection naturally skips missing fields
- **Separator handling**: No trailing separator if phone is absent

**Missing website**:
- **Behavior**: Contact block continues without website entry
- **Implementation**: Array-based collection naturally skips missing fields
- **Separator handling**: No trailing separator if website is absent

**Few contact items** (only email and phone):
- **Behavior**: Contact block shows "email@example.com | +1 555-123-4567"
- **Implementation**: Array-based collection handles any number of items
- **Separator handling**: Only one separator between the two items
- **Visual appearance**: Still centered and readable despite brevity

**Long contact values** (very long email or address):
- **Behavior**: Long values wrap to next line due to `flex-wrap: wrap`
- **Implementation**: No special handling needed
- **Separator handling**: Separators remain with their preceding item
- **Visual appearance**: May look less ideal, but still formal and ATS-friendly

**All contact fields missing**:
- **Behavior**: Contact block is empty (no elements)
- **Implementation**: Array-based collection results in empty array
- **Separator handling**: No separators appended
- **Visual appearance**: Blank line where contact info would be
- **Mitigation**: This is an edge case unlikely in production (users typically have at least email)

**Very long name or title**:
- **Behavior**: Name or title wraps to next line due to center alignment
- **Implementation**: No special handling needed
- **Visual appearance**: Still formal, may affect vertical spacing slightly
- **Note**: This is an existing behavior, not changed by this design

**Address contains commas**:
- **Behavior**: Address displays as-is (e.g., "City, State, Country")
- **Implementation**: No comma injection for address field
- **Separator handling**: Pipes separate address from other fields (e.g., "email@example.com | City, State, Country | portfolio.com")
- **Visual appearance**: Still readable, formal convention

**Email or URL contains pipe character**:
- **Behavior**: Unlikely edge case, would display as-is
- **Implementation**: No special escaping needed
- **Visual appearance**: Could confuse visual parsing, but ATS parsers extract href attributes
- **Mitigation**: This is an extremely rare edge case; unlikely in production data

---

### 6.3 Verification Checklist

After implementation, verify:

**Icon removal**:
- [ ] PDF contains zero icons
- [ ] `@media print` block contains no `svg` references
- [ ] `@media print` block contains no icon-specific width/height constraints

**Contact presentation**:
- [ ] Contact info displays as text-only with pipe separators
- [ ] Contact items are centered across the document width
- [ ] Links are black, no underlining
- [ ] No trailing separator after the last contact item

**Typography**:
- [ ] Name is 14pt, bold, uppercase, centered
- [ ] Title is 9.5pt, italic, centered
- [ ] Section headers are 10pt, bold, uppercase, with bottom border
- [ ] Entry headers are 9.5pt, bold, left-aligned
- [ ] Body text is 9pt, regular, left-aligned

**Spacing**:
- [ ] Section headers have 8pt margin above
- [ ] Entries have 6pt margin below
- [ ] Bullets have 1pt margin below
- [ ] Contact info has 4pt margin above

**Document-first**:
- [ ] PDF feels like a formal document, not a web page
- [ ] No decorative elements (shadows, gradients, text boxes)
- [ ] No content as images
- [ ] ATS-friendly output (selectable text, parseable structure)

**Build and generation**:
- [ ] `pnpm build` passes without errors
- [ ] `pnpm export-pdf` generates valid PDF
- [ ] PDF is A4 format with white background

**Private data**:
- [ ] Contact info injects correctly from `.env`
- [ ] Recommender details inject correctly
- [ ] Footer note injects correctly if present

---

## 7. Rollback Considerations

If issues arise during implementation:

1. **Revert `generate-pdf.js`**:
   - Restore SVG icon constants
   - Restore `createContactItem(svgPath, label, href)` signature
   - Restore SVG injection logic
   - Remove `createSeparator()` function

2. **Revert `cv-pdf.astro`**:
   - Restore `.contact-item svg` CSS rule
   - Restore `gap: 3pt` and `gap: 8pt` on contact-related rules
   - Remove `.contact-separator` CSS rule

3. **No data migration needed**:
   - No changes to `.env` file structure
   - No changes to `cv-processed.json` schema
   - No changes to Go pipeline

4. **Web page unaffected**:
   - `index.astro` remains untouched
   - No impact on public website functionality

---

## 8. Related Artifacts

- Proposal: `openspec/proposals/cv-pdf-harvard-formalization.md`
- Spec: `openspec/specs/cv-pdf-harvard-formalization.md`
- Parent change: `cv-pdf-harvard` — initial Harvard-style PDF implementation
- Original proposal: `openspec/proposals/cv-pdf-harvard.md`
- Verification report: `openspec/verify-reports/cv-pdf-harvard.md`

---

## 9. Success Criteria

The implementation is successful when:

1. PDF contains zero icons; all contact information is text-only with pipe separators
2. Typography and spacing match Harvard GSAS/FAS formal standards
3. Document feels formal and document-first, not web-derived
4. Page count is acceptable (1-2 pages preferred, but formal quality takes priority)
5. ATS systems can reliably parse all content sections
6. Implementation remains within the 400-line review budget
7. No regression in web page functionality or appearance
8. Private data injection continues to work correctly

---

## 10. Open Questions (Deferred)

The following questions are intentionally deferred to after implementation:

1. **Page density**: If the formal PDF exceeds 2 pages, should page density be revisited?
   - **Deferred decision**: Accept expanded page count in first pass; revisit density later if needed by slightly reducing font sizes or spacing within Harvard-acceptable ranges.

2. **Separator character**: Should contact separators use pipes (`|`) or commas (`,`)?
   - **Deferred decision**: Pipes specified in this design as the most common convention for formal CVs. Can be changed to commas in a follow-up change if testing suggests better readability.

3. **Font size tuning**: Should specific font sizes be adjusted based on final page count?
   - **Deferred decision**: Font sizes specified in this design follow Harvard standards. Can be fine-tuned in a follow-up change if page density becomes problematic.

---

## Appendix A: Complete Print CSS After Formalization

```css
@media print {
  @page {
    size: A4;
    margin: 0mm; /* Puppeteer handles 15mm margins via PDF options */
  }

  body {
    font-family: 'Times New Roman', Times, serif;
    font-size: 9pt;
    line-height: 1.15;
    color: #000000;
    background-color: #ffffff;
    margin: 0;
    padding: 0;
    -webkit-print-color-adjust: exact;
    print-color-adjust: exact;
  }

  .cv-header {
    text-align: center;
    margin-bottom: 6pt;
  }

  .cv-header h1 {
    font-family: 'Times New Roman', Times, serif;
    font-size: 14pt;
    font-weight: bold;
    margin: 0;
    padding: 0;
    text-transform: uppercase;
    letter-spacing: 1pt;
  }

  .professional-title {
    font-family: 'Times New Roman', Times, serif;
    font-size: 9.5pt;
    font-weight: normal;
    font-style: italic;
    margin: 2pt 0 4pt 0;
    padding: 0;
    text-align: center;
    color: #333333;
  }

  .education-status {
    font-size: 9pt;
    color: #555555;
    margin: 0;
    padding: 0;
  }

  .contact-info {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    gap: 0;
    margin-top: 4pt;
  }

  .contact-item {
    display: inline-flex;
    align-items: center;
    font-size: 8.5pt;
    color: #000000;
  }

  .contact-item a {
    color: #000000;
    text-decoration: none;
  }

  .contact-separator {
    display: inline-flex;
    font-size: 8.5pt;
    color: #000000;
  }

  .cv-container {
    width: 100%;
    max-width: 100%;
    margin: 0;
    padding: 0;
    background: #ffffff;
  }

  /* Section headers */
  h2.section-header {
    font-family: 'Times New Roman', Times, serif;
    font-size: 10pt;
    font-weight: bold;
    text-transform: uppercase;
    color: #000000;
    margin: 8pt 0 3pt 0;
    padding: 0;
    border-bottom: 1pt solid #000000;
    padding-bottom: 2pt;
    letter-spacing: 0.3pt;
    page-break-after: avoid;
    page-break-inside: avoid;
  }

  /* Entry headers */
  .entry-header {
    display: flex;
    justify-content: space-between;
    align-items: baseline;
    margin-bottom: 4pt;
  }

  h3.company-name,
  h3.project-name,
  h3.institution-name {
    font-family: 'Times New Roman', Times, serif;
    font-size: 9.5pt;
    font-weight: bold;
    color: #000000;
    margin: 0;
    padding: 0;
  }

  .period {
    font-family: 'Times New Roman', Times, serif;
    font-size: 9pt;
    font-weight: normal;
    color: #000000;
    margin: 0;
    padding: 0;
    white-space: nowrap;
  }

  .role-title,
  .degree-title {
    font-family: 'Times New Roman', Times, serif;
    font-size: 9.5pt;
    font-weight: normal;
    font-style: italic;
    color: #000000;
    margin: 0 0 2pt 0;
    padding: 0;
  }

  .experience-entry,
  .project-entry,
  .education-entry {
    margin-bottom: 6pt;
    page-break-inside: avoid;
  }

  .entry-bullets {
    margin: 2pt 0 2pt 0;
    padding-left: 14pt;
    list-style-type: disc;
    color: #000000;
  }

  .entry-bullets li {
    font-family: 'Times New Roman', Times, serif;
    font-size: 9.5pt;
    line-height: 1.2;
    color: #000000;
    margin-bottom: 1pt;
    padding-left: 2pt;
  }

  .technologies,
  .project-tech {
    font-family: 'Times New Roman', Times, serif;
    font-size: 8.5pt;
    font-style: italic;
    color: #333333;
    margin: 2pt 0 0 0;
    padding: 0;
  }

  .skill-category {
    margin-bottom: 4pt;
    page-break-inside: avoid;
  }

  h3.skill-category-header {
    font-family: 'Times New Roman', Times, serif;
    font-size: 9.5pt;
    font-weight: bold;
    color: #000000;
    margin: 0 0 1pt 0;
    padding: 0;
  }

  .skill-list {
    font-family: 'Times New Roman', Times, serif;
    font-size: 9pt;
    line-height: 1.2;
    color: #000000;
    margin: 0;
    padding: 0;
  }

  .summary-text {
    font-family: 'Times New Roman', Times, serif;
    font-size: 9.5pt;
    line-height: 1.25;
    color: #000000;
    margin: 0;
    padding: 0;
    text-align: justify;
  }

  .recommendation-card {
    margin-bottom: 4pt;
    padding: 0;
    border: none;
    background: #ffffff;
    page-break-inside: avoid;
  }

  .recommendation-text {
    font-family: 'Times New Roman', Times, serif;
    font-size: 9pt;
    line-height: 1.2;
    color: #333333;
    margin: 0 0 2pt 0;
    padding: 0 0 0 12pt;
    border-left: 1pt solid #999999;
    font-style: italic;
  }

  .recommendation-text::before {
    content: none;
  }

  .recommendation-text::after {
    content: none;
  }

  .recommendation-author {
    font-family: 'Times New Roman', Times, serif;
    font-size: 8.5pt;
    color: #000000;
    margin: 0;
    padding: 0;
  }

  .author-name {
    font-style: italic;
    font-weight: normal;
  }

  .author-role {
    font-style: normal;
  }

  .author-relation {
    font-style: normal;
  }

  .references-fallback {
    font-family: 'Times New Roman', Times, serif;
    font-size: 9pt;
    color: #000000;
    margin: 0;
    padding: 0;
    font-style: italic;
  }

  /* Page break control */
  section {
    page-break-after: auto;
  }

  section:last-child {
    page-break-after: avoid;
  }

  .section-header {
    page-break-after: avoid;
    page-break-inside: avoid;
  }

  .pdf-note {
    font-family: 'Times New Roman', Times, serif;
    font-size: 8pt;
    color: #666666;
    margin: 8pt 0 0 0;
    padding: 0;
    text-align: center;
    font-style: italic;
  }

  /* Ensure no dark theme remnants */
  *,
  *::before,
  *::after {
    box-sizing: border-box;
  }
}
```

---

## Appendix B: Complete Script Changes After Formalization

```javascript
// In injectPrivateData() function:

// Helper to create a contact item (text-only, no icons)
function createContactItem(label, href) {
  const el = document.createElement(href ? "a" : "span");
  el.className = "contact-item";
  if (href) {
    el.href = href;
    el.target = "_blank";
    el.rel = "noopener noreferrer";
  }
  el.textContent = label;
  contactInfo.appendChild(el);
}

// Helper to create a pipe separator
function createSeparator() {
  const separator = document.createElement("span");
  separator.className = "contact-separator";
  separator.textContent = " | ";
  contactInfo.appendChild(separator);
}

// Build the contact info block
const contactInfo = document.querySelector(".contact-info");
if (!contactInfo) return;

// Clear any existing content (e.g., the Google Form button)
while (contactInfo.firstChild) {
  contactInfo.removeChild(contactInfo.firstChild);
}

// Collect all present contact items into an array
const items = [];
if (data.email) items.push({ label: data.email, href: `mailto:${data.email}` });
if (data.phone) items.push({ label: data.phone, href: null });
if (data.address) items.push({ label: data.address, href: null });
if (data.website) items.push({ label: "Portafolio", href: data.website });
if (data.github) items.push({ label: "GitHub", href: data.github });
if (data.linkedin) items.push({ label: "LinkedIn", href: data.linkedin });

// Append with separators
items.forEach((item, index) => {
  createContactItem(item.label, item.href);
  if (index < items.length - 1) {
    createSeparator();
  }
});

// Note: SVG icon constants (svgEmail, svgPhone, etc.) are removed
```

---

## End of Design Document
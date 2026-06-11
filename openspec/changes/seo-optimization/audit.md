# SEO Audit — CV Site at https://sascodiego.github.io/cv/

**Site Type**: Single-page personal CV (Astro 6.4.5)
**Subject**: Diego Sasco — Software Architect / Senior Software Developer
**Audit Date**: 2025-06-10
**Base URL**: `https://sascodiego.github.io/cv/`

---

## Current State Score: 4/10

The site has basic SEO elements in place but lacks critical technical SEO infrastructure and several on-page optimizations. It's functional for direct sharing but not optimized for search discoverability or social sharing.

---

## Files Retrieved

1. `web/src/pages/index.astro` (lines 1-226) — Main page template with existing meta tags
2. `web/src/styles/global.css` (lines 1-456) — Styling with print optimization for ATS
3. `web/astro.config.mjs` (lines 1-8) — Astro configuration
4. `web/package.json` (lines 1-25) — Dependencies and scripts
5. `web/public/` — Contains favicon.svg and favicon.ico only
6. `web/src/data/cv-processed.json` — CV data source

---

## Key Code — Current SEO Implementation

### Existing Meta Tags (index.astro, lines 13-29)

```astro
<meta charset="utf-8" />
<link rel="icon" type="image/svg+xml" href={`${base}favicon.svg`} />
<meta name="viewport" content="width=device-width, initial-scale=1" />
<meta name="theme-color" content="#09090b" />
<meta name="generator" content={Astro.generator} />
<title>{personalInfo.name} | {personalInfo.title}</title>
<meta name="description" content={`${personalInfo.name} — ${personalInfo.title}. Más de 7 años diseñando sistemas distribuidos, soluciones IoT e integración legacy para el sector financiero, retail e industrial.`} />
<link rel="canonical" href={new URL(base, Astro.site).href} />
```

### Open Graph (index.astro, lines 31-35)

```astro
<meta property="og:title" content={`${personalInfo.name} | ${personalInfo.title}`} />
<meta property="og:description" content={`${personalInfo.name} — ${personalInfo.title}. Más de 7 años diseñando sistemas distribuidos, soluciones IoT e integración legacy para el sector financiero, retail e industrial.`} />
<meta property="og:type" content="website" />
<meta property="og:url" content={new URL(base, Astro.site).href} />
<meta property="og:locale" content="es_ES" />
```

### Schema.org Structured Data (index.astro, lines 37-44)

```astro
<script type="application/ld+json" set:html={JSON.stringify({
  "@context": "https://schema.org",
  "@type": "Person",
  "name": personalInfo.name,
  "jobTitle": personalInfo.title,
  "url": new URL(base, Astro.site).href,
  "knowsAbout": [
    "C#/.NET", "Go", "gRPC", "IoT", "RabbitMQ",
    "integración de sistemas legacy", "automatización RPA",
    "seguridad informática", "desarrollo de software"
  ]
})} />
```

### Astro Configuration (astro.config.mjs, lines 5-7)

```javascript
site: "https://sascodiego.github.io",
base: "/cv/",
```

---

## Specific Issues Found

### Critical Issues

| Issue | Impact | Location |
|-------|--------|----------|
| **No robots.txt** | Search engines may not crawl properly; no control over indexing | `web/public/robots.txt` (missing) |
| **No sitemap.xml** | Search engines cannot discover page structure efficiently | `web/public/sitemap.xml` (missing) |
| **No Open Graph image** | Social links show no preview; poor sharing experience | `index.astro` (missing `og:image`) |
| **No Twitter Card meta tags** | Poor Twitter/X sharing experience | `index.astro` (missing `twitter:card`, `twitter:title`, `twitter:description`, `twitter:image`) |
| **No favicon.ico fallback** | Some browsers/clients may not render SVG favicon | `web/public/favicon.ico` exists but may need optimization |

### High Priority Issues

| Issue | Impact | Location |
|-------|--------|----------|
| **"Built with Astro" footer** | Unnecessary attribution; distracts from personal branding | `index.astro` line 208 |
| **Meta description in Spanish only** | Limits reach to Spanish-language search results | `index.astro` line 18 |
| **Missing alternative language tags** | No `hreflang` for international audiences | `index.astro` (missing) |
| **Incomplete Schema.org data** | Missing `email`, `telephone`, `sameAs`, `worksFor`, `alumniOf` | `index.astro` lines 37-44 |
| **No structured data for experience/education** | ATS parsers may miss structured work history | `index.astro` (missing `JobPosting` or `OrganizationRole` schemas) |
| **No social profile links in Schema** | Search engines cannot verify social profiles | `index.astro` Schema missing `sameAs` array |

### Medium Priority Issues

| Issue | Impact | Location |
|-------|--------|----------|
| **Title tag formulaic** | Could be more descriptive/SEO-friendly | `index.astro` line 17 |
| **No semantic landmarks beyond main** | Could use `header`, `nav`, `aside` landmarks | `index.astro` structure |
| **No article structured data** | Experience items could be marked as articles | `index.astro` lines 84-147 |
| **No BreadcrumbList schema** | No navigation structure for rich snippets | `index.astro` (missing) |
| **No viewport meta tag for accessibility** | Already present but could add `maximum-scale=5` | `index.astro` line 14 |
| **No image alt text strategy** | No images currently, but should plan for it | N/A (no images) |
| **No performance monitoring tags** | No Core Web Vitals or analytics setup | N/A |

---

## Missing Technical SEO Items

1. **robots.txt** — Should allow all crawlers, possibly disallow staging paths
2. **sitemap.xml** — Single entry for main page; could be generated automatically with `@astrojs/sitemap`
3. **Favicon suite** — Currently has SVG and ICO; could add PNG, Apple touch icon
4. **Open Graph image** — Need 1200x630px social sharing image
5. **Twitter Card tags** — Large summary card with image
6. **webapp manifest** — For PWA installation on mobile devices
7. **Structured data completeness** — Expand Schema.org Person with more properties

---

## On-Page SEO Gaps

### Title Tag Optimization

**Current**: `{personalInfo.name} | {personalInfo.title}`
- "Diego Sasco | Desarrollador de Software Senior"

**Issues**:
- Too generic; doesn't emphasize specialties (IoT, distributed systems, C#/.NET, Go)
- Missing location (could target Uruguay/LatAm market)
- No differentiator phrases

**Recommendation**: "Diego Sasco — Software Architect (C#/.NET, Go, IoT) | Montevideo, Uruguay"

### Meta Description Quality

**Current**: Spanish-only, ~180 characters
- "Diego Sasco — Desarrollador de Software Senior. Más de 7 años diseñando sistemas distribuidos..."

**Issues**:
- Single language limits reach
- Could be more keyword-dense for technical SEO
- No call-to-action or contact info

**Recommendation**: Add English version and include key technologies, location, and contact method.

### Heading Hierarchy

**Current structure**:
- H1: Personal name (hero section)
- H2: Section titles ("Sobre Mí", "Conocimientos y Experiencia", "Experiencia Profesional", etc.)
- H3: Skill category titles, project names

**Issues**:
- No H1 with job title/summary for immediate context
- H2 sections are semantically correct but could use landmarks

**Recommendation**: Consider combining name + title in H1, or add subtitle H1.

### Lang Attribute

**Current**: `<html lang="es">` (line 10)

**Issues**:
- Content is mixed (Spanish UI, English tech terms)
- No alternate language links for English version

**Recommendation**: Keep `lang="es"` but add `hreflang="en"` if English version is planned.

### Content Structure for ATS/Crawlers

**Good**:
- Semantic HTML with `<article>`, `<section>`, `<header>`, `<footer>`
- Print stylesheet optimizes for ATS parsers
- Clear role/company structure in work experience

**Gaps**:
- No microdata markup for individual job postings
- Could add `itemscope` and `itemtype` attributes to experience items
- No structured contact info in HTML ( Schema only)

---

## Astro 6.4 Built-in SEO Capabilities

### Available Features

1. **Automatic Sitemap Generation**
   - Package: `@astrojs/sitemap`
   - Configuration: Add to `astro.config.mjs` integrations
   - Benefit: Auto-generates sitemap.xml at build time

2. **Image Optimization**
   - Package: `@astrojs/image`
   - Benefits: Lazy loading, WebP conversion, responsive images
   - Not currently used but available

3. **Built-in Meta Tags**
   - `Astro.generator` meta tag (currently used)
   - Canonical URL generation (currently used)
   - Site-relative URL handling with `import.meta.env.BASE_URL`

4. **Performance Optimization**
   - Zero JS by default (already leveraged)
   - CSS bundling and minification
   - Static site generation (SSG)

5. **i18n Support**
   - Package: `@astrojs/i18n`
   - Benefits: Alternate language pages, automatic hreflang generation
   - Not currently used but available for multi-language CV

### Not Leveraged

- `@astrojs/sitemap` for automatic sitemap generation
- `@astrojs/image` for optimized OG images and favicons
- `@astrojs/i18n` for multi-language support

---

## Conflict Assessment with Parallel PDF Work

### Current State
- **No `cv-pdf.astro` exists yet** in `web/src/pages/`
- Parallel work is planning to add a PDF generation page
- Current site has only one page: `index.astro`

### Potential Conflicts

| Area | Conflict Risk | Mitigation |
|------|---------------|------------|
| **Meta tags** | Low | `cv-pdf.astro` can share same meta structure via component |
| **Sitemap** | Low | Adding new page requires sitemap update (automated if using `@astrojs/sitemap`) |
| **Robots.txt** | Low | May need to allow indexing of PDF page or disallow if it's a utility |
| **Canonical URLs** | Low | Each page will have its own canonical URL |
| **Schema.org** | Low | PDF page may need different schema (e.g., `DigitalDocument`) |
| **Open Graph** | Low | PDF page should have unique OG tags |
| **Shared styles** | Low | `global.css` already handles both, print stylesheet optimized for PDF |

### Coordination Needed

1. **Sitemap generation**: When `cv-pdf.astro` is added, ensure `@astrojs/sitemap` picks it up
2. **Robots.txt**: Decide whether PDF page should be indexed (likely no if it's a utility)
3. **Meta tag structure**: Consider extracting common meta tags into a reusable component
4. **Canonical URLs**: Ensure both pages point to correct URLs
5. **Structured data**: PDF page may need different schema type if it represents a downloadable document

**Recommendation**: No blocking conflicts. SEO changes are safe to proceed. When `cv-pdf.astro` is added, update sitemap and robots.txt accordingly.

---

## Recommended Fixes with File Paths

### Critical Fixes

1. **Add robots.txt**
   - **File**: `web/public/robots.txt`
   - **Content**:
     ```
     User-agent: *
     Allow: /
     Sitemap: https://sascodiego.github.io/cv/sitemap.xml
     ```
   - **Lines**: 4

2. **Add sitemap.xml**
   - **Option A (Manual)**: Create `web/public/sitemap.xml` with single entry
   - **Option B (Recommended)**: Install `@astrojs/sitemap` and configure in `astro.config.mjs`
   - **File**: `web/astro.config.mjs`
   - **Add**: `import sitemap from '@astrojs/sitemap';` and `integrations: [sitemap()]`
   - **Lines**: 2-3

3. **Add Open Graph image**
   - **File**: `web/public/og-image.png` (1200x630px)
   - **Update**: `web/src/pages/index.astro`
   - **Add**: `<meta property="og:image" content={`${base}og-image.png`} />`
   - **Lines**: 1

4. **Add Twitter Card tags**
   - **File**: `web/src/pages/index.astro`
   - **Add**:
     ```astro
     <meta name="twitter:card" content="summary_large_image" />
     <meta name="twitter:title" content={`${personalInfo.name} | ${personalInfo.title}`} />
     <meta name="twitter:description" content={`${personalInfo.name} — ${personalInfo.title}. Más de 7 años diseñando sistemas distribuidos, soluciones IoT e integración legacy para el sector financiero, retail e industrial.`} />
     <meta name="twitter:image" content={`${base}og-image.png`} />
     ```
   - **Lines**: 4

### High Priority Fixes

5. **Remove "Built with Astro" footer**
   - **File**: `web/src/pages/index.astro`
   - **Remove**: Lines 208-210
   - **Lines**: -3

6. **Expand Schema.org Person data**
   - **File**: `web/src/pages/index.astro`
   - **Add to schema**:
     ```javascript
     "email": "contact@example.com",
     "telephone": "+598-XX-XXX-XXXX",
     "sameAs": [
       "https://linkedin.com/in/diegosasco",
       "https://github.com/sascodiego"
     ],
     "address": {
       "@type": "PostalAddress",
       "addressLocality": "Montevideo",
       "addressCountry": "UY"
     }
     ```
   - **Lines**: +8-10

7. **Add JobPosting or OrganizationRole schemas**
   - **File**: `web/src/pages/index.astro`
   - **Add**: Array of `JobPosting` or `OrganizationRole` schemas for each work experience
   - **Lines**: +20-30 (complex)

8. **Optimize title tag**
   - **File**: `web/src/pages/index.astro`
   - **Change line 17**: `<title>Diego Sasco — Software Architect (C#/.NET, Go, IoT) | Montevideo, Uruguay</title>`
   - **Lines**: 1 (changed)

9. **Add English meta description**
   - **File**: `web/src/pages/index.astro`
   - **Add**: `<meta name="description" content="Diego Sasco — Senior Software Developer. 7+ years designing distributed systems, IoT solutions, and legacy integration for finance, retail, and industrial sectors." lang="en" />`
   - **Lines**: 1

### Medium Priority Fixes

10. **Add webapp manifest**
    - **File**: `web/public/manifest.json`
    - **Content**: PWA manifest with name, icons, theme color
    - **Lines**: 15-20

11. **Add hreflang tags**
    - **File**: `web/src/pages/index.astro`
    - **Add**: `<link rel="alternate" hreflang="es" href="..." />` and `<link rel="alternate" hreflang="en" href="..." />`
    - **Lines**: 2

12. **Add BreadcrumbList schema**
    - **File**: `web/src/pages/index.astro`
    - **Add**: Schema for navigation hierarchy
    - **Lines**: 10-15

13. **Add semantic landmarks**
    - **File**: `web/src/pages/index.astro`
    - **Change**: Wrap nav in `<nav>`, add `<header>` and `<footer>` with proper roles
    - **Lines**: 5-10 (refactor)

---

## Estimated Scope

| Priority | Files Changed | Lines Changed | Complexity |
|----------|---------------|---------------|------------|
| **Critical** | 3 files | +15, -0 | Low |
| **High** | 1 file | +20, -3 | Medium |
| **Medium** | 2 files | +25, -0 | Medium |
| **Total** | 5 files | +60, -3 | Medium |

### Time Estimate

- **Critical fixes**: 1-2 hours
- **High priority fixes**: 2-3 hours
- **Medium priority fixes**: 2-3 hours
- **Total**: 5-8 hours

### Dependencies

- **OG image creation**: Requires design work (external task)
- **Contact info**: Need real email/phone/social URLs
- **Testing**: Should verify with rich results test, lighthouse, social share previews

---

## Next Recommended Actions

1. **Immediate (Critical)**:
   - Create `robots.txt` and `sitemap.xml`
   - Install `@astrojs/sitemap` for automatic sitemap generation
   - Add Open Graph and Twitter Card meta tags
   - Create/optimize OG image (1200x630px)

2. **Short-term (High Priority)**:
   - Remove "Built with Astro" footer
   - Expand Schema.org Person data with contact info and social profiles
   - Optimize title tag for better SEO
   - Add English meta description

3. **Medium-term (Medium Priority)**:
   - Add webapp manifest for PWA support
   - Implement structured data for work experience
   - Add hreflang tags for multi-language support
   - Improve semantic landmarks

4. **Post-Implementation**:
   - Test with Google Rich Results Test
   - Validate with Schema.org validator
   - Check social share previews (LinkedIn, Twitter, Facebook)
   - Run Lighthouse audit for performance/SEO/accessibility
   - Monitor Search Console for indexing issues

---

## Astro Configuration Updates

### Recommended `astro.config.mjs`

```javascript
import { defineConfig } from "astro/config";
import sitemap from "@astrojs/sitemap";

export default defineConfig({
  site: "https://sascodiego.github.io",
  base: "/cv/",
  integrations: [sitemap()],
});
```

### Package.json Dependencies to Add

```json
{
  "devDependencies": {
    "@astrojs/sitemap": "^3.0.0"
  }
}
```

---

## Testing Checklist

- [ ] robots.txt accessible at `/robots.txt`
- [ ] sitemap.xml accessible at `/sitemap.xml`
- [ ] Open Graph tags present and correct
- [ ] Twitter Card tags present and correct
- [ ] Schema.org structured data validates
- [ ] Rich results test passes
- [ ] Social share previews render correctly
- [ ] Lighthouse SEO score > 90
- [ ] Canonical URL correct
- [ ] No console errors
- [ ] Mobile-friendly test passes
- [ ] Page speed score acceptable

---

## Conflict Assessment Summary

**No blocking conflicts** with parallel `cv-pdf.astro` work. SEO optimizations can proceed safely. Coordination points:

1. Sitemap will need regeneration when `cv-pdf.astro` is added
2. Robots.txt may need update if PDF page should be disallowed
3. Consider extracting shared meta tags into component for consistency

**Recommendation**: Proceed with SEO optimization. When `cv-pdf.astro` is ready, coordinate on sitemap and robots.txt updates.

---

**Audit Completed**: 2025-06-10
**Auditor**: el Gentleman — Pi Scouting Subagent
**Next Review**: After implementation of critical and high-priority fixes
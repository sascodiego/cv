# Propuesta Técnica: CV Digital con GitHub Pages (Repo Privado) y Exportación Segura

Esta propuesta detalla la arquitectura, el flujo de trabajo y la implementación técnica para crear un currículum vitae (CV) interactivo alojado en **GitHub Pages** desde un **repositorio privado**. El sistema actúa como una "Fuente Única de Verdad" (_Single Source of Truth_), permitiendo mantener una versión web pública libre de datos sensibles, mientras que la generación del PDF con información privada se realiza localmente en tu notebook de desarrollo.

---

## 🏗️ 1. Arquitectura General del Sistema

El proyecto se divide en dos entornos principales con objetivos claramente diferenciados:

```

[ Tu Notebook (Local) ] ──(Genera con Datos Privados)──> Diego_Sasco_CV_Privado.pdf
│
(Push de Código)
▼
[ GitHub (Repo Privado) ] ──(GitHub Actions Workflow)──> [ GitHub Pages (Web Pública) ]
│
(Formulario + CAPTCHA)
▼
[ Solicitud de Contacto ]

```

1. **El Repositorio en GitHub (Privado):** Aloja el código fuente del sitio (desarrollado en Astro). Al ser privado, tu código, tus configuraciones y la lógica de inyección de datos están completamente protegidos de miradas externas.
2. **GitHub Pages (Público):** Servidor estático gratuito que expone la versión web limpia de datos sensibles para reclutadores.
3. **El Entorno Local (Tu Notebook):** El único lugar donde se procesan y fusionan tus datos privados para generar el PDF oficial.
4. **Formulario de Contacto en la Web:** Un botón que despliega un formulario protegido por CAPTCHA para evitar spam, derivando las solicitudes de reclutadores legítimos.

---

## 🔒 2. Estrategia de Datos Libres de Filtraciones

Para garantizar que tus datos privados (teléfono directo, dirección exacta, etc.) nunca se filtren ni queden expuestos en el historial de Git o en la web pública, aplicaremos la siguiente estrategia:

- **Inyección Local mediante Variables de Entorno:** Los datos privados se configuran exclusivamente en tu archivo `.env` local dentro de tu notebook. Este archivo está explícitamente excluido en el `.gitignore`.
- **Componentes Condicionales (Astro):** El código HTML de la web se renderiza condicionalmente. Si las variables privadas no están presentes durante la compilación, esos nodos del DOM ni siquiera se crean en los archivos estáticos finales subidos a GitHub Pages.

---

## 🛠️ 3. Implementación del Proyecto Paso a Paso

### Paso 1: Configuración del Repositorio y Git

1. Crea un repositorio en GitHub con el nombre `cv-portfolio` y asegúrate de marcarlo como **Private**.
2. Clona el repositorio en tu notebook:
   ```bash
   git clone git@github.com:tu-usuario/cv-portfolio.git
   cd cv-portfolio
   ```

````

3. Inicializa el proyecto Astro dentro de la carpeta:
```bash
pnpm create astro@latest . -- --template minimal

````

4. Define el archivo `.gitignore` para asegurar que las variables locales y compilaciones privadas jamás se suban:

```gitignore
# Node modules y build output
node_modules/
dist/
.astro/

# Variables de entorno locales (Contienen tus datos privados)
.env
.env.production
.env.local

# Archivos generados localmente
*.pdf
Diego_Sasco_CV_Privado.pdf

```

### Paso 2: Creación de la Estructura de Datos (`src/data/cv.json`)

Guardaremos toda tu experiencia profesional, educación y tecnologías en un archivo estructurado JSON público. El archivo `.env` local contendrá únicamente lo privado.

**Ejemplo de `.env` (Solo en tu notebook):**

```env
PRIVATE_PHONE="+598 9X XXX XXX"
PRIVATE_ADDRESS="Manga, Montevideo, Uruguay"
PRIVATE_CV_NOTE="Este PDF contiene información confidencial para uso exclusivo de procesos de selección."

```

**Ejemplo del Componente de Contacto (`src/components/ContactInfo.astro`):**

```astro
---
// Astro lee las variables en tiempo de compilación
const privatePhone = process.env.PRIVATE_PHONE || "";
const privateAddress = process.env.PRIVATE_ADDRESS || "";
const publicEmail = "<YOUR_EMAIL>";
---

<div class="contact-card">
  <p><i class="fa-solid fa-envelope"></i> {publicEmail}</p>

  {privatePhone && (
    <p class="private-data"><i class="fa-solid fa-phone"></i> {privatePhone}</p>
  )}
  {privateAddress && (
    <p class="private-data"><i class="fa-solid fa-location-dot"></i> {privateAddress}</p>
  )}
</div>

<style>
  /* Ocultamos visualmente notas exclusivas del PDF en la web si fuese necesario */
  @media screen {
    .pdf-only { display: none; }
  }
</style>

```

### Paso 3: Generación Automática del PDF desde tu Notebook

Para automatizar la creación del PDF exacto sin tener que usar el menú "Imprimir" del navegador, usaremos **Puppeteer** en un script local. Esto correrá 100% en tu máquina.

1. Instala Puppeteer como dependencia de desarrollo:

```bash
pnpm add -D puppeteer dotenv

```

2. Crea el script de automatización `scripts/generate-pdf.js`:

```javascript
import puppeteer from "puppeteer";
import dotenv from "dotenv";
import { spawn } from "child_process";

dotenv.config(); // Carga las variables del .env local

async function main() {
  console.log("🚀 Iniciando servidor local de Astro...");
  // Levanta Astro en modo producción local con tus variables cargadas
  const astroProcess = spawn("pnpm", ["astro", "preview", "--port", "4321"], {
    env: { ...process.env, NODE_ENV: "production" },
    shell: true,
  });

  // Espera 3 segundos a que el servidor esté listo
  await new Promise((resolve) => setTimeout(resolve, 3000));

  console.log("📸 Renderizando PDF con Puppeteer...");
  const browser = await puppeteer.launch({ headless: "new" });
  const page = await browser.newPage();

  // Navega al puerto local
  await page.goto("http://localhost:4321", { waitUntil: "networkidle0" });

  // Genera el PDF aplicando estilos de impresión (@media print)
  await page.pdf({
    path: "./Diego_Sasco_CV_Privado.pdf",
    format: "A4",
    margin: { top: "15mm", bottom: "15mm", left: "12mm", right: "12mm" },
    printBackground: true,
  });

  console.log("✅ PDF generado con éxito: ./Diego_Sasco_CV_Privado.pdf");

  await browser.close();
  astroProcess.kill();
  process.exit(0);
}

main().catch((err) => {
  console.error("❌ Error generando el PDF:", err);
  process.exit(1);
});
```

3. Agrega el comando en tu `package.json`:

```json
"scripts": {
  "dev": "astro dev",
  "build": "astro build",
  "preview": "astro preview",
  "astro": "astro",
  "export-pdf": "pnpm build && node scripts/generate-pdf.js"
}

```

Cada vez que quieras tu PDF actualizado con tus datos privados, simplemente ejecutas en la terminal de tu notebook:

```bash
pnpm export-pdf

```

---

## 🤖 4. Botón de Solicitud de Contacto con CAPTCHA en GitHub Pages

Dado que la web en GitHub Pages es estática y no tiene base de datos ni backend nativo, utilizaremos un servicio de formularios _serverless_ externo (como **Formspree**, **Getform** o **Formbold**) combinado con **Google reCAPTCHA v3** o **Cloudflare Turnstile** (este último es el más moderno, no obliga al usuario a clickear imágenes de semáforos y respeta la privacidad).

### Implementación con Cloudflare Turnstile (Alternativa moderna a reCAPTCHA)

1. Te registras en Cloudflare (es gratis) y creas un sitio en Turnstile para obtener tu `Site Key`.
2. Insertamos el formulario y el widget de validación en tu componente de Astro `src/components/ContactForm.astro`:

```astro
---
// Tu Site Key pública de Turnstile (No es secreta)
const TURNSTILE_SITE_KEY = "3x00000000000000000000FF";
---

<div class="form-container">
  <h3>¿Quieres coordinar una entrevista o solicitar datos adicionales?</h3>
  <p>Completa el formulario para validar tu identidad.</p>

  <form action="[https://formspree.io/f/tu_form_id](https://formspree.io/f/tu_form_id)" method="POST" id="cv-contact-form">
    <div class="input-group">
      <label for="email">Tu Email Corporativo:</label>
      <input type="email" id="email" name="_replyto" required placeholder="ejemplo@empresa.com" />
    </div>

    <div class="input-group">
      <label for="message">Mensaje / Propuesta:</label>
      <textarea id="message" name="message" required placeholder="Hola Diego, nos interesa tu perfil para..."></textarea>
    </div>

    <div class="cf-turnstile" data-sitekey={TURNSTILE_SITE_KEY} data-theme="dark"></div>

    <button type="submit" id="submit-btn" disabled>Enviar Solicitud</button>
  </form>
</div>

<script src="[https://challenges.cloudflare.com/turnstile/v0/api.js](https://challenges.cloudflare.com/turnstile/v0/api.js)" async defer></script>

<script>
  // Script para habilitar el botón solo cuando el CAPTCHA se resuelva con éxito
  const form = document.getElementById('cv-contact-form');
  const submitBtn = document.getElementById('submit-btn');

  window.addEventListener('message', (event) => {
    // Escucha si Turnstile validó correctamente al usuario
    if (event.data && event.data.source === 'cloudflare-turnstile') {
      if (event.data.event === 'success') {
        submitBtn.removeAttribute('disabled');
      }
    }
  });
</script>

<style>
  .form-container {
    background: #1e293b;
    padding: 2rem;
    border-radius: 12px;
    max-width: 500px;
    margin: 2rem auto;
  }
  .input-group {
    margin-bottom: 1.2rem;
    display: flex;
    flex-direction: column;
  }
  input, textarea {
    background: #0f172a;
    border: 1px solid #334155;
    color: #fff;
    padding: 0.7rem;
    border-radius: 6px;
    margin-top: 0.4rem;
  }
  button {
    background: #38bdf8;
    color: #0f172a;
    font-weight: bold;
    padding: 0.8rem;
    border: none;
    border-radius: 6px;
    cursor: pointer;
    width: 100%;
    margin-top: 1rem;
  }
  button:disabled {
    background: #475569;
    color: #94a3b8;
    cursor: not-allowed;
  }
</style>

```

---

## 🤖 5. Flujo de Publicación (GitHub Actions)

Dado que el repositorio es privado, configuraremos un flujo automatizado en `.github/workflows/deploy.yml` que se activará cada vez que hagas `git push` a la rama `main`. Este flujo construirá la versión pública del sitio web (sin datos privados porque las variables de entorno no existen en el servidor de GitHub) y lo desplegará automáticamente.

Crear archivo `.github/workflows/deploy.yml`:

```yaml
name: Deploy Public CV to GitHub Pages

on:
  push:
    branches: ["main"]

# Permisos requeridos para escribir en GitHub Pages desde un repo privado
permissions:
  contents: read
  pages: write
  id-token: write

concurrency:
  group: "pages"
  cancel-in-progress: true

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Setup Node.js
        uses: actions/setup-node@v4
        with:
          node-version: 20

      - name: Install pnpm
        uses: pnpm/action-setup@v3
        with:
          version: 8

      - name: Install dependencies
        run: pnpm install

      - name: Build Static Site (Public Version)
        # Al no pasar variables de entorno aquí, el código de Astro omitirá los bloques privados
        run: pnpm build

      - name: Upload Pages Artifact
        uses: actions/upload-pages-artifact@v3
        with:
          path: ./dist

  deploy:
    environment:
      name: github-pages
      url: ${{ steps: deployment.outputs.page_url }}
    needs: build
    runs-on: ubuntu-latest
    steps:
      - name: Deploy to GitHub Pages
        id: deployment
        uses: actions/deploy-pages@v4

```

---

## 🎯 6. Ventajas del Enfoque Propuesto

1. **Seguridad Absoluta:** Tus datos sensibles jamás salen de tu notebook. El archivo PDF resultante se genera en local y tú decides a qué reclutador enviárselo manualmente.
2. **Web Estática Anti-Spam:** Al utilizar Cloudflare Turnstile junto con un procesador estático como Formspree, tu casilla de correo público está a salvo de scrapers de correos automatizados.
3. **Mantenimiento Simple:** Modificas tu experiencia laboral una sola vez en el archivo `cv.json`. Haces `pnpm export-pdf` para tu archivo personal y `git push` para actualizar tu presencia en la web.

```

```

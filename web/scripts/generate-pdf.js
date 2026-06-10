import puppeteer from "puppeteer";
import dotenv from "dotenv";
import { spawn } from "child_process";
import path from "path";
import { fileURLToPath } from "url";

// Cargar variables del .env local (datos privados, nunca commiteados)
dotenv.config();

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const pdfPath = path.resolve(__dirname, "../../Diego_Sasco_CV_Privado.pdf");

// Datos privados inyectados desde .env
const privateData = {
	name: process.env.PRIVATE_NAME || "",
	email: process.env.PRIVATE_EMAIL || "",
	phone: process.env.PRIVATE_PHONE || "",
	address: process.env.PRIVATE_ADDRESS || "",
	website: process.env.PRIVATE_WEBSITE || "",
	github: process.env.PRIVATE_GITHUB || "",
	linkedin: process.env.PRIVATE_LINKEDIN || "",
	note: process.env.PRIVATE_CV_NOTE || "",
	recommenders: [
		{
			name: process.env.PRIVATE_RECOMMENDER_1_NAME || "",
			role: process.env.PRIVATE_RECOMMENDER_1_ROLE || "",
			relation: process.env.PRIVATE_RECOMMENDER_1_RELATION || "",
		},
		{
			name: process.env.PRIVATE_RECOMMENDER_2_NAME || "",
			role: process.env.PRIVATE_RECOMMENDER_2_ROLE || "",
			relation: process.env.PRIVATE_RECOMMENDER_2_RELATION || "",
		},
		{
			name: process.env.PRIVATE_RECOMMENDER_3_NAME || "",
			role: process.env.PRIVATE_RECOMMENDER_3_ROLE || "",
			relation: process.env.PRIVATE_RECOMMENDER_3_RELATION || "",
		},
	],
};

/**
 * Injects private contact data into the rendered page DOM via Puppeteer.
 * The Astro page has no knowledge of these fields — this is the only place
 * where private data meets the rendered HTML.
 */
async function injectPrivateData(page) {
	await page.evaluate((data) => {
		// Update the hero name if provided
		if (data.name) {
			const h1 = document.querySelector(".hero h1");
			if (h1) h1.textContent = data.name;
		}

		// Build the contact info block
		const contactInfo = document.querySelector(".contact-info");
		if (!contactInfo) return;

		// Clear any existing content (e.g., the Google Form button)
		while (contactInfo.firstChild) {
			contactInfo.removeChild(contactInfo.firstChild);
		}

		// Helper to create a contact item (no innerHTML — safe DOM construction)
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

		const svgEmail =
			'<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path d="M48 64C21.5 64 0 85.5 0 112c0 15.1 7.1 29.3 19.2 38.4L236.8 313.6c11.4 8.5 27 8.5 38.4 0L492.8 150.4c12.1-9.1 19.2-23.3 19.2-38.4c0-26.5-21.5-48-48-48H48zM0 176V384c0 35.3 28.7 64 64 64H448c35.3 0 64-28.7 64-64V176L294.4 339.2c-22.8 17.1-54 17.1-76.8 0L0 176z"/></svg>';
		const svgPhone =
			'<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path d="M164.9 24.6c-7.7-18.6-28-28.5-47.4-23.2l-88 24C12.1 30.2 0 46 0 64C0 311.4 200.6 512 448 512c18 0 33.8-12.1 38.6-29.5l24-88c5.3-19.4-4.6-39.7-23.2-47.4l-96-40c-16.3-6.8-35.2-2.1-46.3 11.6L304.7 368C234.3 334.7 177.3 277.7 144 207.3L193.3 167c13.7-11.2 18.4-30 11.6-46.3l-40-96z"/></svg>';
		const svgLocation =
			'<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 384 512"><path d="M215.7 499.2C267 435 384 279.4 384 192C384 86 298 0 192 0S0 86 0 192c0 87.4 117 243 168.3 307.2c12.3 15.3 35.1 15.3 47.4 0zM192 128a64 64 0 1 1 0 128 64 64 0 1 1 0-128z"/></svg>';
		const svgWeb =
			'<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path d="M352 256c0 22.2-1.2 43.6-3.7 64H163.7c-2.5-20.4-3.7-41.8-3.7-64s1.2-43.6 3.7-64h184.6c2.5 20.4 3.7 41.8 3.7 64zm-3.7 128c-12.3 31.8-30 59.8-51.5 82.5c41.3-11.8 77.4-35.9 103.5-67.9H348.3zm-172.6 0H163.7c26.1 32 62.2 56.1 103.5 67.9c-21.5-22.7-39.2-50.7-51.5-82.5zm0-256H348.3c-12.3-31.8-30-59.8-51.5-82.5c-41.3 11.8-77.4 35.9-103.5 67.9zm-29.3 128H39.2c-5.8 20.3-9.2 41.7-9.2 64s3.4 43.7 9.2 64h107.2c-1.8-20.8-2.8-42.1-2.8-64s1-43.2 2.8-64zm233.2 0c1.8 20.8 2.8 42.1 2.8 64s-1 43.2-2.8 64H472.8c5.8-20.3 9.2-41.7 9.2-64s-3.4-43.7-9.2-64H379.6zM272 48c0 31.4 15.5 67.6 37.6 104H202.4C224.5 115.6 240 79.4 240 48c0-3.3-.3-6.6-.8-9.9c-2.4-.1-4.8-.1-7.2-.1s-4.8 0-7.2 .1c-.5 3.3-.8 6.6-.8 9.9c0 31.4 15.5 67.6 37.6 104H144.3c-2.5-20.4-3.7-41.8-3.7-64s1.2-43.6 3.7-64h15.6C186.1 20 222.2 0 264 0s77.9 20 104.1 52.1h15.6c2.5 20.4 3.7 41.8 3.7 64s-1.2 43.6-3.7 64H367.7c22.1-36.4 37.6-72.6 37.6-104c0-3.3-.3-6.6-.8-9.9c-2.4-.1-4.8-.1-7.2-.1s-4.8 0-7.2 .1c-.5 3.3-.8 6.6-.8 9.9z"/></svg>';
		const svgGithub =
			'<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 496 512"><path d="M165.9 397.4c0 2-2.3 3.6-5.2 3.6-3.3.3-5.6-1.3-5.6-3.6 0-2 2.3-3.6 5.2-3.6 3-.3 5.6 1.3 5.6 3.6zm-31.1-4.5c-.7 2 1.3 4.3 4.3 4.9 2.6 1 5.6 0 6.2-2s-1.3-4.3-4.3-5.2c-2.6-.7-5.5.3-6.2 2.3zm44.2-1.7c-2.9.7-4.9 2.6-4.6 4.9.3 2 2.9 3.3 5.9 2.6 2.9-.7 4.9-2.6 4.6-4.6-.3-1.9-3-3.2-5.9-2.9zM244.8 8C106.1 8 0 113.3 0 252c0 110.9 69.8 205.8 169.5 239.2 12.8 2.3 17.3-5.6 17.3-12.1 0-6.2-.3-40.4-.3-61.4 0 0-70 15-84.7-29.8 0 0-11.4-29.1-27.8-36.6 0 0-22.9-15.7 1.6-15.4 0 0 24.9 2 38.6 25.8 21.9 38.6 58.6 27.5 72.9 20.9 2.3-16 8.8-27.1 16-33.7-55.9-6.2-112.3-14.3-112.3-110.5 0-27.5 7.6-41.3 23.6-58.9-2.6-6.5-11.1-33.3 2.6-67.9 20.9-6.5 69 27 69 27 20-5.6 41.5-8.5 62.8-8.5s42.8 2.9 62.8 8.5c0 0 48.1-33.6 69-27 13.7 34.7 5.2 61.4 2.6 67.9 16 17.7 25.8 31.5 25.8 58.9 0 96.5-58.9 104.2-114.8 110.5 9.2 7.9 17 22.9 17 46.4 0 33.7-.3 75.4-.3 83.6 0 6.5 4.6 14.4 17.3 12.1C428.2 457.8 496 362.9 496 252 496 113.3 383.5 8 244.8 8zM97.2 352.9c-1.3 1-1 3.3.7 5.2 1.6 1.6 3.9 2.3 5.2 1 1.3-1 1-3.3-.7-5.2-1.6-1.6-3.9-2.3-5.2-1zm-10.8-8.1c-.7 1.3.3 2.9 2.3 3.9 1.6 1 3.6.7 4.3-.7.7-1.3-.3-2.9-2.3-3.9-2-.6-3.6-.3-4.3.7zm32.4 35.6c-1.6 1.3-1 4.3 1.3 6.2 2.3 2.3 5.2 2.6 6.5.7 1.3-1.3.7-4.3-1.3-6.2-2.2-2.3-5.2-2.6-6.5-.7zm-11.4-14.7c-1.6 1-1.6 3.6 0 5.9 1.6 2.3 4.3 3.3 5.6 2.3 1.6-1.3 1.6-3.9 0-6.2-1.4-2.3-4-3.3-5.6-2z"/></svg>';
		const svgLinkedin =
			'<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><path d="M416 32H31.9C14.3 32 0 46.5 0 64.3v383.4C0 465.5 14.3 480 31.9 480H416c17.6 0 32-14.5 32-32.3V64.3c0-17.8-14.4-32.3-32-32.3zM135.4 416H69V202.2h66.5V416zm-33.2-243c-21.3 0-38.5-17.3-38.5-38.5S80.9 96 102.2 96c21.2 0 38.5 17.3 38.5 38.5 0 21.3-17.2 38.5-38.5 38.5zm282.1 243h-66.4V312c0-24.8-.5-56.7-34.5-56.7-34.6 0-39.9 27-39.9 54.9V416h-66.4V202.2h63.7v29.2h.9c8.9-16.8 30.6-34.5 62.9-34.5 67.2 0 79.7 44.3 79.7 101.9V416z"/></svg>';

		// Email
		if (data.email) {
			createContactItem(svgEmail, data.email, `mailto:${data.email}`);
		}
		// Phone
		if (data.phone) {
			createContactItem(svgPhone, data.phone, null);
		}
		// Address
		if (data.address) {
			createContactItem(svgLocation, data.address, null);
		}
		// Website
		if (data.website) {
			createContactItem(svgWeb, "Portafolio", data.website);
		}
		// GitHub
		if (data.github) {
			createContactItem(svgGithub, "GitHub", data.github);
		}
		// LinkedIn
		if (data.linkedin) {
			createContactItem(svgLinkedin, "LinkedIn", data.linkedin);
		}

		// Footer note
		if (data.note) {
			const footer = document.createElement("footer");
			footer.className = "pdf-note";
			const p = document.createElement("p");
			p.textContent = data.note;
			footer.appendChild(p);
			document.querySelector("main").appendChild(footer);
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
	}, privateData);
}

async function main() {
	const hasData = Object.values(privateData).some((v) => v);
	if (!hasData) {
		console.warn(
			"⚠️  No PRIVATE_* variables found in .env. The PDF will use placeholder data.",
		);
	}

	console.log("🚀 Iniciando servidor local de Astro (preview)...");

	const astroProcess = spawn("pnpm", ["astro", "preview", "--port", "4321"], {
		env: { ...process.env, NODE_ENV: "production" },
		shell: true,
		cwd: path.resolve(__dirname, ".."),
	});

	astroProcess.stderr.on("data", (data) => {
		console.error(`[Astro Error]: ${data}`);
	});

	// Health-check: poll until the server responds
	const baseUrl = "http://localhost:4321/cv-pipeline/";
	const maxRetries = 15;
	const retryDelay = 500;
	let serverReady = false;

	for (let i = 0; i < maxRetries; i++) {
		try {
			const res = await fetch(baseUrl);
			if (res.ok) {
				serverReady = true;
				break;
			}
		} catch {
			// Server not up yet
		}
		await new Promise((resolve) => setTimeout(resolve, retryDelay));
	}

	if (!serverReady) {
		console.error("❌ Astro preview server did not start in time.");
		astroProcess.kill();
		process.exit(1);
	}

	console.log("📸 Renderizando PDF con Puppeteer...");
	const browser = await puppeteer.launch({
		headless: true,
		args: ["--no-sandbox", "--disable-setuid-sandbox"],
	});

	const page = await browser.newPage();

	try {
		await page.goto(baseUrl, { waitUntil: "networkidle0" });

		// Inject private data into the DOM before capturing
		await injectPrivateData(page);

		await page.pdf({
			path: pdfPath,
			format: "A4",
			margin: {
				top: "15mm",
				bottom: "15mm",
				left: "15mm",
				right: "15mm",
			},
			printBackground: true,
		});

		console.log(`✅ PDF generado con éxito en: ${pdfPath}`);
	} catch (err) {
		console.error("❌ Error durante la generación del PDF:", err);
	} finally {
		await browser.close();
		astroProcess.kill();
		try {
			if (astroProcess.pid) {
				process.kill(-astroProcess.pid, "SIGINT");
			}
		} catch (e) {
			// Proceso ya cerrado, ignorar
		}
		process.exit(0);
	}
}

main().catch((err) => {
	console.error("❌ Error fatal en el script:", err);
	process.exit(1);
});

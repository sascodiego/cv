import puppeteer from "puppeteer";
import dotenv from "dotenv";
import { spawn } from "child_process";
import path from "path";
import { copyFile } from "fs/promises";
import { fileURLToPath } from "url";

// Cargar variables del .env local (datos privados, nunca commiteados)
dotenv.config();

// Technology label normalization mapper
// Converts source IDs (e.g., 'csharp') to human-friendly labels (e.g., 'C#')
const TECH_LABEL_MAP = {
	// Programming languages
	csharp: "C#",
	typescript: "TypeScript",
	javascript: "JavaScript",
	python: "Python",
	go: "Go",
	rust: "Rust",
	java: "Java",
	kotlin: "Kotlin",
	swift: "Swift",
	c: "C",
	cpp: "C++",
	ruby: "Ruby",
	php: "PHP",
	scala: "Scala",
	haskell: "Haskell",
	elixir: "Elixir",
	erlang: "Erlang",

	// Frameworks
	react: "React",
	vue: "Vue.js",
	angular: "Angular",
	svelte: "Svelte",
	nextjs: "Next.js",
	nuxt: "Nuxt.js",
	express: "Express.js",
	fastapi: "FastAPI",
	django: "Django",
	flask: "Flask",
	spring: "Spring",
	springboot: "Spring Boot",
	aspnet: "ASP.NET",
	laravel: "Laravel",

	// Databases
	postgresql: "PostgreSQL",
	mysql: "MySQL",
	sqlite: "SQLite",
	mongodb: "MongoDB",
	redis: "Redis",
	elasticsearch: "Elasticsearch",

	// Cloud and infrastructure
	aws: "AWS",
	azure: "Azure",
	gcp: "GCP",
	kubernetes: "Kubernetes",
	docker: "Docker",
	terraform: "Terraform",
	ansible: "Ansible",

	// DevOps and tools
	git: "Git",
	jenkins: "Jenkins",
	githubactions: "GitHub Actions",
	circleci: "CircleCI",
	prometheus: "Prometheus",
	grafana: "Grafana",

	// Messaging
	kafka: "Apache Kafka",
	rabbitmq: "RabbitMQ",
	activemq: "ActiveMQ",

	// IoT and embedded
	arduino: "Arduino",
	esp32: "ESP32",
	raspberrypi: "Raspberry Pi",
	micropython: "MicroPython",

	// Testing
	jest: "Jest",
	pytest: "PyTest",
	selenium: "Selenium",
	cypress: "Cypress",

	// Build tools
	webpack: "Webpack",
	vite: "Vite",
	babel: "Babel",
	maven: "Maven",
	gradle: "Gradle",
	npm: "npm",
	yarn: "Yarn",
	pnpm: "pnpm",

	// Other
	graphql: "GraphQL",
	grpc: "gRPC",
	rest: "REST",
	soap: "SOAP",
	oauth: "OAuth",
	jwt: "JWT",
	json: "JSON",
	xml: "XML",
	yaml: "YAML",
	markdown: "Markdown",
	latex: "LaTeX",
	vim: "Vim",
	emacs: "Emacs",
	vscode: "VS Code",
	linux: "Linux",
	windows: "Windows",
	macos: "macOS",
	bash: "Bash",
	powershell: "PowerShell",
	zsh: "Zsh",
};

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const pdfPath = path.resolve(__dirname, "../../Diego_Sasco_CV_Privado.pdf");
const publicPdfPath = path.resolve(
	__dirname,
	"../public/Diego_Sasco_CV_Privado.pdf",
);

// Datos privados inyectados desde .env
const privateData = {
	name: process.env.PRIVATE_NAME || "",
	email: process.env.PRIVATE_EMAIL || "",
	phone: process.env.PRIVATE_PHONE || "",
	address: process.env.PRIVATE_ADDRESS || "",
	website: process.env.PRIVATE_WEBSITE || "",
	github: process.env.PRIVATE_GITHUB || "",
	linkedin: process.env.PRIVATE_LINKEDIN || "",
};

/**
 * Injects private contact data into the rendered page DOM via Puppeteer.
 * The Astro page has no knowledge of these fields — this is the only place
 * where private data meets the rendered HTML.
 */
async function injectPrivateData(page) {
	await page.evaluate(
		(data, techMap) => {
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

			// Helper to create a separator
			function createSeparator() {
				const separator = document.createElement("span");
				separator.className = "contact-separator";
				separator.textContent = " · ";
				contactInfo.appendChild(separator);
			}

			// Collect all present contact items into an array
			const items = [];
			if (data.linkedin) items.push({ label: "LinkedIn", href: data.linkedin });
			if (data.email)
				items.push({ label: data.email, href: `mailto:${data.email}` });
			if (data.website) items.push({ label: "Portafolio", href: data.website });
			if (data.github) items.push({ label: "GitHub", href: data.github });

			// Append with separators
			items.forEach((item, index) => {
				createContactItem(item.label, item.href);
				if (index < items.length - 1) {
					createSeparator();
				}
			});

			// Normalize technology labels in skills section
			const skillItems = document.querySelectorAll(".skill-list li");
			skillItems.forEach((item) => {
				const textContent = item.textContent;
				// Technology names appear before parentheses
				const updatedText = textContent.replace(
					/([a-zA-Z0-9_-]+)\s*\(/g,
					(_match, techName) => {
						const normalized = normalizeTechLabel(techName, techMap);
						return `${normalized} (`;
					},
				);
				item.textContent = updatedText;
			});

			// Helper function available in the evaluate context
			function normalizeTechLabel(tag, map) {
				if (!tag || typeof tag !== "string") return "";
				const normalizedTag = tag.toLowerCase().trim();
				return (
					map[normalizedTag] ||
					normalizedTag.charAt(0).toUpperCase() + normalizedTag.slice(1)
				);
			}
		},
		privateData,
		TECH_LABEL_MAP,
	);
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
	const baseUrl = "http://localhost:4321/cv/cv-pdf";
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
				top: "8mm",
				bottom: "8mm",
				left: "8mm",
				right: "8mm",
			},
			printBackground: false,
		});

		await copyFile(pdfPath, publicPdfPath);
		console.log(`✅ PDF generado con éxito en: ${pdfPath}`);
		console.log(`📄 Copia pública actualizada en: ${publicPdfPath}`);
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

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func main() {
	inputPath := flag.String("in", "../data/cv-source.yaml", "Ruta al archivo YAML de origen")
	outputPath := flag.String("out", "../web/src/data/cv-processed.json", "Ruta al archivo JSON de salida")
	flag.Parse()

	if err := run(*inputPath, *outputPath); err != nil {
		fmt.Fprintf(os.Stderr, "❌ Error: %v\n", err)
		os.Exit(1)
	}
}

func run(inputPath, outputPath string) error {
	fmt.Printf("📖 Leyendo origen de datos: %s\n", inputPath)
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("no se pudo abrir el archivo de entrada: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("no se pudo leer el archivo de entrada: %w", err)
	}

	var raw CVRaw
	if err := yaml.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("error al parsear YAML: %w", err)
	}

	fmt.Println("⚙️  Procesando datos y calculando experiencia acumulada...")
	processed := ProcessCV(&raw)

	// Asegurar que el directorio de salida existe
	outDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return fmt.Errorf("no se pudo crear el directorio de salida %s: %w", outDir, err)
	}

	outBytes, err := json.MarshalIndent(processed, "", "  ")
	if err != nil {
		return fmt.Errorf("error al serializar JSON: %w", err)
	}

	if err := os.WriteFile(outputPath, outBytes, 0644); err != nil {
		return fmt.Errorf("no se pudo escribir el archivo de salida: %w", err)
	}

	fmt.Printf("✅ Datos procesados con éxito y guardados en: %s\n", outputPath)
	return nil
}

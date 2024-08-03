package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

func main() {
	jsonFile := flag.String("json", "", "Path to the JSON file")
	templateFile := flag.String("template", "", "Path to the template file")
	outputFile := flag.String("output", "output.md", "Path to the output file")
	showVersion := flag.Bool("version", false, "Show version and exit")

	version := os.Getenv("VERSION")
	if version == "" {
		version = "unknown" // Remplacez cela par la valeur par défaut souhaitée
	}

	flag.Parse()

	if *showVersion {
		fmt.Printf("template_renderer version %s\n", version)
		os.Exit(0)
	}

	if *jsonFile == "" || *templateFile == "" {
		log.Fatal("Both -json and -template flags must be provided")
	}

	// Read the JSON file
	jsonData, err := ioutil.ReadFile(*jsonFile)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}

	// Parse the JSON data
	var data map[string]interface{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		log.Fatalf("Failed to parse JSON data: %v", err)
	}

	// Read the template file
	templateData, err := ioutil.ReadFile(*templateFile)
	if err != nil {
		log.Fatalf("Failed to read template file: %v", err)
	}

	// Parse the template
	tmpl, err := template.New("template").Parse(string(templateData))
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}

	// Execute the template with the data
	output, err := os.Create(*outputFile)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer output.Close()

	if err := tmpl.Execute(output, data); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	fmt.Println("Template successfully executed and output written to", *outputFile)
}

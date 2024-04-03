package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/vqvw/runic"
)

func main() {
	inputFile := flag.String("in", "", "File to parse")
	outputFile := flag.String("out", "", "Output file")

	flag.Parse()

	fileType := filepath.Ext(*outputFile)

	if *inputFile == "" {
		log.Fatalln("No input file provided, use -in [filename]")
	}

	inputBytes, err := os.ReadFile(*inputFile)
	if err != nil {
		log.Fatal("Unable to read input file:\n", err)
	}

	parser := runic.New()

	var output []byte

	switch fileType {
	case "html":
		output = []byte(parser.Html(string(inputBytes)))
	default:
		tree := parser.Parse(string(inputBytes))
		treeJsonBytes, err := json.MarshalIndent(tree, "", "  ")
		if err != nil {
			log.Fatal("Failed to marshal JSON:\n", err)
		}
		output = treeJsonBytes
	}

	if *outputFile == "" {
		fmt.Println(output)
		return
	}

	err = os.WriteFile(*outputFile, output, 0644)
	if err != nil {
		log.Fatal("Failed to write output file:\n", err)
	}
}

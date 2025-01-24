package main

import (
	"encoding/json"
	"flag"
	"fmt"
	v1 "github.com/kyzrfranz/buntesdach/api/v1"
	"github.com/kyzrfranz/interventure-cli/internal/cmd"
	"os"
)

var (
	url     string
	outFile string
)

func main() {
	flag.StringVar(&url, "apiUrl", cmd.EnvOrString("API_URL", "http://localhost:8080"), "buntesdach API URL")
	flag.StringVar(&outFile, "out", cmd.EnvOrString("OUTPUT_FILE", "bio.json"), "name the output json file")

	var bios []v1.Politician

	bios = cmd.FetchPoliticians(url)
	err := writeToJson(bios)
	cmd.NoErrorOrExit(err)

	fmt.Printf("Successfully generated %s\n", outFile)
}

func writeToJson(politicians []v1.Politician) error {
	data, err := json.Marshal(politicians)
	if err != nil {
		return err
	}
	return os.WriteFile(outFile, data, 0644)
}

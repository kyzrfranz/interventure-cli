package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	v1 "github.com/kyzrfranz/buntesdach/api/v1"
	"github.com/kyzrfranz/interventure-cli/internal/cmd"
	"github.com/kyzrfranz/interventure-cli/internal/openai"
	"github.com/kyzrfranz/interventure-cli/pkg/xls"
	"html/template"
	"io"
	"net/http"
	"os"
)

var (
	input          string
	output         string
	tfType         string
	oapiToken      string
	promptTemplate string
)

func main() {
	flag.StringVar(&input, "in", cmd.EnvOrString("INPUT_FILE", "bio.json"), "path to a json file containing politician bios (generated by 'bio' command)")
	flag.StringVar(&output, "out", cmd.EnvOrString("OUTPUT_PATH", ".tmp"), "name of the output file/folder")
	flag.StringVar(&tfType, "type", cmd.EnvOrString("TRANSFORM_TYPE", "xsl"), "can be 'xsl', 'website'")
	flag.StringVar(&oapiToken, "token", cmd.EnvOrString("OPENAI_TOKEN", "123"), "OpenAI API token")
	flag.StringVar(&promptTemplate, "promptTemplate", cmd.EnvOrString("PROMPT_TEMPLATE", "./templates/prompt.tpl"), "OpenAI prompt template")
	flag.Parse()

	data, err := os.ReadFile(input)
	cmd.NoErrorOrExit(err)
	var bios []v1.Politician
	err = json.Unmarshal(data, &bios)
	cmd.NoErrorOrExit(err)

	if tfType == "xsl" {
		err = makeXLS(bios)
	} else if tfType == "website" {
		err = makeWebsiteData(bios)
	} else if tfType == "prompt" {
		fmt.Println(makePrompt(bios))
	} else if tfType == "img" {
		err = downloadImages(bios, output)
	} else {
		fmt.Printf("unknown transform type: %s\n", tfType)
	}
	cmd.NoErrorOrExit(err)
}

func makeXLS(bios []v1.Politician) error {
	xlsGenerator := xls.NewGenerator()
	return xlsGenerator.Generate(output, bios)
}

func makePrompt(bios []v1.Politician) string {

	tpl, err := template.ParseFiles(promptTemplate)
	cmd.NoErrorOrExit(err)
	var buffer bytes.Buffer
	err = tpl.Execute(&buffer, bios)
	cmd.NoErrorOrExit(err)

	return buffer.String()
}

func makeWebsiteData(bios []v1.Politician) error {

	extracts := make([]openai.Extract, 0)

	extractor := openai.NewExtractor(oapiToken)
	for _, pol := range bios {
		if pol.Bio.Homepage != "" {
			xtr, err := extractor.Extract(pol.Bio.Homepage)
			if err != nil {
				fmt.Printf("error extracting data from %s: %v \n", pol.Bio.Homepage, err)
			} else {
				extracts = append(extracts, xtr)
			}
		}
	}

	fmt.Printf("%v", extracts)

	return nil
}

func downloadImages(bios []v1.Politician, tmpDir string) error {

	if _, err := os.Stat(tmpDir); os.IsNotExist(err) {
		if cErr := os.Mkdir(tmpDir, 0755); cErr != nil {
			return cErr
		}
	}

	noImg := 0

	for _, pol := range bios {
		if pol.Media.Foto.URL != "" {
			fmt.Printf("downloading image from %s\n", pol.Media.Foto.URL)
			resp, err := http.Get(pol.Media.Foto.URL)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			img, err := io.ReadAll(resp.Body)
			//display file size
			if err != nil {
				return err
			}
			err = os.WriteFile(fmt.Sprintf("%s/%s.jpg", tmpDir, pol.Bio.Id.Value), img, 0644)
			fmt.Printf("saved image to %s/%s.jpg - size: %d\n", tmpDir, pol.Bio.Id.Value, len(img))
			if len(img) == 1541 {
				noImg++
			}
			if err != nil {
				return err
			}
		}
	}

	fmt.Printf("no image found for %d politicians\n", noImg)

	return nil
}

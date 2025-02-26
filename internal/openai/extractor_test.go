package openai

import (
	"context"
	"fmt"
	"github.com/kyzrfranz/interventure-cli/internal/scraper"
	"github.com/stretchr/testify/assert"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"os"
	"testing"
)

func TestExtractor_Extract(t *testing.T) {
	token := os.Getenv("OPENAI_API")
	e := Extractor{apiToken: token}
	archiveUrl, err := scraper.GetFromArchive("https://www.jensteutrine.de")
	assert.NoError(t, err)
	res, err := e.Extract(archiveUrl)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestGenerateSchema(t *testing.T) {
	llm, err := ollama.New(ollama.WithModel("llama3.3"))

	prompt := "ich gebe dir den content einer webseite eines MdB. Du sollst einen kurzen zusammenfassungstext über die Person insbes. bezüglich arbeits und sozialpolitischen themen vornehmen, sowie folgende kontakldaten sofern vorhanden: email, anschrift bundestag, anschrift wahlkreis, telefonnummer, faxnummer, homepage"
	data := ""

	assert.NoError(t, err)
	cpl, err := llm.Call(context.Background(),
		fmt.Sprintf("%s. Die daten zu analysieren sind: %s", prompt, data),
		llms.WithTemperature(0.1),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			return nil
		}),
	)

	assert.NoError(t, err)
	_ = cpl
}

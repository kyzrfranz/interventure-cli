package openai

import (
	"github.com/kyzrfranz/interventure-cli/internal/scraper"
	"github.com/stretchr/testify/assert"
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

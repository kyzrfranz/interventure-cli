package scraper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScrape(t *testing.T) {

	url := "http://web.archive.org/web/20240719215532/https://sanae-abdi.spd.de/"
	text, err := Scrape(url)
	if err != nil {
		t.Error(err)
	}

	println(text)
}

func TestGetFromArchive(t *testing.T) {

	url := "https://sanae-abdi.spd.de"
	archiveUrl, err := GetFromArchive(url)
	assert.NoError(t, err)
	println(archiveUrl)

}

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

func TestGetProtocol(t *testing.T) {

	url := "https://www.bundestag.de/mediathek?videoid=7615711&url=L21lZGlhdGhla292ZXJsYXk=&mod=mediathek#url=L21lZGlhdGhla292ZXJsYXk/dmlkZW9pZD03NjE1NzExJnVybD1MMjFsWkdsaGRHaGxhMjkyWlhKc1lYaz0mbW9kPW1lZGlhdGhlaw==&mod=mediathek"
	protocol, err := Scrape(url)
	assert.NoError(t, err)
	println(protocol)

}

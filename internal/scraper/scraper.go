package scraper

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"net/http"
	"regexp"
	"strings"
)

type archiveResponse struct {
	Url               string `json:"url"`
	ArchivedSnapshots struct {
		Closest struct {
			Status    string `json:"status"`
			Available bool   `json:"available"`
			Url       string `json:"url"`
			Timestamp string `json:"timestamp"`
		} `json:"closest"`
	} `json:"archived_snapshots"`
}

func Scrape(url string) (string, error) {
	var text string

	c := colly.NewCollector(
		colly.MaxDepth(2),
	)

	// Find and visit all links
	c.OnHTML("html", func(e *colly.HTMLElement) {
		text += e.DOM.Text()
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	err := c.Visit(url)

	return text, err

}

func protocolCollector() *colly.Collector {
	c := colly.NewCollector(
		colly.MaxDepth(2),
	)

	// Find and visit all links
	c.OnHTML(".bt-linkliste", func(e *colly.HTMLElement) {
		fmt.Printf("Found link: %s\n", e.Attr("href"))
		//e.Request.Visit(e.Attr("href")) //NOT TODAY
	})

	return c
}

func dingsCollector(text string) *colly.Collector {
	c := colly.NewCollector(
		colly.MaxDepth(2),
	)

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		//e.Request.Visit(e.Attr("href")) //NOT TODAY
	})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		e.DOM.Find("script, style").Remove()
		html, _ := e.DOM.Html()
		text += html
	})

	c.OnRequest(func(r *colly.Request) {

		//skip non doc urls
		if strings.Contains(r.URL.String(), "javascript") ||
			strings.Contains(r.URL.String(), "https://www.facebook") ||
			strings.Contains(r.URL.String(), "https://www.instagram") ||
			strings.Contains(r.URL.String(), "jpg") ||
			strings.Contains(r.URL.String(), "jpeg") ||
			strings.Contains(r.URL.String(), "png") ||
			strings.Contains(r.URL.String(), "pdf") {
			return
		}

		fmt.Println("Visiting", r.URL)
	})

	return c
}

func trimText(text string) string {
	re := regexp.MustCompile(`\s+`)
	cleaned := re.ReplaceAllString(text, " ")
	return strings.TrimSpace(cleaned)
}

func GetFromArchive(url string) (string, error) {
	baseUrl := fmt.Sprintf("https://archive.org/wayback/available?url=%s", url)
	response, err := http.Get(baseUrl)
	if err != nil {
		return "", err
	}
	if response.StatusCode != 200 {
		return "", fmt.Errorf("could not fetch url: %s", baseUrl)
	}

	var responseJson archiveResponse
	err = json.NewDecoder(response.Body).Decode(&responseJson)
	if err != nil {
		return "", err
	}

	if responseJson.ArchivedSnapshots.Closest.Available == false {
		if strings.HasPrefix(url, "https://") {
			return GetFromArchive(strings.Replace(url, "https://", "http://", 1))
		}
		return "", fmt.Errorf("no archive available for url: %s", url)
	}

	return responseJson.ArchivedSnapshots.Closest.Url, nil
}

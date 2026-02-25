package scrapy

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func FetchData(url string) (string, string, error) {
	return fetchData(url)
}

func fetchData(url string) (string, string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", "", fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", "", err
	}

	title := doc.Find("head title").Text()

	description, exists := doc.Find("head meta[name='description']").Attr("content")
	if exists {
		return title, description, nil
	}

	return title, "", nil
}

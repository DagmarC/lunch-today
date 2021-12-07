package scraper

import (
	"io"
	"log"
	"net/http"

	"github.com/DagmarC/lunch-today/utils"
	"github.com/PuerkitoBio/goquery"
)

func loadDoc(url string) *goquery.Document {

	res, err := http.Get(url)
	utils.Check(err)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		utils.Check(err)
	}(res.Body)

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	utils.Check(err)

	return doc
}

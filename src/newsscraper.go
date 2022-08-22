package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type Content struct {
	Description string
	Link        string
}

// Website 1: https://www.univaq.it/news_archive.php?tipo=In%20evidenza
/*
	It finds all the links within the page by looking at the ones with
	the "avviso_container" class on the respective div
*/
func GetLinksUnivaq() {

	// Scraper created (Collector)
	scraper := colly.NewCollector(colly.AllowedDomains("univaq.it/*"))

	// On every element which has href attribute, fetch it
	scraper.OnHTML("div.avviso_container", func(h *colly.HTMLElement) {

		fmt.Println(h.Text)
	})

	verr := scraper.Visit("https://www.univaq.it/news_archive.php?tipo=In%20evidenza")

	if verr != nil {
		fmt.Printf("Qualcosa è anadato storto1: %v", verr)
	}
}

// Website 2: https://www.disim.univaq.it/news.php?entrant=1

func GetLinksDisim() {

	// Scraper created (Collector)
	scraper := colly.NewCollector(colly.AllowedDomains("www.univaq.it"))

	// On every element which has href attribute, fetch it
	scraper.OnHTML("a[href]", func(h *colly.HTMLElement) {
		l := h.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", h.Text, l)

	})
	verr := scraper.Visit("https://www.disim.univaq.it/news.php?entrant=1")

	if verr != nil {
		fmt.Printf("Qualcosa è anadato storto2: %v", verr)
	}
}

func main() {

	GetLinksDisim()
}

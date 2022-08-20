package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

// Website 1: https://www.univaq.it/news_archive.php?tipo=In%20evidenza

func GetLinksUnivaq() []string {

	var foundLinks []string

	// Scraper created (Collector)
	scraper := colly.NewCollector(colly.AllowedDomains("www.univaq.it"))

	// On every element which has href attribute, fetch it
	scraper.OnHTML("a[href]", func(h *colly.HTMLElement) {
		l := h.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", h.Text, l)

	})

	return foundLinks
}

// Website 2: https://www.disim.univaq.it/news.php?entrant=1
func GetLinksDisim() []string {
	var foundLinks []string

	// Scraper created (Collector)
	scraper := colly.NewCollector(colly.AllowedDomains("www.univaq.it"))

	// On every element which has href attribute, fetch it
	scraper.OnHTML("h5[a[href]]", func(h *colly.HTMLElement) {
		l := h.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", h.Text, l)

	})

	return foundLinks
}

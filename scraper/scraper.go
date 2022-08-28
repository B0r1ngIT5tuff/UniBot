package scraper

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/mymmrac/telego"
)

type UserRequest struct {
	FirstCity  string
	SecondCity string
	FirstDate  [3]string
	SecondDate [3]string
	N_adults   uint
	N_students uint
	N_children uint
	HandheldL  uint
	HoldL      uint
}

func GetPlaneOffers(userData UserRequest) {

	// Scraper created (Collector)
	scraper := colly.NewCollector(colly.AllowedDomains(""))

	// On every element which has href attribute, fetch it
	scraper.OnHTML("shish", func(h *colly.HTMLElement) {

		fmt.Println(h.Text)
	})

	verr := scraper.Visit("")

	if verr != nil {
		fmt.Printf("Qualcosa è andato storto: %v", verr)
	}
}

func GetB_and_B_Offers() {

	// Scraper created (Collector)
	scraper := colly.NewCollector(colly.AllowedDomains(""))

	// On every element which has href attribute, fetch it
	scraper.OnHTML("a[href]", func(h *colly.HTMLElement) {

	})
	verr := scraper.Visit("")

	if verr != nil {
		fmt.Printf("Qualcosa è anadato storto2: %v", verr)
	}
}

func ParseData(m telego.Message) {

	// Parse message with regex
	// Regex for cities: [A-Z]
	// Regex for dates : (\d{4}-\d{2}-\d{2})
	// Regex for adults: (\dadults)
	// Regex for adults: (\dstudents)
	// Regex for adults: (\dchildren)
	// Regex for handheld luggage: (cfc)
	// Regex for hold luggage: (bfc)
}

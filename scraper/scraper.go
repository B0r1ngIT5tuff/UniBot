package scraper

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/gocolly/colly/v2"
)

type UserNews struct {
	Text  string `json:"text"`
	Title string `json:"title"`
}

/*
This method writes all the data included in the struct in a json file.
Its purpose is to be used by the CheckNews() function to see if there are any changes.
*/
func WriteToJson(filename string, u []UserNews) (int, error) {
	jFile, cerr := os.Create(filename) // Creates a new file

	// Error checking
	if cerr != nil {
		fmt.Println("There has been an error while creating the json file : " + cerr.Error())
		os.Exit(1)
	}
	jcontent, err := json.Marshal(u) // Converts the struct into json format

	// Error checking
	if err != nil {
		fmt.Println("There has been an error while converting to json format : " + err.Error())
		os.Exit(1)
	}

	n, ferr := jFile.Write(jcontent) // Writes the struct in json format in the json file

	return n, ferr
}

/*
This function Fetches all the news from the website.
It returns all the information with a slice of UserNews ([]UserNews)
and writes the news to a file for further news checking.
*/
func GetNews() []UserNews {

	// Scraper created (Collector)
	scraper := colly.NewCollector(
		colly.AllowedDomains("www.univaq.it"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246"),
	)

	var news []UserNews // News for the user

	// On every element which has href attribute, fetch it
	scraper.OnHTML("div.full-container-20 div.avviso_container", func(h *colly.HTMLElement) {

		userNews := UserNews{
			Text:  h.ChildAttr("a", "href"),          // Fetches the links of the articles in the page
			Title: h.ChildText("div.avviso a[href]"), // Fetches the Day of publication
		}

		news = append(news, userNews)
	})

	verr := scraper.Visit("https://www.univaq.it/news_archive.php?tipo=In%20evidenza")

	if verr != nil {
		fmt.Printf("Something went wrong: %v", verr)
	}

	_, ferr := WriteToJson("news.json", news)
	// Error checking
	if ferr != nil {
		fmt.Println("There has been an error while writing to the file : " + ferr.Error())
		os.Exit(1)
	}

	return news
}

/*
It checks if there are new articles on the page.
The check is performed by calling the GetNews() function and comparing
the structs fetched with those written in the json file.
If the result of the comparison is true nothing happens (are identical, no new articles),
instead if the result is false (the structs are different) the new article will be sent to the user.
*/
func CheckNews() bool {

	var newsSt []UserNews
	checkN := GetNews()                // Fetches for new infromations
	jFile, err := os.Open("news.json") // Opens the file

	// Error checking
	if err != nil {
		fmt.Println("There has been an error while opening the file in read mode: " + err.Error())
		os.Exit(1)
	}
	fileN, ferr := io.ReadAll(jFile) // Reads the content of the file

	// Error checking
	if ferr != nil {
		fmt.Println("There has been an error while reading the file : " + ferr.Error())
		os.Exit(1)
	}

	uerr := json.Unmarshal(fileN, &newsSt)
	// Error checking
	if uerr != nil {
		fmt.Println("There has been an error while parsing the file to []UserNews: " + uerr.Error())
		os.Exit(1)
	}

	if !compareNews(newsSt, checkN) {
		return false
	}
	return true
}

/*
This function compares two slice of type UserNews.
*/
func compareNews(a, b []UserNews) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

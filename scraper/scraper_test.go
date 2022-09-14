package scraper

import (
	"testing"
)

func TestGetNews(t *testing.T) {
	u := GetNews()

	for i, v := range u {
		if string(v.Text[i]) == "" || string(v.Title[i]) == "" {
			t.Failed()
		}
	}
}

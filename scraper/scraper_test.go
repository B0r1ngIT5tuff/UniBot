package scraper

import (
	"testing"

	"github.com/mymmrac/telego"
)

func TestGetPLaneOffers(t *testing.T) {

}

func TestGetB_and_B_Offers(t *testing.T) {

}

func TestParseData(t *testing.T) {

	ur := UserRequest{}

	m := telego.Message{Text: "CIA-PSR/2022-08-07/2022-08-29/2adults"}

	m1 := telego.Message{Text: "Ciao bro"}
	_, err := ur.ParseData(m)

	if err != nil {
		t.Failed()
	}

	_, err1 := ur.ParseData(m1)

	if err1 != nil {
		t.Failed()
	}

}

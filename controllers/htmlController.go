package controllers

import (
	"fmt"
	"log"
	"net/http"

	"../models"
	"github.com/PuerkitoBio/goquery"
)

//GetValuesHTML permit to use the HTML from the specific page and select the need data
//using a query from JS
func GetValuesHTML(page string, domain *models.Domain) {
	response, err := http.Get("http://" + page)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", response.StatusCode, response.Status)
	}

	// Load the HTML document
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	domain.Title = document.Find("title").Text()

	logo, exist := document.Find("link[rel=\"shortcut icon\"]").First().Attr("href")
	if exist {
		fmt.Println("href: ", logo)
		domain.Logo = logo
	}
}

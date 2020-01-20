package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"../models"
)

//ReadDomainInfo access to an API and use the data from JSON
func ReadDomainInfo(page string) models.Domain {

	var domain models.Domain
	response, err := http.Get("https://api.ssllabs.com/api/v3/analyze?host=" + page)
	if err != nil {
		domain.IsDown = true
		log.Fatal(err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		//fmt.Println(string(data))

		domain = models.Domain{}
		_ = json.Unmarshal([]byte(data), &domain)
	}

	ReadWhoisIP(&domain)
	CalculateServersGrade(&domain)
	GetValuesHTML(page, &domain)

	return domain
}

//CalculateServersGrade is a function tha calculate the less SSLGrade
func CalculateServersGrade(domain *models.Domain) {
	if len(domain.Servers) > 0 {
		domain.SSLGrade = "F"
		for i := 0; i < len(domain.Servers); i++ {
			if domain.Servers[i].SSLGrade < domain.SSLGrade {
				domain.SSLGrade = domain.Servers[i].SSLGrade
			}
		}
	}
}

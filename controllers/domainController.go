package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

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
	domain.LastQuery = time.Now()

	//fmt.Println("json:....", domain)

	return domain
}

//CalculateServersGrade is a function tha calculate the less SSLGrade
func CalculateServersGrade(domain *models.Domain) {
	var grades = []string{"A-", "A", "A+", "B-", "B", "B+", "C-", "C", "C+", "D-", "D", "D+", "E-", "E", "E+", "F-", "F", "F+"}
	if len(domain.Servers) > 0 {
		domain.SSLGrade = "F+"

		var grade, serverGrade int
		for i := 0; i < len(domain.Servers); i++ {
			for j := 0; j < len(grades); j++ {
				if domain.Servers[i].SSLGrade == grades[j] {
					serverGrade = j
				}
				if domain.SSLGrade == grades[j] {
					grade = j
				}
			}
			if serverGrade < grade {
				domain.SSLGrade = domain.Servers[i].SSLGrade
			}
		}
	}
}

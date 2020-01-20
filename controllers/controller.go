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
		log.Fatal(err)
		domain.IsDown = true
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		//fmt.Println(string(data))

		domain = models.Domain{}
		_ = json.Unmarshal([]byte(data), &domain)
	}

	ReadWhoisIP(&domain)
	return domain
}

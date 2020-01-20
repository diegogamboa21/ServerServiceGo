package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"../models"
)

//ReadWhoisIP is a function to find data about the server IP using to an API
func ReadWhoisIP(domain *models.Domain) {
	for i := 0; i < len(domain.Servers); i++ {
		response, err := http.Get("http://ip-api.com/json/" + domain.Servers[i].Address)
		if err != nil {
			log.Fatal(err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			//fmt.Println(string(data))

			type WhoisIP struct {
				Country string `json:"countryCode"`
				Owner   string `json:"isp"`
			}

			whoisIP := WhoisIP{}
			_ = json.Unmarshal([]byte(data), &whoisIP)

			domain.Servers[i].Country = whoisIP.Country
			domain.Servers[i].Owner = whoisIP.Owner
		}
	}
}

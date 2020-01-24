package utils

import (
	"fmt"

	"../models"
)

//PrintDomainInfo show the domain info
func PrintDomainInfo(domain models.Domain) {

	for _, s := range domain.Servers {
		PrintServerInfo(s)
	}

	fmt.Println("serversChanged: ", domain.ServersChanged)
	fmt.Println("SSLGrade: ", domain.SSLGrade)
	fmt.Println("PreviusSSLGrade: ", domain.PreviusSSLGrade)
	fmt.Println("Logo: ", domain.Logo)
	fmt.Println("title: ", domain.Title)
	fmt.Println("IsDown: ", domain.IsDown)
	fmt.Println("LastQuery", domain.LastQuery)
}

//PrintServerInfo show the domain info
func PrintServerInfo(server models.Server) {
	fmt.Println()
	fmt.Println("Address: ", server.Address)
	fmt.Println("SSLGrade: ", server.SSLGrade)
	fmt.Println("Country: ", server.Country)
	fmt.Println("Owner: ", server.Owner)
}

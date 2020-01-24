package main

import (
	"./controllers"
)

func main() {
	page := "trello.com"
	domain := controllers.ReadDomainInfo(page)
	controllers.ConnectDB()
	controllers.InsertDomainOnDB(page, &domain)
	controllers.DisconnectDB()
}

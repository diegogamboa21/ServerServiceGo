package main

import (
	"fmt"

	"./controllers"
)

func main() {
	domain := controllers.ReadDomainInfo("google.com")
	//utils.PrintDomainInfo(domain)
	controllers.ConnectDB()
	controllers.InsertDomainOnDB(&domain)
	controllers.DisconnectDB()
	fmt.Println("Finished")
}

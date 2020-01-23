package main

import (
	"fmt"

	"./controllers"
)

func main() {
	controllers.ConnectDB()
	fmt.Println("Init Data Base")
	domain := controllers.ReadDomainInfo("amazon.com")
	//utils.PrintDomainInfo(domain)
	controllers.InsertDomainOnDB(&domain)
	fmt.Println("Finished")
}

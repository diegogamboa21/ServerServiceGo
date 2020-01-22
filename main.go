package main

import (
	"fmt"

	"./controllers"
)

func main() {
	controllers.InitDB()
	fmt.Println("Init Data Base")
	//domain := controllers.ReadDomainInfo("truora.com")
	//utils.PrintDomainInfo(domain)
	fmt.Println("Finished")
}

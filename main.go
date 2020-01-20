package main

import (
	"./controllers"
	"./utils"
)

func main() {
	domain := controllers.ReadDomainInfo("truora.com")
	utils.PrintDomainInfo(domain)
}

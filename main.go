package main

import (
	"fmt"

	"./controllers"
)

func main() {
	domain := controllers.ReadDomainInfo("amazon.com")
	fmt.Println(domain)
}

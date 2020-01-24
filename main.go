package main

import (
	"./controllers"
)

func main() {

	mux := controllers.Routes()
	server := controllers.NewServer(mux)
	server.Run()
}

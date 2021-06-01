package main

import (
	"log"

	"github.com/angeldhakal/tv-tracker/routes"
)

func main() {
	router := routes.MainRouter()
	log.Fatalln(router.Run(":8080"))
}

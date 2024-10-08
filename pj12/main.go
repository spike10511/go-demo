package main

import (
	"log"
	"pj12/routes"
)

func main() {
	err := routes.InitRoutes()
	if err != nil {
		log.Fatalf("Failed to initialize routes: %v", err)
	}
}

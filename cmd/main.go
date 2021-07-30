package main

import (
	"fmt"
	"log"
	"net/http"

	"go-grafana/internal/router"
)

// main
func main() {
	var port = "3000"
	fmt.Println("Initializing routes..")
	// initialize routes
	r := router.InitRoutes()

	fmt.Println("Server started. Listening to port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

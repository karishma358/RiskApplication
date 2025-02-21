package main

import (
	"log"
	"net/http"
	"riskapp/routes"
)

func main() {
	r := routes.SetupRouter()

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

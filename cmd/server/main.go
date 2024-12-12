package main

import (
	"HealthHub-Connect/internal/db"
	"HealthHub-Connect/routes"
	"log"
	"net/http"
)

func main() {
	db.InitDatabase()
	db.Migrate()

	routes.HomeRouters()

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

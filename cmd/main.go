package main

import (
	"log"
	"net/http"
	"os"

	"ip-malicious-db/controllers"
	"ip-malicious-db/routes"
	"ip-malicious-db/services"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	port := os.Getenv("SERVER_PORT")
	databaseURL := os.Getenv("NEO4J_DB")
	user := os.Getenv("NEO4J_USER")
	password := os.Getenv("NEO4J_PASSWORD")

	neo4jService := services.NewNeo4jService(databaseURL, user, password)
	defer neo4jService.Close()

	githubService := services.NewGithubIpService()
	ipMaliciousService := services.NewMaliciousIpService(neo4jService, githubService)

	metasploitController := controllers.NewMaliciousIpController(ipMaliciousService)
	routes.RegisterRoutes(router, metasploitController)

	log.Printf("Server running on %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/mastertinner/handler"
	"gitlab.com/klud/graphql-docker-api/schema"
)

// Gets endpoint from API_ENDPOINT env var, else assigns one by default
func apiEndpoint() string {
	end := os.Getenv("API_ENDPOINT")
	if end == "" {
		end = "/graphql"
	}
	return end
}

// Gets the port to serve on from GQL_PORT env var, else assigns one by default
func getPort() string {
	port := os.Getenv("GQL_PORT")
	if port == "" {
		port = ":8080"
	}
	return port
}

// Enables GraphiQL by default, if 0 is set GraphiQL is disabled.
func enableUI() bool {
	env := os.Getenv("GRAPHIQL")
	if env == "0" {
		return false
	}
	return true
}

func main() {
	var uiStatus string
	ui := enableUI()
	h := handler.New(&handler.Config{
		Schema:   &schema.Schema,
		Pretty:   true,
		GraphiQL: ui,
	})
	http.Handle(apiEndpoint(), h)
	if ui {
		uiStatus = "enabled"
	} else {
		uiStatus = "disabled"
	}
	log.Printf("Endpoint: %v | GraphiQL: %v | Serving on %v\n", apiEndpoint(), uiStatus, getPort())
	log.Fatalln(http.ListenAndServe(getPort(), handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)))
}

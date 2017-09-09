package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
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
	h := handler.New(&handler.Config{
		Schema:   &schema.Schema,
		Pretty:   true,
		GraphiQL: enableUI(),
	})
	cyanBold := color.New(color.FgCyan, color.Bold)
	if enableUI() {
		cyanBold.Printf("GraphiQL is enabled by default.\nSet GRAPHIQL=0 in order to disable, this is recommended when using in production.\n")
	} else {
		cyanBold.Println("GraphiQL has been disabled.")
	}
	bold := color.New(color.FgYellow, color.Bold)
	bold.Printf("Endpoint: %v | Serving on %v\n", apiEndpoint(), getPort())
	http.Handle(apiEndpoint(), h)
	log.Fatalln(http.ListenAndServe(getPort(), handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)))
}

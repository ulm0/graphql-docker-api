package main

import (
	"log"
	"os"

	"github.com/urfave/negroni"

	"github.com/gorilla/mux"
	"gitlab.com/ulm0/graphql-docker-api/resolver"
	"gitlab.com/ulm0/graphql-docker-api/schema"
	graphql "gitlab.com/ulm0/graphql-go"
	"gitlab.com/ulm0/graphql-go/relay"
)

var scheme *graphql.Schema

// Gets the api route from the API_ENDPOINT env var, else assigns one by default
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

func init() {
	scheme = graphql.MustParseSchema(schema.Schema, &resolver.Resolver{})
}

func main() {
	h := relay.New(&relay.Config{
		Schema:   scheme,
		GraphiQL: enableUI(),
	})

	if enableUI() {
		log.Printf("Endpoint: %v | GraphiQL: %v", apiEndpoint(), enableUI())
	} else {
		log.Printf("Endpoint: %v | GraphiQL: %v", apiEndpoint(), enableUI())
	}

	r := mux.NewRouter()
	r.Handle(apiEndpoint(), h).Methods("GET", "POST")
	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(getPort())
}

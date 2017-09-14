package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	graphql "github.com/neelance/graphql-go"
	"github.com/neelance/graphql-go/relay"
	"gitlab.com/klud/graphql-docker-api/resolver"
	"gitlab.com/klud/graphql-docker-api/schema"
)

var scheme *graphql.Schema

// Gets GraphiQL route from UI_ENDPOINT env var, else assigns one by default
func uiEndpoint() string {
	end := os.Getenv("UI_ENDPOINT")
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
	var uiStatus string

	http.Handle("/query", &relay.Handler{Schema: scheme})

	if enableUI() {
		uiStatus = "enabled"
		http.Handle(uiEndpoint(), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write(page)
		}))
		log.Printf("Endpoint: /query | GraphiQL: %v | UI route: %v | Serving on %v\n", uiStatus, uiEndpoint(), getPort())
	} else {
		uiStatus = "disabled"
		log.Printf("Endpoint: /query | GraphiQL: %v | Serving on %v\n", uiStatus, getPort())
	}

	log.Fatalln(http.ListenAndServe(getPort(), handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)))
}

var page = []byte(`
	<!DOCTYPE html>
	<html>
		<head>
			<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.css" />
			<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/1.1.0/fetch.min.js"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.js"></script>
		</head>
		<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
			<div id="graphiql" style="height: 100vh;">Loading...</div>
			<script>
				function graphQLFetcher(graphQLParams) {
					return fetch("/query", {
						method: "post",
						body: JSON.stringify(graphQLParams),
						credentials: "include",
					}).then(function (response) {
						return response.text();
					}).then(function (responseBody) {
						try {
							return JSON.parse(responseBody);
						} catch (error) {
							return responseBody;
						}
					});
				}
	
				ReactDOM.render(
					React.createElement(GraphiQL, {fetcher: graphQLFetcher}),
					document.getElementById("graphiql")
				);
			</script>
		</body>
	</html>
`)

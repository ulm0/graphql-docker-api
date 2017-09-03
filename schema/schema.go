package schema

import (
	"golang.org/x/net/context"

	"github.com/docker/docker/client"
	"github.com/graphql-go/graphql"
)

var (
	ctx = context.Background()

	cli, _ = client.NewEnvClient()

	queryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		// TODO implement queries logic
		// List:
		// - Containers (args: all bool, limit int, size bool, filter ContainerFilter)
		// - Images (args: all bool, digest bool, filter ImageFilter)
		// - Networks (args: filter NetworkFilter)
		// - SearchImagesOnDockerHub (args: term string!)
		// - System
		// - Volumes (args: filter VolumeFilter)

		Fields: graphql.Fields{
			// // List containers matching various filters
			// "containers": &graphql.Field{
			// 	Type: container,
			// 	// This must return a container slice
			// 	// It can receive (args: all bool, limit int, size bool, filter containerFilter)
			// 	// Resolver
			// },
			// // List images matching various filters
			// "images": &graphql.Field{
			// 	Type: image,
			// 	// This must return an image slice
			// 	// It can receive (args: all bool, digests bool, filter imageFilter)
			// 	// Resolver
			// },
			// // List networks matching various filters
			// "networks": &graphql.Field{
			// 	Type: network,
			// 	// This must return a network slice
			// 	// It can receive (args: filter networFilter)
			// 	// Resolver
			// },
			// // Search images on Docker Hub
			// "searchImagesOnDockerHub": &graphql.Field{
			// 	Type: imageSearchResult,
			// 	// It must return a imageSearchResult slice
			// 	// It MUST receive (args: term string!)
			// 	// Resolver
			// },
			// // Return more details about the Swarm. Requires Docker API >= 1.24
			// "swarm": &graphql.Field{
			// 	Type: swarm,
			// 	// Resolver
			// },
			// System-wide information
			"system": &graphql.Field{
				Type: system,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return cli, nil
				},
			},
			// // List volumes matching various filters
			// "volumes": &graphql.Field{
			// 	Type: volumeList,
			// 	// It can receive (args: filter volumeFilter)
			// 	// Resolver
			// },
		},
	})
	// Schema for queries and mutations
	Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
)

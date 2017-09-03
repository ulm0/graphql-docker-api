package schema

import "github.com/graphql-go/graphql"

var (
	// Image related custom object types
	image = graphql.NewObject(graphql.ObjectConfig{
		Name: "Image",
		Fields: graphql.Fields{
			"created": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"details": &graphql.Field{
				Type: imageDetails,
				// Resolver
			},
			"history": &graphql.Field{
				Type: imageHistory,
				// Side Notes: this needs to return an slice of imageHistory
				// so probably the type needs to be tweaked
				// Resolver
			},
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
				// Resolver
			},
			"labels": &graphql.Field{
				Type: stringAnyMap,
				// Resolver
			},
			"parentId": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"repoDigests": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"repoTags": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"size": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"virtualSize": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
		},
	})
	imageDetails = graphql.NewObject(graphql.ObjectConfig{
		Name: "ImageDetails",
		Fields: graphql.Fields{
			"architecture": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"author": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"comment": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"config": &graphql.Field{
				Type: containerConfig,
				// Resolver
			},
			"container": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"containerConfig": &graphql.Field{
				Type: containerConfig,
				// Resolver
			},
			"created": &graphql.Field{
				Type: graphql.DateTime,
				// Resolver
			},
			"dockerVersion": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
				// Resolver
			},
			"os": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"parent": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"rootFs": &graphql.Field{
				Type: rootFs,
				// Resolver
			},
			"size": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"virtualSize": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
		},
	})
	imageFilter = graphql.NewObject(graphql.ObjectConfig{
		Name: "ImageFilter",
		Fields: graphql.Fields{
			"before": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"dangling": &graphql.Field{
				Type: graphql.Boolean,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"label": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"reference": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"since": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
		},
	})
	imageHistory = graphql.NewObject(graphql.ObjectConfig{
		Name: "ImageHistory",
		Fields: graphql.Fields{
			"comment": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"created": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"createdBy": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
				// Resolver
			},
			"size": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"tags": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
		},
	})
	imageSearchResult = graphql.NewObject(graphql.ObjectConfig{
		Name: "ImageSearchResult",
		Fields: graphql.Fields{
			"automated": &graphql.Field{
				Type: graphql.Boolean,
				// Resolver
			},
			"description": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"name": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"official": &graphql.Field{
				Type: graphql.Boolean,
				// Resolver
			},
			"starCount": &graphql.Field{
				Type: graphql.Int,
				// Resolver
			},
		},
	})
	rootFs = graphql.NewObject(graphql.ObjectConfig{
		Name: "RootFs",
		Fields: graphql.Fields{
			"layers": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"type": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
		},
	})
)

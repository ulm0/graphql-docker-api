package schema

import "github.com/graphql-go/graphql"

var (
	// Volume related custom object types
	volume = graphql.NewObject(graphql.ObjectConfig{
		Name: "Volume",
		Fields: graphql.Fields{
			"driver": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"driverOpts": &graphql.Field{
				Type: stringAnyMap,
				// Resolver
			},
			"labels": &graphql.Field{
				Type: stringAnyMap,
				// Resolver
			},
			"mountpoint": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"name": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"scope": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"status": &graphql.Field{
				Type: stringAnyMap,
				// Resolver
			},
		},
	})
	volumeFilter = graphql.NewObject(graphql.ObjectConfig{
		Name: "VolumeFilter",
		Fields: graphql.Fields{
			"dangling": &graphql.Field{
				Type: graphql.Boolean,
				// Side Notes: this needs to return an slice of Boolean
				// so probably the type needs to be tweaked
				// Resolver
			},
			"driver": &graphql.Field{
				Type: graphql.String,
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
			"name": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
		},
	})
	volumeList = graphql.NewObject(graphql.ObjectConfig{
		Name: "VolumeList",
		Fields: graphql.Fields{
			"volumes": &graphql.Field{
				Type: volume,
				// Side Notes: this needs to return an slice of volume
				// so probably the type needs to be tweaked
				// Resolver
			},
			"warnings": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
		},
	})
)

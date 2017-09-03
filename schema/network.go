package schema

import "github.com/graphql-go/graphql"

var (
	// Network related custom object types
	ipam = graphql.NewObject(graphql.ObjectConfig{
		Name: "Ipam",
		Fields: graphql.Fields{
			"config": &graphql.Field{
				Type: ipamConfig,
				// Side Notes: this needs to return an slice of ipamConfig
				// so probably the type needs to be tweaked
				// Resolver
			},
			"driver": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
		},
	})
	ipamConfig = graphql.NewObject(graphql.ObjectConfig{
		Name: "IpamConfig",
		Fields: graphql.Fields{
			"gateway": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"ipRange": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"subnet": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
		},
	})
	network = graphql.NewObject(graphql.ObjectConfig{
		Name: "Network",
		Fields: graphql.Fields{
			"attachable": &graphql.Field{
				Type: graphql.Boolean,
				// Resolver
			},
			"containers": &graphql.Field{
				Type: stringAnyMap,
				// Resolver
			},
			"driver": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"enableIPv6": &graphql.Field{
				Type: graphql.Boolean,
				// Resolver
			},
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
				// Resolver
			},
			"internal": &graphql.Field{
				Type: graphql.Boolean,
				// Resolver
			},
			"ipam": &graphql.Field{
				Type: ipam,
				// Resolver
			},
			"labels": &graphql.Field{
				Type: stringAnyMap,
				// Resolver
			},
			"name": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"options": &graphql.Field{
				Type: stringAnyMap,
				// Resolver
			},
			"scope": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
		},
	})
	networkFilter = graphql.NewObject(graphql.ObjectConfig{
		Name: "NetworkFilter",
		Fields: graphql.Fields{
			"driver": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of strings
				// so probably the type needs to be tweaked
				// Resolver
			},
			"id": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of strings
				// so probably the type needs to be tweaked
				// Resolver
			},
			"label": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of strings
				// so probably the type needs to be tweaked
				// Resolver
			},
			"name": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of strings
				// so probably the type needs to be tweaked
				// Resolver
			},
			"scope": &graphql.Field{
				Type: networkScope,
				// Side Notes: this needs to return an slice of strings
				// so probably the type needs to be tweaked
				// Resolver
			},
			"type": &graphql.Field{
				Type: networkType,
				// Side Notes: this needs to return an slice of strings
				// so probably the type needs to be tweaked
				// Resolver
			},
		},
	})
	networkScope = graphql.NewEnum(graphql.EnumConfig{
		Name: "ContainerHealth",
		Values: graphql.EnumValueConfigMap{
			"global": &graphql.EnumValueConfig{
				Value: "global",
			},
			"local": &graphql.EnumValueConfig{
				Value: "local",
			},
			"swarm": &graphql.EnumValueConfig{
				Value: "swarm",
			},
		},
	})
	networkType = graphql.NewEnum(graphql.EnumConfig{
		Name: "ContainerHealth",
		Values: graphql.EnumValueConfigMap{
			"builtin": &graphql.EnumValueConfig{
				Value: "builtin",
			},
			"custom": &graphql.EnumValueConfig{
				Value: "custom",
			},
		},
	})
)

package schema

import "github.com/graphql-go/graphql"

var (
	// Scalar custom types
	date         = graphql.NewScalar(graphql.ScalarConfig{})
	stringAnyMap = graphql.NewScalar(graphql.ScalarConfig{})
	stringSet    = graphql.NewScalar(graphql.ScalarConfig{})
)

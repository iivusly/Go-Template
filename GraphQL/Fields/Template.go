package Fields

import (
	// GraphQL
	"github.com/graphql-go/graphql"
)

var Template = &graphql.Field{
	Type: graphql.String,
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		return "Hello, world!", nil
	},
}

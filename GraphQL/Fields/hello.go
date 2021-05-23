package Fields

import (
	// GraphQL
	"github.com/graphql-go/graphql"
)

var Hello = &graphql.Field{
	Type: graphql.String,
	Resolve: func (params graphql.ResolveParams) (interface{}, error) {
		return "world", nil
	},
}
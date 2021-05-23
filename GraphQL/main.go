package GraphQL

import (
	"log"

	"github.com/Calpico-Drink/Go-Template/GraphQL/Fields"

	// GraphQL
	"github.com/graphql-go/graphql"
)

func Init() graphql.Schema {

	// Change here
	fields := graphql.Fields{
		"hello": Fields.Hello,
	}


	rootQuery := graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: fields,
	}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	return schema
}
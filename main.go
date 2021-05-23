package main

import (
	"log"

	// GraphQL
	"github.com/graphql-go/graphql"

	"github.com/Calpico-Drink/Go-Template/util"
)

func main() {
	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	Serve(&schema, util.GetEnv("PORT", "3000"))
}

/*
Docker

--Build
docker build -t graphql .


*/

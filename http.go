package main

import (
	"fmt"

	// GraphQL
	"github.com/graphql-go/graphql"

	// GraphQL HTTP
	"net/http"
	"github.com/graphql-go/handler"
)

func Serve(Schema *graphql.Schema, port string) {
	nHandler := handler.New(&handler.Config{
		Schema: Schema,
	})

	http.Handle("/api", nHandler)

	fmt.Println("Serving on port ", port)

	http.ListenAndServe(port, nil)
}
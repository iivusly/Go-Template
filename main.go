package main

import (
	"github.com/Calpico-Drink/Go-Template/GraphQL"
	"github.com/Calpico-Drink/Go-Template/util"
)

func main() {
	schema := GraphQL.Init()

	GraphQL.Serve(&schema, util.GetEnv("PORT", "3000"))
}

/*
Docker

--Build
docker build -t graphql .

*/

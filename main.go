package main

import (
	"github.com/Calpico-Drink/Go-Template/util"
	"github.com/Calpico-Drink/Go-Template/GraphQL"
)

func main() {
	schema := GraphQL.Init()

	Serve(&schema, util.GetEnv("PORT", "3000"))
}

/*
Docker

--Build
docker build -t graphql .

*/

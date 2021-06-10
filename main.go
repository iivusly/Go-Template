package main

import (
	"github.com/Calpico-Drink/Go-Template/Database"
	"github.com/Calpico-Drink/Go-Template/GraphQL"
	"github.com/Calpico-Drink/Go-Template/util"
)

func main() {
	Database.Init("localhost", "5432", "postgres", "", "")

	schema := GraphQL.Init()

	GraphQL.Serve(&schema, util.GetEnv("PORT", "3333"))
}

/*
Docker

--Build
docker build -t graphql .

*/

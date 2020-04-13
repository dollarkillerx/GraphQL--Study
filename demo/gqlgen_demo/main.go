package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"graphql_demo/demo/gqlgen_demo/graph"
	"graphql_demo/demo/gqlgen_demo/resolver"
	"log"
	"net/http"
)

func main() {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%v/ for GraphQL playground", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"log"
	"testing"
)

func TestDemo1(t *testing.T) {
	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				fmt.Println("============")
				fmt.Println(p.Args)
				fmt.Println(p.Info)
				fmt.Println(p.Source)
				return "world", nil
			},
		},
	}

	// Root节点
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalln(err)
	}

	// Query
	query := `
	{
		hello
	}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("Errors: %+v\n", r.Errors)
	}
	marshal, err := json.Marshal(r)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(marshal))
}

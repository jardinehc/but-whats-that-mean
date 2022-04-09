package main

import (
	_ "embed"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

//go:embed schema.graphql
var schema string

type resolver struct{}

func (_ *resolver) Hello() string { return "Hello, world!" }

func main() {

	schema := graphql.MustParseSchema(schema, &resolver{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

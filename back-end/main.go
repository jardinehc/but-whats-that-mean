package main

import (
	_ "embed"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	log "github.com/sirupsen/logrus"
)

//go:embed schema.graphql
var schema string

type resolver struct{}

func (_ *resolver) Hello() string { return "Hello, world!" }

func main() {

	schema := graphql.MustParseSchema(schema, &resolver{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	http.Handle("/", playground.Handler("playground", "/query"))
	log.Fatal(http.ListenAndServe(":8080", nil))
	log.Info("listening on port 8080...")

}

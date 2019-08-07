package main

import (
	"github.com/friendsofgo/graphiql"
	gql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/lucasew/golog"
	"github.com/lucasew/graphql_postgres/graph"
	"io/ioutil"
	"net/http"
)

var log = golog.Default.NewLogger("main")

func main() {
	f, err := ioutil.ReadFile("./schema.graphql")
	if err != nil {
		panic(err)
	}
	schema := gql.MustParseSchema(string(f), &graph.Resolver{})
	http.Handle("/graphql", &relay.Handler{Schema: schema})
	ui, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		log.Error("Erro ao aplicar graphiql: %s", err.Error())
	}
	http.Handle("/graphiql", ui)
	log.Info("Iniciando...")
	http.ListenAndServe(":8080", nil)
}

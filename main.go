package main

import (
    "github.com/lucasew/graphql_postgress/graph"
    "net/http"
    gql "github.com/graph-gophers/graphql-go"
    "github.com/graph-gophers/graphql-go/relay"
    "io/ioutil"
    "github.com/lucasew/golog"
    "github.com/friendsofgo/graphiql"
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

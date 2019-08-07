package graph

import (
    "github.com/lucasew/graphql_postgres/db"
    "github.com/lucasew/graphql_postgres/gather"
    "github.com/lucasew/golog"
)


var log = golog.Default.NewLogger("graph")
var dummy int32 = 0
func (_ Resolver) GetPessoa(v struct{Id *int32}) (*db.Pessoa, error) {
    log.NewLogger("pessoa").Info("getPessoa: %d", *v.Id)
    ret, err := gather.Pessoa{}.GetByID(*v.Id)
    return &ret, err
}

func (_ Resolver) AddPessoa(v struct{P *db.PessoaInput}) (*db.Pessoa, error) {
    log.NewLogger("pessoa").Info("addPessoa: %s %s", v.P.Nome, v.P.Sobrenome)
    ret, err := gather.Pessoa{}.AddPessoa(*v.P)
    return &ret, err
}

func (_ Resolver) Pessoas() (*[]*db.Pessoa, error) {
    log.NewLogger("pessoa").Info("pessoas")
    res, err := gather.Pessoa{}.Pessoas()
    ret := make([]*db.Pessoa, len(res))
    for i := 0; i < len(res); i++ {
        ret[i] = &res[i]
    }
    return &ret, err
}

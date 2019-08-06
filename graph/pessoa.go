package graph

import (
    "github.com/lucasew/graphql_postgress/db"
    "errors"
)

func (_ Resolver) GetPessoa(v struct{Id *int32}) (*db.Pessoa, error) {
    var ret []db.Pessoa
    db.DB.Where("id = ?", *v.Id).Find(&ret)
    if (len(ret) == 0) {
        return nil, errors.New("registro não encontrado")
    }
    return &ret[0], nil
}

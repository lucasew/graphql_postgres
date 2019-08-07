package graph

import (
    "github.com/lucasew/graphql_postgress/db"
    "github.com/lucasew/graphql_postgress/gather"
)

func (_ Resolver) GetCasa(v struct{ID *int32}) (*db.Casa, error) {
    log.NewLogger("casa").Info("getCasa: %d", *v.ID)
    ret, err := gather.Casa{}.GetCasa(*v.ID)
    return &ret, err
}

func (_ Resolver) Casas() (*[]*db.Casa, error) {
    c, err := gather.Casa{}.Casas()
    log.Info("%#v", c)
    if err != nil {
        return nil, err
    }
    ret := make([]*db.Casa, len(c))
    for i := 0; i < len(c); i++ {
        ret[i] = &c[i]
    }
    return &ret, nil
}

func (_ Resolver) AddCasa(v struct {C *db.CasaInput} ) (*db.Casa, error) {
    c, err := gather.Casa{}.AddCasa(*v.C)
    log.Info("%d", len(*v.C.Pessoas))
    return &c, err
}

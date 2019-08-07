package graph

import (
	"github.com/lucasew/graphql_postgres/db"
	"github.com/lucasew/graphql_postgres/gather"
)

// GetCasa graphql method
func (Resolver) GetCasa(v struct{ ID *int32 }) (*db.Casa, error) {
	log.NewLogger("casa").Info("getCasa: %d", *v.ID)
	ret, err := gather.Casa{}.GetCasa(*v.ID)
	return &ret, err
}

// Casas graphql method
func (Resolver) Casas() (*[]*db.Casa, error) {
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

// AddCasa graphql method
func (Resolver) AddCasa(v struct{ C *db.CasaInput }) (*db.Casa, error) {
	c, err := gather.Casa{}.AddCasa(*v.C)
	log.Info("%d", len(*v.C.Pessoas))
	return &c, err
}

package gather_test

import (
	"fmt"
	"github.com/lucasew/graphql_postgres/db"
	"github.com/lucasew/graphql_postgres/gather"
	"testing"
)

func TestCasaCasas(t *testing.T) {
	v, err := gather.Casa{}.Casas()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Casa.Casas: %#v\n", v)
}

func TestCasaGetCasa(t *testing.T) {
	v, err := gather.Casa{}.GetCasa(1)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Casa.GetCasa: %#v\n", v)
}

func TestCasaAddCasa(t *testing.T) {
	ca := db.Casa{
		Numero: 1234,
		Rua:    "dos bobos",
		Pessoas: []db.Pessoa{
			{
				Nome:      "teste",
				Sobrenome: "eoq",
			},
			{
				Nome:      "João",
				Sobrenome: "Santos",
			},
		},
	}
	v := db.CasaInput{
		Numero: &ca.Numero,
		Rua:    &ca.Rua,
		Pessoas: &[]*db.PessoaInput{
			&db.PessoaInput{Nome: &ca.Pessoas[0].Nome, Sobrenome: &ca.Pessoas[0].Sobrenome},
			&db.PessoaInput{Nome: &ca.Pessoas[1].Nome, Sobrenome: &ca.Pessoas[1].Sobrenome},
		},
	}
	ca, err := gather.Casa{}.AddCasa(v)
	t.Logf("%#v\n", ca)
	if err != nil {
		t.Error(err)
	}
}

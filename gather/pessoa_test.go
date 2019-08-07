package gather_test

import (
    "testing"
    "github.com/lucasew/graphql_postgress/gather"
    "github.com/lucasew/graphql_postgress/db"
    "fmt"
)


func TestPessoaPessoas(t *testing.T) {
    v, err := gather.Pessoa{}.Pessoas()
    if (err != nil) {
        t.Error(err)
    }
    fmt.Printf("Pessoas: %#v\n", v)
}

func TestPessoaByID(t *testing.T) {
    v, err := gather.Pessoa{}.GetByID(1)
    if (err != nil) {
        t.Error(err)
    }
    fmt.Printf("Pessoa by id: %#v\n", v)
}

func TestAddPessoa(t *testing.T) {
    p := db.Pessoa{
        Nome: "John",
        Sobrenome: "Doe",
    }
    in := db.PessoaInput{
        Nome: &p.Nome,
        Sobrenome: &p.Sobrenome,
    }
    p, err := gather.Pessoa{}.AddPessoa(in)
    if err != nil {
        t.Error(err)
    }
    fmt.Printf("AddPessoa: %#v\n", p)
    // limpeza
    err = db.DB.Where("nome = 'John'").Delete(&db.Pessoa{}).Error
    if err != nil {
        t.Error(err)
    }
}


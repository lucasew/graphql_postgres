package gather

import (
    "github.com/lucasew/graphql_postgres/db"
)

type Pessoa struct {}

func (_ Pessoa) Pessoas() ([]db.Pessoa, error) {
    var ret []db.Pessoa
    err := db.DB.Find(&ret).Error
    return ret, err
}

func (_ Pessoa) GetByID(id int32) (db.Pessoa, error) {
    var p db.Pessoa
    return p, db.DB.Where("id = ?", id).Find(&p).Error
}

func (self Pessoa) AddPessoa(p db.PessoaInput) (db.Pessoa, error) {
    inserir := p.Convert()
    return inserir, db.DB.Create(&inserir).Error
}

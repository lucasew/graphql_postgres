package gather

import (
	"github.com/lucasew/graphql_postgres/db"
)

// Pessoa container for pessoa related methods
type Pessoa struct{}

// Pessoas get all pessoas on the BD
func (Pessoa) Pessoas() ([]db.Pessoa, error) {
	var ret []db.Pessoa
	err := db.DB.Find(&ret).Error
	return ret, err
}

// GetByID get a pessoa object by id
func (Pessoa) GetByID(id int32) (db.Pessoa, error) {
	var p db.Pessoa
	return p, db.DB.Where("id = ?", id).Find(&p).Error
}

// AddPessoa inserts a new pessoa on BD
func (Pessoa) AddPessoa(p db.PessoaInput) (db.Pessoa, error) {
	inserir := p.Convert()
	return inserir, db.DB.Create(&inserir).Error
}

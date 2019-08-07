package gather

import (
	"github.com/lucasew/graphql_postgres/db"
)

// Casa container for casa related operations on DB
type Casa struct{}

// Casas get all objects from db
func (s Casa) Casas() ([]db.Casa, error) {
	var c []db.Casa
	db.DB.Find(&c)
	for i := 0; i < len(c); i++ {
		err := s.fillPessoa(&c[i])
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

// GetCasa get the casa of given ID
func (s Casa) GetCasa(id int32) (db.Casa, error) {
	var c db.Casa
	db.DB.Where("id = ?", id).Find(&c)
	return c, s.fillPessoa(&c)
}

// fillPessoa fills the pessoas field of casa struct
func (Casa) fillPessoa(c *db.Casa) error {
	return db.DB.
		Where("id in (?)", db.DB.
			Select("pessoa_id").
			Table("casa_pessoa").
			Where("casa_id = ?", c.Id).
			SubQuery()).
		Find(&c.Pessoas).Error
}

// AddCasa insert a new casa on the DB
func (Casa) AddCasa(in db.CasaInput) (db.Casa, error) {
	c := in.Convert()
	return c, db.DB.Create(&c).Error
}

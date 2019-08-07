package gather

import (
    "github.com/lucasew/graphql_postgress/db"
)

type Casa struct {}

func (this Casa) Casas() ([]db.Casa, error) {
    var c []db.Casa
    db.DB.Find(&c)
    for i:=0; i < len(c); i++ {
        err := this.fillPessoa(&c[i])
        if err != nil {
            return nil, err
        }
    }
    return c, nil
}

func (this Casa) GetCasa(id int32) (db.Casa, error) {
    var c db.Casa
    db.DB.Where("id = ?", id).Find(&c)
    return c, this.fillPessoa(&c)
}

func (_ Casa) fillPessoa(c *db.Casa) error {
    return db.DB.
        Where("id in (?)", db.DB.
            Select("pessoa_id").
            Table("casa_pessoa").
            Where("casa_id = ?", c.Id).
            SubQuery()).
        Find(&c.Pessoas).Error
}

func (_ Casa) AddCasa(in db.CasaInput) (db.Casa, error) {
    c := in.Convert()
    return c, db.DB.Create(&c).Error;
}

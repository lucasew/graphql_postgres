package db

type Pessoa struct {
    Id int32
    Nome string
    Sobrenome string
}

func (p Pessoa) ID() *int32 {
    return &p.Id
}

func (p Pessoa) NOME() *string {
    return &p.Nome
}

func (p Pessoa) SOBRENOME() *string {
    return &p.Sobrenome
}

func init() {
    <-dbDone
    DB.Debug().AutoMigrate(&Pessoa{})
}
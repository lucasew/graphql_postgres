package db

// PessoaInput the object that is given by the api user
type PessoaInput struct {
	Nome      *string
	Sobrenome *string
}

// Convert Convert PessoaInput to a DB processable object
func (p PessoaInput) Convert() Pessoa {
	return Pessoa{
		Nome:      *p.Nome,
		Sobrenome: *p.Sobrenome,
	}
}

// Pessoa The DB processable Pessoa object
type Pessoa struct {
	Id        int32
	Nome      string
	Sobrenome string
}

// ID api getter
func (p Pessoa) ID() *int32 {
	return &p.Id
}

// NOME api getter
func (p Pessoa) NOME() *string {
	return &p.Nome
}

// SOBRENOME api getter
func (p Pessoa) SOBRENOME() *string {
	return &p.Sobrenome
}

func init() {
	go func() {
		<-dbDone
		DB.Debug().AutoMigrate(&Pessoa{})
	}()
}

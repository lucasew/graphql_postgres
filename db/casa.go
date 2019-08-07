package db

// CasaInput object received from graphql api.
type CasaInput struct {
	Numero  *int32
	Rua     *string
	Pessoas *[]*PessoaInput
}

// Convert convert the input to the prefered type for database
func (c CasaInput) Convert() Casa {
	pessoas := make([]Pessoa, len(*c.Pessoas))
	for i := 0; i < len(*c.Pessoas); i++ {
		pessoas[i] = (*c.Pessoas)[i].Convert()
	}
	return Casa{
		Numero:  *c.Numero,
		Rua:     *c.Rua,
		Pessoas: pessoas,
	}
}

// Casa Casa type to working with gorm and database
type Casa struct {
	Id      int32
	Numero  int32
	Rua     string
	Pessoas []Pessoa `gorm:"many2many:casa_pessoa;association_jointable_foreignkey:pessoa_id"`
}

// ID getter when returing to graphql api
func (c Casa) ID() *int32 {
	return &c.Id
}

// NUMERO getter when returning to graphql api
func (c Casa) NUMERO() *int32 {
	return &c.Numero
}

// RUA getter when returning to graphql api
func (c Casa) RUA() *string {
	return &c.Rua
}

// PESSOAS getter when returning to graphql api
func (c Casa) PESSOAS() *[]*Pessoa {
	p := make([]*Pessoa, len(c.Pessoas))
	for i := 0; i < len(c.Pessoas); i++ {
		p[i] = &c.Pessoas[i]
	}
	return &p
}

func init() {
	go func() {
		<-dbDone
		DB.Debug().AutoMigrate(&Casa{})
	}()
}

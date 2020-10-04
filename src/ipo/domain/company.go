package domain

type CompanyId string

type Company struct {
	symbol string
	name   string
}

func NewCompany(symbol string, name string) Company {
	return Company{symbol, name}
}

func (c Company) Symbol() string {
	return c.symbol
}

func (c Company) Name() string {
	return c.name
}

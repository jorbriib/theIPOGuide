package domain

type CompanyId string

type Company struct{
	id CompanyId
	symbol string
	name string
}

func NewCompany(id CompanyId, symbol string, name string ) Company{
	return Company{id, symbol, name}
}

func (c Company) Id() CompanyId {
	return c.id
}

func (c Company) Symbol() string {
	return c.symbol
}

func (c Company) Name() string {
	return c.name
}
package domain

type IpoRepository interface {
	Find() ([]Ipo, error)
	GetByAlias(alias string) (*Ipo, error)
}

type MarketRepository interface {
	FindByIds([]MarketId) ([]Market, error)
	GetById(MarketId) (*Market, error)
}

type CompanyRepository interface {
	FindByIds([]CompanyId) ([]Company, error)
	GetById(CompanyId) (*Company, error)
}

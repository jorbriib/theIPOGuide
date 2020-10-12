package domain

type IpoRepository interface {
	Find() ([]Ipo, error)
}

type MarketRepository interface {
	FindByIds([]MarketId) ([]Market, error)
}

type CompanyRepository interface {
	FindByIds([]CompanyId) ([]Company, error)
}

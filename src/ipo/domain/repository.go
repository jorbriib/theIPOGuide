package domain

type IpoRepository interface {
	Find(marketId MarketId, countryId CountryId, sectorId SectorId, industryId IndustryId, blackList []IpoId, offset int, limit int) ([]Ipo, error)
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

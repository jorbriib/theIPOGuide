package domain

type IpoRepository interface {
	Find(marketIds []MarketId, countryIds []CountryId, sectorIds []SectorId, industryIds []IndustryId, blackList []IpoId, sortBy string, offset uint, limit uint) ([]Ipo, error)
	Count(marketIds []MarketId, countryIds []CountryId, sectorIds []SectorId, industryIds []IndustryId, blackList []IpoId) (uint, error)
	SearchByText(text string) ([]Ipo, error)
	GetByAlias(alias string) (*Ipo, error)
}

type CompanyRepository interface {
	FindByIds([]CompanyId) ([]Company, error)
	GetById(CompanyId) (*Company, error)
}


type MarketRepository interface {
	All() ([]Market, error)
	FindByIds([]MarketId) ([]Market, error)
	GetById(MarketId) (*Market, error)
	FindByCodes(codes []string) ([]Market, error)
}

type SectorRepository interface {
	All() ([]Sector, error)
	FindByAliases(aliases []string) ([]Sector, error)
}

type CountryRepository interface {
	All() ([]Country, error)
	FindByCodes(codes []string) ([]Country, error)
}

package application

import (
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
	"log"
)

// SearchByTextService is the service to manage the IPOs
type SearchByTextService struct {
	ipoRepository     domain.IpoRepository
	marketRepository  domain.MarketRepository
	companyRepository domain.CompanyRepository
	countryRepository domain.CountryRepository
	sectorRepository  domain.SectorRepository
}

// NewSearchByTextService returns the SearchByTextService
func NewSearchByTextService(
	ipoRepository domain.IpoRepository,
	marketRepository domain.MarketRepository,
	companyRepository domain.CompanyRepository,
	countryRepository domain.CountryRepository,
	sectorRepository domain.SectorRepository,
) SearchByTextService {
	return SearchByTextService{ipoRepository, marketRepository, companyRepository, countryRepository, sectorRepository}
}

// SearchByTextQuery returns a struct
type SearchByTextQuery struct {
	text string
}

// NewSearchByTextQuery returns the query used by GetIPOs method
func NewSearchByTextQuery(text string) SearchByTextQuery {
	return SearchByTextQuery{text}
}

// SearchByTextResponse is the response from GetIPOs method
type SearchByTextResponse struct {
	ipos      []domain.Ipo
	markets   []domain.Market
	companies []domain.Company
}

// Get returns the response data
func (r SearchByTextResponse) Get() ([]domain.Ipo, []domain.Market, []domain.Company) {
	return r.ipos, r.markets, r.companies
}

// Run obtains IPOs and related data
func (h SearchByTextService) Run(query SearchByTextQuery) (*SearchByTextResponse, error) {

	ipos, err := h.ipoRepository.SearchByText(query.text)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	mapMarketIds := make(map[domain.MarketId]domain.MarketId)
	mapCompanyIds := make(map[domain.CompanyId]domain.CompanyId)
	for _, ipo := range ipos {
		mapMarketIds[ipo.MarketId()] = ipo.MarketId()
		mapCompanyIds[ipo.CompanyId()] = ipo.CompanyId()
	}

	marketIds := make([]domain.MarketId, len(mapMarketIds))
	i := 0
	for _, marketId := range mapMarketIds {
		marketIds[i] = marketId
		i++
	}
	markets, err := h.marketRepository.FindByIds(marketIds)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	companyIds := make([]domain.CompanyId, len(mapCompanyIds))
	i = 0
	for _, companyId := range mapCompanyIds {
		companyIds[i] = companyId
		i++
	}

	companies, err := h.companyRepository.FindByIds(companyIds)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &SearchByTextResponse{
		ipos,
		markets,
		companies,
	}, nil
}

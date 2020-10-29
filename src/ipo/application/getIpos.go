package application

import (
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
	"log"
)

const defaultLimit uint = 20

// GetIposService is the service to manage the IPOs
type GetIposService struct {
	ipoRepository     domain.IpoRepository
	marketRepository  domain.MarketRepository
	companyRepository domain.CompanyRepository
	countryRepository domain.CountryRepository
	sectorRepository  domain.SectorRepository
}

// NewGetIposService returns the GetIposService
func NewGetIposService(
	ipoRepository domain.IpoRepository,
	marketRepository domain.MarketRepository,
	companyRepository domain.CompanyRepository,
	countryRepository domain.CountryRepository,
	sectorRepository domain.SectorRepository,
) GetIposService {
	return GetIposService{ipoRepository, marketRepository, companyRepository, countryRepository, sectorRepository}
}

// GetIposQuery returns a struct
type GetIposQuery struct {
	marketCodes   []string
	countryCodes  []string
	sectorAliases []string
	page uint
}

// NewGetIposQuery returns the query used by GetIPOs method
func NewGetIposQuery(marketsCode []string, countriesCode []string, sectorsAlias []string, page uint) GetIposQuery {
	return GetIposQuery{marketsCode, countriesCode, sectorsAlias, page}
}

// GetIposResponse is the response from GetIPOs method
type GetIposResponse struct {
	total     uint
	ipos      []domain.Ipo
	markets   []domain.Market
	companies []domain.Company
}

// Get returns the response data
func (r GetIposResponse) Get() (uint, []domain.Ipo, []domain.Market, []domain.Company) {
	return r.total, r.ipos, r.markets, r.companies
}

// Run obtains IPOs and related data
func (h GetIposService) Run(query GetIposQuery) (*GetIposResponse, error) {
	selectedMarkets, err := h.marketRepository.FindByCodes(query.marketCodes)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	selectedMarketIds := make([]domain.MarketId, len(selectedMarkets))
	i := 0
	for _, selectedMarketId := range selectedMarkets {
		selectedMarketIds[i] = selectedMarketId.Id()
		i++
	}

	selectedCountries, err := h.countryRepository.FindByCodes(query.countryCodes)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	selectedCountryIds := make([]domain.CountryId, len(selectedCountries))
	i = 0
	for _, selectedCountryId := range selectedCountries {
		selectedCountryIds[i] = selectedCountryId.Id()
		i++
	}

	selectedSectors, err := h.sectorRepository.FindByAliases(query.sectorAliases)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	selectedSectorIds := make([]domain.SectorId, len(selectedSectors))
	i = 0
	for _, selectedSectorId := range selectedSectors {
		selectedSectorIds[i] = selectedSectorId.Id()
		i++
	}

	offset := query.page * defaultLimit
	ipos, err := h.ipoRepository.Find(selectedMarketIds, selectedCountryIds, selectedSectorIds, []domain.IndustryId{}, []domain.IpoId{}, "", offset, defaultLimit)
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
	i = 0
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

	total, err := h.ipoRepository.Count(selectedMarketIds, selectedCountryIds, selectedSectorIds, []domain.IndustryId{}, []domain.IpoId{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &GetIposResponse{
		total,
		ipos,
		markets,
		companies,
	}, nil
}

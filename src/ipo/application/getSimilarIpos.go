package application

import (
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
	"log"
)

// GetSimilarIposService is the service to manage the IPOs
type GetSimilarIposService struct {
	ipoRepository     domain.IpoRepository
	marketRepository  domain.MarketRepository
	companyRepository domain.CompanyRepository
}

// NewGetSimilarIposService returns the IPoService
func NewGetSimilarIposService(
	ipoRepository domain.IpoRepository,
	marketRepository domain.MarketRepository,
	companyRepository domain.CompanyRepository,
) GetSimilarIposService {
	return GetSimilarIposService{ipoRepository, marketRepository, companyRepository}
}

// GetSimilarIposQuery returns a struct
type GetSimilarIposQuery struct {
	alias string
}

// Alias returns the response data
func (q GetSimilarIposQuery) Alias() string {
	return q.alias
}

// NewGetSimilarIposQuery returns the query used by GetIPO method
func NewGetSimilarIposQuery(alias string) GetSimilarIposQuery {
	return GetSimilarIposQuery{alias}
}

// GetIpoResponse is the response from GetIPO method
type GetSimilarIposResponse struct {
	ipos      []domain.Ipo
	markets   []domain.Market
	companies []domain.Company
}

// Get returns the response data
func (r GetSimilarIposResponse) Get() ([]domain.Ipo, []domain.Market, []domain.Company) {
	return r.ipos, r.markets, r.companies
}

// Run obtains a IPO and related data
func (h GetSimilarIposService) Run(query GetSimilarIposQuery) (*GetSimilarIposResponse, error) {
	ipo, err := h.ipoRepository.GetByAlias(query.alias)
	if err != nil {
		return nil, err
	}
	if ipo == nil {
		return nil, nil
	}

	company, err := h.companyRepository.GetById(ipo.CompanyId())
	if err != nil {
		return nil, err
	}
	if company == nil {
		return nil, nil
	}

	marketId := ipo.MarketId()
	sectorId := company.Sector().Id()

	blackList := []domain.IpoId{ipo.Id()}
	similarIpos, err := h.ipoRepository.Find(
		[]domain.MarketId{marketId},
		[]domain.CountryId{},
		[]domain.SectorId{sectorId},
		[]domain.IndustryId{},
		blackList,
		"random", 0, 5)
	if err != nil {
		return nil, err
	}

	if len(similarIpos) < 5 {
		for _, ipo := range similarIpos {
			blackList = append(blackList, ipo.Id())
		}
		similarSectorIpos, err := h.ipoRepository.Find(
			[]domain.MarketId{},
			[]domain.CountryId{},
			[]domain.SectorId{sectorId},
			[]domain.IndustryId{},
			blackList,
			"random",0, uint(5-len(similarIpos)))
		if err != nil {
			return nil, err
		}
		similarIpos = append(similarIpos, similarSectorIpos...)
	}

	if len(similarIpos) < 5 {
		for _, ipo := range similarIpos {
			blackList = append(blackList, ipo.Id())
		}
		similarMarketIpos, err := h.ipoRepository.Find(
			[]domain.MarketId{marketId},
			[]domain.CountryId{},
			[]domain.SectorId{},
			[]domain.IndustryId{},
			blackList,
			"random",0, uint(5-len(similarIpos)))
		if err != nil {
			return nil, err
		}
		similarIpos = append(similarIpos, similarMarketIpos...)
	}

	mapMarketIds := make(map[domain.MarketId]domain.MarketId)
	mapCompanyIds := make(map[domain.CompanyId]domain.CompanyId)
	for _, ipo := range similarIpos {
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

	return &GetSimilarIposResponse{
		ipos:      similarIpos,
		markets:   markets,
		companies: companies,
	}, nil
}

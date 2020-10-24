package application

import (
	"fmt"
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
)

// IpoService is the service to manage the IPOs
type IpoService struct {
	ipoRepository     domain.IpoRepository
	marketRepository  domain.MarketRepository
	companyRepository domain.CompanyRepository
}

// NewService returns the IPoService
func NewService(ipoRepository domain.IpoRepository, marketRepository domain.MarketRepository, companyRepository domain.CompanyRepository) IpoService {
	return IpoService{ipoRepository, marketRepository, companyRepository}
}

// GetIposQuery returns a struct
type GetIposQuery struct {
}

// NewGetIposQuery returns the query used by GetIPOs method
func NewGetIposQuery() GetIposQuery {
	return GetIposQuery{}
}

// GetIposResponse is the response from GetIPOs method
type GetIposResponse struct {
	ipos      []domain.Ipo
	markets   []domain.Market
	companies []domain.Company
}

// Get returns the response data
func (r GetIposResponse) Get() ([]domain.Ipo, []domain.Market, []domain.Company) {
	return r.ipos, r.markets, r.companies
}

// GetIPOS obtains IPOs and related data
func (h IpoService) GetIPOs(query GetIposQuery) (*GetIposResponse, error) {

	ipos, err := h.ipoRepository.Find("", "", "", "", nil, 0, 20)
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
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
		fmt.Println(err)
		return nil, err
	}

	return &GetIposResponse{
		ipos,
		markets,
		companies,
	}, nil
}

// GetIposQuery returns a struct
type GetIpoQuery struct {
	alias string
}

// Alias returns the response data
func (q GetIpoQuery) Alias() string {
	return q.alias
}

// NewGetIpoQuery returns the query used by GetIPO method
func NewGetIpoQuery(alias string) GetIpoQuery {
	return GetIpoQuery{alias}
}

// GetIpoResponse is the response from GetIPO method
type GetIpoResponse struct {
	ipo     *domain.Ipo
	market  *domain.Market
	company *domain.Company
}

// Get returns the response data
func (r GetIpoResponse) Get() (*domain.Ipo, *domain.Market, *domain.Company) {
	return r.ipo, r.market, r.company
}

// GetIPO obtains a IPO and related data
func (h IpoService) GetIPO(query GetIpoQuery) (*GetIpoResponse, error) {
	ipo, err := h.ipoRepository.GetByAlias(query.alias)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if ipo == nil {
		return nil, nil
	}

	market, err := h.marketRepository.GetById(ipo.MarketId())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	company, err := h.companyRepository.GetById(ipo.CompanyId())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &GetIpoResponse{
		ipo,
		market,
		company,
	}, nil
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

// GetIPO obtains a IPO and related data
func (h IpoService) GetSimilarIPOs(query GetSimilarIposQuery) (*GetSimilarIposResponse, error) {
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
	similarIpos, err := h.ipoRepository.Find(marketId, "", sectorId, "", blackList, 0, 5)
	if err != nil {
		return nil, err
	}

	if len(similarIpos) < 5 {
		for _, ipo :=range similarIpos{
			blackList = append(blackList, ipo.Id())
		}
		similarSectorIpos, err := h.ipoRepository.Find("", "", sectorId, "", blackList, 0, 5-len(similarIpos))
		if err != nil {
			return nil, err
		}
		similarIpos = append(similarIpos, similarSectorIpos...)
	}

	if len(similarIpos) < 5 {
		for _, ipo :=range similarIpos{
			blackList = append(blackList, ipo.Id())
		}
		similarMarketIpos, err := h.ipoRepository.Find(marketId, "", "", "", blackList, 0, 5-len(similarIpos))
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
		fmt.Println(err)
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
		fmt.Println(err)
		return nil, err
	}

	return &GetSimilarIposResponse{
		ipos:      similarIpos,
		markets:   markets,
		companies: companies,
	}, nil
}

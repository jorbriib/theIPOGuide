package application

import (
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
)

// Service is the interface where the methods are declared
type Service interface {
	GetIPOs(query GetIposQuery) (*GetIposResponse, error)
}

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

func (r GetIposResponse) Get() ([]domain.Ipo, []domain.Market, []domain.Company) {
	return r.ipos, r.markets, r.companies
}

// GetIPOS obtains IPOs and related data
func (h IpoService) GetIPOs(query GetIposQuery) (*GetIposResponse, error) {
	ipos, err := h.ipoRepository.Find()
	if err != nil {
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
		return nil, err
	}

	return &GetIposResponse{
		ipos,
		markets,
		companies,
	}, nil
}

package application

import (
	"github.com/jorbriib/theIPOGuide/backend/src/domain"
	"log"
)

// IpoService is the service to manage the IPOs
type GetIpoService struct {
	ipoRepository     domain.IpoRepository
	marketRepository  domain.MarketRepository
	companyRepository domain.CompanyRepository
}

// NewService returns the IPoService
func NewGetIpoService(
	ipoRepository domain.IpoRepository,
	marketRepository domain.MarketRepository,
	companyRepository domain.CompanyRepository,
) GetIpoService {
	return GetIpoService{ipoRepository, marketRepository, companyRepository}
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

// Run obtains a IPO and related data
func (h GetIpoService) Run(query GetIpoQuery) (*GetIpoResponse, error) {
	ipo, err := h.ipoRepository.GetByAlias(query.alias)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if ipo == nil {
		return nil, nil
	}

	market, err := h.marketRepository.GetById(ipo.MarketId())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	company, err := h.companyRepository.GetById(ipo.CompanyId())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &GetIpoResponse{
		ipo,
		market,
		company,
	}, nil
}

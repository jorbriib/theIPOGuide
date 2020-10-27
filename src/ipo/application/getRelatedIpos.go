package application

import (
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
	"log"
)

// IpoService is the service to manage the IPOs
type GetRelatedIposService struct {
	marketRepository  domain.MarketRepository
	countryRepository domain.CountryRepository
	sectorRepository  domain.SectorRepository
}

// NewService returns the IPoService
func NewGetRelatedIposService(
	marketRepository domain.MarketRepository,
	countryRepository domain.CountryRepository,
	sectorRepository domain.SectorRepository,
) GetRelatedIposService {
	return GetRelatedIposService{marketRepository, countryRepository, sectorRepository}
}

// GetRelatedIposQuery returns a struct
type GetRelatedIposQuery struct{}

// NewGetRelatedIposQuery returns the query used by GetRelatedIpos method
func NewGetRelatedIposQuery() GetRelatedIposQuery {
	return GetRelatedIposQuery{}
}

// GetRelatedIposResponse is the response from GetRelatedIpos method
type GetRelatedIposResponse struct {
	markets   []domain.Market
	countries []domain.Country
	sectors   []domain.Sector
}

// Get returns the response data
func (r GetRelatedIposResponse) Get() ([]domain.Market, []domain.Country, []domain.Sector) {
	return r.markets, r.countries, r.sectors
}

// Run obtains a IPO and related data
func (h GetRelatedIposService) Run(query GetRelatedIposQuery) (*GetRelatedIposResponse, error) {

	markets, err := h.marketRepository.All()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	countries, err := h.countryRepository.All()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	sectors, err := h.sectorRepository.All()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &GetRelatedIposResponse{
		markets:   markets,
		countries: countries,
		sectors:   sectors,
	}, nil
}

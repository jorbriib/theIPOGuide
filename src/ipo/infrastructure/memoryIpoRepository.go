package infrastructure

import (
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
	"time"
)

type MemoryIpoRepository struct {
	ipos []*domain.Ipo
}

func NewMemoryIpoRepository() MemoryIpoRepository {
	ipos := make([]*domain.Ipo, 2)

	company1 := domain.HydrateCompany("PINS", "Pinterest")
	country := domain.HydrateCountry("US", "USA")
	market := domain.HydrateMarket("NQ", "Nasdaq", country)
	now := time.Now()
	ipo1 := domain.HydrateIpo("1", company1, market, &now)

	company2 := domain.HydrateCompany("NIO", "NIO Cars")
	ipo2 := domain.HydrateIpo("2", company2, market, &now)

	ipos[0] = &ipo1
	ipos[1] = &ipo2

	return MemoryIpoRepository{ipos: ipos}
}

func (r MemoryIpoRepository) Find() []*domain.Ipo {
	return r.ipos
}

package infrastructure

import (
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
)

type MemoryIpoRepository struct {
	ipos []*domain.Ipo
}

func NewMemoryIpoRepository() MemoryIpoRepository {
	ipos := make([]*domain.Ipo, 2)

	company1 := domain.NewCompany("11", "PINS", "Pinterest")
	country := domain.NewCountry("US", "USA")
	market := domain.NewMarket("NQ", "Nasdaq", country)
	ipo1 := domain.NewIpo("1", company1, market, "2020-05-05")

	company2 := domain.NewCompany("12", "NIO", "NIO Cars")
	ipo2 := domain.NewIpo("2", company2, market, "2020-20-05")

	ipos[0] = &ipo1
	ipos[1] = &ipo2

	return MemoryIpoRepository{ipos: ipos}
}

func (r MemoryIpoRepository) Find() []*domain.Ipo {
	return r.ipos
}

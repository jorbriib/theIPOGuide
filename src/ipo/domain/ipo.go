package domain

import "time"

type ID string

type Ipo struct {
	id           ID
	company      Company
	market       Market
	expectedDate *time.Time
}

func HydrateIpo(id ID, company Company, market Market, expectedDate *time.Time) Ipo {
	return Ipo{id: id, company: company, market: market, expectedDate: expectedDate}
}

func (i Ipo) Id() ID {
	return i.id
}

func (i Ipo) Company() Company {
	return i.company
}

func (i Ipo) Market() Market {
	return i.market
}

func (i Ipo) ExpectedDate() *time.Time {
	return i.expectedDate
}

func (i Ipo) ToString() string {
	return i.company.name+" ("+i.company.symbol+") in "+i.market.name
}


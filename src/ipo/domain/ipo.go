package domain

type ID string

type Ipo struct {
	id ID
	company Company
	market Market
	expectedDate string
}
func NewIpo(id ID, company Company, market Market, expectedDate string) Ipo {
	return Ipo{id: id, company: company, market: market, expectedDate: expectedDate}
}

func (i Ipo) Id() ID{
	return i.id
}

func (i Ipo) Company() Company{
	return i.company
}

func (i Ipo) Market() Market{
	return i.market
}

func (i Ipo) ExpectedDate() string{
	return i.expectedDate
}
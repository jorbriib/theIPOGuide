package domain

import "time"

// IpoId represents the IPO Id
type IpoId string

// Ipo represents the IPO entity
type Ipo struct {
	id             IpoId
	alias		   string
	intro		  string
	marketId       MarketId
	companyId      CompanyId
	priceCentsFrom uint32
	priceCentsTo   uint32
	shares         uint32
	expectedDate   *time.Time
}

// HydrateIpo hydrates the IPO struct
func HydrateIpo(
	id IpoId,
	alias string,
	intro string,
	marketId MarketId,
	companyId CompanyId,
	priceCentsFrom uint32,
	priceCentsTo uint32,
	shares uint32,
	expectedDate *time.Time,
) Ipo {
	return Ipo{
		id,
		alias,
		intro,
		marketId,
		companyId,
		priceCentsFrom,
		priceCentsTo,
		shares,
		expectedDate,
	}
}

// IpoId returns the IPO id as string
func (i Ipo) Id() IpoId {
	return i.id
}

// Alias returns the IPO alias as string
func (i Ipo) Alias() string {
	return i.alias
}

// Intro returns the IPO intro as string
func (i Ipo) Intro() string {
	return i.intro
}

// Market returns the IPO market as struct
func (i Ipo) MarketId() MarketId {
	return i.marketId
}

// Company returns the IPO market as struct
func (i Ipo) CompanyId() CompanyId {
	return i.companyId
}

// PriceCentsFrom returns the IPO price from as unsigned integer
func (i Ipo) PriceCentsFrom() uint32 {
	return i.priceCentsFrom
}

// PriceCentsTo returns the IPO price to as unsigned integer
func (i Ipo) PriceCentsTo() uint32 {
	return i.priceCentsTo
}

// Shares returns the IPO shares as unsigned integer
func (i Ipo) Shares() uint32 {
	return i.shares
}

// ExpectedDate returns the IPO expected date as Time
func (i Ipo) ExpectedDate() *time.Time {
	return i.expectedDate
}

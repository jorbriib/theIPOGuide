package domain

import "time"

// ID represents the IPO Id
type ID string

// Ipo represents the IPO entity
type Ipo struct {
	id             ID
	market         Market
	company        Company
	priceCentsFrom uint32
	priceCentsTo   uint32
	shares         uint32
	expectedDate   *time.Time
}

// HydrateIpo hydrates the IPO struct
func HydrateIpo(
	id ID,
	market Market,
	company Company,
	priceCentsFrom uint32,
	priceCentsTo uint32,
	shares uint32,
	expectedDate *time.Time,
) Ipo {
	return Ipo{
		id,
		market,
		company,
		priceCentsFrom,
		priceCentsTo,
		shares,
		expectedDate,
	}
}

// ID returns the IPO id as string
func (i Ipo) Id() ID {
	return i.id
}

// Market returns the IPO market as struct
func (i Ipo) Market() Market {
	return i.market
}

// Company returns the IPO market as struct
func (i Ipo) Company() Company {
	return i.company
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

// ToString returns a string representation of a IPO
func (i Ipo) ToString() string {
	return i.company.name + " (" + i.company.symbol + ") in " + i.market.name
}

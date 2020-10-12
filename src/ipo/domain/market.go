package domain

// MarketId represents the Market Id
type MarketId string

// Market represents the market entity
type Market struct {
	id       MarketId
	code     string
	name     string
	currency Currency
}

// HydrateMarket hydrates the market struct
func HydrateMarket(id MarketId, code string, name string, currency Currency) Market {
	return Market{
		id,
		code,
		name,
		currency,
	}
}

// Id returns the market id as string
func (m Market) Id() MarketId {
	return m.id
}

// Code returns the market code as string
func (m Market) Code() string {
	return m.code
}

// Name returns the market name as string
func (m Market) Name() string {
	return m.name
}

/// Currency returns the market currency struct
func (m Market) Currency() Currency {
	return m.currency
}

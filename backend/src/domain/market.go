package domain

// MarketId represents the Market Id
type MarketId string

// Market represents the market entity
type Market struct {
	id        MarketId
	code      string
	name      string
	currency  Currency
	image     string
	totalIpos int
}

// HydrateMarket hydrates the market struct
func HydrateMarket(id MarketId, code string, name string, currency Currency, image string, totalIpos int) Market {
	return Market{
		id,
		code,
		name,
		currency,
		image,
		totalIpos,
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

// Image returns the market image as string
func (m Market) Image() string {
	return m.image
}

// TotalIpos returns the number of ipos related to a market
func (m Market) TotalIpos() int {
	return m.totalIpos
}
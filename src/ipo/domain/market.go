package domain

// Market represents the market entity
type Market struct {
	code    string
	name    string
	currency Currency
}

// HydrateMarket hydrates the market struct
func HydrateMarket(code string, name string, currency Currency) Market {
	return Market{
		code:    code,
		name:    name,
		currency: currency,
	}
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
func (m Market) Currency() Currency{
	return m.currency
}
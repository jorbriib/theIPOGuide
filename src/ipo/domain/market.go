package domain

type Market struct {
	code    string
	name    string
	country Country
}

func HydrateMarket(code string, name string, country Country) Market {
	return Market{
		code:    code,
		name:    name,
		country: country,
	}
}

func (m Market) Code() string {
	return m.code
}

func (m Market) Name() string {
	return m.name
}

func (m Market) Country() Country{
	return m.country
}
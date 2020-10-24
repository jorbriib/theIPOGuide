package domain

// CountryId represents the Country Id
type CountryId string

// Country represents the company's country
type Country struct {
	id   CountryId
	code string
	name string
}

// HydrateCountry hydrates the country struct
func HydrateCountry(id CountryId, code string, name string) Country {
	return Country{id: id, code: code, name: name}
}

// Id returns the country code as string
func (c Country) Id() CountryId {
	return c.id
}

// Code returns the country code as string
func (c Country) Code() string {
	return c.code
}

// Name returns the country name as string
func (c Country) Name() string {
	return c.name
}

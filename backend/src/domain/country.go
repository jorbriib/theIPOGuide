package domain

// CountryId represents the Country Id
type CountryId string

// Country represents the company's country
type Country struct {
	id        CountryId
	code      string
	name      string
	image     string
	totalIpos int
}

// HydrateCountry hydrates the country struct
func HydrateCountry(id CountryId, code string, name string, image string, totalIpos int) Country {
	return Country{id: id, code: code, name: name, image: image, totalIpos: totalIpos}
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

// Image returns the country image as string
func (c Country) Image() string {
	return c.image
}

// TotalIpos returns the number of ipos related to a country
func (c Country) TotalIpos() int {
	return c.totalIpos
}

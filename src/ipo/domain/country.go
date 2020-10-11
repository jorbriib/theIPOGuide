package domain

// Country represents the company's country
type Country struct{
	code string
	name string
}

// HydrateCountry hydrates the country struct
func HydrateCountry(code string, name string) Country {
	return Country{code: code, name: name}
}

// Code returns the country code as string
func (c Country) Code() string {
	return c.code
}

// Name returns the country name as string
func (c Country) Name() string {
	return c.name
}
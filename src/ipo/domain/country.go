package domain

type Country struct{
	code string
	name string
}

func NewCountry(code string, name string) Country {
	return Country{code: code, name: name}
}

func (c Country) Code() string {
	return c.code
}

func (c Country) Name() string {
	return c.name
}
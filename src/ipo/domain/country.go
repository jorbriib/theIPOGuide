package domain

type Country struct{
	code string
	name string
}

func (c Country) Code() string {
	return c.code
}

func (c Country) Name() string {
	return c.name
}
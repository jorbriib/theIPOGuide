package domain

type Market struct{
	code string
	name string
	country Country
}

func (m Market) Code() string {
	return m.code
}

func (m Market) Name() string {
	return m.name
}
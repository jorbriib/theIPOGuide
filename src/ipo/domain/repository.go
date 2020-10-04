package domain

type IpoRepository interface {
	find() []*Ipo
}
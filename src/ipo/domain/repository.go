package domain

type IpoRepository interface {
	Find() ([]Ipo, error)
}

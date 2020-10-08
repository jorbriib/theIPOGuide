package application

import (
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
)

type Service interface {
	GetIPOs(query GetIposQuery) ([]domain.Ipo, error)
}

type GetIposQuery struct {
}

func NewGetIposQuery() GetIposQuery {
	return GetIposQuery{}
}

type IpoService struct {
	ipoRepository domain.IpoRepository
}

func NewService(ipoRepository domain.IpoRepository) IpoService {
	return IpoService{ipoRepository: ipoRepository}
}

func (h IpoService) GetIPOs(query GetIposQuery) ([]domain.Ipo, error) {
	ipos, err := h.ipoRepository.Find()
	if err != nil {
		return []domain.Ipo{}, err
	}

	return ipos, nil
}

package application

import (
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
)

type Service struct {
	ipoRepository domain.IpoRepository
}

func NewService(ipoRepository domain.IpoRepository) Service {
	return Service{ipoRepository: ipoRepository}
}

type GetIposQuery struct {
}

func NewGetIposQuery() GetIposQuery {
	return GetIposQuery{}
}

type GetIposResponse struct {
	ipos []string
}

func (h Service) GetIPOs(query GetIposQuery) (GetIposResponse, error) {
	ipos, err := h.ipoRepository.Find()
	if err != nil {
		return GetIposResponse{}, err
	}

	iposName := make([]string, len(ipos))
	for _, ipo := range ipos {
		iposName = append(iposName, ipo.ToString())
	}

	return GetIposResponse{ipos: iposName}, nil
}

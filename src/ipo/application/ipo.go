package application

import (
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
)

type Service interface {
	GetIPOs(query GetIposQuery) (GetIposResponse, error)
}

type GetIposQuery struct {
}

func NewGetIposQuery() GetIposQuery {
	return GetIposQuery{}
}

type GetIposResponse struct {
	ipos []string
}

func (r GetIposResponse) GetIpos() []string{
	return r.ipos
}

type IpoService struct {
	ipoRepository domain.IpoRepository
}

func NewService(ipoRepository domain.IpoRepository) IpoService {
	return IpoService{ipoRepository: ipoRepository}
}

func (h IpoService) GetIPOs(query GetIposQuery) (GetIposResponse, error) {
	ipos, err := h.ipoRepository.Find()
	if err != nil {
		return GetIposResponse{}, err
	}

	var iposName = make([]string, len(ipos))
	for k, ipo := range ipos {
		iposName[k] = ipo.ToString()
	}

	return GetIposResponse{ipos: iposName}, nil
}

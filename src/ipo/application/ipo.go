package application

import (
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
)

type Handler struct {
	ipoRepository domain.IpoRepository
}

func NewHandler(ipoRepository domain.IpoRepository) Handler {
	return Handler{ipoRepository: ipoRepository}
}

type getIposQuery struct {
}

func NewGetIposQuery() getIposQuery {
	return getIposQuery{}
}

type GetIposResponse struct {
	ipos []string
}

func (h Handler) GetIPOs(query getIposQuery) (GetIposResponse, error) {
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

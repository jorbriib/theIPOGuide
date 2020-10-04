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

func (h Handler) GetIPOs(query getIposQuery) GetIposResponse {
	ipos := h.ipoRepository.Find()
	iposName := make([]string, len(ipos))
	for _, ipo := range ipos {
		iposName = append(iposName, ipo.Company().Name())
	}

	return GetIposResponse{ipos: iposName}
}

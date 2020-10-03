package api

import (
	"fmt"
	"github.com/golossus/routing"
	"net/http"
)

func AddRoutes(router *routing.Router) {
	controller := ipoController{}
	_ = router.Get("/", controller.GetIPOs)
}

type ipoController struct {

}

func (c ipoController) GetIPOs(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprint(writer, "Welcome to the API")
}
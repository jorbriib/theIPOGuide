package main

import (
	"fmt"
	"github.com/golossus/routing"
	"github.com/jorbriib/theIPOGuide/src/ipo/application"
	"github.com/jorbriib/theIPOGuide/src/ipo/infrastructure"
	ipo_public_api "github.com/jorbriib/theIPOGuide/src/ipo/ui/public/api"
	"net/http"
)

func main() {
	r := routing.NewRouter()

	repository := infrastructure.NewMemoryIpoRepository()
	handler := application.NewHandler(repository)
	controller := ipo_public_api.NewController(handler)
	_ = r.Get("/", controller.GetIpos)

	fmt.Println("Server listening")
	err := http.ListenAndServe(":80", &r)
	if err != nil{
		panic(err)
	}
}

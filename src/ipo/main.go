package main

import (
	"fmt"
	"github.com/golossus/routing"
	ipo_public_api "github.com/jorbriib/theIPOGuide/src/ipo/ui/public/api"
	"net/http"
)

func main() {
	r := routing.NewRouter()

	aipo_public_apipi.AddRoutes(&r)


	fmt.Println("Server listening")

	_ = http.ListenAndServe(":80", &r)

}

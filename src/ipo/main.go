package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golossus/routing"
	"github.com/jorbriib/theIPOGuide/src/ipo/application"
	"github.com/jorbriib/theIPOGuide/src/ipo/infrastructure"
	"github.com/jorbriib/theIPOGuide/src/ipo/ui/public/api"
	"log"
	"net/http"
	"os"
)

func main() {
	r := routing.NewRouter()

	db := getDB()

	ipoRepository := infrastructure.NewMySQLIpoRepository(db)
	marketRepository := infrastructure.NewMySQLMarketRepository(db)
	companyRepository := infrastructure.NewMySQLCompanyRepository(db)
	service := application.NewService(ipoRepository, marketRepository, companyRepository)
	controller := api.NewController(service)

	_ = r.Get("/v1/ipos", api.ContentTypeMiddleware(api.EnableCorsMiddleware(controller.GetIpos)))
	_ = r.Get("/v1/ipos/{alias}", api.ContentTypeMiddleware(api.EnableCorsMiddleware(controller.GetIpo)))

	_ = r.Get("/v1/{notFound}", api.ContentTypeMiddleware(api.EnableCorsMiddleware(notFound)))
	_ = r.Get("/{notFound}", api.ContentTypeMiddleware(api.EnableCorsMiddleware(notFound)))

	fmt.Println("Server listening")
	errServer := http.ListenAndServe(":80", &r)
	if errServer != nil {
		log.Fatal(errServer)
	}
}

func notFound(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusNotFound)
}


func getDB() *sql.DB {
	mysqlAddr := os.Getenv("MYSQL_ADDR")
	mysqlDBName := os.Getenv("MYSQL_DATABASE")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")

	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s", mysqlUser, mysqlPassword, mysqlAddr, mysqlDBName)

	db, errorDB := sql.Open("mysql", conn)
	if errorDB != nil {
		log.Fatal(errorDB)
	}
	return db
}

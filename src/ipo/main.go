package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golossus/routing"
	"github.com/jorbriib/theIPOGuide/src/ipo/application"
	"github.com/jorbriib/theIPOGuide/src/ipo/infrastructure"
	ipo_public_api "github.com/jorbriib/theIPOGuide/src/ipo/ui/public/api"
	"log"
	"net/http"
	"os"
)

func main() {
	r := routing.NewRouter()

	db := getDB()

	repository := infrastructure.NewMySQLIpoRepository(db)
	service := application.NewService(repository)
	controller := ipo_public_api.NewController(service)
	_ = r.Get("/v1/ipos", controller.GetIpos)

	fmt.Println("Server listening")
	errServer := http.ListenAndServe(":80", &r)
	if errServer != nil {
		log.Fatal(errServer)
	}
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

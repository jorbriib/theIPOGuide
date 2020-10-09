package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jorbriib/theIPOGuide/src/ipo/application"
	"github.com/jorbriib/theIPOGuide/src/ipo/infrastructure"
	ipo_public_api "github.com/jorbriib/theIPOGuide/src/ipo/ui/public/api"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetIpos(t *testing.T) {
	assertion := assert.New(t)
	db := getDbTest()

	repository := infrastructure.NewMySQLIpoRepository(db)
	service := application.NewService(repository)
	controller := ipo_public_api.NewController(service)

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	controller.GetIpos(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assertion.Equal(http.StatusOK, resp.StatusCode)
	assertion.Equal("application/json; charset=UTF-8", resp.Header.Get("Content-Type"))
	assertion.JSONEq("[{\"companyName\":\"Pinterest\",\"marketName\":\"Nasdaq\"}]", string(body))

}



func getDbTest() *sql.DB {
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
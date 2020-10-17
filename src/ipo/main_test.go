package main_test

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jorbriib/theIPOGuide/src/ipo/application"
	"github.com/jorbriib/theIPOGuide/src/ipo/infrastructure"
	"github.com/jorbriib/theIPOGuide/src/ipo/ui/public/api"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var db *sql.DB

func TestMain(m *testing.M) {
	mysqlAddr := os.Getenv("MYSQL_ADDR")
	mysqlDBName := os.Getenv("MYSQL_DATABASE")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")

	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s", mysqlUser, mysqlPassword, mysqlAddr, mysqlDBName)

	var errorDB error
	db, errorDB = sql.Open("mysql", conn)
	if errorDB != nil {
		log.Fatal(errorDB)
	}

	code := m.Run()
	os.Exit(code)
}

func TestGetIpos(t *testing.T) {
	assertion := assert.New(t)

	ipoRepository := infrastructure.NewMySQLIpoRepository(db)
	marketRepository := infrastructure.NewMySQLMarketRepository(db)
	companyRepository := infrastructure.NewMySQLCompanyRepository(db)
	service := application.NewService(ipoRepository, marketRepository, companyRepository)
	controller := api.NewController(service)

	r := httptest.NewRequest("GET", "/v1/ipos", nil)
	w := httptest.NewRecorder()
 

	controller.GetIpos(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assertion.Equal(http.StatusOK, resp.StatusCode)
	assertion.JSONEq("[{\"alias\": \"pinterest\",\"company\": {\"symbol\": \"PINS\",\"name\": \"Pinterest\",\"sector\": \"Communication Services\",\"country\": \"USA\",\"logo\": \"/assets/images/pinterest-logo.jpg\"},\"market\": {\"name\": \"Nasdaq\"},\"priceFrom\": \"$24.5\",\"priceTo\": \"$25.8\",\"expectedDate\": \"2020-10-10 00:00:00 +0000 UTC\"}]", string(body))

}

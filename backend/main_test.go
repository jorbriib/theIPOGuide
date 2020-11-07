package main_test

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jorbriib/theIPOGuide/backend/src/application"
	"github.com/jorbriib/theIPOGuide/backend/src/infrastructure"
	"github.com/jorbriib/theIPOGuide/backend/src/ui/public/api"
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
	countryRepository := infrastructure.NewMySQLCountryRepository(db)
	sectorRepository := infrastructure.NewMySQLSectorRepository(db)
	service := application.NewGetIposService(ipoRepository, marketRepository, companyRepository, countryRepository, sectorRepository)
	controller := api.NewGetIposController(service)

	r := httptest.NewRequest("GET", "/v1/ipos", nil)
	w := httptest.NewRecorder()
 

	controller.Run(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assertion.Equal(http.StatusOK, resp.StatusCode)
	assertion.JSONEq("{\"total\":2,\"list\":[{\"alias\":\"array-technologies\",\"company\":{\"symbol\":\"ARRY\",\"name\":\"Array Technologies\",\"sector\":\"Technology\",\"country\":\"United States of America\",\"logo\":\"/assets/images/array-technologies-logo.jpg\"},\"market\":{\"code\":\"NQGB\",\"name\":\"Nasdaq Global\",\"currency\":\"USD\"},\"priceFrom\":\"$19\",\"priceTo\":\"\",\"expectedDate\":\"2020-10-15 00:00:00 +0000 UTC\"},{\"alias\":\"pinterest\",\"company\":{\"symbol\":\"PINS\",\"name\":\"Pinterest\",\"sector\":\"Communication Services\",\"country\":\"United States of America\",\"logo\":\"/assets/images/pinterest-logo.jpg\"},\"market\":{\"code\":\"NQGB\",\"name\":\"Nasdaq Global\",\"currency\":\"USD\"},\"priceFrom\":\"$22\",\"priceTo\":\"\",\"expectedDate\":\"2019-04-18 00:00:00 +0000 UTC\"}]}", string(body))

}

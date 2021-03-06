package infrastructure_test

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jorbriib/theIPOGuide/backend/src/domain"
	"github.com/jorbriib/theIPOGuide/backend/src/infrastructure"
	"github.com/stretchr/testify/assert"
	"log"
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

func TestNewMySQLCompanyRepository(t *testing.T) {
	r := infrastructure.NewMySQLCompanyRepository(db)
	assert.NotNil(t, r)
}

func TestMySQLCompanyRepository_FindByIds_ReturnsSliceLength0_WhenNoCompanyIds(t *testing.T) {
	r := infrastructure.NewMySQLCompanyRepository(db)

	var ids []domain.CompanyId
	response, err := r.FindByIds(ids)

	assert.Nil(t, err)
	assert.Equal(t, 0, len(response))
}

func TestMySQLCompanyRepository_FindByIds(t *testing.T) {
	r := infrastructure.NewMySQLCompanyRepository(db)

	ids := []domain.CompanyId{
		domain.CompanyId("4293f9f9-c2b7-1e7b-8271-77a4ce70c6f0"),
	}
	response, err := r.FindByIds(ids)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(response))
	assert.Equal(t, ids[0], response[0].Id())
}

func TestMySQLCompanyRepository_GetById(t *testing.T) {
	r := infrastructure.NewMySQLCompanyRepository(db)

	id := domain.CompanyId("4293f9f9-c2b7-1e7b-8271-77a4ce70c6f0")

	response, err := r.GetById(id)

	assert.Nil(t, err)
	assert.Equal(t, id, response.Id())
}

func TestMySQLCompanyRepository_GetById_ReturnsNilWhenNotFound(t *testing.T) {
	r := infrastructure.NewMySQLCompanyRepository(db)

	id := domain.CompanyId("1293f9f9-c2b7-1e7b-8271-77a4ce70c6f0")

	response, err := r.GetById(id)

	assert.Nil(t, err)
	assert.Nil(t, response)
}

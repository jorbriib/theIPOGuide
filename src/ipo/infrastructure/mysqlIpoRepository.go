package infrastructure

import (
	"database/sql"
	"fmt"
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
	"time"
)

type ipoSQL struct {
	Id            string     `db:"id"`
	CompanySymbol string     `db:"companySymbol"`
	CompanyName   string     `db:"companyName"`
	MarketCode    string     `db:"marketCode"`
	MarketName    string     `db:"marketName"`
	CountryCode   string     `db:"countryCode"`
	CountryName   string     `db:"countryName"`
	ExpectedDate  *time.Time `db:"expectedDate"`
}

type MySQLIpoRepository struct {
	table string
	db    *sql.DB
}

func NewMySQLIpoRepository(db *sql.DB) MySQLIpoRepository {
	return MySQLIpoRepository{table: "ipos", db: db}
}

func (r MySQLIpoRepository) Find() ([]domain.Ipo, error) {

	query := `
    SELECT i.uuid as id, 
           c.symbol as companySymbol, c.name as companyName, 
           m.code as marketCode, m.name as marketName, 
           ct.code as countryCode, ct.name as countryName,
           i.expected_date as expectedDate
	FROM ipos i
	INNER JOIN companies c ON c.uuid = i.company_id
	INNER JOIN markets m ON m.uuid = i.market_id
	INNER JOIN countries ct ON ct.uuid = m.country_id
  `

	rows, err := r.db.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var result []domain.Ipo
	for rows.Next() {
		ipoSql := &ipoSQL{}

		_ = rows.Scan(
			&ipoSql.Id,
			&ipoSql.CompanySymbol,
			&ipoSql.CompanyName,
			&ipoSql.MarketCode,
			&ipoSql.MarketName,
			&ipoSql.CountryCode,
			&ipoSql.CountryName,
			&ipoSql.ExpectedDate,
		)

		company := domain.HydrateCompany(ipoSql.CompanySymbol, ipoSql.CompanyName)
		country := domain.HydrateCountry(ipoSql.CountryCode, ipoSql.CountryName)
		market := domain.HydrateMarket(ipoSql.MarketCode, ipoSql.MarketName, country)
		ipo := domain.HydrateIpo(domain.ID(ipoSql.Id), company, market, ipoSql.ExpectedDate)

		result = append(result, ipo)
	}


	return result, nil
}

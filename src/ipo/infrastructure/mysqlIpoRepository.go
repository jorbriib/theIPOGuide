package infrastructure

import (
	"database/sql"
	"fmt"
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
	"time"
)

type ipoSQL struct {
	Id             string `db:"id"`
	Alias          string `db:"alias"`
	MarketId       string `db:"marketId"`
	CompanyId      string `db:"companyId"`
	PriceCentsFrom uint32 `db:"priceCentsFrom"`
	PriceCentsTo   uint32 `db:"priceCentsTo"`
	Shares         uint32 `db:"shares"`
	ExpectedDate   string `db:"expectedDate"`
}

type MySQLIpoRepository struct {
	table string
	db    *sql.DB
}

func NewMySQLIpoRepository(db *sql.DB) MySQLIpoRepository {
	return MySQLIpoRepository{table: "ipos", db: db}
}

func (r MySQLIpoRepository) GetByAlias(alias string) (*domain.Ipo, error) {

	query := `
    SELECT BIN_TO_UUID(i.uuid) AS id, i.alias as alias, BIN_TO_UUID(i.market_id) AS marketId, BIN_TO_UUID(i.company_id) AS companyId, 
           i.price_cents_from AS priceCentsFrom, i.price_cents_to AS priceCentsTo, 
           i.shares as shares, i.expected_date as expectedDate
	FROM ipos i
	WHERE alias=?
  `

	row := r.db.QueryRow(query, alias)

	ipoSql := &ipoSQL{}
	_ = row.Scan(
		&ipoSql.Id,
		&ipoSql.Alias,
		&ipoSql.MarketId,
		&ipoSql.CompanyId,
		&ipoSql.PriceCentsFrom,
		&ipoSql.PriceCentsTo,
		&ipoSql.Shares,
		&ipoSql.ExpectedDate,
	)

	layout := "2006-01-02"
	timeExpectedDate, err := time.Parse(layout, ipoSql.ExpectedDate)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	ipo := domain.HydrateIpo(
		domain.IpoId(ipoSql.Id),
		ipoSql.Alias,
		domain.MarketId(ipoSql.MarketId),
		domain.CompanyId(ipoSql.CompanyId),
		ipoSql.PriceCentsFrom,
		ipoSql.PriceCentsTo,
		ipoSql.Shares,
		&timeExpectedDate,
	)

	return &ipo, nil
}

func (r MySQLIpoRepository) Find() ([]domain.Ipo, error) {

	query := `
    SELECT BIN_TO_UUID(i.uuid) AS id, i.alias as alias, BIN_TO_UUID(i.market_id) AS marketId, BIN_TO_UUID(i.company_id) AS companyId, 
           i.price_cents_from AS priceCentsFrom, i.price_cents_to AS priceCentsTo, 
           i.shares as shares, i.expected_date as expectedDate
	FROM ipos i
  `

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var result []domain.Ipo
	for rows.Next() {
		ipoSql := &ipoSQL{}

		_ = rows.Scan(
			&ipoSql.Id,
			&ipoSql.Alias,
			&ipoSql.MarketId,
			&ipoSql.CompanyId,
			&ipoSql.PriceCentsFrom,
			&ipoSql.PriceCentsTo,
			&ipoSql.Shares,
			&ipoSql.ExpectedDate,
		)

		layout := "2006-01-02"
		timeExpectedDate, err := time.Parse(layout, ipoSql.ExpectedDate)
		if err != nil {
			fmt.Println(err)
			continue
		}
		ipo := domain.HydrateIpo(
			domain.IpoId(ipoSql.Id),
			ipoSql.Alias,
			domain.MarketId(ipoSql.MarketId),
			domain.CompanyId(ipoSql.CompanyId),
			ipoSql.PriceCentsFrom,
			ipoSql.PriceCentsTo,
			ipoSql.Shares,
			&timeExpectedDate,
		)

		result = append(result, ipo)
	}

	return result, nil
}

package infrastructure

import (
	"database/sql"
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
	"time"
)

type ipoSQL struct {
	Id             string     `db:"id"`
	MarketId       string     `db:"marketId"`
	CompanyId      string     `db:"companyId"`
	PriceCentsFrom uint32     `db:"priceCentsFrom"`
	PriceCentsTo   uint32     `db:"priceCentsTo"`
	Shares         uint32     `db:"shares"`
	ExpectedDate   *time.Time `db:"expectedDate"`
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
    SELECT BIN_TO_UUID(i.uuid) AS id, BIN_TO_UUID(i.market_id) AS marketId, BIN_TO_UUID(i.company_id) AS companyId, 
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
			&ipoSql.MarketId,
			&ipoSql.CompanyId,
			&ipoSql.PriceCentsFrom,
			&ipoSql.PriceCentsTo,
			&ipoSql.Shares,
			&ipoSql.ExpectedDate,
		)

		ipo := domain.HydrateIpo(
			domain.IpoId(ipoSql.Id),
			domain.MarketId(ipoSql.MarketId),
			domain.CompanyId(ipoSql.CompanyId),
			ipoSql.PriceCentsFrom,
			ipoSql.PriceCentsTo,
			ipoSql.Shares,
			ipoSql.ExpectedDate,
		)

		result = append(result, ipo)
	}

	return result, nil
}

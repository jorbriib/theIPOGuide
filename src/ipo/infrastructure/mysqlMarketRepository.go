package infrastructure

import (
	"bytes"
	"database/sql"
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
	"strings"
)

type marketSQL struct {
	Id              string `db:"id"`
	Code            string `db:"code"`
	Name            string `db:"name"`
	CurrencyCode    string `db:"currencyCode"`
	CurrencyName    string `db:"currencyName"`
	CurrencyDisplay string `db:"currencyDisplay"`
}

type MySQLMarketRepository struct {
	table string
	db    *sql.DB
}

func NewMySQLMarketRepository(db *sql.DB) MySQLMarketRepository {
	return MySQLMarketRepository{table: "markets", db: db}
}

func (r MySQLMarketRepository) FindByIds(ids []domain.MarketId) ([]domain.Market, error) {

	if len(ids) == 0 {
		return []domain.Market{}, nil
	}

	inQuery := make([]string, len(ids))

	var uuidToBinQuery bytes.Buffer
	for k, id := range ids {
		uuidToBinQuery.WriteString("UUID_TO_BIN('")
		uuidToBinQuery.WriteString(string(id))
		uuidToBinQuery.WriteString("')")
		inQuery[k] = uuidToBinQuery.String()

		uuidToBinQuery.Reset()
	}

	query := `
    SELECT BIN_TO_UUID(m.uuid) AS id, m.code as code, m.name as name,
	c.code as currencyCode, c.name as currencyName, c.display as currencyDisplay
	FROM markets m
	INNER JOIN currencies c ON c.uuid = m.currency_id
	WHERE m.uuid IN (` + strings.Join(inQuery, ",") + `)
  `

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var result []domain.Market
	for rows.Next() {
		marketSql := &marketSQL{}

		_ = rows.Scan(
			&marketSql.Id,
			&marketSql.Code,
			&marketSql.Name,
			&marketSql.CurrencyCode,
			&marketSql.CurrencyName,
			&marketSql.CurrencyDisplay,
		)

		currency := domain.HydrateCurrency(marketSql.CurrencyCode, marketSql.CurrencyName, marketSql.CurrencyDisplay)
		market := domain.HydrateMarket(
			domain.MarketId(marketSql.Id),
			marketSql.Code,
			marketSql.Name,
			currency,
		)

		result = append(result, market)
	}

	return result, nil
}

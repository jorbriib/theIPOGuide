package infrastructure

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
	"strconv"
	"strings"
	"time"
)

type ipoSQL struct {
	Id             string        `db:"id"`
	Alias          string        `db:"alias"`
	MarketId       string        `db:"marketId"`
	CompanyId      string        `db:"companyId"`
	PriceCentsFrom uint32        `db:"priceCentsFrom"`
	PriceCentsTo   sql.NullInt32 `db:"priceCentsTo"`
	Shares         uint32        `db:"shares"`
	ExpectedDate   string        `db:"expectedDate"`
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
	err := row.Scan(
		&ipoSql.Id,
		&ipoSql.Alias,
		&ipoSql.MarketId,
		&ipoSql.CompanyId,
		&ipoSql.PriceCentsFrom,
		&ipoSql.PriceCentsTo,
		&ipoSql.Shares,
		&ipoSql.ExpectedDate,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	layout := "2006-01-02"
	timeExpectedDate, err := time.Parse(layout, ipoSql.ExpectedDate)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var priceTo uint32
	if ipoSql.PriceCentsTo.Valid {
		priceTo = uint32(ipoSql.PriceCentsTo.Int32)
	}
	ipo := domain.HydrateIpo(
		domain.IpoId(ipoSql.Id),
		ipoSql.Alias,
		domain.MarketId(ipoSql.MarketId),
		domain.CompanyId(ipoSql.CompanyId),
		ipoSql.PriceCentsFrom,
		priceTo,
		ipoSql.Shares,
		&timeExpectedDate,
	)

	return &ipo, nil
}

func (r MySQLIpoRepository) Find(
	marketId domain.MarketId,
	countryId domain.CountryId,
	sectorId domain.SectorId,
	industryId domain.IndustryId,
	blackList []domain.IpoId,
	offset int,
	limit int,
) ([]domain.Ipo, error) {

	query := `
    SELECT BIN_TO_UUID(i.uuid) AS id, i.alias as alias, BIN_TO_UUID(i.market_id) AS marketId, BIN_TO_UUID(i.company_id) AS companyId, 
           i.price_cents_from AS priceCentsFrom, i.price_cents_to AS priceCentsTo, 
           i.shares as shares, i.expected_date as expectedDate
	FROM ipos i  
`
	if countryId != "" || sectorId != "" || industryId != "" {
		query = query + " INNER JOIN companies c ON c.uuid = i.company_id "
		if countryId != "" {
			query = query + " AND c.country_id = UUID_TO_BIN('" + string(countryId) + "') "
		}
		if sectorId != "" {
			query = query + " AND c.sector_id = UUID_TO_BIN('" + string(sectorId) + "') "
		}
		if industryId != "" {
			query = query + " AND c.industry_id = UUID_TO_BIN('" + string(industryId) + "') "
		}
	}

	if marketId != "" || len(blackList) > 0 {
		query = query + " WHERE 1 = 1 "
		if marketId != "" {
			query = query + " AND i.market_id = UUID_TO_BIN('" + string(marketId) + "') "
		}
		if len(blackList) > 0 {
			var uuidToBinQuery bytes.Buffer
			inQuery := make([]string, len(blackList))
			for k, id := range blackList {

				uuidToBinQuery.WriteString("UUID_TO_BIN('")
				uuidToBinQuery.WriteString(string(id))
				uuidToBinQuery.WriteString("')")
				inQuery[k] = uuidToBinQuery.String()
				uuidToBinQuery.Reset()
			}
			query = query + " AND i.uuid NOT IN ("+strings.Join(inQuery, ",")+") "
		}
	}

	query = query + `ORDER BY i.expected_date DESC LIMIT ` + strconv.Itoa(offset) + `, ` + strconv.Itoa(limit)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var result []domain.Ipo
	for rows.Next() {
		ipoSql := &ipoSQL{}

		err = rows.Scan(
			&ipoSql.Id,
			&ipoSql.Alias,
			&ipoSql.MarketId,
			&ipoSql.CompanyId,
			&ipoSql.PriceCentsFrom,
			&ipoSql.PriceCentsTo,
			&ipoSql.Shares,
			&ipoSql.ExpectedDate,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}

		layout := "2006-01-02"
		timeExpectedDate, err := time.Parse(layout, ipoSql.ExpectedDate)
		if err != nil {
			fmt.Println(err)
			continue
		}

		var priceTo uint32
		if ipoSql.PriceCentsTo.Valid {
			priceTo = uint32(ipoSql.PriceCentsTo.Int32)
		}
		ipo := domain.HydrateIpo(
			domain.IpoId(ipoSql.Id),
			ipoSql.Alias,
			domain.MarketId(ipoSql.MarketId),
			domain.CompanyId(ipoSql.CompanyId),
			ipoSql.PriceCentsFrom,
			priceTo,
			ipoSql.Shares,
			&timeExpectedDate,
		)

		result = append(result, ipo)
	}

	return result, nil
}

package infrastructure

import (
	"bytes"
	"database/sql"
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
	"log"
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
	ipo, err := r.buildIpo(ipoSql)

	return ipo, nil
}

func (r MySQLIpoRepository) Find(
	marketIds []domain.MarketId,
	countryIds []domain.CountryId,
	sectorIds []domain.SectorId,
	industryIds []domain.IndustryId,
	blackList []domain.IpoId,
	offset uint,
	limit uint,
) ([]domain.Ipo, error) {

	query := `
    SELECT BIN_TO_UUID(i.uuid) AS id, i.alias as alias, BIN_TO_UUID(i.market_id) AS marketId, BIN_TO_UUID(i.company_id) AS companyId, 
           i.price_cents_from AS priceCentsFrom, i.price_cents_to AS priceCentsTo, 
           i.shares as shares, i.expected_date as expectedDate
	FROM ipos i  
`
	query = r.prepareQuery(marketIds, countryIds, sectorIds, industryIds, blackList, offset, limit, query)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var result []domain.Ipo
	for rows.Next() {
		ipoSql := &ipoSQL{}
		err := rows.Scan(
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
			log.Println(err)
			return nil, err
		}
		ipo, err := r.buildIpo(ipoSql)
		if err != nil {
			continue
		}
		result = append(result, *ipo)
	}

	return result, nil
}


func (r MySQLIpoRepository) Count(
	marketIds []domain.MarketId,
	countryIds []domain.CountryId,
	sectorIds []domain.SectorId,
	industryIds []domain.IndustryId,
	blackList []domain.IpoId,
) (uint, error) {
	var count uint

	query := `
    SELECT COUNT(*) 
	FROM ipos i  
`
	query = r.prepareQuery(marketIds, countryIds, sectorIds, industryIds, blackList, 0, 0, query)
	row := r.db.QueryRow(query)

	err := row.Scan(&count)
	if err == sql.ErrNoRows {
		return count, nil
	}
	if err != nil {
		log.Println(err)
		return count, err
	}

	return count, nil
}

func (r MySQLIpoRepository) prepareQuery(marketIds []domain.MarketId, countryIds []domain.CountryId, sectorIds []domain.SectorId, industryIds []domain.IndustryId, blackList []domain.IpoId, offset uint, limit uint, query string) string {
	if len(countryIds) != 0 || len(sectorIds) != 0 || len(industryIds) != 0 {
		query = query + " INNER JOIN companies c ON c.uuid = i.company_id "

		if len(countryIds) > 0 {
			var uuidToBinQuery bytes.Buffer
			inQuery := make([]string, len(countryIds))
			for k, id := range countryIds {

				uuidToBinQuery.WriteString("UUID_TO_BIN('")
				uuidToBinQuery.WriteString(string(id))
				uuidToBinQuery.WriteString("')")
				inQuery[k] = uuidToBinQuery.String()
				uuidToBinQuery.Reset()
			}
			query = query + " AND c.country_id IN (" + strings.Join(inQuery, ",") + ") "
		}
		if len(sectorIds) > 0 {
			var uuidToBinQuery bytes.Buffer
			inQuery := make([]string, len(sectorIds))
			for k, id := range sectorIds {

				uuidToBinQuery.WriteString("UUID_TO_BIN('")
				uuidToBinQuery.WriteString(string(id))
				uuidToBinQuery.WriteString("')")
				inQuery[k] = uuidToBinQuery.String()
				uuidToBinQuery.Reset()
			}
			query = query + " AND c.sector_id IN (" + strings.Join(inQuery, ",") + ") "
		}
		if len(industryIds) > 0 {
			var uuidToBinQuery bytes.Buffer
			inQuery := make([]string, len(industryIds))
			for k, id := range industryIds {

				uuidToBinQuery.WriteString("UUID_TO_BIN('")
				uuidToBinQuery.WriteString(string(id))
				uuidToBinQuery.WriteString("')")
				inQuery[k] = uuidToBinQuery.String()
				uuidToBinQuery.Reset()
			}
			query = query + " AND c.industry_id IN (" + strings.Join(inQuery, ",") + ") "
		}
	}

	if len(marketIds) != 0 || len(blackList) > 0 {
		query = query + " WHERE 1 = 1 "
		if len(marketIds) > 0 {
			var uuidToBinQuery bytes.Buffer
			inQuery := make([]string, len(marketIds))
			for k, id := range marketIds {

				uuidToBinQuery.WriteString("UUID_TO_BIN('")
				uuidToBinQuery.WriteString(string(id))
				uuidToBinQuery.WriteString("')")
				inQuery[k] = uuidToBinQuery.String()
				uuidToBinQuery.Reset()
			}
			query = query + " AND i.market_id IN (" + strings.Join(inQuery, ",") + ") "
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
			query = query + " AND i.uuid NOT IN (" + strings.Join(inQuery, ",") + ") "
		}
	}

	query = query + ` ORDER BY i.expected_date DESC `
	if limit > 0 {
		query = query + ` LIMIT ` + strconv.Itoa(int(offset)) + `, ` + strconv.Itoa(int(limit))
	}
	return query
}

func (r MySQLIpoRepository) buildIpo(ipoSql *ipoSQL) (*domain.Ipo, error){


	layout := "2006-01-02"
	timeExpectedDate, err := time.Parse(layout, ipoSql.ExpectedDate)
	if err != nil {
		log.Println(err)
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
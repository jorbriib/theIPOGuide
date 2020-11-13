package infrastructure

import (
	"database/sql"
	"github.com/jorbriib/theIPOGuide/backend/src/domain"
	"strings"
)

type countrySQL struct {
	Id        string `db:"id"`
	Code      string `db:"code"`
	Name      string `db:"name"`
	ImageUrl  string `db:"imageUrl"`
	TotalIpos int    `db:"totalIpos"`
}

type MySQLCountryRepository struct {
	table string
	db    *sql.DB
}

func NewMySQLCountryRepository(db *sql.DB) MySQLCountryRepository {
	return MySQLCountryRepository{table: "countries", db: db}
}

func (r MySQLCountryRepository) All() ([]domain.Country, error) {
	query := `
    SELECT BIN_TO_UUID(c.uuid) AS id, c.code as code, c.name as name, 
	c.image_url as imageUrl, c.total_ipos as totalIpos
	FROM countries c`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	result := r.resultFromRows(rows)
	return result, nil
}

func (r MySQLCountryRepository) FindByCodes(codes []string) ([]domain.Country, error) {
	args := make([]interface{}, len(codes))
	for i, id := range codes {
		args[i] = id
	}

	query := `
    SELECT BIN_TO_UUID(c.uuid) AS id, c.code as code, c.name as name, 
	c.image_url as imageUrl, c.total_ipos as totalIpos
	FROM countries c
	WHERE c.code IN (?` + strings.Repeat(",?", len(args)-1) + `)`

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	result := r.resultFromRows(rows)
	return result, nil
}

func (r MySQLCountryRepository) resultFromRows(rows *sql.Rows) []domain.Country {
	var result []domain.Country
	for rows.Next() {
		countrySQL := &countrySQL{}

		_ = rows.Scan(
			&countrySQL.Id,
			&countrySQL.Code,
			&countrySQL.Name,
			&countrySQL.ImageUrl,
			&countrySQL.TotalIpos,
		)

		country := domain.HydrateCountry(
			domain.CountryId(countrySQL.Id),
			countrySQL.Code,
			countrySQL.Name,
			countrySQL.ImageUrl,
			countrySQL.TotalIpos,
		)

		result = append(result, country)
	}
	return result
}

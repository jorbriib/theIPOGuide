package infrastructure

import (
	"database/sql"
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
	"strings"
)

type countrySQL struct {
	Id   string `db:"id"`
	Code string `db:"code"`
	Name string `db:"name"`
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
    SELECT BIN_TO_UUID(c.uuid) AS id, c.code as code, c.name as name 
	FROM countries c`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var result []domain.Country
	for rows.Next() {
		countrySQL := &countrySQL{}

		_ = rows.Scan(
			&countrySQL.Id,
			&countrySQL.Code,
			&countrySQL.Name,
		)

		country := domain.HydrateCountry(
			domain.CountryId(countrySQL.Id),
			countrySQL.Code,
			countrySQL.Name,
		)

		result = append(result, country)
	}

	return result, nil
}

func (r MySQLCountryRepository) FindByCodes(codes []string) ([]domain.Country, error) {
	args := make([]interface{}, len(codes))
	for i, id := range codes {
		args[i] = id
	}

	query := `
    SELECT BIN_TO_UUID(c.uuid) AS id, c.code as code, c.name as name 
	FROM countries c
	WHERE c.code IN (?`+strings.Repeat(",?", len(args)-1)+`)`

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var result []domain.Country
	for rows.Next() {
		countrySQL := &countrySQL{}

		_ = rows.Scan(
			&countrySQL.Id,
			&countrySQL.Code,
			&countrySQL.Name,
		)

		country := domain.HydrateCountry(
			domain.CountryId(countrySQL.Id),
			countrySQL.Code,
			countrySQL.Name,
		)

		result = append(result, country)
	}

	return result, nil
}
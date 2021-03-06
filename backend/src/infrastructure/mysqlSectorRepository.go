package infrastructure

import (
	"database/sql"
	"github.com/jorbriib/theIPOGuide/backend/src/domain"
	"strings"
)

type sectorSQL struct {
	Id   string `db:"id"`
	Alias string `db:"alias"`
	Name string `db:"name"`
}

type MySQLSectorRepository struct {
	table string
	db    *sql.DB
}

func NewMySQLSectorRepository(db *sql.DB) MySQLSectorRepository {
	return MySQLSectorRepository{table: "countries", db: db}
}

func (r MySQLSectorRepository) All() ([]domain.Sector, error) {
	query := `
    SELECT BIN_TO_UUID(s.uuid) AS id, s.alias as alias, s.name as name 
	FROM sectors s`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var result []domain.Sector
	for rows.Next() {
		sectorSQL := &sectorSQL{}

		_ = rows.Scan(
			&sectorSQL.Id,
			&sectorSQL.Alias,
			&sectorSQL.Name,
		)

		sector := domain.HydrateSector(
			domain.SectorId(sectorSQL.Id),
			sectorSQL.Alias,
			sectorSQL.Name,
		)

		result = append(result, sector)
	}

	return result, nil
}


func (r MySQLSectorRepository) FindByAliases(aliases []string) ([]domain.Sector, error) {
	args := make([]interface{}, len(aliases))
	for i, id := range aliases {
		args[i] = id
	}

	query := `
    SELECT BIN_TO_UUID(s.uuid) AS id, s.alias as alias, s.name as name 
	FROM sectors s
	WHERE s.alias IN (?`+strings.Repeat(",?", len(args)-1)+`)`

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var result []domain.Sector
	for rows.Next() {
		sectorSQL := &sectorSQL{}

		_ = rows.Scan(
			&sectorSQL.Id,
			&sectorSQL.Alias,
			&sectorSQL.Name,
		)

		sector := domain.HydrateSector(
			domain.SectorId(sectorSQL.Id),
			sectorSQL.Alias,
			sectorSQL.Name,
		)

		result = append(result, sector)
	}

	return result, nil
}

package infrastructure

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/jorbriib/theIPOGuide/src/ipo/domain"
	"strings"
)

type companySQL struct {
	Id                    string `db:"id"`
	Symbol                string `db:"symbol"`
	Name                  string `db:"name"`
	SectorName            string `db:"sectorName"`
	Address               string `db:"address"`
	CountryCode           string `db:"countryCode"`
	CountryName           string `db:"countryName"`
	Phone                 string `db:"phone"`
	Email                 sql.NullString `db:"email"`
	Website               string `db:"website"`
	Employees             uint32 `db:"employees"`
	Description           string `db:"description"`
	Founded               uint16 `db:"founded"`
	Ceo                   string `db:"ceo"`
	FiscalYearEnd         string `db:"fiscalYearEnd"`
	IpoUrl                string `db:"ipoUrl"`
	ExchangeCommissionUrl string `db:"exchangeCommissionUrl"`
	LogoUrl               string `db:"logoUrl"`
}

// MySQLCompanyRepository is the repository to manage companies
type MySQLCompanyRepository struct {
	table string
	db    *sql.DB
}

// NewMySQLCompanyRepository returns the repository
func NewMySQLCompanyRepository(db *sql.DB) MySQLCompanyRepository {
	return MySQLCompanyRepository{table: "companies", db: db}
}

// GetById returns a company by id
func (r MySQLCompanyRepository) GetById(id domain.CompanyId) (*domain.Company, error) {
	ipos, err := r.FindByIds([]domain.CompanyId{id})
	if err != nil{
		return nil, err
	}
	if len(ipos) < 1{
		return nil, nil
	}
	return &ipos[0], nil
}


// FindByIds returns the companies filtering by companyIds
func (r MySQLCompanyRepository) FindByIds(ids []domain.CompanyId) ([]domain.Company, error) {

	if len(ids) == 0 {
		return []domain.Company{}, nil
	}

	var uuidToBinQuery bytes.Buffer
	inQuery := make([]string, len(ids))
	for k, id := range ids {

		uuidToBinQuery.WriteString("UUID_TO_BIN('")
		uuidToBinQuery.WriteString(string(id))
		uuidToBinQuery.WriteString("')")
		inQuery[k] = uuidToBinQuery.String()
		uuidToBinQuery.Reset()
	}

	query := `
    SELECT BIN_TO_UUID(c.uuid) AS id, c.symbol as symbol, c.name as name,
	s.name as sectorName, c.address as address, ct.code as countryCode, ct.name as countryName,
	c.phone as phone, c.email as email, c.website as website, c.employees as employees,
	c.description as description, c.founded as founded, c.ceo as ceo, c.fiscal_year_end as fiscalYearEnd,
	c.ipo_url as ipoUrl, c.exchange_commission_url as exchangeCommissionUrl, c.logo_url as logoUrl
	FROM companies c
	INNER JOIN countries ct ON ct.uuid = c.country_id
	INNER JOIN sectors s ON s.uuid = c.sector_id
	WHERE c.uuid IN (` + strings.Join(inQuery, ",") + `)
  `

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var result []domain.Company
	for rows.Next() {
		companySql := &companySQL{}

		err = rows.Scan(
			&companySql.Id,
			&companySql.Symbol,
			&companySql.Name,
			&companySql.SectorName,
			&companySql.Address,
			&companySql.CountryCode,
			&companySql.CountryName,
			&companySql.Phone,
			&companySql.Email,
			&companySql.Website,
			&companySql.Employees,
			&companySql.Description,
			&companySql.Founded,
			&companySql.Ceo,
			&companySql.FiscalYearEnd,
			&companySql.IpoUrl,
			&companySql.ExchangeCommissionUrl,
			&companySql.LogoUrl,
		)
		if err != nil{
			fmt.Println(err)
			continue
		}

		sector := domain.HydrateSector(companySql.SectorName)
		country := domain.HydrateCountry(companySql.CountryCode, companySql.CountryName)

		var email string
		if companySql.Email.Valid {
			email = companySql.Email.String
		}
		company := domain.HydrateCompany(
			domain.CompanyId(companySql.Id),
			companySql.Symbol,
			companySql.Name,
			sector,
			companySql.Address,
			country,
			companySql.Phone,
			email,
			companySql.Website,
			companySql.Employees,
			companySql.Description,
			companySql.Founded,
			companySql.Ceo,
			companySql.FiscalYearEnd,
			companySql.IpoUrl,
			companySql.ExchangeCommissionUrl,
			companySql.LogoUrl,
		)

		result = append(result, company)
	}

	return result, nil
}


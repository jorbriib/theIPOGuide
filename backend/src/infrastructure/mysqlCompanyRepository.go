package infrastructure

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/jorbriib/theIPOGuide/backend/src/domain"
	"strings"
)

type companySQL struct {
	Id                    string         `db:"id"`
	Symbol                string         `db:"symbol"`
	Name                  string         `db:"name"`
	SectorId              string         `db:"sectorId"`
	SectorAlias           string         `db:"sectorAlias"`
	SectorName            string         `db:"sectorName"`
	IndustryId            string         `db:"industryId"`
	IndustryAlias         string         `db:"industryAlias"`
	IndustryName          string         `db:"industryName"`
	Address               string         `db:"address"`
	CountryId             string         `db:"countryId"`
	CountryCode           string         `db:"countryCode"`
	CountryName           string         `db:"countryName"`
	Phone                 string         `db:"phone"`
	Email                 sql.NullString `db:"email"`
	Website               string         `db:"website"`
	Facebook              sql.NullString `db:"facebook"`
	Twitter               sql.NullString `db:"twitter"`
	Linkedin              sql.NullString `db:"linkedin"`
	Pinterest             sql.NullString `db:"pinterest"`
	Instagram             sql.NullString `db:"instagram"`
	Employees             uint32         `db:"employees"`
	Description           string         `db:"description"`
	Founded               uint16         `db:"founded"`
	Ceo                   string         `db:"ceo"`
	FiscalYearEnd         string         `db:"fiscalYearEnd"`
	IpoUrl                string         `db:"ipoUrl"`
	ExchangeCommissionUrl string         `db:"exchangeCommissionUrl"`
	LogoUrl               string         `db:"logoUrl"`
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
	if err != nil {
		return nil, err
	}
	if len(ipos) < 1 {
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
	BIN_TO_UUID(s.uuid) as sectorId, s.alias as sectorAlias,  s.name as sectorName, 
	BIN_TO_UUID(i.uuid) as industryId,  i.alias as industryAlias, i.name as industryName,
	c.address as address, BIN_TO_UUID(ct.uuid) as countryId, ct.code as countryCode, ct.name as countryName,
	c.phone as phone, c.email as email, c.website as website, c.employees as employees,
	c.facebook as facebook, c.twitter as twitter, c.linkedin as linkedin, c.pinterest as pinterest, c.instagram as instagram,
	c.description as description, c.founded as founded, c.ceo as ceo, c.fiscal_year_end as fiscalYearEnd,
	c.ipo_url as ipoUrl, c.exchange_commission_url as exchangeCommissionUrl, c.logo_url as logoUrl
	FROM companies c
	INNER JOIN countries ct ON ct.uuid = c.country_id
	INNER JOIN sectors s ON s.uuid = c.sector_id
	INNER JOIN industries i ON i.uuid = c.industry_id
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
			&companySql.SectorId,
			&companySql.SectorAlias,
			&companySql.SectorName,
			&companySql.IndustryId,
			&companySql.IndustryAlias,
			&companySql.IndustryName,
			&companySql.Address,
			&companySql.CountryId,
			&companySql.CountryCode,
			&companySql.CountryName,
			&companySql.Phone,
			&companySql.Email,
			&companySql.Website,
			&companySql.Employees,
			&companySql.Facebook,
			&companySql.Twitter,
			&companySql.Linkedin,
			&companySql.Pinterest,
			&companySql.Instagram,
			&companySql.Description,
			&companySql.Founded,
			&companySql.Ceo,
			&companySql.FiscalYearEnd,
			&companySql.IpoUrl,
			&companySql.ExchangeCommissionUrl,
			&companySql.LogoUrl,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}

		sector := domain.HydrateSector(domain.SectorId(companySql.SectorId), companySql.SectorAlias, companySql.SectorName)
		industry := domain.HydrateIndustry(domain.IndustryId(companySql.IndustryId), companySql.IndustryAlias, companySql.IndustryName)
		country := domain.HydrateCountry(domain.CountryId(companySql.CountryId), companySql.CountryCode, companySql.CountryName)

		var email string
		if companySql.Email.Valid {
			email = companySql.Email.String
		}

		var facebook string
		if companySql.Facebook.Valid {
			facebook = companySql.Facebook.String
		}
		var twitter string
		if companySql.Twitter.Valid {
			twitter = companySql.Twitter.String
		}
		var linkedin string
		if companySql.Linkedin.Valid {
			linkedin = companySql.Linkedin.String
		}
		var pinterest string
		if companySql.Pinterest.Valid {
			pinterest = companySql.Pinterest.String
		}
		var instagram string
		if companySql.Instagram.Valid {
			instagram = companySql.Instagram.String
		}

		company := domain.HydrateCompany(
			domain.CompanyId(companySql.Id),
			companySql.Symbol,
			companySql.Name,
			sector,
			industry,
			companySql.Address,
			country,
			companySql.Phone,
			email,
			companySql.Website,
			companySql.Employees,
			companySql.Description,
			facebook,
			twitter,
			linkedin,
			pinterest,
			instagram,
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

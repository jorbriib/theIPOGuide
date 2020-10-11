package domain

// Company represents the IPO's company
type Company struct {
	symbol               string
	name                 string
	sector               Sector
	address              string
	country              Country
	phone                string
	email                string
	website              string
	employees            int
	description          string
	founded              int
	ceo                  string
	fiscalYearEnd        string
	ipoUrl               string
	exchangeComissionUrl string
}

// HydrateCompany hydrates the company struct
func HydrateCompany(
	symbol string,
	name string,
	sector Sector,
	address string,
	country Country,
	phone string,
	email string,
	website string,
	employees int,
	description string,
	founded int,
	ceo string,
	fiscalYearEnd string,
	ipoUrl string,
	exchangeCommissionUrl string,
) Company {
	return Company{
		symbol,
		name,
		sector,
		address,
		country,
		phone,
		email,
		website,
		employees,
		description,
		founded,
		ceo,
		fiscalYearEnd,
		ipoUrl,
		exchangeCommissionUrl,
	}
}

// Symbol returns the company symbol string
func (c Company) Symbol() string {
	return c.symbol
}

// Name returns the company name as string
func (c Company) Name() string {
	return c.name
}

// Sector returns the company sector as sector struct
func (c Company) Sector() Sector {
	return c.sector
}

// Address returns the company address as string
func (c Company) Address() string {
	return c.address
}

// Country returns the company country as country struct
func (c Company) Country() Country {
	return c.country
}
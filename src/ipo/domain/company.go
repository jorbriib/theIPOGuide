package domain

// CompanyId represents the Company Id
type CompanyId string

// Company represents the IPO's company
type Company struct {
	id                   CompanyId
	symbol               string
	name                 string
	sector               Sector
	address              string
	country              Country
	phone                string
	email                string
	website              string
	employees            uint32
	description          string
	founded              uint16
	ceo                  string
	fiscalYearEnd        string
	ipoUrl               string
	exchangeComissionUrl string
	logoUrl              string
}

// HydrateCompany hydrates the company struct
func HydrateCompany(
	id CompanyId,
	symbol string,
	name string,
	sector Sector,
	address string,
	country Country,
	phone string,
	email string,
	website string,
	employees uint32,
	description string,
	founded uint16,
	ceo string,
	fiscalYearEnd string,
	ipoUrl string,
	exchangeCommissionUrl string,
	logoUrl string,
) Company {
	return Company{
		id,
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
		logoUrl,
	}
}

// Id returns the company id as string
func (c Company) Id() CompanyId {
	return c.id
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


// LogoUrl returns the company country as country struct
func (c Company) LogoUrl() string {
	return c.logoUrl
}
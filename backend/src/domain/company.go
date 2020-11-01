package domain

// CompanyId represents the Company Id
type CompanyId string

// Company represents the IPO's company
type Company struct {
	id                   CompanyId
	symbol               string
	name                 string
	sector               Sector
	industry             Industry
	address              string
	country              Country
	phone                string
	email                string
	website              string
	employees            uint32
	description           string
	facebook              string
	twitter               string
	linkedin              string
	pinterest             string
	instagram             string
	founded               uint16
	ceo                   string
	fiscalYearEnd         string
	ipoUrl                string
	exchangeCommissionUrl string
	logoUrl               string
}

// HydrateCompany hydrates the company struct
func HydrateCompany(
	id CompanyId,
	symbol string,
	name string,
	sector Sector,
	industry Industry,
	address string,
	country Country,
	phone string,
	email string,
	website string,
	employees uint32,
	description string,
	facebook string,
	twitter string,
	linkedin string,
	pinterest string,
	instagram string,
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
			industry,
		address,
		country,
		phone,
		email,
		website,
		employees,
		description,
		facebook,
		twitter,
		linkedin,
		pinterest,
		instagram,
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

// Industry returns the company sector industry sector struct
func (c Company) Industry() Industry {
	return c.industry
}

// Address returns the company address as string
func (c Company) Address() string {
	return c.address
}

// Country returns the company country as country struct
func (c Company) Country() Country {
	return c.country
}

// LogoUrl returns the company logoUrl
func (c Company) LogoUrl() string {
	return c.logoUrl
}

// Facebook returns the company facebook
func (c Company) Facebook() string {
	return c.facebook
}

// Twitter returns the company twitter
func (c Company) Twitter() string {
	return c.twitter
}

// Linkedin returns the company linkedin
func (c Company) Linkedin() string {
	return c.linkedin
}

// Pinterest returns the company pinterest
func (c Company) Pinterest() string {
	return c.pinterest
}

// Instagram returns the company instagram
func (c Company) Instagram() string {
	return c.instagram
}

// Phone returns the company phone
func (c Company) Phone() string {
	return c.phone
}

// Email returns the company email
func (c Company) Email() string {
	return c.email
}

// Website returns the company website
func (c Company) Website() string {
	return c.website
}

// Description returns the company description
func (c Company) Description() string {
	return c.description
}

// Description returns the number of employees
func (c Company) Employees() uint32 {
	return c.employees
}

// Founded returns when the company was founded
func (c Company) Founded() uint16 {
	return c.founded
}

// Ceo returns the company ceo
func (c Company) Ceo() string {
	return c.ceo
}

// FiscalYearEnd returns when the ends the fiscal year
func (c Company) FiscalYearEnd() string {
	return c.fiscalYearEnd
}

// IpoUrl returns the IPO url
func (c Company) IpoUrl() string {
	return c.ipoUrl
}

// ExchangeCommissionUrl returns the exchange commisssion url
func (c Company) ExchangeCommissionUrl() string {
	return c.exchangeCommissionUrl
}
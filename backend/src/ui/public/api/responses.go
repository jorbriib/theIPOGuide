package api

type GetIposJsonResponse struct {
	Total uint                  `json:"total"`
	List  []IpoListJsonResponse `json:"list"`
}

type SearchByTextJsonResponse struct {
	List []IpoListJsonResponse `json:"list"`
}
type IpoListJsonResponse struct {
	Alias        string                   `json:"alias"`
	Company      *CompanyListJsonResponse `json:"company"`
	Market       *MarketJsonResponse      `json:"market"`
	PriceFrom    string                   `json:"priceFrom"`
	PriceTo      string                   `json:"priceTo"`
	ExpectedDate string                   `json:"expectedDate"`
}

type CompanyListJsonResponse struct {
	Symbol  string `json:"symbol"`
	Name    string `json:"name"`
	Sector  string `json:"sector"`
	Country string `json:"country"`
	Logo    string `json:"logo"`
}

type MarketJsonResponse struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Currency string `json:"currency"`
}

type ExtendedMarketJsonResponse struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	Currency  string `json:"currency"`
	Image     string `json:"image"`
	TotalIpos int    `json:"totalIpos"`
}

type CountryJsonResponse struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type ExtendedCountryJsonResponse struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	Image     string `json:"image"`
	TotalIpos int    `json:"totalIpos"`
}

type SectorJsonResponse struct {
	Alias string `json:"alias"`
	Name  string `json:"name"`
}

type ExtendedSectorJsonResponse struct {
	Alias     string `json:"alias"`
	Name      string `json:"name"`
	Image     string `json:"image"`
	TotalIpos int    `json:"totalIpos"`
}

type IpoViewJsonResponse struct {
	Alias        string                   `json:"alias"`
	Intro        string                   `json:"intro"`
	Company      *CompanyViewJsonResponse `json:"company"`
	Market       *MarketJsonResponse      `json:"market"`
	PriceFrom    string                   `json:"priceFrom"`
	PriceTo      string                   `json:"priceTo"`
	Shares       uint32                   `json:"shares"`
	ExpectedDate string                   `json:"expectedDate"`
}

type CompanyViewJsonResponse struct {
	Symbol                string `json:"symbol"`
	Name                  string `json:"name"`
	Sector                string `json:"sector"`
	Industry              string `json:"industry"`
	Address               string `json:"address"`
	Country               string `json:"country"`
	Phone                 string `json:"phone"`
	Email                 string `json:"email"`
	Website               string `json:"website"`
	Description           string `json:"description"`
	Facebook              string `json:"facebook"`
	Twitter               string `json:"twitter"`
	Linkedin              string `json:"linkedin"`
	Pinterest             string `json:"pinterest"`
	Instagram             string `json:"instagram"`
	Employees             uint32 `json:"employees"`
	Founded               uint16 `json:"founded"`
	Ceo                   string `json:"ceo"`
	FiscalYearEnd         string `json:"fiscalYearEnd"`
	IpoUrl                string `json:"ipoUrl"`
	ExchangeCommissionUrl string `json:"exchangeCommissionUrl"`
	Logo                  string `json:"logo"`
}

type IpoRelations struct {
	Markets   []*ExtendedMarketJsonResponse  `json:"markets"`
	Countries []*ExtendedCountryJsonResponse `json:"countries"`
	Sectors   []*ExtendedSectorJsonResponse  `json:"sectors"`
}

// ErrorMessage is the error message to send to the clients
type ErrorMessage struct {
	Message string `json:"message"`
}

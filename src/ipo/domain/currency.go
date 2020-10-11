package domain

import (
	"fmt"
)

type Currency struct {
	code    string
	name    string
	display string
}

// HydrateCurrency hydrates struct
func HydrateCurrency(code string, name string, display string) Currency {
	return Currency{code: code, name: name, display: display}
}

// Code returns the currency code as string
func (c Currency) Code() string {
	return c.code
}

// Name returns the currency name as string
func (c Currency) Name() string {
	return c.name
}

// DisplayFromCents returns the price transformed in unit values
func (c Currency) DisplayFromCents(cents uint32) string {
	return fmt.Sprintf(c.display, fmt.Sprint(float64(cents)/100.0))
}

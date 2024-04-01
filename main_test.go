package main

import (
	"testing"
)

func createTestCountries() []Country {
    return []Country{
        {
            Name: struct{
                Common string `json:"common"`
            }{
                Common: "France",
            },
            Currencies: map[string]struct{
                Name   string `json:"name"`
                Symbol string `json:"symbol"`
            }{
                "EUR": {Name: "Euro", Symbol: "€"},
            },
        },
        {
            Name: struct{
                Common string `json:"common"`
            }{
                Common: "Germany",
            },
            Currencies: map[string]struct{
                Name   string `json:"name"`
                Symbol string `json:"symbol"`
            }{
                "EUR": {Name: "Euro", Symbol: "€"},
            },
        },
        {
            Name: struct{
                Common string `json:"common"`
            }{
                Common: "Australia",
            },
            Currencies: map[string]struct{
                Name   string `json:"name"`
                Symbol string `json:"symbol"`
            }{
                "AUD": {Name: "Australian Dollar", Symbol: "$"},
            },
        },
    }
}

// Test sorting by country name in ascending order
func TestSortCountriesByNameAscending(t *testing.T) {
	countries := createTestCountries()
	sortCountries(countries, "name", true)
	if countries[0].Name.Common != "Australia" || countries[1].Name.Common != "France" || countries[2].Name.Common != "Germany" {
		t.Errorf("Countries are not sorted by name in ascending order")
	}
}

// Test sorting by country name in descending order
func TestSortCountriesByNameDescending(t *testing.T) {
	countries := createTestCountries()
	sortCountries(countries, "name", false)
	if countries[0].Name.Common != "Germany" || countries[1].Name.Common != "France" || countries[2].Name.Common != "Australia" {
		t.Errorf("Countries are not sorted by name in descending order")
	}
}

// Test sorting by currency in ascending order
func TestSortCountriesByCurrencyAscending(t *testing.T) {
	countries := createTestCountries()
	sortCountries(countries, "currency", true)
	if countries[0].Currencies["AUD"].Name != "Australian Dollar" || countries[1].Currencies["EUR"].Name != "Euro" {
		t.Errorf("Countries are not sorted by currency in ascending order")
	}
}


package utilities

import "time"

// Represents a university
type Uni struct {
	Country   string            `json:"country"`
	IsoCode   string            `json:"alpha_two_code"`
	Name      string            `json:"name"`
	WebPages  []string          `json:"web_pages"`
	Languages map[string]string `json:"languages,omitempty"`
	// TODO: Change to a struct and get desired map from it instead
	Map string `json:"maps,omitempty"`
}

// Extra country information for university
type MissingFieldsFromCountry struct {
	Languages map[string]string
	Maps      map[string]string
}

// Border codes for neighboring countries
type BorderCountries struct {
	BorderingCodes []string `json:"borders,omitempty"`
}

// The different names for a country
type CountryNames struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
	// AltSpellings []string `json:"altSpellings"`
}

// Diagnostics for current api and foreign apis.
type DiagInfo struct {
	UniApiStatus     string        `json:"universitiesapi"`
	CountryApiStatus string        `json:"countriesapi"`
	Version          string        `json:"version"`
	Uptime           time.Duration `json:"uptime"`
}

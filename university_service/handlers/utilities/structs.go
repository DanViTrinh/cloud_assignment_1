package utilities

import "time"

// Represents a university
type Uni struct {
	Country   string            `json:"country"`
	IsoCode   string            `json:"alpha_two_code"`
	Name      string            `json:"name"`
	WebPages  []string          `json:"web_pages"`
	Languages map[string]string `json:"languages,omitempty"`
	Map       string            `json:"maps,omitempty"`
}

// The fields used for country
type Country struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
	// AltSpellings []string `json:"altSpellings"`
	Languages map[string]string
	Maps      struct {
		Map string `json:"openStreetMaps"`
	}
	BorderingCodes []string `json:"borders,omitempty"`
	Cca2           string   `json:"cca2"`
}

// Diagnostics for current api and foreign apis.
type DiagInfo struct {
	UniApiStatus     string        `json:"universitiesapi"`
	CountryApiStatus string        `json:"countriesapi"`
	Version          string        `json:"version"`
	Uptime           time.Duration `json:"uptime"`
}

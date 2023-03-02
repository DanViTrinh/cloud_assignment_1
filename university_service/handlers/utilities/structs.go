package utilities

import "time"

// TODO: the name in the example and the name in this struct is not the same
type University struct {
	Country   string            `json:"country"`
	IsoCode   string            `json:"alpha_two_code"`
	Name      string            `json:"name"`
	WebPages  []string          `json:"web_pages"`
	Languages map[string]string `json:"languages,omitempty"`
	Map       string            `json:"maps,omitempty"`
}

type MissingFieldsFromCountry struct {
	Languages map[string]string
	Maps      map[string]string
}

type BorderingCountries struct {
	BorderingCodes []string `json:"borders,omitempty"`
}

type CountryNames struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
	// AltSpellings []string `json:"altSpellings"`
}

//diag
type DiagInfo struct {
	UniApiStatus     string        `json:"universitiesapi"`
	CountryApiStatus string        `json:"countriesapi"`
	Version          string        `json:"version"`
	Uptime           time.Duration `json:"uptime"`
}

package utilities

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

/*
TODO: figure out:
not sure if subregion or bordering countries should be used for
neighbouring countries.
Japan for instance has no bordering countries
*/
type SubRegion struct {
	SubRegion string `json:"subregion"`
}

type BorderingCountries struct {
	BorderingCodes []string `json:"borders,omitempty"`
}

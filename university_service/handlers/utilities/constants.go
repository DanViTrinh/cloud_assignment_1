package utilities

const Version = "v1"

// URL paths
const DefaultPath = "/"
const UniInfoPath = "/unisearcher/v1/uniinfo"
const NeighbourUnisPath = "/unisearcher/v1/neighbourunis"
const DiagPath = "/unisearcher/v1/diag"

// universities api

// real
// const UniversitiesAPIurl = "http://universities.hipolabs.com"

// mock
const UniversitiesAPIurl = "http://localhost:8082"

const UniversitiesSearch = "/search"

// countries api

// real
// const CountriesAPIurl = "https://restcountries.com/v3.1"

// mock
const CountriesAPIurl = "http://localhost:8081/v3.1"

const CountriesName = "/name"
const CountriesAlphaCode = "/alpha"
const CountriesSubregion = "/subregion"
const DesiredMap string = "openStreetMaps"

const StandardInternalServerErrorMsg = "internal server error, refer to logs"
const NotImplementedMsg = "method not yet supported"

const NanoSecondsInAsecond = 1000000000

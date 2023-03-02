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
const UniAPI = "http://localhost:8082"

const UniSearch = "/search"

// countries api

// real
// const CountriesAPIurl = "https://restcountries.com/v3.1"

// mock
const CountryAPI = "http://localhost:8081/v3.1"

const CountryName = "/name"
const CountryCode = "/alpha"
const CountrySubregion = "/subregion"
const DesiredMap string = "openStreetMaps"

const InternalErrMsg = "internal server error, refer to logs"
const NotImplementedMsg = "method not yet supported"

const NanoSecondsInAsecond = 1000000000

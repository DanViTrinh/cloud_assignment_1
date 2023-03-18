package utilities

const Version = "v1"

// URL paths
const DefaultPath = "/"
const UniInfoPath = "/unisearcher/v1/uniinfo"
const NeighborUnisPath = "/unisearcher/v1/neighbourunis"
const DiagPath = "/unisearcher/v1/diag"

// universities api

// real
const UniAPI = "http://universities.hipolabs.com"

// mock
// const UniAPI = "http://localhost:8082"

const UniSearch = "/search"
const UniNameParam = "name_contains"

// countries api

// NTNU country api
const CountryAPI = "http://129.241.150.113:8080/v3.1"

// real
// const CountryAPI = "https://restcountries.com/v3.1"

// mock
// const CountryAPI = "http://localhost:8081/v3.1"

const CountryName = "/name"
const CountryCode = "/alpha"
const CountrySubregion = "/subregion"
const DesiredMap = "openStreetMaps"
const TestCountryCode = "/NO"

const InternalErrMsg = "internal server error, refer to logs"
const NotImplementedMsg = "method not yet supported"
const UnmarshalErrMsg = "error during marshalling data"
const ResponseErrMsg = "error during writing response"
const OutputErrMsg = "error when returning output"

// amount of nano seconds in a second
const NanoSecInSec = 1000000000

# University info service
The service provides three different endpoints
```
.../unisearcher/v1/uniinfo/
.../unisearcher/v1/neighbourunis/
.../unisearcher/v1/diag/
```
The service works by uses two foreign api's:
```
Uni api: http://universities.hipolabs.com/
Country api: https://restcountries.com/ 
```

## .../unisearcher/v1/uniinfo/
This endpoint retrieves university information from the foreign api's. 
The service then uses the retrieved information to combine the two api's.
The endpoint only accepts GET requests. Any other request will be received with
501 Not implemented status code. If no university was found it will display an 
empty array.

Using the endpoint:
```
Method: GET
Path: uniinfo/{:partial_or_complete_university_name}/
``` 
Example request:
```
uniinfo/norwegian%20university%20of%20science%20and%20technology/
```
Response: 
```json
[
  {
      "country": "Norway",
      "alpha_two_code": "NO",
      "name": "Norwegian University of Science and Technology", 
      "web_pages": ["http://www.ntnu.no/"],
      "languages": {"nno": "Norwegian Nynorsk",
                    "nob": "Norwegian Bokmål",
                    "smi": "Sami"},
      "map": "https://www.openstreetmap.org/relation/2978650"
  },
  ... 
]
```

### How it works
The endpoint calls the uni api first to retrieve the initial university name. 
It then calls the country api to fill in the rest of the information. There
is of course a local country cache that caches country that has already been 
found.

## .../unisearcher/v1/neighbourunis/
This endpoint retrieves university information from the neighbor countries of 
the given country. The university from the country itself will not be shown. 

Using the endpoint:
```
Method: GET
Path: neighbourunis/{:country_name}/{:partial_or_complete_university_name}{?limit={:number}}
``` 
Example request:
```
unisearcher/v1/neighbourunis/sweden/science?limit=2
```
Response: 
```json
[
    {
        "country": "Norway",
        "alpha_two_code": "NO",
        "name": "Norwegian University of Life Sciences",
        "web_pages": [
            "http://www.umb.no/"
        ],
        "languages": {
            "nno": "Norwegian Nynorsk",
            "nob": "Norwegian Bokmål",
            "smi": "Sami"
        },
        "map": "https://www.openstreetmap.org/relation/2978650"
    },
    {
        "country": "Finland",
        "alpha_two_code": "FI",
        "name": "Häme University of Applied Sciences",
        "web_pages": [
            "https://www.hamk.fi/"
        ],
        "languages": {
            "fin": "Finnish",
            "swe": "Swedish"
        },
        "map": "openstreetmap.org/relation/54224"
    }
]
```
The limit parameter is an optional parameter that can be omitted.
The limit parameter limits the output to the value given in the limit.

### How it works
1. Calls the university api to get all the universities with the given name 
2. Calls the country api to get the neighboring countries. 
3. Calls the neighbor api to get the alpha two code of each country. 
4. Filters after the alpha two codes. 
5. Display the output. 

Why i have decided to solve it like this is further explained under reflection.

## .../unisearcher/v1/diag/
This endpoint's purpose is to get a diagnostics of the server. It shows the 
uptime, version and status of the foreign api's.

Example response:
```json
{
    "universitiesapi": "200 OK",
    "countriesapi": "200 OK",
    "version": "v1",
    "uptime": 488
}
```

## Package structure
```bash
. 
├───cmd
├───handlers
└───utilities
```
cmd is where the main is located.

handlers are the different handlers for the endpoints.

Utilities is the folder for code that is shared over the entire projects. 
Things like structs, const and other general utilities.

##  Reflection
### Error handling 
The error handling inn this service is inspired by:
> https://medium.com/@ozdemir.zynl/rest-api-error-handling-in-go-behavioral-type-assertion-509d93636afd

One of the principles from this blog is to handle errors only once. 
Which is done in this service, all errors are returned to a root caller that 
handles all the errors. 

There is three possible error types in the service. 

1. Client error
2. Server error 
3. non http error 

All errors in this code is wrapped with either ClientError or ServerError, the 
purpose of this is to be able to send a correct status code to the client.

The non http error is due to an error not being wrapped, this is not supposed to
occur in this code. 

### Over fetching in neighbourunis endpoint
The name of the country is sometimes different in the foreign api's. The only 
field that is guaranteed to be the same between the api's is the alpha two code.

To solve this issue there are several solutions the two main solutions that was 
suggested was to either over fetch or to use the alternate names in the country 
api. 

Using the alternate names in the country to call the university api increases 
the amount of api calls. Which increases delay. 

Over fetching reduced the amount of api calls, but wastes more ram and cpu 
because the service have to set this data into a struct. 

The solution that is used in the service is over fetching, the service has 
therefore prioritized reducing delay rather than reducing resource usage. 

### Duplicate universities from university api
The university api that is used sometimes give duplicate universities when 
searching with "name". This is somehow solved when using their "name_contains" 
instead. The reason for this unknown and an issue has been posted on their 
github.

## Improvements
### Caching 
The api makes a lot of the same calls to the countries api. It does have a local
for each request, but not a persistent cache that can be used for all the 
requests. Introducing a cache would of course be more taxing on the server 
resources, but it would improve performance. 

# Countries mock 
This mock server is supposed to imitate:

https://restcountries.com/ 

## To build
Navigate into countries_mock and run this command:
```bash
go run ./main.go
``` 
Remember to navigate into the countries_mock otherwise the application won't
be able to find its resource file. 

## How to use the api
Requests: 
```
Method: GET
Path: .../v3.1/name/{name}
``` 

Or

```
Method: GET
Path: .../v3.1/name/{name}?fullText=true
```

Or 

```
Method: GET
Path: .../v3.1/alpha/{code}
```

Or

```
Method: GET
Path: .../v3.1/subregion/{region}
```


The mock server will give out two different responses depending on the request
used. The mock server will give out the same output as long as the url is valid. 

If fullText, alpha or subregion is used the mock server will give out the json 
file named norway.json located under res.

If only name is given the mock server will return s_countries.json with one 
exception. 

The exception is if the name in the mock server is "empty" it will return 404.



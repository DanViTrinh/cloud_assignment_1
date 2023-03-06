# Countries mock 
This mock server is supposed to imitate:

http://universities.hipolabs.com

## To build
Navigate into universities_mock and run this command:
```bash
go run ./main.go
``` 
Remember to navigate into the universities_mock otherwise the application won't
be able to find its resource file. 

## How to use the api
Requests: 
```
Method: GET
Path: .../search?name={name_of_uni}
``` 

Or

```
Method: GET
Path: .../search?name={name_of_uni}&country={country_name}
```
The mock server will give out two different responses depending on the request
used. The country name and university name does not matter. 
The mock server will give the same output as long as the url is valid.

If country is given the mock server will return the json file named 
university_with_country.json under res.

If only name is given it will return the json file named university.json under 
res. 
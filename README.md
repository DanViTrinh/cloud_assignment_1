# Cloud assignment 1 
The goal of this assignment is to make a REST application that can retrieve
university information with extra country information. 

The project uses two foreign api's:

http://universities.hipolabs.com/

https://restcountries.com/

Since the project uses foreign api's I have decided to use mock servers.

## Directory structure
```bash
├───.github
│   └───workflows
├───mocks
│   ├───countries_mock
│   │   └───res
│   └───universities_mock
│       └───res
└───university_service
    ├───cmd
    ├───handlers
    └───utilities
```
In .github is where all workflows (github actions) is located.

The project is split in two main directories. One mocks and the other is 
university_service. 

There is a README for each of the main directories which gives more detailed
information about the project.





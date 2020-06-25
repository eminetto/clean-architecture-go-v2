# Clean Architecture in Go

[![Build Status](https://travis-ci.org/eminetto/clean-architecture-go.svg?branch=master)](https://travis-ci.org/eminetto/clean-architecture-go)

Clean Architecture sample

## Post

[https://medium.com/@eminetto/clean-architecture-using-golang-b63587aa5e3f](https://medium.com/@eminetto/clean-architecture-using-golang-b63587aa5e3f)


## Build

  make

## Run tests

  make test

## API requests 

### Add a book

```
curl -X "POST" "http://localhost:8080/v1/book" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json' \
     -d $'{
  "title": "I Am Ozzy",
  "author": "Ozzy Osbourne",
  "pages": 294,
  "quantity":1
}'
```
### Search a book

```
curl "http://localhost:8080/v1/book?title=ozzy" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```

### Show all books

```
curl "http://localhost:8080/v1/book" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```

## CMD 

### Search for a book

```
./bin/search ozzy
```
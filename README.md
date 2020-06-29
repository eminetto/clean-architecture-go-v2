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

### Add book

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
### Search book

```
curl "http://localhost:8080/v1/book?title=ozzy" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```

### Show books

```
curl "http://localhost:8080/v1/book" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```

### Add user

```
curl -X "POST" "http://localhost:8080/v1/user" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json' \
     -d $'{
  "email": "ozzy@metal.net",
  "fist_name": "Ozzy",
  "last_name": "Osbourne",
  "password": "bateater666",
  "quantity":1
}'

```
### Search user

```
curl "http://localhost:8080/v1/user?name=ozzy" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```

### Show users

```
curl "http://localhost:8080/v1/user" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```

## CMD 

### Search for a book

```
./bin/search ozzy
```
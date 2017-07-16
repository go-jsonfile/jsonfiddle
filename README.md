
# jsonfiddle

[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/go-jsonfile/jsonfiddle?status.svg)](http://godoc.org/github.com/go-jsonfile/jsonfiddle)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-jsonfile/jsonfiddle)](https://goreportcard.com/report/github.com/go-jsonfile/jsonfiddle)
[![travis Status](https://travis-ci.org/go-jsonfile/jsonfiddle.svg?branch=master)](https://travis-ci.org/go-jsonfile/jsonfiddle)

## TOC
- [jsonfiddle - JSON Fiddling](#jsonfiddle---json-fiddling)
- [Usage](#usage)
  - [$ jsonfiddle](#-jsonfiddle)
  - [$ jsonfiddle fmt](#-jsonfiddle-fmt)
  - [$ jsonfiddle sort](#-jsonfiddle-sort)
  - [$ jsonfiddle j2s](#-jsonfiddle-j2s)
- [Examples](#examples)
  - [Format with `jsonfiddle fmt`](#format-with-`jsonfiddle-fmt`)
    - [Pretty print](#pretty-print)
      - [> test/CustomerI.ref](#-testcustomeriref)
    - [Compact](#compact)
      - [> test/Customer.ref](#-testcustomerref)
    - [Sort fields with `jsonfiddle sort`](#sort-fields-with-`jsonfiddle-sort`)
      - [Sort with pretty print](#sort-with-pretty-print)
        - [> test/CustomerSI.ref](#-testcustomersiref)
      - [Sort in compact](#sort-in-compact)
        - [> test/CustomerSC.ref](#-testcustomerscref)
      - [JSON to struct via `jsonfiddle j2s`](#json-to-struct-via-`jsonfiddle-j2s`)
        - [> test/CustomerJ2S.ref](#-testcustomerj2sref)
- [Purpose](#purpose)
- [Installation](#installation)
  - [Credits](#credits)
  - [Similar Projects](#similar-projects)

## jsonfiddle - JSON Fiddling

The `jsonfiddle` makes it easy to look at the JSON data from different aspects. 

- **`jsonfiddle fmt`** will format the JSON data, either compact it or pretty printing it. The order of fields are intact. 
- **`jsonfiddle sort`** will sort the JSON data fields recursively, so that the attributes at any level are in sorted order.
- **`jsonfiddle j2s`** means json to struct. It will extract the structure of JSON data as Go struct.

# Usage

### $ jsonfiddle
```sh
JSON Fiddling
built on 2017-07-16

Tool to fiddle with json strings

Options:

  -h, --help         display help information
      --prefix       prefix for json string output
  -d, --indent[= ]   indent for json string output
  -c, --compact      Compact JSON data, remove all whitespaces
  -v, --verbose      Verbose mode (Multiple -v options increase the verbosity.)

Commands:

  fmt    Format json string
  sort   Sort json fields recursively
  j2s    json to struct
```

### $ jsonfiddle fmt
```sh
Format json string

Options:

  -h, --help         display help information
      --prefix       prefix for json string output
  -d, --indent[= ]   indent for json string output
  -c, --compact      Compact JSON data, remove all whitespaces
  -v, --verbose      Verbose mode (Multiple -v options increase the verbosity.)
  -i, --input       *the source to get json string from (mandatory)
  -o, --output       the output (default: stdout)
```

### $ jsonfiddle sort
```sh
Sort json fields recursively

Options:

  -h, --help         display help information
      --prefix       prefix for json string output
  -d, --indent[= ]   indent for json string output
  -c, --compact      Compact JSON data, remove all whitespaces
  -v, --verbose      Verbose mode (Multiple -v options increase the verbosity.)
  -i, --input       *the source to get json string from (mandatory)
  -o, --output       the output (default: stdout)
```

### $ jsonfiddle j2s
```sh
json to struct

Options:

  -h, --help         display help information
      --prefix       prefix for json string output
  -d, --indent[= ]   indent for json string output
  -c, --compact      Compact JSON data, remove all whitespaces
  -v, --verbose      Verbose mode (Multiple -v options increase the verbosity.)
  -f, --fmt[=json]   the structural format of the input data (json/yaml)
  -i, --input       *the source of the input JSON (mandatory)
  -o, --output       the output (default: stdout)
      --name         the name of the root struct (default: as input file name)
      --pkg[=main]   the name of the package for the generated code
      --subStruct    create types for sub-structs
```

# Examples

## Format with `jsonfiddle fmt`

### Pretty print

	$ jsonfiddle fmt -i test/Customer.json

#### > test/CustomerI.ref
```json
{
 "firstName": "John",
 "lastName": "Smith",
 "age": 25,
 "address": {
  "streetAddress": "21 2nd Street",
  "city": "New York",
  "state": "NY",
  "postalCode": "10021"
 },
 "phoneNumber": [
  {
   "type": "home",
   "number": "212 555-1234"
  },
  {
   "type": "fax",
   "number": "646 555-4567"
  }
 ]
}
```

### Compact

	$ jsonfiddle fmt -c -i test/Customer.json

#### > test/Customer.ref
```json
{"firstName":"John","lastName":"Smith","age":25,"address":{"streetAddress":"21 2nd Street","city":"New York","state":"NY","postalCode":"10021"},"phoneNumber":[{"type":"home","number":"212 555-1234"},{"type":"fax","number":"646 555-4567"}]}
```

You can also do,

	$ cat Customer.json | jsonfiddle fmt -c -i

and the result is the same (and for all other examples using `-i` as well). 

## Sort fields with `jsonfiddle sort`

### Sort with pretty print

	$ jsonfiddle sort -i test/Customer.json

#### > test/CustomerSI.ref
```json
{
 "address": {
  "city": "New York",
  "postalCode": "10021",
  "state": "NY",
  "streetAddress": "21 2nd Street"
 },
 "age": 25,
 "firstName": "John",
 "lastName": "Smith",
 "phoneNumber": [
  {
   "number": "212 555-1234",
   "type": "home"
  },
  {
   "number": "646 555-4567",
   "type": "fax"
  }
 ]
}
```

### Sort in compact

	$ jsonfiddle sort -c -i test/Customer.json

#### > test/CustomerSC.ref
```json
{"address":{"city":"New York","postalCode":"10021","state":"NY","streetAddress":"21 2nd Street"},"age":25,"firstName":"John","lastName":"Smith","phoneNumber":[{"number":"212 555-1234","type":"home"},{"number":"646 555-4567","type":"fax"}]}
```

## JSON to struct via `jsonfiddle j2s`

	$ jsonfiddle j2s -i test/Customer.json

#### > test/CustomerJ2S.ref
```go
package main

type Customer struct {
	Address struct {
		City          string `json:"city"`
		PostalCode    string `json:"postalCode"`
		State         string `json:"state"`
		StreetAddress string `json:"streetAddress"`
	} `json:"address"`
	Age         int64  `json:"age"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber []struct {
		Number string `json:"number"`
		Type   string `json:"type"`
	} `json:"phoneNumber"`
}
```


# Purpose

A few more words on why I'm writing the tool -- because I need to compare JSON string that are roughly close and very complicated in the mean time -- sometimes even less than 30% of fields are the same, of course, this is after their having been sorted, otherwise, it'd be 100% different.

Thus all the JSON comparison tools I found are failing under such hash request. So far, I personally find that

- Sorting the JSON data fields recursively and producing plain text file (via `jsonfiddle sort`), then use the state-of-the-art text comparison tools to compare them is the best approach, for my above scenario.
- For extremely long and very complicated JSONs, converting json to abstract Go struct (via `jsonfiddle j2s`) is the quickest approach to compare them at higher level.


# Installation

```
go get github.com/go-jsonfile/jsonfiddle
```

All patches welcome.


## Credits

- [Ladicle/gojson](https://github.com/Ladicle/gojson) forked source for JSON to struct
- [ChimeraCoder/gojson](https://github.com/ChimeraCoder/gojson) the original source of [Ladicle/gojson](https://github.com/Ladicle/gojson).

## Similar Projects

All the following similar projects have been considered before writing one on my own instead.

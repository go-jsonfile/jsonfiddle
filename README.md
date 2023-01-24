
# jsonfiddle

[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/go-jsonfile/jsonfiddle?status.svg)](http://godoc.org/github.com/go-jsonfile/jsonfiddle)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-jsonfile/jsonfiddle)](https://goreportcard.com/report/github.com/go-jsonfile/jsonfiddle)
[![travis Status](https://travis-ci.org/go-jsonfile/jsonfiddle.svg?branch=master)](https://travis-ci.org/go-jsonfile/jsonfiddle)
[![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-R.svg)](http://godoc.org/github.com/go-easygen/wireframe)

## TOC
- [jsonfiddle - JSON Fiddling](#jsonfiddle---json-fiddling)
- [Usage](#usage)
  - [$ jsonfiddle || true](#-jsonfiddle--true)
  - [$ jsonfiddle esc || true](#-jsonfiddle-esc--true)
  - [$ jsonfiddle sort || true](#-jsonfiddle-sort--true)
  - [$ jsonfiddle x2j || true](#-jsonfiddle-x2j--true)

## jsonfiddle - JSON Fiddling

The `jsonfiddle` makes it easy to look at the JSON data from different aspects. 

- **`jsonfiddle esc`** will escape any arbitrary string so as to embed it as content of json string.
- **`jsonfiddle fmt`** will format the JSON data, either compact it or pretty printing it. The order of fields are intact. 
- **`jsonfiddle sort`** will sort the JSON data fields recursively, so that the attributes at any level are in sorted order.
- **`jsonfiddle x2j`** means xml to json. It will convert data from xml format into json.
- **`jsonfiddle j2s`** means json to struct. It will extract the structure of JSON data as Go struct.

# Usage

```
$ jsonfiddle -V
jsonfiddle - JSON Fiddling, version 0.5.0
Built on 2023-01-22
Copyright (C) 2017-2023, Tong Sun

Tool to fiddle with json strings
```

### $ jsonfiddle || true
```sh
Please specify one command of: esc, fmt, j2s, sort or x2j

Usage:
  jsonfiddle [OPTIONS] <command>

Application Options:
  -c, --compact  Compact JSON data, remove all whitespaces
      --prefix=  prefix for json string output
  -d, --indent=  indent for json string output (default:  )
  -p, --protect  protect {{TEMPLATE}} in JSON data
  -v, --verbose  Verbose mode (Multiple -v options increase the verbosity)
  -V, --version  Show program version and exit

Help Options:
  -h, --help     Show this help message

Available commands:
  esc   Escape json string
  fmt   Format json string
  j2s   JSON to struct
  sort  Sort json fields recursively
  x2j   XML to JSON
```

### $ jsonfiddle esc || true
```sh
the required flag `-i, --input' was not specified

Usage:
  jsonfiddle [OPTIONS] esc [esc-OPTIONS]

Application Options:
  -c, --compact     Compact JSON data, remove all whitespaces
      --prefix=     prefix for json string output
  -d, --indent=     indent for json string output (default:  )
  -p, --protect     protect {{TEMPLATE}} in JSON data
  -v, --verbose     Verbose mode (Multiple -v options increase the verbosity)
  -V, --version     Show program version and exit

Help Options:
  -h, --help        Show this help message

[esc command options]
      -i, --input=  the source to get json string from (or "-" for stdin)
                    (mandatory)
      -o, --output= the output, default to stdout (default: -)
```

### $ jsonfiddle fmt || true
```sh
the required flag `-i, --input' was not specified

Usage:
  jsonfiddle [OPTIONS] fmt [fmt-OPTIONS]

Application Options:
  -c, --compact     Compact JSON data, remove all whitespaces
      --prefix=     prefix for json string output
  -d, --indent=     indent for json string output (default:  )
  -p, --protect     protect {{TEMPLATE}} in JSON data
  -v, --verbose     Verbose mode (Multiple -v options increase the verbosity)
  -V, --version     Show program version and exit

Help Options:
  -h, --help        Show this help message

[fmt command options]
      -i, --input=  the source to get json string from (mandatory)
      -o, --output= the output, default to stdout (default: -)
```

### $ jsonfiddle sort || true
```sh
the required flag `-i, --input' was not specified

Usage:
  jsonfiddle [OPTIONS] sort [sort-OPTIONS]

Application Options:
  -c, --compact     Compact JSON data, remove all whitespaces
      --prefix=     prefix for json string output
  -d, --indent=     indent for json string output (default:  )
  -p, --protect     protect {{TEMPLATE}} in JSON data
  -v, --verbose     Verbose mode (Multiple -v options increase the verbosity)
  -V, --version     Show program version and exit

Help Options:
  -h, --help        Show this help message

[sort command options]
      -i, --input=  the source to get json string from (mandatory)
      -o, --output= the output, default to stdout (default: -)
```

### $ jsonfiddle j2s || true
```sh
the required flag `-i, --input' was not specified

Usage:
  jsonfiddle [OPTIONS] j2s [j2s-OPTIONS]

JSON convert to Go struct

Application Options:
  -c, --compact        Compact JSON data, remove all whitespaces
      --prefix=        prefix for json string output
  -d, --indent=        indent for json string output (default:  )
  -p, --protect        protect {{TEMPLATE}} in JSON data
  -v, --verbose        Verbose mode (Multiple -v options increase the verbosity)
  -V, --version        Show program version and exit

Help Options:
  -h, --help           Show this help message

[j2s command options]
      -f, --fmt=       the structural format of the input data (json/yaml)
                       (default: json)
      -i, --input=     the source of the input JSON (mandatory)
      -o, --output=    the output, default to stdout (default: -)
          --name=      the name of the root struct (default: as input file name)
          --pkg=       the name of the package for the generated code (default:
                       main)
          --subStruct  create types for sub-structs
```

### $ jsonfiddle x2j || true
```sh
the required flag `-i, --input' was not specified

Usage:
  jsonfiddle [OPTIONS] x2j [x2j-OPTIONS]

Application Options:
  -c, --compact     Compact JSON data, remove all whitespaces
      --prefix=     prefix for json string output
  -d, --indent=     indent for json string output (default:  )
  -p, --protect     protect {{TEMPLATE}} in JSON data
  -v, --verbose     Verbose mode (Multiple -v options increase the verbosity)
  -V, --version     Show program version and exit

Help Options:
  -h, --help        Show this help message

[x2j command options]
      -i, --input=  the source of the input JSON (mandatory)
      -o, --output= the output, default to stdout (default: -)
```

# Examples

## Escape with `jsonfiddle esc`

### $ jsonfiddle esc -i test/Customer.ref 2>/dev/null
```json
"{\n \"firstName\": \"John\",\n \"lastName\": \"Smith\",\n \"age\": 25,\n \"address\": {\n  \"streetAddress\": \"21 2nd Street\",\n  \"city\": \"New York\",\n  \"state\": \"NY\",\n  \"postalCode\": \"10021\"\n },\n \"phoneNumber\": [\n  {\n   \"type\": \"home\",\n   \"number\": \"212 555-1234\"\n  },\n  {\n   \"type\": \"fax\",\n   \"number\": \"646 555-4567\"\n  }\n ]\n}\n\n"
```

### Usage

`jsonfiddle esc` will escape any arbitrary string so as to embed it as content of json string. This seems useless at first, but it actually allows you to embed any arbitrary file into [GitHub Gists JSON API](https://developer.github.com/v3/gists/), so as to post any arbitrary file onto GitHub Gist:


    echo '{"description":"SmartyStreets API Demo","public":true,"files":{"SmartyStreets.json":{"content":'"`jsonfiddle fmt -i test/SmartyStreets.json | jsonfiddle esc -i`"'}}}' | curl --data @- https://api.github.com/gists

This will give you
https://gist.github.com/anonymous/1423d4768dd9b88262ca513626e68d8e


By "_arbitrary file_" I do mean arbitrary file. Check this out:
https://gist.github.com/anonymous/a51798ce99ff59d8d4ba536cbf4b6996

This is why `jsonfiddle esc` is a command on its own, instead of being part of functionalities of `jsonfiddle fmt` or `jsonfiddle sort`.

## Format with `jsonfiddle fmt`

### Pretty print

	$ jsonfiddle fmt -i test/Customer.json

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

### Protect templates in json data

There are times that json data may contain templates, i.e., strings like `{{VARIABLE}}`. Some of the pretty printing tools, like the json plugin in Notepad++, cannot handle such template forms well, and will turn `{{VARIABLE}}` into:

```json
{
  {
    VARIABLE
  }
}
```

What's worse is that when such template variables are for `int`, e.g.: `"age":{{Var_Age}}`, they then wouldn't be able to handle it, for inputs like
test/CustomerP.json
```sh
{"firstName":"{{C_firstName}}","lastName":"{{C_lastName}}","age":{{C_age}},"address":{"streetAddress":"{{C_address1}}","city":"{{C_city}}","state":"NY","postalCode":"10021"}}
```

To make such template variables work for those tools, the `-p,--protect` option is introduced:

	$ jsonfiddle fmt -p -i test/CustomerP.json

#### > test/CustomerP.ref
```json
{
 "firstName": "<<C_firstName>>",
 "lastName": "<<C_lastName>>",
 "age": "<<C_age>>",
 "address": {
  "streetAddress": "<<C_address1>>",
  "city": "<<C_city>>",
  "state": "NY",
  "postalCode": "10021"
 }
}
```

### Compact

	$ jsonfiddle fmt -c -i test/Customer.json

#### > test/CustomerC.ref
```json
{"firstName":"John","lastName":"Smith","age":25,"address":{"streetAddress":"21 2nd Street","city":"New York","state":"NY","postalCode":"10021"},"phoneNumber":[{"type":"home","number":"212 555-1234"},{"type":"fax","number":"646 555-4567"}]}
```

You can also do,

	$ cat Customer.json | jsonfiddle fmt -c -i -

and the result is the same (and for all other examples using `-i` as well). 

## Sort fields with `jsonfiddle sort`

### Sort in compact

	$ jsonfiddle sort -c -i test/Customer.json

#### > test/CustomerSC.ref
```json
{"address":{"city":"New York","postalCode":"10021","state":"NY","streetAddress":"21 2nd Street"},"age":25,"firstName":"John","lastName":"Smith","phoneNumber":[{"number":"212 555-1234","type":"home"},{"number":"646 555-4567","type":"fax"}]}
```

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

### XML to JSON, sort then pretty print

```
$ jsonfiddle x2j -i test/Books.xml | jsonfiddle sort -i - | jsonfiddle fmt -i -
XML to JSON
jsonfiddle v0.5.0. x2j - XML to JSON
Sort json fields recursively
Format json string
{
 "catalog": {
  "book": [
   {
    "-id": "bk101",
    "author": "Gambardella, Matthew",
    "description": "An in-depth look at creating applications \n      with XML.",
    "genre": "Computer",
    "price": "44.95",
    "publish_date": "2000-10-01",
    "title": "XML Developer's Guide"
   },
   {
    "-id": "bk102",
    "author": "Ralls, Kim",
    "description": "A former architect battles corporate zombies, \n      an evil sorceress, and her own childhood to become queen \n      of the world.",
    "genre": "Fantasy",
    "price": "5.95",
    "publish_date": "2000-12-16",
    "title": "Midnight Rain"
   },
   {
    "-id": "bk103",
    "author": "Corets, Eva",
    "description": "After the collapse of a nanotechnology \n      society in England, the young survivors lay the \n      foundation for a new society.",
    "genre": "Fantasy",
    "price": "5.95",
    "publish_date": "2000-11-17",
    "title": "Maeve Ascendant"
   }
  ]
 }
}
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

# Download binaries

- The latest binary executables are available under  
https://github.com/go-jsonfile/jsonfiddle/releases  
as the result of the Continuous-Integration process.
- I.e., they are built right from the source code during every git tagging commit automatically.
- Pick & choose the binary executable that suits your OS and its architecture. E.g., for Linux, it would most probably be the `jsonfiddle_linux_VER_amd64` file. If your OS and its architecture is not available in the download list, please let me know and I'll add it.
- You may want to rename it to a shorter name instead, e.g., `jsonfiddle`, after downloading it. 


# Debian package

Available at the above releases url as well.

# Install Source

To install the source code instead:

```
go get github.com/go-jsonfile/jsonfiddle
```


## Credits

- [Ladicle/gojson](https://github.com/Ladicle/gojson) forked source for JSON to struct
- [ChimeraCoder/gojson](https://github.com/ChimeraCoder/gojson) the original source of [Ladicle/gojson](https://github.com/Ladicle/gojson).

## Similar Projects

All the following similar projects have been considered before writing one on my own instead.

. . . to be filled . . .

## Author(s) & Contributor(s)

Tong SUN  
![suntong from cpan.org](https://img.shields.io/badge/suntong-%40cpan.org-lightgrey.svg "suntong from cpan.org")

_Powered by_ [**WireFrame**](https://github.com/go-easygen/wireframe),  [![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-Y.svg)](http://godoc.org/github.com/go-easygen/wireframe), the _one-stop wire-framing solution_ for Go cli based projects, from start to deploy.

All patches welcome. 

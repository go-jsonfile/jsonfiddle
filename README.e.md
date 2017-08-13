
# {{.Name}}

{{render "license/shields" . "License" "MIT"}}
{{template "badge/godoc" .}}
{{template "badge/goreport" .}}
{{template "badge/travis" .}}

## {{toc 5}}

## {{.Name}} - JSON Fiddling

The `jsonfiddle` makes it easy to look at the JSON data from different aspects. 

- **`jsonfiddle fmt`** will format the JSON data, either compact it or pretty printing it. The order of fields are intact. 
- **`jsonfiddle sort`** will sort the JSON data fields recursively, so that the attributes at any level are in sorted order.
- **`jsonfiddle j2s`** means json to struct. It will extract the structure of JSON data as Go struct.

# Usage

### $ {{exec "jsonfiddle" | color "sh"}}

### $ {{shell "jsonfiddle fmt" | color "sh"}}

### $ {{shell "jsonfiddle sort" | color "sh"}}

### $ {{shell "jsonfiddle j2s" | color "sh"}}

# Examples

## Format with `jsonfiddle fmt`

### Pretty print

	$ jsonfiddle fmt -i test/Customer.json

#### > {{cat "test/CustomerSI.ref" | color "json"}}

### Compact

	$ jsonfiddle fmt -c -i test/Customer.json

#### > {{cat "test/Customer.ref" | color "json"}}

You can also do,

	$ cat Customer.json | jsonfiddle fmt -c -i

and the result is the same (and for all other examples using `-i` as well). 

## Sort fields with `jsonfiddle sort`

### Sort with pretty print

	$ jsonfiddle sort -i test/Customer.json

#### > {{cat "test/CustomerSI.ref" | color "json"}}

### Sort in compact

	$ jsonfiddle sort -c -i test/Customer.json

#### > {{cat "test/CustomerSC.ref" | color "json"}}

## JSON to struct via `jsonfiddle j2s`

	$ jsonfiddle j2s -i test/Customer.json

#### > {{cat "test/CustomerJ2S.ref" | color "go"}}


# Purpose

A few more words on why I'm writing the tool -- because I need to compare JSON string that are roughly close and very complicated in the mean time -- sometimes even less than 30% of fields are the same, of course, this is after their having been sorted, otherwise, it'd be 100% different.

Thus all the JSON comparison tools I found are failing under such hash request. So far, I personally find that

- Sorting the JSON data fields recursively and producing plain text file (via `jsonfiddle sort`), then use the state-of-the-art text comparison tools to compare them is the best approach, for my above scenario.
- For extremely long and very complicated JSONs, converting json to abstract Go struct (via `jsonfiddle j2s`) is the quickest approach to compare them at higher level.

# Download binaries

- The latest binary executables are available under  
https://github.com/go-jsonfile/{{.Name}}/releases  
as the result of the Continuous-Integration process.
- I.e., they are built right from the source code during every git tagging commit automatically by [travis-ci](https://travis-ci.org/).
- Pick & choose the binary executable that suits your OS and its architecture. E.g., for Linux, it would most probably be the `{{.Name}}_linux_VER_amd64` file. If your OS and its architecture is not available in the download list, please let me know and I'll add it.
- You may want to rename it to a shorter name instead, e.g., `{{.Name}}`, after downloading it. 


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

## Author(s) & Contributor(s)

Tong SUN  
![suntong from cpan.org](https://img.shields.io/badge/suntong-%40cpan.org-lightgrey.svg "suntong from cpan.org")

All patches welcome. 

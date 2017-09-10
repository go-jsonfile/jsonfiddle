
# {{.Name}}

{{render "license/shields" . "License" "MIT"}}
{{template "badge/godoc" .}}
{{template "badge/goreport" .}}
{{template "badge/travis" .}}

## {{toc 5}}

## {{.Name}} - JSON Fiddling

The `jsonfiddle` makes it easy to look at the JSON data from different aspects. 

- **`jsonfiddle esc`** will escape any arbitrary string so as to embed it as content of json string.
- **`jsonfiddle fmt`** will format the JSON data, either compact it or pretty printing it. The order of fields are intact. 
- **`jsonfiddle sort`** will sort the JSON data fields recursively, so that the attributes at any level are in sorted order.
- **`jsonfiddle j2s`** means json to struct. It will extract the structure of JSON data as Go struct.

# Usage

### $ {{exec "jsonfiddle" | color "sh"}}

### $ {{shell "jsonfiddle esc" | color "sh"}}

### $ {{shell "jsonfiddle fmt" | color "sh"}}

### $ {{shell "jsonfiddle sort" | color "sh"}}

### $ {{shell "jsonfiddle j2s" | color "sh"}}

# Examples

## Escape with `jsonfiddle esc`

### $ {{shell "jsonfiddle esc -i test/Customer.ref" | color "json"}}

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

#### > {{cat "test/CustomerSI.ref" | color "json"}}

### Protect templates in json data

There are times that json data may contain templates, i.e., strings like `{{"{{VARIABLE}}"}}`. Some of the pretty printing tools, like the json plugin in Notepad++, cannot handle such template forms well, and will turn `{{"{{VARIABLE}}"}}` into:

```json
{
  {
    VARIABLE
  }
}
```

What's worse is that when such template variables are for `int`, e.g.: `"age":{{"{{Var_Age}}"}}`, they then wouldn't be able to handle it.

To make such template variables work for those tools, the `-p,--protect` option is introduced:

	$ jsonfiddle fmt -p -i test/CustomerP.json

#### > {{cat "test/CustomerP.ref" | color "json"}}

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

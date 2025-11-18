////////////////////////////////////////////////////////////////////////////
// Program: jsonfiddle
// Purpose: JSON Fiddling
// Authors: Tong Sun (c) 2017-2023, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"

	"github.com/go-easygen/go-flags/clis"
)

// *** Sub-command: fmt ***
// Exec implements the business logic of command `fmt`
func (x *FmtCommand) Exec(args []string) error {
	fileI := clis.GetInputStream(x.Filei)
	defer fileI.Close()
	data := readJson(fileI)
	if x.Unescape {
		data = unescapeUnicode(data)
	}

	var out bytes.Buffer
	var err error
	if opts.Compact {
		err = json.Compact(&out, data)
	} else {
		err = json.Indent(&out, data, opts.Prefix, opts.Indent)
	}
	clis.AbortOn("Formatting input", err)
	fileO := clis.GetOutputStream(x.Fileo)
	defer fileO.Close()
	out.WriteTo(fileO)
	fmt.Fprintln(fileO)
	return nil
}

func unescapeUnicode(b []byte) []byte {
	// Unescape Unicode escape sequences like \u003c to actual characters
	re := regexp.MustCompile(`\\u[0-9a-fA-F]{4}`)
	return re.ReplaceAllFunc(b, func(match []byte) []byte {
		// Convert \uXXXX to actual character
		r, err := strconv.ParseInt(string(match[2:]), 16, 32)
		if err != nil {
			return match
		}
		return []byte(string(rune(r)))
	})
}

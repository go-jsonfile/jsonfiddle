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

	"github.com/go-easygen/go-flags/clis"
)

// *** Sub-command: fmt ***
// Exec implements the business logic of command `fmt`
func (x *FmtCommand) Exec(args []string) error {
	fileI := clis.GetInputStream(x.Filei)
	defer fileI.Close()
	data := readJson(fileI)

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

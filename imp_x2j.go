////////////////////////////////////////////////////////////////////////////
// Program: jsonfiddle
// Purpose: JSON Fiddling
// Authors: Tong Sun (c) 2019-2023, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"

	xj "github.com/basgys/goxml2json"

	"github.com/go-easygen/go-flags/clis"
)

////////////////////////////////////////////////////////////////////////////
// x2j

// *** Sub-command: x2j ***
// Exec implements the business logic of command `x2j`
func (x *X2jCommand) Exec(args []string) error {
	fmt.Fprintf(os.Stderr, "%s v%s. x2j - XML to JSON\n", progname, version)
	// err := ...
	// clis.WarnOn("X2j, Exec", err)
	fileI := clis.GetInputStream(x.Filei)
	defer fileI.Close()
	fileO := clis.GetOutputStream(x.Fileo)
	defer fileO.Close()

	json, err := xj.Convert(fileI)
	clis.AbortOn("Convert from xml to json", err)

	fmt.Fprintln(fileO, json.String())
	return nil
}

////////////////////////////////////////////////////////////////////////////
// Program: jsonfiddle
// Purpose: JSON Fiddling
// Authors: Tong Sun (c) 2017-2023, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"path/filepath"

	"github.com/go-easygen/go-flags/clis"
	"github.com/go-jsonfile/gojson"
	"github.com/suntong/enum"
)

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	fmtTypeEn = enum.NewEnum()
	fmtJson   = fmtTypeEn.Iota("json")
	fmtYaml   = fmtTypeEn.Iota("yaml")

	gjparser = []gojson.Parser{gojson.ParseJson, gojson.ParseYaml}
)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// *** Sub-command: j2s ***
// Exec implements the business logic of command `j2s`
func (x *J2sCommand) Exec(args []string) error {
	fmtType, ok := fmtTypeEn.Get(x.FmtType)
	if !ok {
		clis.AbortOn("Get FmtType", fmt.Errorf("Invalid format string: '%s'\n", x.FmtType))
	}
	nameRoot := clis.Basename(filepath.Base(x.Filei))
	if len(x.Name) != 0 {
		nameRoot = x.Name
	}

	fileI := clis.GetInputStream(x.Filei)
	defer fileI.Close()
	output, err := gojson.Generate(fileI, gjparser[fmtType],
		nameRoot, x.Pkg, []string{x.FmtType}, x.SubStruct)
	clis.AbortOn("Gojson Parsing", err)

	fileO := clis.GetOutputStream(x.Fileo)
	defer fileO.Close()
	fmt.Fprint(fileO, string(output))
	// clis.WarnOn("J2s, Exec", err)
	return nil
}

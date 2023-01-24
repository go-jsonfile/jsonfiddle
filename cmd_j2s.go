////////////////////////////////////////////////////////////////////////////
// Program: jsonfiddle
// Purpose: JSON Fiddling
// Authors: Tong Sun (c) 2017-2023, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"

	"github.com/go-easygen/go-flags/clis"
)

// *** Sub-command: j2s ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The J2sCommand type defines all the configurable options from cli.
type J2sCommand struct {
	FmtType   string `short:"f" long:"fmt" description:"the structural format of the input data (json/yaml)" default:"json"`
	Filei     string `short:"i" long:"input" description:"the source of the input JSON (mandatory)" required:"true"`
	Fileo     string `short:"o" long:"output" description:"the output, default to stdout" default:"-"`
	Name      string `long:"name" description:"the name of the root struct (default: as input file name)"`
	Pkg       string `long:"pkg" description:"the name of the package for the generated code" default:"main"`
	SubStruct bool   `long:"subStruct" description:"create types for sub-structs"`
}

var j2sCommand J2sCommand

func init() {
	parser.AddCommand("j2s",
		"JSON to struct",
		"JSON convert to Go struct",
		&j2sCommand)
}

func (x *J2sCommand) Execute(args []string) error {
	fmt.Fprintf(os.Stderr, "JSON to struct\n")
	// fmt.Fprintf(os.Stderr, "Copyright (C) 2017-2023, Tong Sun\n\n")
	clis.Setup("jsonfiddle::j2s", opts.Verbose)
	clis.Verbose(1, "Doing J2s, with %+v, %+v", opts, args)
// 	fmt.Println(x.FmtType, x.Filei, x.Fileo, x.Name, x.Pkg, x.SubStruct)
	return x.Exec(args)
}

// Exec implements the business logic of command `j2s`
// func (x *J2sCommand) Exec(args []string) error {
// 	// err := ...
// 	// clis.WarnOn("J2s, Exec", err)
// 	// or,
// 	// clis.AbortOn("J2s, Exec", err)
// 	return nil
// }

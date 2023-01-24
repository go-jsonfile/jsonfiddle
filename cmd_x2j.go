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

// *** Sub-command: x2j ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The X2jCommand type defines all the configurable options from cli.
type X2jCommand struct {
	Filei string `short:"i" long:"input" description:"the source of the input JSON (mandatory)" required:"true"`
	Fileo string `short:"o" long:"output" description:"the output, default to stdout" default:"-"`
}

var x2jCommand X2jCommand

func init() {
	parser.AddCommand("x2j",
		"XML to JSON",
		"",
		&x2jCommand)
}

func (x *X2jCommand) Execute(args []string) error {
	fmt.Fprintf(os.Stderr, "XML to JSON\n")
	// fmt.Fprintf(os.Stderr, "Copyright (C) 2017-2023, Tong Sun\n\n")
	clis.Setup("jsonfiddle::x2j", opts.Verbose)
	clis.Verbose(1, "Doing X2j, with %+v, %+v", opts, args)
// 	fmt.Println(x.Filei, x.Fileo)
	return x.Exec(args)
}

// Exec implements the business logic of command `x2j`
// func (x *X2jCommand) Exec(args []string) error {
// 	// err := ...
// 	// clis.WarnOn("X2j, Exec", err)
// 	// or,
// 	// clis.AbortOn("X2j, Exec", err)
// 	return nil
// }

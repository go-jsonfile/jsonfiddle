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

// *** Sub-command: esc ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The EscCommand type defines all the configurable options from cli.
type EscCommand struct {
	Filei string `short:"i" long:"input" description:"the source to get json string from (or \"-\" for stdin) (mandatory)" required:"true"`
	Fileo string `short:"o" long:"output" description:"the output, default to stdout" default:"-"`
}

var escCommand EscCommand

func init() {
	parser.AddCommand("esc",
		"Escape json string",
		"",
		&escCommand)
}

func (x *EscCommand) Execute(args []string) error {
	fmt.Fprintf(os.Stderr, "Escape json string\n")
	// fmt.Fprintf(os.Stderr, "Copyright (C) 2017-2023, Tong Sun\n\n")
	clis.Setup("jsonfiddle::esc", opts.Verbose)
	clis.Verbose(1, "Doing Esc, with %+v, %+v", opts, args)
// 	fmt.Println(x.Filei, x.Fileo)
	return x.Exec(args)
}

// Exec implements the business logic of command `esc`
// func (x *EscCommand) Exec(args []string) error {
// 	// err := ...
// 	// clis.WarnOn("Esc, Exec", err)
// 	// or,
// 	// clis.AbortOn("Esc, Exec", err)
// 	return nil
// }

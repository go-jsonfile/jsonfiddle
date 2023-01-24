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

// *** Sub-command: sort ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The SortCommand type defines all the configurable options from cli.
type SortCommand struct {
	Filei string `short:"i" long:"input" description:"the source to get json string from (mandatory)" required:"true"`
	Fileo string `short:"o" long:"output" description:"the output, default to stdout" default:"-"`
}

var sortCommand SortCommand

////////////////////////////////////////////////////////////////////////////
// Function definitions

func init() {
	gfParser.AddCommand("sort",
		"Sort json fields recursively",
		"",
		&sortCommand)
}

func (x *SortCommand) Execute(args []string) error {
	fmt.Fprintf(os.Stderr, "Sort json fields recursively\n")
	// fmt.Fprintf(os.Stderr, "Copyright (C) 2017-2023, Tong Sun\n\n")
	clis.Setup("jsonfiddle::sort", opts.Verbose)
	clis.Verbose(1, "Doing Sort, with %+v, %+v", opts, args)
	// fmt.Println(x.Filei, x.Fileo)
	return x.Exec(args)
}

// // Exec implements the business logic of command `sort`
// func (x *SortCommand) Exec(args []string) error {
// 	// err := ...
// 	// clis.WarnOn("sort::Exec", err)
// 	// or,
// 	// clis.AbortOn("sort::Exec", err)
// 	return nil
// }

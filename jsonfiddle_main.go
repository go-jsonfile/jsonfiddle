// jsonfiddle - JSON Fiddling

// Tool to fiddle with json strings

package main

////////////////////////////////////////////////////////////////////////////
// Program: jsonfiddle
// Purpose: JSON Fiddling
// Authors: Tong Sun (c) 2017-2025, All rights reserved
////////////////////////////////////////////////////////////////////////////

//go:generate sh jsonfiddle_cliGen.sh
//go:generate emd gen -in README.e.md -out README.md

import (
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/go-easygen/go-flags"
	"github.com/go-easygen/go-flags/clis"
)

//////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	progname = "jsonfiddle"
	version  = "0.6.0"
	date     = "2025-11-18"

	// Opts store all the configurable options
	Opts OptsT
)

var gfParser = flags.NewParser(&Opts, flags.Default)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// ==========================================================================
// Function main
func main() {
	Opts.Version = showVersion
	Opts.Verbflg = func() {
		Opts.Verbose++
	}

	if _, err := gfParser.Parse(); err != nil {
		fmt.Println()
		gfParser.WriteHelp(os.Stdout)
		os.Exit(1)
	}
	fmt.Println()
	//DoJsonfiddle()
}

//==========================================================================
// support functions

func showVersion() {
	fmt.Fprintf(os.Stderr, "jsonfiddle - JSON Fiddling, version %s\n", version)
	fmt.Fprintf(os.Stderr, "Built on %s\n", date)
	fmt.Fprintf(os.Stderr, "Copyright (C) 2017-2025, Tong Sun\n\n")
	fmt.Fprintf(os.Stderr, "Tool to fiddle with json strings\n")
	os.Exit(0)
}

// readJson reads the given json file as []byte.
func readJson(r io.Reader) []byte {
	data, err := io.ReadAll(r)
	clis.AbortOn("Reading json input", err)

	if Opts.Protect {
		data = regexp.MustCompile(`({{)([^ }]+)(}})`).
			ReplaceAll(data, []byte(`<<${2}>>`))
		// "age":<<C_age>> => "age":"<<C_age>>"
		data = regexp.MustCompile(`(:)(<<[^>]+>>)([]},])`).
			ReplaceAll(data, []byte(`${1}"${2}"${3}`))
	}
	clis.Verbose(2, "%s", string(data))
	return data
}

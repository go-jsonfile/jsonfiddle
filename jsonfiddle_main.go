////////////////////////////////////////////////////////////////////////////
// Program: jf
// Purpose: JSON Fiddling
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

//go:generate sh -v jsonfiddle_cliGen.sh

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/mkideal/cli"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The OptsT type defines all the configurable options for jsonfiddle.
type OptsT struct {
	Prefix  string
	Indent  string
	Compact bool
	Protect bool
	Verbose int
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	progname = "jsonfiddle"
	// version tracks the release version.
	version = "0.5.0"
	date    = "2019-05-30"
)

var (
	rootArgv *rootT
	// Opts store all the configurable options for jsonfiddle.
	Opts OptsT
)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {
	//NOTE: You can set any writer implements io.Writer
	// default writer is os.Stdout
	if err := cli.Root(root,
		cli.Tree(escDef),
		cli.Tree(fmtDef),
		cli.Tree(sortDef),
		cli.Tree(j2sDef),
		cli.Tree(x2jDef)).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("")
}

//==========================================================================
// Main dispatcher

func jsonfiddle(ctx *cli.Context) error {
	ctx.JSON(ctx.RootArgv())
	ctx.JSON(ctx.Argv())
	fmt.Println()

	return nil
}

//==========================================================================
// support functions

// Basename returns the file name without extension.
func readJson(r io.Reader) []byte {
	data, err := ioutil.ReadAll(r)
	abortOn("Reading json input", err)

	if Opts.Protect {
		data = regexp.MustCompile(`({{)([^ }]+)(}})`).
			ReplaceAll(data, []byte(`<<${2}>>`))
		// "age":<<C_age>> => "age":"<<C_age>>"
		data = regexp.MustCompile(`(:)(<<[^>]+>>)([]},])`).
			ReplaceAll(data, []byte(`${1}"${2}"${3}`))
	}
	verbose(2, "%s", string(data))
	return data
}

// Basename returns the file name without extension.
func Basename(s string) string {
	n := strings.LastIndexByte(s, '.')
	if n > 0 {
		return s[:n]
	}
	return s
}

// abortOn will quit on anticipated errors gracefully without stack trace
func abortOn(errCase string, e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "[%s] %s error: %v\n", progname, errCase, e)
		os.Exit(1)
	}
}

// verbose will print info to stderr according to the verbose level setting
func verbose(levelSet int, format string, args ...interface{}) {
	if Opts.Verbose >= levelSet {
		fmt.Fprintf(os.Stderr, "["+progname+"] "+format+"\n", args...)
	}
}

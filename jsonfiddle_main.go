////////////////////////////////////////////////////////////////////////////
// Program: jf
// Purpose: JSON Fiddling
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

//go:generate sh -v jsonfiddle_cliGen.sh

import (
	"fmt"
	"os"
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
	Verbose int
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	progname = "jsonfiddle"
	// version tracks the release version.
	version = "v0.3.0"
	date    = "2017-08-16"
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
		cli.Tree(j2sDef)).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
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
func verbose(levelSet, levelNow int, format string, args ...interface{}) {
	if levelNow >= levelSet {
		fmt.Fprintf(os.Stderr, "["+progname+"] "+format+"\n", args...)
	}
}

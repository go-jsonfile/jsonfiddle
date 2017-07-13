////////////////////////////////////////////////////////////////////////////
// Program: jsonfiddle
// Purpose: JSON Fiddling
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"github.com/mkideal/cli"
	clix "github.com/mkideal/cli/ext"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

//==========================================================================
// jsonfiddle

type rootT struct {
	cli.Helper
	Prefix  string      `cli:"prefix" usage:"prefix for json string output"`
	Indent  string      `cli:"d,indent" usage:"indent for json string output" dft:" "`
	Compact bool        `cli:"c,compact" usage:"Compact JSON data, remove all whitespaces"`
	Verbose cli.Counter `cli:"v,verbose" usage:"Verbose mode (Multiple -v options increase the verbosity.)"`
}

var root = &cli.Command{
	Name:   "jsonfiddle",
	Desc:   "JSON Fiddling\nbuilt on " + buildTime,
	Text:   "Tool to fiddle with json strings",
	Global: true,
	Argv:   func() interface{} { return new(rootT) },
	Fn:     jsonfiddle,

	NumArg: cli.AtLeast(1),
}

// Template for main starts here
////////////////////////////////////////////////////////////////////////////
// Global variables definitions

//  var (
//          progname  = "jsonfiddle"
//          VERSION   = "0.1.0"
//          buildTime = "2017-07-13"
//  )

//  var rootArgv *rootT

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
//  func main() {
//  	cli.SetUsageStyle(cli.ManualStyle) // up-down, for left-right, use NormalStyle
//  	//NOTE: You can set any writer implements io.Writer
//  	// default writer is os.Stdout
//  	if err := cli.Root(root,
//  		cli.Tree(sortDef)).Run(os.Args[1:]); err != nil {
//  		fmt.Fprintln(os.Stderr, err)
//  	}
//  	fmt.Println("")
//  }

// Template for main dispatcher starts here
//==========================================================================
// Main dispatcher

//  func jsonfiddle(ctx *cli.Context) error {
//  	ctx.JSON(ctx.RootArgv())
//  	ctx.JSON(ctx.Argv())
//  	fmt.Println()

//  	return nil
//  }

// Template for CLI handling starts here

////////////////////////////////////////////////////////////////////////////
// sort

//  func sortCLI(ctx *cli.Context) error {
//  	rootArgv = ctx.RootArgv().(*rootT)
//  	argv := ctx.Argv().(*sortT)
//  	fmt.Printf("[sort]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
//  	return nil
//  }

type sortT struct {
	Filei *clix.Reader `cli:"*i,input" usage:"The source to get json string from (mandatory)"`
	Fileo *clix.Writer `cli:"o,output" usage:"The output (default: stdout)"`
}

var sortDef = &cli.Command{
	Name: "sort",
	Desc: "Sort json fields recursive",
	Argv: func() interface{} { return new(sortT) },
	Fn:   sortCLI,

	NumOption: cli.AtLeast(1),
}

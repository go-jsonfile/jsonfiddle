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
	Compact bool        `cli:"c,compact" usage:"Compact JSON data, remove all whitespaces"`
	Prefix  string      `cli:"prefix" usage:"prefix for json string output"`
	Indent  string      `cli:"d,indent" usage:"indent for json string output" dft:" "`
	Verbose cli.Counter `cli:"v,verbose" usage:"Verbose mode (Multiple -v options increase the verbosity.)"`
}

var root = &cli.Command{
	Name:   "jsonfiddle",
	Desc:   "JSON Fiddling\nVersion " + version + " built on " + date,
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
//          version   = "0.1.0"
//          date = "2017-08-14"
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
//  		cli.Tree(escDef),
//  		cli.Tree(fmtDef),
//  		cli.Tree(sortDef),
//  		cli.Tree(j2sDef)).Run(os.Args[1:]); err != nil {
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
// esc

//  func escCLI(ctx *cli.Context) error {
//  	rootArgv = ctx.RootArgv().(*rootT)
//  	argv := ctx.Argv().(*escT)
//  	fmt.Printf("[esc]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
//  	return nil
//  }

type escT struct {
	Filei *clix.Reader `cli:"*i,input" usage:"the source to get json string from (mandatory)"`
	Fileo *clix.Writer `cli:"o,output" usage:"the output (default: stdout)"`
}

var escDef = &cli.Command{
	Name: "esc",
	Desc: "Escape json string",

	Argv: func() interface{} { return new(escT) },
	Fn:   escCLI,

	NumOption: cli.AtLeast(1),
}

////////////////////////////////////////////////////////////////////////////
// fmt

//  func fmtCLI(ctx *cli.Context) error {
//  	rootArgv = ctx.RootArgv().(*rootT)
//  	argv := ctx.Argv().(*fmtT)
//  	fmt.Printf("[fmt]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
//  	return nil
//  }

type fmtT struct {
	Filei *clix.Reader `cli:"*i,input" usage:"the source to get json string from (mandatory)"`
	Fileo *clix.Writer `cli:"o,output" usage:"the output (default: stdout)"`
}

var fmtDef = &cli.Command{
	Name: "fmt",
	Desc: "Format json string",

	Argv: func() interface{} { return new(fmtT) },
	Fn:   fmtCLI,

	NumOption: cli.AtLeast(1),
}

////////////////////////////////////////////////////////////////////////////
// sort

//  func sortCLI(ctx *cli.Context) error {
//  	rootArgv = ctx.RootArgv().(*rootT)
//  	argv := ctx.Argv().(*sortT)
//  	fmt.Printf("[sort]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
//  	return nil
//  }

type sortT struct {
	Filei *clix.Reader `cli:"*i,input" usage:"the source to get json string from (mandatory)"`
	Fileo *clix.Writer `cli:"o,output" usage:"the output (default: stdout)"`
}

var sortDef = &cli.Command{
	Name: "sort",
	Desc: "Sort json fields recursively",

	Argv: func() interface{} { return new(sortT) },
	Fn:   sortCLI,

	NumOption: cli.AtLeast(1),
}

////////////////////////////////////////////////////////////////////////////
// j2s

//  func j2sCLI(ctx *cli.Context) error {
//  	rootArgv = ctx.RootArgv().(*rootT)
//  	argv := ctx.Argv().(*j2sT)
//  	fmt.Printf("[j2s]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
//  	return nil
//  }

type j2sT struct {
	FmtType   string       `cli:"f,fmt" usage:"the structural format of the input data (json/yaml)" dft:"json"`
	Filei     *clix.Reader `cli:"*i,input" usage:"the source of the input JSON (mandatory)"`
	Fileo     *clix.Writer `cli:"o,output" usage:"the output (default: stdout)"`
	Name      string       `cli:"name" usage:"the name of the root struct (default: as input file name)"`
	Pkg       string       `cli:"pkg" usage:"the name of the package for the generated code" dft:"main"`
	SubStruct bool         `cli:"subStruct" usage:"create types for sub-structs"`
}

var j2sDef = &cli.Command{
	Name: "j2s",
	Desc: "JSON to struct",

	Argv: func() interface{} { return new(j2sT) },
	Fn:   j2sCLI,

	NumOption: cli.AtLeast(1),
}

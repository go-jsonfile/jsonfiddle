////////////////////////////////////////////////////////////////////////////
// Program: jsonfiddle
// Purpose: JSON Fiddling
// Authors: Tong Sun (c) 2019, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	//  	"fmt"
	//  	"os"

	"github.com/mkideal/cli"
	//  	"github.com/mkideal/cli/clis"
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
	Protect bool        `cli:"p,protect" usage:"protect {{TEMPLATE}} in JSON data"`
	Verbose cli.Counter `cli:"v,verbose" usage:"Verbose mode (Multiple -v options increase the verbosity.)"`
}

var root = &cli.Command{
	Name: "jsonfiddle",
	Desc: "JSON Fiddling\nVersion " + version + " built on " + date +
		"\nCopyright (C) 2019, Tong Sun",
	Text:   "Tool to fiddle with json strings",
	Global: true,
	Argv:   func() interface{} { return new(rootT) },
	Fn:     jsonfiddle,

	NumArg: cli.AtLeast(1),
}

// Template for main starts here
////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The OptsT type defines all the configurable options from cli.
//  type OptsT struct {
//  	Compact	bool
//  	Prefix	string
//  	Indent	string
//  	Protect	bool
//  	Verbose	cli.Counter
//  	Verbose int
//  }

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

//  var (
//          progname  = "jsonfiddle"
//          version   = "0.1.0"
//          date = "2019-05-30"

//  	rootArgv *rootT
//  	// Opts store all the configurable options
//  	Opts OptsT
//  )

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
//  func main() {
//  	cli.SetUsageStyle(cli.DenseNormalStyle) // left-right, for up-down, use ManualStyle
//  	//NOTE: You can set any writer implements io.Writer
//  	// default writer is os.Stdout
//  	if err := cli.Root(root,
//  		cli.Tree(escDef),
//  		cli.Tree(fmtDef),
//  		cli.Tree(sortDef),
//  		cli.Tree(j2sDef),
//  		cli.Tree(x2jDef)).Run(os.Args[1:]); err != nil {
//  		fmt.Fprintln(os.Stderr, err)
//  		os.Exit(1)
//  	}
//  	fmt.Println("")
//  }

// Template for main dispatcher starts here
//==========================================================================
// Dumb root handler

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
//  	clis.Setup(progname, rootArgv.Verbose.Value())
//  	clis.Verbose(2, "<%s> -\n  %+v\n  %+v\n  %v\n", ctx.Path(), rootArgv, argv, ctx.Args())
//  	Opts.Compact, Opts.Prefix, Opts.Indent, Opts.Protect, Opts.Verbose, Opts.Verbose =
//  		rootArgv.Compact, rootArgv.Prefix, rootArgv.Indent, rootArgv.Protect, rootArgv.Verbose, rootArgv.Verbose.Value()
//  	// argv.Filei, argv.Fileo,
//  	//return nil
//  	return DoEsc()
//  }
//
// DoEsc implements the business logic of command `esc`
//  func DoEsc() error {
//  	fmt.Fprintf(os.Stderr, "%s v%s. esc - Escape json string\n", progname, version)
//  	// fmt.Fprintf(os.Stderr, "Copyright (C) 2019, Tong Sun\n\n")
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
//  	clis.Setup(progname, rootArgv.Verbose.Value())
//  	clis.Verbose(2, "<%s> -\n  %+v\n  %+v\n  %v\n", ctx.Path(), rootArgv, argv, ctx.Args())
//  	Opts.Compact, Opts.Prefix, Opts.Indent, Opts.Protect, Opts.Verbose, Opts.Verbose =
//  		rootArgv.Compact, rootArgv.Prefix, rootArgv.Indent, rootArgv.Protect, rootArgv.Verbose, rootArgv.Verbose.Value()
//  	// argv.Filei, argv.Fileo,
//  	//return nil
//  	return DoFmt()
//  }
//
// DoFmt implements the business logic of command `fmt`
//  func DoFmt() error {
//  	fmt.Fprintf(os.Stderr, "%s v%s. fmt - Format json string\n", progname, version)
//  	// fmt.Fprintf(os.Stderr, "Copyright (C) 2019, Tong Sun\n\n")
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
//  	clis.Setup(progname, rootArgv.Verbose.Value())
//  	clis.Verbose(2, "<%s> -\n  %+v\n  %+v\n  %v\n", ctx.Path(), rootArgv, argv, ctx.Args())
//  	Opts.Compact, Opts.Prefix, Opts.Indent, Opts.Protect, Opts.Verbose, Opts.Verbose =
//  		rootArgv.Compact, rootArgv.Prefix, rootArgv.Indent, rootArgv.Protect, rootArgv.Verbose, rootArgv.Verbose.Value()
//  	// argv.Filei, argv.Fileo,
//  	//return nil
//  	return DoSort()
//  }
//
// DoSort implements the business logic of command `sort`
//  func DoSort() error {
//  	fmt.Fprintf(os.Stderr, "%s v%s. sort - Sort json fields recursively\n", progname, version)
//  	// fmt.Fprintf(os.Stderr, "Copyright (C) 2019, Tong Sun\n\n")
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
//  	clis.Setup(progname, rootArgv.Verbose.Value())
//  	clis.Verbose(2, "<%s> -\n  %+v\n  %+v\n  %v\n", ctx.Path(), rootArgv, argv, ctx.Args())
//  	Opts.Compact, Opts.Prefix, Opts.Indent, Opts.Protect, Opts.Verbose, Opts.Verbose =
//  		rootArgv.Compact, rootArgv.Prefix, rootArgv.Indent, rootArgv.Protect, rootArgv.Verbose, rootArgv.Verbose.Value()
//  	// argv.FmtType, argv.Filei, argv.Fileo, argv.Name, argv.Pkg, argv.SubStruct,
//  	//return nil
//  	return DoJ2s()
//  }
//
// DoJ2s implements the business logic of command `j2s`
//  func DoJ2s() error {
//  	fmt.Fprintf(os.Stderr, "%s v%s. j2s - JSON to struct\n", progname, version)
//  	// fmt.Fprintf(os.Stderr, "Copyright (C) 2019, Tong Sun\n\n")
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

////////////////////////////////////////////////////////////////////////////
// x2j

//  func x2jCLI(ctx *cli.Context) error {
//  	rootArgv = ctx.RootArgv().(*rootT)
//  	argv := ctx.Argv().(*x2jT)
//  	clis.Setup(progname, rootArgv.Verbose.Value())
//  	clis.Verbose(2, "<%s> -\n  %+v\n  %+v\n  %v\n", ctx.Path(), rootArgv, argv, ctx.Args())
//  	Opts.Compact, Opts.Prefix, Opts.Indent, Opts.Protect, Opts.Verbose, Opts.Verbose =
//  		rootArgv.Compact, rootArgv.Prefix, rootArgv.Indent, rootArgv.Protect, rootArgv.Verbose, rootArgv.Verbose.Value()
//  	// argv.Filei, argv.Fileo,
//  	//return nil
//  	return DoX2j()
//  }
//
// DoX2j implements the business logic of command `x2j`
//  func DoX2j() error {
//  	fmt.Fprintf(os.Stderr, "%s v%s. x2j - XML to JSON\n", progname, version)
//  	// fmt.Fprintf(os.Stderr, "Copyright (C) 2019, Tong Sun\n\n")
//  	return nil
//  }

type x2jT struct {
	Filei *clix.Reader `cli:"*i,input" usage:"the source of the input JSON (mandatory)"`
	Fileo *clix.Writer `cli:"o,output" usage:"the output (default: stdout)"`
}

var x2jDef = &cli.Command{
	Name: "x2j",
	Desc: "XML to JSON",

	Argv: func() interface{} { return new(x2jT) },
	Fn:   x2jCLI,

	NumOption: cli.AtLeast(1),
}

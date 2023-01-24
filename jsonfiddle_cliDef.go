// jsonfiddle - JSON Fiddling
//
// Tool to fiddle with json strings

package main

////////////////////////////////////////////////////////////////////////////
// Program: jsonfiddle
// Purpose: JSON Fiddling
// Authors: Tong Sun (c) 2017-2023, All rights reserved
////////////////////////////////////////////////////////////////////////////

import (
//  	"fmt"
//  	"os"

// "github.com/go-easygen/go-flags"
)

// Template for main starts here

//////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

//  var (
//          progname  = "jsonfiddle"
//          version   = "0.1.0"
//          date = "2023-01-23"

//  	// opts store all the configurable options
//  	opts optsT
//  )
//
//  var parser = flags.NewParser(&opts, flags.Default)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
//  func main() {
//  	opts.Version = showVersion
//  	opts.Verbflg = func() {
//  		opts.Verbose++
//  	}
//
//  	if _, err := parser.Parse(); err != nil {
//  		fmt.Println()
//  		parser.WriteHelp(os.Stdout)
//  		os.Exit(1)
//  	}
//  	fmt.Println()
//  	//DoJsonfiddle()
//  }
//
//  func showVersion() {
//   	fmt.Fprintf(os.Stderr, "jsonfiddle - JSON Fiddling, version %s\n", version)
//  	fmt.Fprintf(os.Stderr, "Built on %s\n", date)
//   	fmt.Fprintf(os.Stderr, "Copyright (C) 2017-2023, Tong Sun\n\n")
//  	fmt.Fprintf(os.Stderr, "Tool to fiddle with json strings\n")
//  	os.Exit(0)
//  }
// Template for main ends here

// DoJsonfiddle implements the business logic of command `jsonfiddle`
//  func DoJsonfiddle() error {
//  	return nil
//  }

// Template for type define starts here

// The optsT type defines all the configurable options from cli.
type optsT struct {
	Compact bool   `short:"c" long:"compact" description:"Compact JSON data, remove all whitespaces"`
	Prefix  string `long:"prefix" description:"prefix for json string output"`
	Indent  string `short:"d" long:"indent" description:"indent for json string output" default:" "`
	Protect bool   `short:"p" long:"protect" description:"protect {{TEMPLATE}} in JSON data"`
	Verbflg func() `short:"v" long:"verbose" description:"Verbose mode (Multiple -v options increase the verbosity)"`
	Verbose int
	Version func() `short:"V" long:"version" description:"Show program version and exit"`
}

// Template for type define ends here

// Template for "esc" CLI handling starts here
////////////////////////////////////////////////////////////////////////////
// Program: jsonfiddle
// Purpose: JSON Fiddling
// Authors: Tong Sun (c) 2017-2023, All rights reserved
////////////////////////////////////////////////////////////////////////////

//  package main

//  import (
//  	"fmt"
//  	"os"
//
//  	"github.com/go-easygen/go-flags/clis"
//  )

// *** Sub-command: esc ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The EscCommand type defines all the configurable options from cli.
//  type EscCommand struct {
//  	Filei	string	`short:"i" long:"input" description:"the source to get json string from (or \"-\" for stdin) (mandatory)" required:"true"`
//  	Fileo	string	`short:"o" long:"output" description:"the output, default to stdout" default:"-"`
//  }

//
//  var escCommand EscCommand
//
//  func init() {
//  	parser.AddCommand("esc",
//  		"Escape json string",
//  		"",
//  		&escCommand)
//  }
//
//  func (x *EscCommand) Execute(args []string) error {
//   	fmt.Fprintf(os.Stderr, "Escape json string\n")
//   	// fmt.Fprintf(os.Stderr, "Copyright (C) 2017-2023, Tong Sun\n\n")
//   	clis.Setup("jsonfiddle::esc", opts.Verbose)
//   	clis.Verbose(1, "Doing Esc, with %+v, %+v", opts, args)
//   	fmt.Println(x.Filei, x.Fileo)
//  	return x.Exec(args)
//  }
//
// Exec implements the business logic of command `esc`
// func (x *EscCommand) Exec(args []string) error {
// 	// err := ...
// 	// clis.WarnOn("Esc, Exec", err)
// 	// or,
// 	// clis.AbortOn("Esc, Exec", err)
// 	return nil
// }
// Template for "esc" CLI handling ends here

// Template for "fmt" CLI handling starts here
////////////////////////////////////////////////////////////////////////////
// Program: jsonfiddle
// Purpose: JSON Fiddling
// Authors: Tong Sun (c) 2017-2023, All rights reserved
////////////////////////////////////////////////////////////////////////////

//  package main

//  import (
//  	"fmt"
//  	"os"
//
//  	"github.com/go-easygen/go-flags/clis"
//  )

// *** Sub-command: fmt ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The FmtCommand type defines all the configurable options from cli.
//  type FmtCommand struct {
//  	Filei	string	`short:"i" long:"input" description:"the source to get json string from (mandatory)" required:"true"`
//  	Fileo	string	`short:"o" long:"output" description:"the output, default to stdout" default:"-"`
//  }

//
//  var fmtCommand FmtCommand
//
//  func init() {
//  	parser.AddCommand("fmt",
//  		"Format json string",
//  		"",
//  		&fmtCommand)
//  }
//
//  func (x *FmtCommand) Execute(args []string) error {
//   	fmt.Fprintf(os.Stderr, "Format json string\n")
//   	// fmt.Fprintf(os.Stderr, "Copyright (C) 2017-2023, Tong Sun\n\n")
//   	clis.Setup("jsonfiddle::fmt", opts.Verbose)
//   	clis.Verbose(1, "Doing Fmt, with %+v, %+v", opts, args)
//   	fmt.Println(x.Filei, x.Fileo)
//  	return x.Exec(args)
//  }
//
// Exec implements the business logic of command `fmt`
// func (x *FmtCommand) Exec(args []string) error {
// 	// err := ...
// 	// clis.WarnOn("Fmt, Exec", err)
// 	// or,
// 	// clis.AbortOn("Fmt, Exec", err)
// 	return nil
// }
// Template for "fmt" CLI handling ends here

// Template for "sort" CLI handling starts here
////////////////////////////////////////////////////////////////////////////
// Program: jsonfiddle
// Purpose: JSON Fiddling
// Authors: Tong Sun (c) 2017-2023, All rights reserved
////////////////////////////////////////////////////////////////////////////

//  package main

//  import (
//  	"fmt"
//  	"os"
//
//  	"github.com/go-easygen/go-flags/clis"
//  )

// *** Sub-command: sort ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The SortCommand type defines all the configurable options from cli.
//  type SortCommand struct {
//  	Filei	string	`short:"i" long:"input" description:"the source to get json string from (mandatory)" required:"true"`
//  	Fileo	string	`short:"o" long:"output" description:"the output, default to stdout" default:"-"`
//  }

//
//  var sortCommand SortCommand
//
//  func init() {
//  	parser.AddCommand("sort",
//  		"Sort json fields recursively",
//  		"",
//  		&sortCommand)
//  }
//
//  func (x *SortCommand) Execute(args []string) error {
//   	fmt.Fprintf(os.Stderr, "Sort json fields recursively\n")
//   	// fmt.Fprintf(os.Stderr, "Copyright (C) 2017-2023, Tong Sun\n\n")
//   	clis.Setup("jsonfiddle::sort", opts.Verbose)
//   	clis.Verbose(1, "Doing Sort, with %+v, %+v", opts, args)
//   	fmt.Println(x.Filei, x.Fileo)
//  	return x.Exec(args)
//  }
//
// Exec implements the business logic of command `sort`
// func (x *SortCommand) Exec(args []string) error {
// 	// err := ...
// 	// clis.WarnOn("Sort, Exec", err)
// 	// or,
// 	// clis.AbortOn("Sort, Exec", err)
// 	return nil
// }
// Template for "sort" CLI handling ends here

// Template for "j2s" CLI handling starts here
////////////////////////////////////////////////////////////////////////////
// Program: jsonfiddle
// Purpose: JSON Fiddling
// Authors: Tong Sun (c) 2017-2023, All rights reserved
////////////////////////////////////////////////////////////////////////////

//  package main

//  import (
//  	"fmt"
//  	"os"
//
//  	"github.com/go-easygen/go-flags/clis"
//  )

// *** Sub-command: j2s ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The J2sCommand type defines all the configurable options from cli.
//  type J2sCommand struct {
//  	FmtType	string	`short:"f" long:"fmt" description:"the structural format of the input data (json/yaml)" default:"json"`
//  	Filei	string	`short:"i" long:"input" description:"the source of the input JSON (mandatory)" required:"true"`
//  	Fileo	string	`short:"o" long:"output" description:"the output, default to stdout" default:"-"`
//  	Name	string	`long:"name" description:"the name of the root struct (default: as input file name)"`
//  	Pkg	string	`long:"pkg" description:"the name of the package for the generated code" default:"main"`
//  	SubStruct	bool	`long:"subStruct" description:"create types for sub-structs"`
//  }

//
//  var j2sCommand J2sCommand
//
//  func init() {
//  	parser.AddCommand("j2s",
//  		"JSON to struct",
//  		"JSON convert to Go struct",
//  		&j2sCommand)
//  }
//
//  func (x *J2sCommand) Execute(args []string) error {
//   	fmt.Fprintf(os.Stderr, "JSON to struct\n")
//   	// fmt.Fprintf(os.Stderr, "Copyright (C) 2017-2023, Tong Sun\n\n")
//   	clis.Setup("jsonfiddle::j2s", opts.Verbose)
//   	clis.Verbose(1, "Doing J2s, with %+v, %+v", opts, args)
//   	fmt.Println(x.FmtType, x.Filei, x.Fileo, x.Name, x.Pkg, x.SubStruct)
//  	return x.Exec(args)
//  }
//
// Exec implements the business logic of command `j2s`
// func (x *J2sCommand) Exec(args []string) error {
// 	// err := ...
// 	// clis.WarnOn("J2s, Exec", err)
// 	// or,
// 	// clis.AbortOn("J2s, Exec", err)
// 	return nil
// }
// Template for "j2s" CLI handling ends here

// Template for "x2j" CLI handling starts here
////////////////////////////////////////////////////////////////////////////
// Program: jsonfiddle
// Purpose: JSON Fiddling
// Authors: Tong Sun (c) 2017-2023, All rights reserved
////////////////////////////////////////////////////////////////////////////

//  package main

//  import (
//  	"fmt"
//  	"os"
//
//  	"github.com/go-easygen/go-flags/clis"
//  )

// *** Sub-command: x2j ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The X2jCommand type defines all the configurable options from cli.
//  type X2jCommand struct {
//  	Filei	string	`short:"i" long:"input" description:"the source of the input JSON (mandatory)" required:"true"`
//  	Fileo	string	`short:"o" long:"output" description:"the output, default to stdout" default:"-"`
//  }

//
//  var x2jCommand X2jCommand
//
//  func init() {
//  	parser.AddCommand("x2j",
//  		"XML to JSON",
//  		"",
//  		&x2jCommand)
//  }
//
//  func (x *X2jCommand) Execute(args []string) error {
//   	fmt.Fprintf(os.Stderr, "XML to JSON\n")
//   	// fmt.Fprintf(os.Stderr, "Copyright (C) 2017-2023, Tong Sun\n\n")
//   	clis.Setup("jsonfiddle::x2j", opts.Verbose)
//   	clis.Verbose(1, "Doing X2j, with %+v, %+v", opts, args)
//   	fmt.Println(x.Filei, x.Fileo)
//  	return x.Exec(args)
//  }
//
// Exec implements the business logic of command `x2j`
// func (x *X2jCommand) Exec(args []string) error {
// 	// err := ...
// 	// clis.WarnOn("X2j, Exec", err)
// 	// or,
// 	// clis.AbortOn("X2j, Exec", err)
// 	return nil
// }
// Template for "x2j" CLI handling ends here

////////////////////////////////////////////////////////////////////////////
// Program: jf
// Purpose: JSON Fiddling
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

//go:generate sh -v jsonfiddleCLIGen.sh

import (
	"fmt"
	"os"

	"github.com/mkideal/cli"
)

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	progname  = "jf"
	VERSION   = "0.1.0"
	buildTime = "2017-07-13"
)

var rootArgv *rootT

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {
	//NOTE: You can set any writer implements io.Writer
	// default writer is os.Stdout
	if err := cli.Root(root,
		cli.Tree(sortDef)).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println("")
}

//==========================================================================
// Main dispatcher

func jf(ctx *cli.Context) error {
	ctx.JSON(ctx.RootArgv())
	ctx.JSON(ctx.Argv())
	fmt.Println()

	return nil
}

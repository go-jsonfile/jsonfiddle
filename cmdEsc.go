////////////////////////////////////////////////////////////////////////////
// Program: jsonfiddle
// Purpose: JSON Fiddling
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"

	"github.com/mkideal/cli"
)

func escCLI(ctx *cli.Context) error {
	rootArgv = ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*escT)
	fmt.Printf("[esc]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
	return nil
}

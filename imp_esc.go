////////////////////////////////////////////////////////////////////////////
// Program: jsonfiddle
// Purpose: JSON Fiddling
// Authors: Tong Sun (c) 2017-2023, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-easygen/go-flags/clis"
)

// *** Sub-command: esc ***
// Exec implements the business logic of command `esc`
func (x *EscCommand) Exec(args []string) error {
	data := clis.ReadInput(x.Filei)

	out, err := json.Marshal(string(data))
	clis.AbortOn("Formatting input", err)

	fileO := clis.GetOutputStream(x.Fileo)
	defer fileO.Close()

	fmt.Fprintf(fileO, "%s", out)
	return nil
}

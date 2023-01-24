////////////////////////////////////////////////////////////////////////////
// Program: jsonfiddle
// Purpose: JSON Fiddling
// Authors: Tong Sun (c) 2017-2023, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-easygen/go-flags/clis"
)

// *** Sub-command: sort ***
// Exec implements the business logic of command `sort`
func (x *SortCommand) Exec(args []string) error {
	fileI := clis.GetInputStream(x.Filei)
	defer fileI.Close()
	fileO := clis.GetOutputStream(x.Fileo)
	defer fileO.Close()

	return cmdSort(fileI, fileO)
}

//==========================================================================
// support functions

func cmdSort(r io.Reader, w io.Writer) error {
	var res interface{}
	content := readJson(r)
	json.Unmarshal(content, &res)
	var js []byte
	var err error
	if opts.Compact {
		js, err = json.Marshal(res)
	} else {
		js, err = json.MarshalIndent(res, opts.Prefix, opts.Indent)
	}
	clis.AbortOn("[::sort] Marshaling input", err)
	fmt.Fprintln(w, string(js))
	return nil
}

package main_test

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
)

const (
	cmdTest = "jsonfiddle"
	dirTest = "test/"
	extRef  = ".ref" // extension for reference file
	extGot  = ".got" // extension for generated file
)

// testIt runs @cmdEasyGen with @argv and compares the generated
// output for @name with the corresponding @extRef
func testIt(t *testing.T, name string, argv ...string) {
	var (
		diffOut         bytes.Buffer
		generatedOutput = name + extGot
		cmd             = exec.Command(cmdTest, argv...)
	)

	t.Logf("Testing %s:\n\t%s %s\n", name, cmdTest, strings.Join(argv, " "))

	// open the out file for writing
	outfile, err := os.Create(generatedOutput)
	if err != nil {
		t.Errorf("write error [%s: %s] %s.", name, argv, err)
	}
	defer outfile.Close()
	cmd.Stdout = outfile

	err = cmd.Start()
	if err != nil {
		t.Errorf("start error [%s: %s] %s.", name, argv, err)
	}
	err = cmd.Wait()
	if err != nil {
		t.Errorf("exit error [%s: %s] %s.", name, argv, err)
	}

	cmd = exec.Command("diff", "-U1", name+extRef, generatedOutput)
	cmd.Stdout = &diffOut

	err = cmd.Start()
	if err != nil {
		t.Errorf("start error %s [%s: %s]", err, name, argv)
	}
	err = cmd.Wait()
	if err != nil {
		t.Errorf("cmp error %s [%s: %s]\n%s", err, name, argv, diffOut.String())
	}
	os.Remove(generatedOutput)
}

func TestExec(t *testing.T) {
	os.Chdir(dirTest)

	// == Test Basic Functions
	// -- fmt
	testIt(t, "CustomerC", "fmt", "-c", "-i", "Customer.json")
	testIt(t, "Customer", "fmt", "-i", "Customer.json")
	testIt(t, "Schedules", "fmt", "-i", "Schedules.json")
	// -- sort
	testIt(t, "CustomerSI", "sort", "-i", "Customer.json")
	testIt(t, "CustomerSC", "sort", "-c", "-i", "Customer.json")
	testIt(t, "SchedulesSI", "sort", "-i", "Schedules.json")
	// -- j2s
	testIt(t, "CustomerJ2S", "j2s", "-i", "Customer.json")
	testIt(t, "GoodsJ2S", "j2s", "-i", "Goods.json")
	testIt(t, "MenuItemsJ2S", "j2s", "-i", "MenuItems.json")
	testIt(t, "SmartyStreetsJ2S", "j2s", "-i", "SmartyStreets.json")
	testIt(t, "SchedulesJ2S", "j2s", "-i", "Schedules.json")

	//Test String Functions

	//Test Bigger files
}

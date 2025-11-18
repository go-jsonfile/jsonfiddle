package main

import (
	"testing"
)

// NOTE: The FormatJSONMixed function and the Element/Root structs
// are assumed to be in a file named 'json_formatter.go' in the same directory.

// Test data used for the Object-Wrapped Array case
type testObjectData struct {
	Users []struct {
		ID     int  `json:"id"`
		Active bool `json:"active"`
	} `json:"users"`
}

// Test data used for the Pure Array case
type testArrayData []struct {
	Name string `json:"name"`
	City string `json:"city"`
}

func TestFormatJSONMixed_Object(t *testing.T) {
	// GIVEN a structured Go object that will serialize to {"key": [elements]}
	testData := testObjectData{
		Users: []struct {
			ID     int  `json:"id"`
			Active bool `json:"active"`
		}{
			{ID: 1, Active: true},
			{ID: 2, Active: false},
		},
	}

	// The expected output for this structure
	expected := `{ "users": [
    {"active":true,"id":1},
    {"active":false,"id":2}
  ]
}`

	// WHEN the formatting function is called
	formatted, err := FormatJSONMixed(testData)

	// THEN it should format correctly without error
	if err != nil {
		t.Fatalf("FormatJSONMixed failed with error: %v", err)
	}

	if formatted != expected {
		t.Errorf("Object Test Failed.\nExpected:\n%s\nGot:\n%s", expected, formatted)
	}
}

func TestFormatJSONMixed_Array(t *testing.T) {
	// GIVEN a structured Go array that will serialize to [elements]
	testData := testArrayData{
		{Name: "Alice", City: "NYC"},
		{Name: "Bob", City: "LA"},
		{Name: "Charlie", City: "SF"},
	}

	// The expected output for this structure
	expected := `[
  {"city":"NYC","name":"Alice"},
  {"city":"LA","name":"Bob"},
  {"city":"SF","name":"Charlie"}
]`

	// WHEN the formatting function is called
	formatted, err := FormatJSONMixed(testData)

	// THEN it should format correctly without error
	if err != nil {
		t.Fatalf("FormatJSONMixed failed with error: %v", err)
	}

	if formatted != expected {
		t.Errorf("Array Test Failed.\nExpected:\n%s\nGot:\n%s", expected, formatted)
	}
}

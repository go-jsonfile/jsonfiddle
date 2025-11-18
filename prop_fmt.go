package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// FormatJSONMixed handles two top-level JSON structures:
// 1. A single-key object containing an array: {"key": [elements]}
// 2. A pure array of objects: [elements]
// It formats the output with compact, single-line array elements.
func FormatJSONMixed(input interface{}) (string, error) {
	var rawJSON []byte
	var err error

	// 1. Ensure we have raw JSON bytes to work with.
	switch v := input.(type) {
	case []byte:
		rawJSON = v
	case string:
		rawJSON = []byte(v)
	default:
		rawJSON, err = json.Marshal(v)
		if err != nil {
			return "", fmt.Errorf("failed to marshal input: %v", err)
		}
	}

	// 2. Unmarshal into a generic interface{} to check its type.
	var genericData interface{}
	if err := json.Unmarshal(rawJSON, &genericData); err != nil {
		return "", fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	var sb strings.Builder

	// Determine the array source based on the top-level structure
	switch v := genericData.(type) {
	case map[string]interface{}:
		// --- Case 1: Root is a single-key object ---

		// Validation check ensures only one key exists
		if len(v) != 1 {
			return "", fmt.Errorf("input JSON object must have exactly one root key, found %d", len(v))
		}

		sb.WriteString("{ ")

		// Use a composite loop/assignment to get the single key/value pair
		// because Go maps are unordered and cannot be accessed by index or
		// a known key without knowing the key name first
		var key string
		var value interface{}
		for k, val := range v {
			key = k
			value = val
			break // geet the 1st (and only) pair, see earlier strong validation check
		}

		// Validate the value is an array (slice)
		elementArray, ok := value.([]interface{})
		if !ok {
			return "", fmt.Errorf("value for key %q is not a JSON array", key)
		}

		keyBytes, _ := json.Marshal(key)
		// Write the indented key and array start
		//sb.WriteString("  ") // 2-space indent
		sb.Write(keyBytes)
		sb.WriteString(": [\n")

		// Format the elements
		if err := formatArrayElements(&sb, elementArray, 4); err != nil {
			return "", err
		}

		// Write the closing array and object
		sb.WriteString("  ]\n")
		sb.WriteString("}")

	case []interface{}:
		// --- Case 2: Root is a pure array ---

		sb.WriteString("[\n")

		// Format the elements
		if err := formatArrayElements(&sb, v, 2); err != nil {
			return "", err
		}

		// Write the closing array bracket
		sb.WriteString("]")

	default:
		return "", fmt.Errorf("unsupported root JSON structure. Must be an array object or an array")
	}

	return sb.String(), nil
}

// formatArrayElements contains the core logic for marshalling and indenting elements.
func formatArrayElements(sb *strings.Builder, elements []interface{}, indentSpaces int) error {
	indent := strings.Repeat(" ", indentSpaces)
	numElements := len(elements)

	for i, element := range elements {
		// Marshal the single element (compact, one-line string)
		elemBytes, err := json.Marshal(element)
		if err != nil {
			return fmt.Errorf("failed to marshal element %d: %v", i, err)
		}

		// Write the indented element
		sb.WriteString(indent)
		sb.Write(elemBytes)

		// Add comma and newline
		if i < numElements-1 {
			sb.WriteString(",")
		}
		sb.WriteString("\n")
	}
	return nil
}

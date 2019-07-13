package utils

import (
	"bytes"
	"encoding/json"
)

// PrettyPrintJSON is a helper for easily pretty printing JSON
func PrettyPrintJSON(jsonBytes []byte) *bytes.Buffer {
	buf := new(bytes.Buffer)
	json.Indent(buf, jsonBytes, "", "  ")
	return buf
}

package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
)

// PrettyPrintJSON is a helper for easily pretty printing JSON
func PrettyPrintJSON(jsonBytes []byte) *bytes.Buffer {
	buf := new(bytes.Buffer)
	json.Indent(buf, jsonBytes, "", "  ")
	return buf
}

// PrintSortedMap sorts a map of strings based on the top level keys
func PrintSortedMap(unsorted map[string]string) {
	var sortedKeys []string

	for i := range unsorted {
		sortedKeys = append(sortedKeys, i)
	}

	sort.Strings(sortedKeys)
	for i := range sortedKeys {
		j := sortedKeys[i]
		fmt.Printf("%s. %s\n", j, unsorted[j])
	}
}

package utils_test

import (
	"testing"

	"github.com/jsncmgs1/hue_cli/lib/utils"
)

func TestPrettyPrintJSON(t *testing.T) {
	json := []byte("{\"foo\": \"bar\"}")
	got := utils.PrettyPrintJSON(json)
	expected := "{\n  \"foo\": \"bar\"\n}"
	if got.String() != expected {
		t.Errorf("PrettyPrint did not correctly format, got: \n%s, expected: %s", got, expected)
	}
}

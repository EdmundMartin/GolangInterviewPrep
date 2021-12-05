package ch6_strings

import "testing"

func TestSpreadsheetEncoding(t *testing.T) {
	test := SpreadsheetEncoding("ZZ")

	if test != 702 {
		t.Error("Unexpected result")
	}
}
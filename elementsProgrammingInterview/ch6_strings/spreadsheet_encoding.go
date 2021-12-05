package ch6_strings



func SpreadsheetEncoding(col string) int {
	var result int32
	for _, el := range col {
		result = result * 26 + el - 'A' + 1
	}
	return int(result)
}

package firefly

import "strconv"

func saveParseInt(s string) int {
	i, _ := strconv.ParseInt(s, 10, 32)
	return int(i)
}

func saveParseBool(s string) bool {
	b, _ := strconv.ParseBool(s)
	return b
}

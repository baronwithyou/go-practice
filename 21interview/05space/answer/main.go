package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "ab saDQWE sf"

	fmt.Print(checkAndReplace(s))
}

func checkAndReplace(s string) (string, bool) {
	for _, v := range s {
		if string(v) != " " && !unicode.IsLetter(v) {
			return s, false
		}
	}

	return strings.ReplaceAll(s, " ", "%20"), true
}

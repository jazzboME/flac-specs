package main

import (
	"fmt"
	"unicode"
)

func caseInsensitive(path string) string {
	// cf. https://wenzr.wordpress.com/2018/04/09/go-glob-case-insensitive/

	p := ""
	for _, r := range path {
		if unicode.IsLetter(r) {
			p += fmt.Sprintf("[%c%c]", unicode.ToLower(r), unicode.ToUpper(r))
		} else {
			p += string(r)
		}
	}
	return p
}

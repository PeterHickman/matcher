package matcher

import (
	"strings"
)

func ParsePattern(text string) []string {
	l := ""
	t := []string{}

	for i := 0; i < len(text); i++ {
		c := text[i : i+1]

		if c == "*" || c == "?" {
			if l != "" {
				t = append(t, l)
				l = ""
			}

			if len(t) == 0 {
				// List is empty
				t = append(t, c)
			} else if c != "*" {
				// Character is not "*"
				t = append(t, c)
			} else if t[len(t)-1] != "*" {
				// List is not empty and the character is "*"
				// but the last element of the list is not a "*"
				t = append(t, c)
			} else {
				// Discard the duplicate "*"
			}
		} else {
			l += c
		}
	}

	if l != "" {
		t = append(t, l)
	}

	return t
}

func MatchPattern(pattern []string, text string) bool {
	pos := 0
	wild := false
	tl := len(text)

	for _, sym := range pattern {
		if sym == "?" {
			pos += 1
		} else if sym == "*" {
			wild = true
		} else {
			i := strings.Index(text[pos:tl], sym)
			if i == -1 {
				// Not found
				return false
			} else if i == 0 {
				// Start of the string
				pos += len(sym)
			} else {
				// Beyond the start
				if wild {
					pos += i + len(sym)
				} else {
					return false
				}
			}

			wild = false
		}
	}

	// Star as the last argument captures any and all
	if pattern[len(pattern)-1] == "*" {
		return true
	}

	return pos == tl
}

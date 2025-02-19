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

			t = append(t, c)
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
				if wild == true {
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

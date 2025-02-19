package matcher

import (
	"testing"
)

func compare(t *testing.T, pattern, path string, expected bool) {
	p := ParsePattern(pattern)

	actual := MatchPattern(p, path)

	if expected != actual {
		t.Fatalf("%s did not match %s", pattern, path)
	}
}

func TestExactMatch(t *testing.T) {
	compare(t, "fred.bak", "fred.bak", true)
}

func TestNotSimpleMatch(t *testing.T) {
	compare(t, "fred.bak", "fred.old", false)
}

func TestSlightlyLonger(t *testing.T) {
	compare(t, "fred.bak", "fred.baka", false)
}

func TestSlightlyShorter(t *testing.T) {
	compare(t, "fred.bak", "fred.ba", false)
}

func TestStartsWithStarMatches(t *testing.T) {
	compare(t, "*.bak", "fred.bak", true)
}

func TestStartsWithStarFails(t *testing.T) {
	compare(t, "*.bak", "fred.baka", false)
}

func TestEndsWithStarMatches(t *testing.T) {
	compare(t, "fred.*", "fred.bak", true)
}

func TestEndsWithQuestionMatches(t *testing.T) {
	compare(t, "fred.ba?", "fred.bak", true)
}

func TestEndsWithQuestionFails(t *testing.T) {
	compare(t, "fred.ba?", "fred.baka", false)
}

func TestTwoStars(t *testing.T) {
	compare(t, "**.bak", "fred.bak", true)
}

func TestUnusedStar(t *testing.T) {
	compare(t, "*fred.bak", "fred.bak", true)
}

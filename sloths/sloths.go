package sloths

import (
	"strings"
)

func IsSlothful(s string) bool {
	s = strings.ToLower(s)
	slothsLikeThis := strings.Contains(s, "🌺") &&
		!strings.Contains(s, "🏎️")

	if strings.Contains(s, "sloth") {
		return true
	} else if slothsLikeThis {
		return true
	}
	return false
}

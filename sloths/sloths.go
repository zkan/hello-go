package sloths

import (
	"strings"
)

func IsSlothful(s string) bool {
	s = strings.ToLower(s)

	if strings.Contains(s, "sloth") {
		return true
	}
	return false
}

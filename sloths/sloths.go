package sloths

import (
	"strings"
)

func IsSlothful(s string) bool {
	if strings.Contains(s, "sloth") {
		return true
	}
	return false
}

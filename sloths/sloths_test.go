package sloths

import (
	"testing"
)

func TestIsSlothful(t *testing.T) {
	if IsSlothful("hello, world!") {
		t.Error("hello, world! is not supposed to be slothful")
	}
}

package sloths

import (
	"testing"
)

func TestIsSlothful(t *testing.T) {
	if IsSlothful("hello, world!") {
		t.Error("hello, world! is not supposed to be slothful")
	}

	if !IsSlothful("hello, slothful world!") {
		t.Error("hello, slothful world! is supposed to be slothful")
	}

	if !IsSlothful("Sloths rule!") {
		t.Error("Sloths rule! is supposed to be slothful")
	}

}

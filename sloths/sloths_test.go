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

func TestHibiscusEmoji(t *testing.T) {
	if !IsSlothful("Nothing like an iced hibiscus tea! 🌺") {
		t.Error("Nothing like an iced hibiscus tea! 🌺 " +
			"is supposed to be slothful")
	}

	if IsSlothful("Get your 🌺 flowers! They're going fast! 🏎️") {
		t.Error("Get your 🌺 flowers! They're going fast! 🏎️ " +
			"is not supposed to be slothful")
	}
}

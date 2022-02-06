package utils

import (
	"strings"
	"unicode"
)

// GenerateSlug converts a string into a lowercase dasherized slug
// For example: GenerateSlug("My cool object") returns "my-cool-object"
func GenerateSlug(str string) (slug string) {
	return strings.Map(func(r rune) rune {
		switch {
		case r == ' ', r == '-':
			return '-'
		case r == '_', unicode.IsLetter(r), unicode.IsDigit(r):
			return r
		default:
			return -1
		}
	}, strings.ToLower(strings.TrimSpace(str)))
}

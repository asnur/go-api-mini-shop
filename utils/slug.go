package utils

import "strings"

func SlugString(s string) string {
	// Make text lowercase
	s = strings.ToLower(s)

	// Replace all spaces with hyphens
	s = strings.ReplaceAll(s, " ", "-")

	return s
}

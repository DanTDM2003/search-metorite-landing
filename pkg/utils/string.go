package utils

import (
	"regexp"
	"strings"
)

func Slugify(title string) string {
	// Convert the title to lowercase
	slug := strings.ToLower(title)

	// Replace non-letter or non-digit characters with spaces
	reg := regexp.MustCompile(`[^a-z0-9\s-]`)
	slug = reg.ReplaceAllString(slug, "")

	// Replace sequences of spaces or hyphens with a single hyphen
	slug = strings.ReplaceAll(slug, " ", "-")
	reg = regexp.MustCompile(`-+`)
	slug = reg.ReplaceAllString(slug, "-")

	// Trim hyphens from the beginning and end of the string
	slug = strings.Trim(slug, "-")

	return slug
}

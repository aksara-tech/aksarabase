package regex

import "regexp"

func RemoveMultiString(s string) string {
	pattern := regexp.MustCompile(`\s+`)
	res := pattern.ReplaceAllString(s, " ")
	return res
}

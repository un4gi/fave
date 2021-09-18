package filters

import (
	"strings"
)

const (
	keyAPI = "keyword="
	match  = "&isExactMatch=true"
)

func FilterKeyWords(k string, m bool) string {
	k = strings.ReplaceAll(k, " ", "+")
	var query string
	if m {
		query = keyAPI + k + match
	} else {
		query = keyAPI + k
	}
	return query
}

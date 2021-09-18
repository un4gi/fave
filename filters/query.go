package filters

import "strings"

func QueryBuilder(date, key, cwe, cvss string) string {

	var q []string
	for _, s := range []string{date, key, cwe, cvss} {
		if strings.TrimSpace(s) != "" {
			q = append(q, s)
		}
	}

	return strings.Join(q, "&")
}

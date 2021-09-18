package filters

import "fmt"

const cvssAPI = "cvssV3Severity="

func FilterCVSS(c string) string {
	var query string
	if c == "CRITICAL" || c == "HIGH" || c == "MEDIUM" || c == "LOW" {
		query = cvssAPI + c
	} else {
		fmt.Println("Please enter a valid CVSS severity (CRITICAL, HIGH, MEDIUM, or LOW")
	}
	return query
}

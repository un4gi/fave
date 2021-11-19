package main

import (
	"flag"
	"strings"

	"github.com/un4gi/fave/api"
	"github.com/un4gi/fave/filters"
)

func main() {

	filterCWE := flag.Int("cwe", 0, "Search for CVEs based on a CWE number.")
	filterKey := flag.String("k", "", "Search for CVEs based on a keyword (or words).")
	filterExact := flag.Bool("exact", false, "Return only items matching the exact keyword(s) specified with -k")
	filterDays := flag.Int("fd", 0, "Number of days prior to today to filter") // Maximum of 120 days
	// filterMonths := flag.Int("fm", 0, "Number of months prior to today to filter") // This was made obsolete by a change to the NIST API.
	// filterYears := flag.Int("fy", 0, "Number of years prior to today to filter") // This was made obsolete by a change to the NIST API.
	filterCVSS := flag.String("s", "", "Search for CVEs based on the CVSS V3 severity rating.\n"+
		"(CRITICAL, HIGH, MEDIUM, or LOW)\n")

	flag.Parse()

	var cvss string
	if *filterCVSS != "" {
		cvss = filters.FilterCVSS(strings.ToUpper(*filterCVSS))
	} else {
		cvss = ""
	}

	var cwe string
	if *filterCWE > 0 {
		cwe = filters.FilterCWE(*filterCWE)
	} else {
		cwe = ""
	}

	var key string
	if *filterKey != "" {
		key = filters.FilterKeyWords(*filterKey, *filterExact)
	} else {
		key = ""
	}

	var date string
	if *filterDays != 0 {
		date = filters.FilterDate(*filterDays)
	} else {
		date = ""
	}

	query := filters.QueryBuilder(date, key, cwe, cvss)
	api.QueryAPI(query)

}

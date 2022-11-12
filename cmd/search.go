/*
Copyright © 2022 Tony West

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0
	
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/un4gi/fave/api"
	"github.com/un4gi/fave/filters"
)

var (
	cwe       int
	cweString string
	date      string
	days      int
	exact     bool
	key       string
	severity  string
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search the NIST CVE database using the provided data.",
	Long: banner + `
The search command allows you to search the CVE database while filtering on a number of different things, such as:

	- keyword/words
	- age (maximum of 120 days)
	- CWE ID
	- CVSS Severity
`,
	Run: func(cmd *cobra.Command, args []string) {
		if severity != "" {
			severity = filters.FilterCVSS(strings.ToUpper(severity))
		} else {
			severity = ""
		}

		if cwe > 0 {
			cweString = filters.FilterCWE(cwe)
		} else {
			cweString = ""
		}

		if key != "" {
			key = filters.FilterKeyWords(key, exact)
		} else {
			key = ""
		}

		if days != 0 {
			date = filters.FilterDate(days)
		} else {
			date = ""
		}

		query := filters.QueryBuilder(date, key, cweString, severity)
		api.BriefAPIQuery(query)
	},
}

func init() {
	searchCmd.Flags().BoolVar(&exact, "exact", false, "returns only items matching the exact keyword(s) specified.")
	searchCmd.Flags().IntVarP(&days, "days", "d", 0, "the number of days prior to today to filter (maximum of 120 days).")
	searchCmd.Flags().IntVarP(&cwe, "cwe", "c", 0, "Search for CVEs based on a CWE number.")
	searchCmd.Flags().StringVarP(&key, "key", "k", "", "a word or phrase to search - this is required.")
	searchCmd.Flags().StringVarP(&severity, "severity", "s", "", "filters based on the CVSS v3 severity rating (CRITICAL, HIGH, MEDIUM, or LOW).")
}

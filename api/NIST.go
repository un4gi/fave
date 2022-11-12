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
package api

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/un4gi/fave/models"
	"github.com/un4gi/fave/requests"
)

const (
	// nistAPI    = "https://services.nvd.nist.gov/rest/json/cves/1.0" -- version 1 of the API is being retired in 2023.
	nistAPI    = "https://services.nvd.nist.gov/rest/json/cves/2.0"
	numResults = 1000
)

var (
	cveResults  models.CVEResults
	numPages    uint
	score       float64
	scoreString string
)

func getPagination(url string) uint {

	req := requests.MakeGetRequest(url)

	err := json.Unmarshal(req, &cveResults)
	if err != nil && !strings.Contains(fmt.Sprint(err), "unexpected end of JSON input") {
		fmt.Println("Error:", err)
	}

	bodyString := string(req)

	if len(bodyString) > 0 {
		totalResults := cveResults.TotalResults

		if math.Remainder(float64(totalResults), numResults) >= 0.01 {
			floatPages := float64(totalResults) / numResults
			numPages = uint(floatPages) + 1
		} else {
			numPages = totalResults / numResults
		}
	}
	return numPages
}

func BriefAPIQuery(q string) {
	url := nistAPI + "?" + q
	numPages := getPagination(url)

	for startIndex := uint(0); startIndex <= numPages; startIndex = startIndex + numResults {

		if url != nistAPI+"?" {
			url = url + "&"
		}
		query := url + fmt.Sprintf("startIndex=%d&resultsPerPage=%d", startIndex, numResults)

		req := requests.MakeGetRequest(query)
		err := json.Unmarshal(req, &cveResults)
		if err != nil && !strings.Contains(fmt.Sprint(err), "unexpected end of JSON input") {
			fmt.Println("Error:", err)
		}

		w := tabwriter.NewWriter(os.Stdout, 17, 4, 0, ' ', tabwriter.Debug)

		fmt.Fprintln(w, "CVE ID", "\t", "CVSS", "\t", "Date Published", "\t", "Description")

		for i := range cveResults.Vulnerabilities {
			cveID := string(cveResults.Vulnerabilities[i].Cve.ID)
			description := string(cveResults.Vulnerabilities[i].Cve.Descriptions[0].Value)

			if len(cveResults.Vulnerabilities[i].Cve.Metrics.CvssMetrics) == 0 {
				score = 0
			} else {
				score = cveResults.Vulnerabilities[i].Cve.Metrics.CvssMetrics[0].CvssData.BaseScore
			}

			if len(description) > 120 {
				description = description[0:120] + "..."
			}

			fmt.Fprint(w, string(cveID), "\t", fmt.Sprint(score), "\t", strings.Replace(cveResults.Vulnerabilities[i].Cve.Published, "T", " ", 1)[0:10], "\t", description+"\r\n", "\t")
			w.Flush()
		}
	}
}

func DescribeCVE(cveID string) {
	url := nistAPI + "?cveId=" + cveID

	req := requests.MakeGetRequest(url)

	err := json.Unmarshal(req, &cveResults)
	if err != nil && !strings.Contains(fmt.Sprint(err), "unexpected end of JSON input") {
		fmt.Println("Error:", err)
	}

	if len(cveResults.Vulnerabilities[0].Cve.Metrics.CvssMetrics) == 0 {
		score = 0
		scoreString = "N/A"
	} else {
		score = cveResults.Vulnerabilities[0].Cve.Metrics.CvssMetrics[0].CvssData.BaseScore
		scoreString = fmt.Sprint(score)
	}

	fmt.Println("CVE ID:", cveID)
	fmt.Println("CVSS Score:", scoreString)
	fmt.Println("Date Published:", strings.Replace(cveResults.Vulnerabilities[0].Cve.Published, "T", " ", 1)[0:10])
	fmt.Println("Last Modified:", strings.Replace(cveResults.Vulnerabilities[0].Cve.LastModified, "T", " ", 1)[0:10])
	fmt.Println("\r\nDescription:", cveResults.Vulnerabilities[0].Cve.Descriptions[0].Value)
	fmt.Println("\r\nReferences:")

	for i := range cveResults.Vulnerabilities[0].Cve.References {
		fmt.Println(cveResults.Vulnerabilities[0].Cve.References[i].URL)
	}
}

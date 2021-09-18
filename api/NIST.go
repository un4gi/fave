package api

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/un4gi/fave/config"
	"github.com/un4gi/fave/requests"
)

const (
	nistAPI    = "https://services.nvd.nist.gov/rest/json/cves/1.0"
	numResults = 1000
)

func getPagination(url string) uint {

	req := requests.MakeGetRequest(url)

	var cveResults config.CVEResults
	err := json.Unmarshal(req, &cveResults)
	if err != nil {
		fmt.Println("Error unmarshalling data:", err)
	}

	bodyString := string(req)

	var numPages uint
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

func QueryAPI(q string) {
	url := nistAPI + "?" + q
	numPages := getPagination(url)
	var cveResults config.CVEResults

	for startIndex := uint(0); startIndex <= numPages; startIndex = startIndex + numResults {

		if url != nistAPI+"?" {
			url = url + "&"
		}
		query := url + fmt.Sprintf("startIndex=%d&resultsPerPage=%d", startIndex, numResults)

		req := requests.MakeGetRequest(query)
		err := json.Unmarshal(req, &cveResults)
		if err != nil {
			fmt.Println(err)
		}

		n := 1
		for i := range cveResults.Results.CVEItems {
			cveID := string(cveResults.Results.CVEItems[i].CVE.DataMeta.CVEID)
			description := string(cveResults.Results.CVEItems[i].CVE.Description.DescriptionData[0].Value)
			score := cveResults.Results.CVEItems[i].Impact.BaseMetricV3.CVSSV3.BaseScore

			d := "*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*."
			sd := "*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*"

			if n == 1 {
				fmt.Println(d)
				fmt.Println(sd, "RESULT", fmt.Sprint(i+1), sd)
				fmt.Println(d)
				n--
			} else {
				fmt.Println(sd, "RESULT", fmt.Sprint(i+1), sd)
				fmt.Println(d)
			}
			fmt.Println("\r\nCVE ID:", string(cveID))
			if score >= 0.1 {
				fmt.Println("CVSS V3:", fmt.Sprint(score))
			} else {
				fmt.Println("CVSS V3: N/A")
			}
			fmt.Println("Date Published:", strings.Replace(cveResults.Results.CVEItems[i].PublishedDate, "T", " ", 1))
			fmt.Println("Last Modified:", strings.Replace(cveResults.Results.CVEItems[i].LastModified, "T", " ", 1))
			fmt.Println("Description:", string(description)+"\r\n")
			fmt.Println("References:")
			for j := range cveResults.Results.CVEItems[i].CVE.References.ReferenceData {
				reference := cveResults.Results.CVEItems[i].CVE.References.ReferenceData[j].URL
				fmt.Println(reference)
			}
			fmt.Println(d)
		}
	}
}

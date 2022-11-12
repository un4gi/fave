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
package filters

import (
	"fmt"
	"time"
)

const (
	formatTime = "2006-01-02" // Do not need the second/minute/hour.
	//example of the resulting format: "2019-01-01T00:00:00:000 UTC-05:00"

	modEnd   = "lastModEndDate="
	modStart = "lastModStartDate="
	pubEnd   = "pubEndDate="
	pubStart = "pubStartDate="
)

func FilterDate(d int) string {

	if d < 0 {
		fmt.Println("I can't predict future CVEs :(")
		fmt.Println("Please use positive numbers when filtering date.")
	}

	if d > 120 {
		fmt.Println("The maximum supported amount of days is 120.")
	}
	// TODO: Quick maths to stack multiple queries for larger windows than 120 days

	// Get local time and format it.
	now := time.Now()
	t := now.AddDate(0, 0, -d)
	dt := t.Format(formatTime) + "T00:00:00.000"
	nt := now.Format(formatTime) + "T00:00:00.000"
	query := pubStart + dt + "&" + pubEnd + nt + "&" + modStart + dt + "&" + modEnd + nt
	return query
}

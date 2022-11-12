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

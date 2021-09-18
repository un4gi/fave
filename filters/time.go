package filters

import (
	"fmt"
	"net/url"
	"time"
)

const (
	formatTime = "2006-01-02" // Do not need the second/minute/hour.
	//example of the resulting format: "2019-01-01T00:00:00:000 UTC-05:00"

	modEnd   = "modEndDate="
	modStart = "modStartDate="
	pubEnd   = "pubEndDate="
	pubStart = "pubStartDate="
)

func FilterDate(y int, m int, d int) string {

	if y < 0 || m < 0 || d < 0 {
		fmt.Println("I can't predict future CVEs :(")
		fmt.Println("Please use positive numbers when filtering date.")
	}

	// Get local time and format it.
	t := time.Now()
	t = t.AddDate(-y, -m, -d)
	dt := url.QueryEscape(t.Format(formatTime) + "T00:00:00:000 UTC-05:00")
	query := pubStart + dt + "&" + modStart + dt
	return query
}

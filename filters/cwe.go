package filters

import "fmt"

const cweAPI = "cweId=CWE-"

func FilterCWE(c int) string {
	query := cweAPI + fmt.Sprint(c)
	return query
}

package tushare

import (
	"regexp"
)

// IsDateFormat Check date format YYYYMMDD
func IsDateFormat(dates ...string) bool {
	re := regexp.MustCompile(`^\d{4}\d{2}\d{2}$`)
	for _, date := range dates {
		if date == "" {
			continue
		}
		if ret := re.MatchString(date); !ret {
			return false
		}

	}
	return true
}

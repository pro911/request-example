package comm

import "regexp"

func IsMobile(mobile string) bool {
	result, _ := regexp.MatchString(`^(1[3456789][0-9]\d{9})`, mobile)
	return result
}

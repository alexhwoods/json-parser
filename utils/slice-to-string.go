package utils

import "strconv"

func ConvertSliceToString(slice []interface{}) string {
	var result string
	for _, item := range slice {
		switch v := item.(type) {
		case rune:
			result += string(v)
		case int:
			result += strconv.Itoa(v)
		}
	}
	return result
}

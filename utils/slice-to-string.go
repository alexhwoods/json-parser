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
		case string:
			result += v
		case bool:
			result += strconv.FormatBool(v)
		case nil:
			result += "null"
		}
	}
	return result
}

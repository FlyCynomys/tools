package format

import (
	"strconv"
)

func ToInt(v interface{}) int {
	if v == nil {
		return 0
	}
	num, err := strconv.ParseInt(v.(string), 10, 64)
	if err != nil {
		return 0
	}
	return int(num)
}

func ToInt64(v interface{}) int64 {
	if v == nil {
		return 0
	}
	num, err := strconv.ParseInt(v.(string), 10, 64)
	if err != nil {
		return 0
	}
	return num
}

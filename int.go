package gtype

import "strconv"

func Int64ToString(v int64) string {
	s := strconv.FormatInt(v, 10)
	return s
}

package gtype

import (
	"errors"
	"reflect"
	"strconv"
)

func Int64ToString(v int64) string {
	s := strconv.FormatInt(v, 10)
	return s
}
func toInt64(v interface{}) (int64, error) {
	var r int64
	r = 0
	switch reflect.TypeOf(v).Kind() {
	case reflect.Int64:
		r = v.(int64)
		break
	case reflect.String:
		s := v.(string)
		r, _ = strconv.ParseInt(s, 10, 64)
		break
	default:
		return 0, errors.New("not found relation type")
	}

	return r, nil
}

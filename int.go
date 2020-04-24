package gtype

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func Int64ToString(v int64) string {
	s := strconv.FormatInt(v, 10)
	return s
}
func ToInt64(v interface{}) int64 {
	var r int64
	var err error
	r = 0
	if v == nil {
		return r
	}
	switch reflect.TypeOf(v).Kind() {
	case reflect.Int64:
		r = v.(int64)
		break
	case reflect.String:
		s := v.(string)
		r, _ = strconv.ParseInt(s, 10, 64)
		break
	default:
		err = errors.New("not found relation type")
		return 0
	}
	if err != nil {
		_ = fmt.Errorf(err.Error())
	}
	return r
}
func Float64ToInt64(v float64) int64 {
	y := int64(v)
	return y
}

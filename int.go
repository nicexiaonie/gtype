package gtype

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"unsafe"
)

func Int64ToString(v int64) string {
	s := strconv.FormatInt(v, 10)
	return s
}

func IntTo64(v int) int64 {
	return int64(v)
}

func Int64ToInt(v int64) int {
	r := *(*int)(unsafe.Pointer(&v))
	return r
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
	case reflect.Int:
		r = int64(v.(int))
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

func ToInt(v interface{}) int {
	var r int
	var err error
	if v == nil {
		return 0
	}
	switch reflect.TypeOf(v).Kind() {
	case reflect.String:
		r, _ = strconv.Atoi(v.(string))
		break
	case reflect.Int64:
		iv6 := v.(int64)
		r = *(*int)(unsafe.Pointer(&iv6))
		break
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

func Float64Format(v float64, n int) float64 {
	x := math.Pow(10, float64(n))
	z := int64(v * x)
	v = float64(z) / x
	return v
}
func Float32Format(v float32, n int) float32 {
	x := float32(math.Pow(10, float64(n)))
	z := int32(v * x)
	v = float32(z) / x
	return v
}

package gtype

import (
	"fmt"
	"reflect"
)

type slice struct {
}

var Slice slice

func (*slice) InSlice(array interface{}, v interface{}) bool {
	r := false

	f := reflect.TypeOf(array).Kind()
	fmt.Println(111)
	fmt.Println(fmt.Sprintf("%v", f))
	fmt.Println(111)

	//for _, sv := range slice {
	//	if sv == v {
	//		r = true
	//		break
	//	}
	//}
	return r
}

package main

import (
	"fmt"
	"github.com/nicexiaonie/gtype"
)

func main() {

	var a float32

	a = 78.34543
	fmt.Println(a)
	fmt.Println(gtype.Float32Format(a, 2))

}

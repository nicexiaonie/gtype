package main

import (
	"fmt"
	"github.com/nicexiaonie/gtype"
)

func main() {

	var aa float64
	aa = 345.435

	fmt.Println(gtype.Float64Format(aa, 1))
}

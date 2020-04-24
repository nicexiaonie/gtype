package main

import (
	"fmt"
	"github.com/nicexiaonie/gtype"
)

func main()  {

	a := make(map[string]interface{})

	b :=gtype.ToString(a["a"])

fmt.Println(b)





}

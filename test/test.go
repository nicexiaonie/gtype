package main

import (
	"fmt"
	"github.com/nicexiaonie/gtype"
	"time"
)

func main() {

	a := make(map[string]interface{})
	a["a"] = 23
	a["b"] = time.Now()
	a["c"] = "adfs"

	c := gtype.MapTimeToStringFormat(a, "")

	b := gtype.ToString(c)

	fmt.Println(b)

}

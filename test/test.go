package main

import (
	"fmt"
)

func main()  {

	a := map[int]string{
		1:"a",
		4:"d",
	}

	if a[10] == ""{
		fmt.Print(111)
	}
fmt.Println(a[10])

}

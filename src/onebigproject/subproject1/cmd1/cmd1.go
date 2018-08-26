package main

import (
	"fmt"
	"onebigproject/commonpackages/stringutil"
)

func main() {
	s := "Hello world !!"
	s = stringutil.Reverse(s)

	fmt.Println(s)
}

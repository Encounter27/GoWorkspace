package main

import (
	"fmt"
	"project1/pkg/com/company/projectname/domain/stringutil"
)

func main() {
	s := "Hello world !!"
	s = stringutil.Reverse(s)

	fmt.Println(s)
}

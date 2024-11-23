package main

import (
	"fmt"
)

func main() {
	var p = new([]int)
	fmt.Println("p=", p)
	fmt.Println("*p=", *p)
	//*p = 1
	//fmt.Println("p=", p)
	//fmt.Println("*p=", *p)

	var q = make([]int, 5, 10)
	fmt.Println("q=", q)
}

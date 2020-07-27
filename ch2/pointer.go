package main

import (
	"fmt"
)

func f() *int {
	x := 1
	return &x
}

func incr(p *int) int {
	*p++
	return *p
}

func main() {
	p := f()
	fmt.Println(*p)
	fmt.Println(incr(p))
	fmt.Println(*p)
}

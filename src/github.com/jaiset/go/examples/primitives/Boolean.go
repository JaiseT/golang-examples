package main

import (
	"fmt"
)

func main() {

	var n bool = true

	fmt.Printf("%v, %T\n", n, n)

	m := 1 == 1 //return true
	o := 1 == 2 // return false

	fmt.Printf("%v,%T\n", m, m)
	fmt.Printf("%v,%T\n", o, o)

	var p bool //return false by default

	fmt.Printf("%v, %T\n", p, p)
}

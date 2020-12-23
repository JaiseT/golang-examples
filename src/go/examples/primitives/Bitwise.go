package main

import (
	"fmt"
)

func main() {

	a := 10 //1010
	b := 3  //0011

	fmt.Println(a & b)  //AND  0010
	fmt.Println(a | b)  //OR   1011
	fmt.Println(a ^ b)  //NOT  1001
	fmt.Println(a &^ b) //XOR  0100

	c := 8
	fmt.Println(c << 3) //2^3 * 2^3 = 2^6 = 64
	fmt.Println(c >> 3) //2^3 / 2^ 3 = 1

}

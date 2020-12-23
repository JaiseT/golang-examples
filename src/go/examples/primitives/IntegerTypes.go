package main

import (
	"fmt"
)

//Depending on the env, int will be atleast 32 bits long
//Zero value for all numeric type is 0
func main() {

	n := 42 //default int

	fmt.Printf("%v, %T\n", n, n)

	var m int8 = 127 //int16,int32,int64 - signed int
	fmt.Printf("%v, %T\n", m, m)
	var u uint = 42 // unsigned int , uint8,unit16,unit32,unit64
	fmt.Printf("%v, %T\n", u, u)

	fmt.Println("Now do some arithmatic with numbers")

	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b) // return 3. reminder is gone
	fmt.Println(a % b) //reminder 1
	//Call a function to do this
	sum := doCalc(a, b)
	fmt.Println("sum ", sum)

	var x int = 10
	var y int8 = 3

	//fmt.Println(x + y) //uncomment to see errors on compile
	fmt.Println(x + int(y)) //cast with user consent
}

func doCalc(a int, b int) int {
	fmt.Println(a + b)
	c := a + b
	return c
}

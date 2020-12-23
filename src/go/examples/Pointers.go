package main

import "fmt"

func main() {
	x := 10

	fmt.Println(x)
	fmt.Println(&x)

	//changeVal(x)

	//x is not changed here !
	fmt.Println(x)

	changeVal(&x)

	//x has the new value from fn
	fmt.Println(x)

}

func changeVal(x *int) {
	*x = 7
}

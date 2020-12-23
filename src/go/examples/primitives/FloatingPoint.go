package main

import "fmt"

//IEEE 764 standard
func main() {

	var f float32 = 3.4 //+/- 1.18E-38  -  +/- 3.4E38  //float64
	//f = 13.7e72
	f = 2.1E14
	fmt.Printf("%v,%T\n", f, f)

}

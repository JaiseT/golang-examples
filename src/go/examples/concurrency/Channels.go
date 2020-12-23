package main

import "fmt"

func foo(c chan int, someVal int) {
	c <- someVal * 5
}

func main() {

	fmt.Println("concurrency is fun")

	fooVal := make(chan int)

	go foo(fooVal, 5)
	go foo(fooVal, 3)

	v1 := <-fooVal
	v2 := <-fooVal

	fmt.Println(v1, v2)

}

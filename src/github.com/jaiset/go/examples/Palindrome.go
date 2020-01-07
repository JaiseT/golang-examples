package main

import "fmt"

func main() {
	var number, reminder, temp int

	var reverse int = 0

	fmt.Print("enter a positive integer")
	fmt.Scan(&number)

	temp = number

	for {

		reminder = number % 10
		reverse = reverse*10 + reminder
		number /= 10

		if number == 0 {
			break
		}

	}

	if temp == reverse {
		fmt.Printf("%d is a palindrome", temp)
	} else {
		fmt.Printf("%d is not a palindrome", temp)
	}
}

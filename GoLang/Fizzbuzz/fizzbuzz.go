package main

import "fmt"

// This program will count from 1 to 100 but every factor of 3 prints 'Fizz' and every factor of 5 prints 'Buzz'
// Factors of both 3 and 5 print 'FizzBuzz'
func main() {

	for i := 1; i <= 100; i++ {
		if i%3 != 0 && i%5 != 0 {
			fmt.Printf("%d", i)
		}
		if i%3 == 0 {
			fmt.Printf("Fizz")
		}
		if i%5 == 0 {
			fmt.Printf("Buzz")
		}
		fmt.Printf("\n")
	}
}

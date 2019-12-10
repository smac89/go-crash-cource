package main

import "fmt"

func main() {
	// Long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// Short method
	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	// FizzBuzz
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println(i, "fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "buzz")
		} else {
			fmt.Println(i)
		}
	}
}

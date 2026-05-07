package main

import "fmt"

func main() {

	var numbers [6]int

	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 33

	fmt.Println(numbers)

	fmt.Println(len(numbers))

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}

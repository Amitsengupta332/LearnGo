package main

import "fmt"

func main() {

	// makeCoffee := func() {
	// 	fmt.Println("Making coffee...")
	// }

	// makeCoffee()

	// variable-scope

	makeCoffee := func() {
		coffee := "Espresso"
		fmt.Println("Making coffee...")
		fmt.Println("Coffee is ready:", coffee)
	}

	makeCoffee()

	// fmt.Println(Coffee)

}

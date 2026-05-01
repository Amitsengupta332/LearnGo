package main

import "fmt"

func main() {
	var name string = "Hello";

	//  grouped variable

	var (
		firstName string = "John";
		lastName string = "Doe";
	)

	fmt.Println(firstName);
	fmt.Println(lastName);
	fmt.Println(name);
}
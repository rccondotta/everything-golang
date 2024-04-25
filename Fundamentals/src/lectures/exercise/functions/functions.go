//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:

package main

import "fmt"

//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
func greetPerson(name string) {
	fmt.Println("Hello ", name)
}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
func greet() string {
	return "hello"
}

//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
func add(x, y, z int) int {
	return x + y + z
}

// * Write a function that returns any number
func anyNumber() int {
	return 25
}

// * Write a function that returns any two numbers
func anyTwoNumber() (int, int) {
	return 26, 28
}

//* Add three numbers together using any combination of the existing functions.
//  * Print the result

// * Call every function at least once
func main() {
	greetPerson("Tom")
	fmt.Println(greet())

	number := add(1, 2, 3)
	fmt.Println(number)
	fmt.Println(anyNumber())
	fmt.Println(anyTwoNumber())
}

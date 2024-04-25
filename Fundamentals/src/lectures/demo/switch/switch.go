package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {

	switch p := price(); {

	case p < 2:
		fmt.Println("cheap item")
	case p < 2:
		fmt.Println("moderately priced")
	default:
		fmt.Println("Expensive item")

	}

	ticket := Economy
	switch ticket {

	case Economy:
		fmt.Println("economy item")
	case Business:
		fmt.Println("business priced")
	case FirstClass:
		fmt.Println("first class item")
	default:
		fmt.Println("other item")

	}

}

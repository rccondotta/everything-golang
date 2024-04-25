package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkCleanlinees(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i] // always make a copy

		if room.cleaned {
			fmt.Println(room.name, "is clean")
		} else {
			fmt.Println(room.name, "is dirty")
		}

	}
}

func main() {

	rooms := [...]Room{
		{name: "Office"},
		{name: "Warehouse"},
		{name: "Reception"},
		{name: "Ops"},
	}

	checkCleanlinees(rooms)

	fmt.Println("Performing cleaning...")
	rooms[2].cleaned = true
	rooms[3].cleaned = true

	checkCleanlinees(rooms)

}

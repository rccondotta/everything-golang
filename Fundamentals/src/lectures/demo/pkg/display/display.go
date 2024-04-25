package display

import "fmt"

// Capitalize D to make available outside of package
func Display(msg string) {
	fmt.Println(msg)
}

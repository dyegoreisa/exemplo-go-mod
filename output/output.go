package output

import "fmt"

// Show func
func Show(text string) bool {
	fmt.Printf("Showing: %v\n", text)
	return true
}

package main

import "fmt"

func main() {
	fruits := []string{"Apple", "Banana", "Cherry", "Date"}

	// Iterate through the slice using a traditional for loop
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Index %d: %s\n", i, fruits[i])
	}
}

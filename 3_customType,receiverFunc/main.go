package main

import (
	"fmt"
)

// array: phan tu co dinh
// slice: phan tu co the mo rong

func main() {
	//khoi tao 1 slice
	sliceArray := make([]string, 3)
	sliceArray[0] = "A"
	sliceArray[1] = "B"
	sliceArray[2] = "C"

	// them 1 ptu vao slice
	sliceArray = append(sliceArray, "D")
	fmt.Println(sliceArray)

	// for loop
	for index, item := range sliceArray {
		fmt.Println(index, item)
	}

	decks := cards{"A", "B", "C", "D"}
	decks.print()

}

// func newItem() string {
// 	return "E"
// }

package main

import (
	"fmt"
)

// array: phan tu co dinh
// slice: phan tu co the mo rong

func main() {
	var test string = "test"

	fmt.Println(test)

	//khoi tao 1 array
	// array := [3]int{1, 2, 3}

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

}

// func newItem() string {
// 	return "E"
// }

package main

import "fmt"

type cards []string

func (c cards) print() {
	fmt.Println(c)
}

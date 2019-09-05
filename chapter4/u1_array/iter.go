package main

import "fmt"

type MyType int

const (
	V1 MyType = iota
	V2
	V3
	V4
)

func main() {
	//var s := [...]int{1,2,3}
	s := [...]int{V1: 111, 2: 222, 3: 333, 10: 101010}

	for i, v := range s {
		fmt.Printf("%d : %d\n", i, v)
	}
	println("next")
	for _, v := range s {
		fmt.Printf("%d\n", v)
	}
}

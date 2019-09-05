package main

import "fmt"

func reverse(s []int) {
	println(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

}

func main() {
	//s := [...]string{"qqq", "www", "eee"}
	//
	//qw := s[0:2]
	//for i, v := range qw {
	//	fmt.Printf("%d:%s\n", i, v)
	//}
	//
	//v2 := qw
	//
	//for i, v := range v2 {
	//	fmt.Printf("%d:%s\n", i, v)
	//}

	tsint := []int{1, 2, 3, 4, 5}

	reverse(tsint[:2])
	reverse(tsint[2:])
	reverse(tsint)
	for i, v := range tsint {
		fmt.Printf("%d:%d\n", i, v)
	}
}

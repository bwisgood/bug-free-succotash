package main

import "fmt"

func reverse1(p *[]int) {
	s := *p
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int, k int) {
	lens := len(s)
	res := make([]int, lens)
	for i := 0; i < lens; i++ {
		index := i + k
		if index >= lens {
			index = index - lens
		}
		res[i] = s[index]
	}

}

//func eraseDuplicate(s []int) {
//	k := 0
//	for i:=0;i<
//}

func main() {
	l := []int{1, 2, 3, 4, 5}
	p := &l
	reverse1(p)
	rotate(l, 3)
	fmt.Print(l)
}

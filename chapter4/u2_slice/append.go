package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}

	z[len(x)] = y
	return z

}

func appendSlice(x []int, y ...int) {
	var z []int
	var zlen int
	zlen = len(x) + len(y)

	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
}

func main() {
	s := []int{1, 2, 3}
	s1 := appendInt(s, 4)

	for _, v := range s1 {
		fmt.Printf("%d", v)
	}
	fmt.Printf("cap=%d\t%v\n", cap(s), s)
	fmt.Printf("cap1=%d\t%v\n", cap(s1), s1)

}

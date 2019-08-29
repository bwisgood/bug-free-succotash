package main

import "learn/chapter2/u3_variable"

func main() {
	a := 1
	//println(a)
	change(a)
	//fmt.Println(a)
	println(u3_variable.AbsoluteZeroC)
}

func change(a int) {
	println(a)
	a += 1
	println(a)
}

package main

import (
	"fmt"
	"learn/chapter3/u5_string"
	"strconv"
)

func main() {
	//fmt.Println(u5_string.Basename("a/b/c.go"))
	//fmt.Println(u5_string.Basename2("a/b/c.go"))
	//fmt.Println(u5_string.Comma("123456789"))
	fmt.Println(u5_string.IterComma("123456789"))
	fmt.Println(u5_string.IterComma(string(123456789)))
	fmt.Println(u5_string.IterComma(strconv.Itoa(123456789)))
	//fmt.Println(u5_string.SupportFloatComma("123456789.12345"))
	//fmt.Println(u5_string.SupportFloatComma("+123456789.12345"))
	//fmt.Println(u5_string.SupportFloatComma("-123456789.12345"))
	//fmt.Println(u5_string.HasSameCharacters("北京我爱心", "京爱心我北123"))

	//fmt.Println(u5_string.IntToString([]int{1, 2, 3}))
}

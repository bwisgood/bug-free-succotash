package main

import (
	"fmt"
	"os"
)

func main(){
	var s = os.Args[0]
	fmt.Println(s)
	fmt.Println(os.Args[0])
	// todo 测试echo1 和 echo3的时间差异
}
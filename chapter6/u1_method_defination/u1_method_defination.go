package main

import (
	"fmt"
	"strconv"
)

type A struct {
	string
	name string
}

type student struct {
	id   int
	name string
	sex  bool
}

func (s *student) modifyName(name string) bool {
	s.name = name
	return true
}

func (s student) getFullName() string {
	return strconv.Itoa(s.id) + s.name + strconv.FormatBool(s.sex)
}

func (s student) modifyNameWithoutPtr(name string) bool {
	s.name = name

	fmt.Println(s.name)
	return true
}

func (a A) getName() string {
	return a.name
}

func main() {

	//a := A{"hehe", "baiwei"}
	//println(a.getName())

	s1 := student{id: 1, name: "baiwei", sex: true}
	s2 := student{id: 2, name: "zhouzhengz", sex: false}

	s1.modifyName("bw")
	s2.modifyNameWithoutPtr("zz")

	println(s1.getFullName())
	println(s2.getFullName())
}

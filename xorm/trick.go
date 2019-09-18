package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Account struct {
	Id      int64  `xorm:"pk"`
	Name    string `xorm:"unique"`
	Balance float64
	Version int `xorm:"version"` //乐观锁
}

type Student struct {
	Id       int64  `xorm:"INT(11) autoincr pk"`
	Username string `xorm:"VARCHAR(64)"`
	Address  string `xorm:"VARCHAR(256)"`
}

func main() {

	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "root", "mysql", "localhost:3306", "test")
	engine, err := xorm.NewEngine("mysql", params)
	if err != nil {
		fmt.Println(err)
		return
	}
	engine.ShowSQL(true)
	engine.ShowExecTime(true)

	err = engine.Sync2(new(Student), new(Account))
	if err != nil {
		panic(err)
	}

	st1 := new(Student)
	st1.Username = "baiwei"
	st1.Address = "China"
	affected, err := engine.Insert(st1)
	fmt.Println(affected)

	st2 := new(Student)
	result, err := engine.Where("id=?", 1).Get(st2)
	fmt.Println(result)

	fmt.Println(st2.Username)
	fmt.Println(st2.Address)

	st3 := new(Student)
	st3.Username = "zhang"
	st3.Address = "beijing"
	_, _ = engine.InsertOne(st3)

	results, err := engine.QueryString("select * from student")
	for k, v := range results {
		fmt.Println(k)
		for m, n := range v {
			fmt.Println(m)
			fmt.Println(n)
		}
		fmt.Println("===")
	}
}

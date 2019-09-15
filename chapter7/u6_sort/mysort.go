package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Student struct {
	id   int
	name string
	sex  bool
}

var students = []*Student{
	{id: 1, name: "Tom", sex: true},
	{id: 2, name: "Amy", sex: false},
	{id: 3, name: "John", sex: true},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)

	if err != nil {
		panic(s)
	}
	return d

}

func printStudents(students []*Student) {
	const format = "%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	_, _ = fmt.Fprintf(tw, format, "ID", "Name", "Sex")
	_, _ = fmt.Fprintf(tw, format, "------", "------", "------")

	for _, t := range students {
		_, _ = fmt.Fprintf(tw, format, t.id, t.name, t.sex)
	}

	_ = tw.Flush()
}

type byName []*Student

func (x byName) Len() int           { return len(x) }
func (x byName) Less(i, j int) bool { return x[i].name < x[j].name }
func (x byName) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {
	sort.Sort(byName(students))
	printStudents(students)
	sort.Sort(sort.Reverse(byName(students)))
	fmt.Println()
	printStudents(students)
}

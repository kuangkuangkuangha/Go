package main

import "fmt"

type student struct {
	id   int
	name string
}

func main() {
	var clus []student

	stu1 := student{
		id:   1,
		name: "zhangkuang",
	}

	clus = append(clus, stu1)

	fmt.Println(clus[0])
}

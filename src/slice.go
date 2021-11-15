package main

import "fmt"

func main() {
	var a []int
	a = []int{1, 2, 3}
	a[1] = 10
	fmt.Println(a)

	s := make([]int, 5, 10)
	println(len(s), cap(s))
	fmt.Println(s)

	var e []int
	if e == nil {
		println("Nil Slice")
	}
	println(len(s), cap(s))
}

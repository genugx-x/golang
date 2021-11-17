package main

import "fmt"

type dict struct {
	data map[int]string
}

func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d
}

func main() {
	dic := newDict()
	dic.data[1] = "A"
	fmt.Println(dic)
}

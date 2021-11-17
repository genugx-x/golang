package main

import "fmt"

// type Person struct { 패키지 외부의 사용 경우 Person으로 입력
type person struct {
	name string
	age  int
}

func main() {
	// person 객체 생성
	p := person{}

	p.name = "Kang"
	p.age = 32

	fmt.Println(p)

	// var p1 person
	// p1 = person{"Bob", 20}
	// p2 := person{name: "Sean", age: 50}
	// p3 := person{"Nami", 17}

	p4 := new(person)
	p4.name = "Lee"
	p4.age = 32
	fmt.Println(p4)
	p4 = new(person)
	fmt.Println(p4)
}

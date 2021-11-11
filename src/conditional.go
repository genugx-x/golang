package main

import "fmt"

func conditional() {
	var i = 1
	var n int = 2

	if i == 1 {
		println("One")
	}

	if n == 1 {
		println("One")
	} else if n == 2 {
		println("Two")
	} else {
		println("Other")
	}

	var max int = 13
	if val := i * 2; val < max {
		println("i =", i)
		println("val =", val)
		i++
	} else {
		println("else")
	}

	var name string
	var category = 4
	switch category {
	case 1:
		name = "Paper Book"
	case 2:
		name = "eBook"
	case 3, 4:
		name = "Blog"
	default:
		name = "Other"
	}
	println(name)

	switch x := category << 2; x - 1 {
	case 3:
		println("x = 4? : ", x)
	default:
		println("x =", x)
	}

	grade(68)

	// var v int
	// switch v.(type) {
	// case int:
	// 	println("int")
	// case bool:
	// 	println("bool")
	// case string:
	// 	println("string")
	// default:
	// 	println("unknown")
	// }
	check(1)
}

func check(val int) {
	switch val {
	case 1:
		fmt.Println("1 이하")
		fallthrough
	case 2:
		println("2 이하")
	case 3:
		fmt.Println("3 이하")
		fallthrough
	default:
		println("default 도달")
	}
}

func grade(score int) {
	switch {
	case score >= 90:
		println("A")
	case score >= 80:
		println("B")
	case score >= 70:
		println("C")
	case score >= 60:
		println("D")
	default:
		print("No Hope")
	}
}

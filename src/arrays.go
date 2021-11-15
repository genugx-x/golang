package main

func main() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	println(a[1])

	//var a1 = [3]int{1, 2, 3}
	//var a2 = [...]int{1, 2, 3}

	var multiArray [3][4][5]int
	multiArray[0][1][2] = 10

	var a4 = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	println(a4[1][2])
}

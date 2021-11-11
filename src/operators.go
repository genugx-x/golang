package main

func operators() {
	var a = 5
	var b = 20
	var c = (a + b) / 5
	// i++

	println(a == b, a != c, a >= b)

	/*
		A && B
		A || !(C && B)
	*/

	c = (a & b) << 5
	a = 100
	a *= 10
	a >>= 2
	a |= 1

	var k int = 10
	var p = &k // k의 주소를 할당

	println(p, *p) // p가 가리키는 주소에 있는 실제 내용을 출력

}

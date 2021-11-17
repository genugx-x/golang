package main

type Rect struct {
	width, height int
}

func (r Rect) area() int {
	r.width++
	return r.width * r.height
}

func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}

func main() {
	rect := Rect{10, 20}
	rect2 := Rect{10, 20}
	area := rect.area()    // 메서드 호출
	area2 := rect2.area2() // 메서드 호출
	println(rect.width, area)
	println(rect2.width, area2)
}

package main

func nextValue() func() int {
	i := 0
	println("i :", 0)
	return func() int {
		i++
		return i
	}
}

func main() {
	next := nextValue()
	println("--------")
	println(next())
	println(next())
	println(next())
	println()
	println()
	anotherNext := nextValue()
	println("--------")
	println(anotherNext())
	println(anotherNext())
	println()
	println(next())
	println(next())
}

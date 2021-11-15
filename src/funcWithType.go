package main

func main() {
	calc(calculator, 5, 6)
}

type calculator func(int, int) int

func calc(f calculator, a int, b int) int {
	result := f(a, b)
	return result
}

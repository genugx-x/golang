package main

import "fmt"
import "rsc.io/quote"

func main() {
	fmt.Println("Hello World!")
	fmt.Println(quote.Go())

	var a int = 1
	var f float32 = 11.

	println(a)
	println(f)

	a = 10
	f = 12.0

	fmt.Println(a)
	fmt.Println(f)

	var b int
	println(b)

	var c, d, e int
	c, d, e = 1, 2, 3
	println(c)
	println(d)
	println(e)

}

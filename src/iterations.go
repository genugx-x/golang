package main

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println(sum)

	n := 1
	for n < 100 {
		n *= 2
		print(n, " ")

	}
	println(n)

	t := 1
	for ; t < 100; t *= 3 {
		print(t, " ")
	}
	println()

	// for {
	// 	println("Infinite loop")
	// }

	names := []string{"A", "B", "C"}
	for index, name := range names {
		println(index, name)
	}

	a := 1
	for a < 15 {
		if a == 5 {
			a += a
			continue
		}
		a++
		if a > 10 {
			break
		}
	}
	if a == 11 {
		goto END
	}
	println("a = ", a)

END:
	println("End")

	b := 0
L1:
	for {
		println("for 진입")
		if b == 0 {
			break L1
		}
	}
	println("OK")
}

package main

func main() {
	msg := "Hello"
	say(msg)
	sayWithPointer(&msg)
	println(msg)

	sayWithVariadicFunction("This ", "is ", "a ", "book ")
	sayWithVariadicFunction("Hello World!")

	println("Total :", sum(1, 7, 3, 5, 9))

	count, total := sumWithCount(1, 7, 3, 5, 9)
	println(count, total)

}

func say(msg string) {
	println(msg)
}

func sayWithPointer(msg *string) {
	println(*msg, msg)
	*msg = "Changed"
}

func sayWithVariadicFunction(msg ...string) {
	for _, s := range msg {
		print(s)
	}
	println()
}

func sum(nums ...int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

func sumWithCount(nums ...int) (int, int) {
	s := 0
	count := 0
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}

func sumWithNamedReturnParameters(nums ...int) (count int, total int) {
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	return
}

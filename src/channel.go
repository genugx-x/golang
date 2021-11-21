package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 123
	}()

	// var i int
	t := <-ch
	println(t)

	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		done <- true
	}()
	<-done
}

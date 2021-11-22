package main

import (
	"fmt"
	"time"
)

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
	// b := <-done
	println(<-done)

	// c := make(chan int)
	// c <- 1 // 수신루틴이 없으므로 데드락
	// fmt.Println(<-c) // 코멘트해도 데드락 (별도의 Go루틴이 없기 때문)

	c := make(chan int, 1)
	c <- 101
	fmt.Println(<-c)

	ch2 := make(chan string, 1)
	sendChan(ch2)
	receiveChan(ch2)

	ch3 := make(chan int, 2)
	ch3 <- 1
	ch3 <- 2

	close(ch3)

	println(<-ch3)
	println(<-ch3)

	if _, success := <-ch3; !success {
		println("더이상 데이터 없음.")
	}

	ch4 := make(chan int, 2)
	ch4 <- 1
	ch4 <- 2

	close(ch4)
	/*
		for {
			if i, success := <-ch4; success {
				println(i)
			} else {
				break
			}
		}
	*/

	for i := range ch4 {
		println(i)
	}

	done1 := make(chan bool)
	done2 := make(chan bool)
	go run1(done1)
	go run2(done2)

EXIT:
	for {
		select {
		case <-done1:
			println("run1 완료")
		case <-done2:
			println("run2 완료")
			break EXIT
		}
	}

}

func run1(done chan bool) {
	println("run1 start!")
	time.Sleep(2 * time.Second)
	done <- true
	// 채널에 수신한 이후 아래 코드 동작 안함
	// time.Sleep(2 * time.Second)
	// println("run1 end!")
}

func run2(done chan bool) {
	println("run2 start!")
	time.Sleep(4 * time.Second)
	done <- true
}

func sendChan(ch chan<- string) {
	ch <- "Data"
	// x := <-ch // 에러발생
}

func receiveChan(ch <-chan string) {
	data := <-ch
	fmt.Println(data)
}

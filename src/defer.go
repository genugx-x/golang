package main

import "os"

func main() {
  f, err := os.Open("1.txt")
  if err != nil {
    panic(err)
  }
  // panic 실행 후 defer의 모든 함수가 실행 되는 것이라 생각 했지만 실행 안됨
  defer println("defer test") // 실행 안됨.. 왜?
  defer f.Close() // finally 블럭처럼 마지막에 Clean-up 작업을 위해 사용된다.

  bytes := make([]byte, 1024)
  f.Read(bytes)
  println(len(bytes))
}

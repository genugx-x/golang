package main

import (
  "fmt"
  "sync"
)

func main() {
  // WaitGroup 생성. 2개의 Go루틴을 기다림.
  var wait sync.WaitGroup
  wait.Add(2)

  go func() {
    defer wait.Done()
    fmt.Println("Hello")
  }()

  go func(msg string) {
    defer wait.Done()
    fmt.Println(msg)
  }("Hi")

  wait.Wait()
}

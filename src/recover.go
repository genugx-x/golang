package main

import (
  "os"
  "fmt"
)

func main() {
  openFile("Invalid.txt")

  println("Done")
}

func openFile(filename string) {
  defer func() {
    if r := recover(); r != nil {
      fmt.Println("OPEN ERROR", r)
    }
  }()

  f, err := os.Open(filename)
  if err != nil {
    panic(err)
  }

  defer f.Close()

}

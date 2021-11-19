package main

import "os"

func main() {
  openFile("Invalid.txt")
  println("Done")
}

func openFile(filename string) {
  f, err := os.Open(filename)
  println("Checking error....")
  if err != nil {
    panic(err)
  }

  defer f.Close()
}

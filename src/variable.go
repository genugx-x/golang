package main

func main() {
  var a int
  var f float32 = 11.
  println("before edit : ", a, f);
  a = 10
  f = 12.0

  var x, z, y int
  var i, j, k int = 1, 2, 3
  println("after edit : ", a, f);
  println(x, z, y);
  println(i, j, k);

  var b int

  var c bool
  var d string = "123"
  println(b, c, d);

  var g = 1
  var h ="Hi"
  println(g, h)

  m := 1
  println(m)

  const t int = 10
  const s string = "Hi"

  const (
      Visa = "Visa"
      Master = "MasterCard"
      Amex = "American Express"
  )

  const (
    Apple = iota
    Grape
    Orange
  )

}

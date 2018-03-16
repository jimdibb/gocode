package main

import "fmt"

func add( x int, count int) func(int) (int,int) {
  return func(y int) (int,int) {
    count += 1
    return x+y, count
  }
}

func main(){
  y := add(200, 100) // func y will add 200 to it's parameter and return how
  // many times it's been called, resetting to 100
  z := add(2, 0) // z will add 2 to it's parameter and return it's call count from 0

  c, v := z(1)
  fmt.Println(c,v)
  c, v = z(2)
  fmt.Println(c,v)

  y = add(50,50)
  c,v = y(20)
  fmt.Println(c, v)

}

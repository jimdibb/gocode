package main

import "fmt"

func post(start int, end int, c chan int, sem chan bool) {
	for i := start; i < end; i++ {
		c <- i
	}
	sem <- true
}

func main() {
	c := make(chan int)
	done := make(chan bool)

	go post(32, 1000, c, done)
	go post(2000, 2300, c, done)
	/*
		go func(start, end ) {
			for i := 120 ; i < 90; i++ {
				c <- string(i)
			}
			done <- true
		}()

		go func() {
			for i := 91; i < 120; i++ {
				c <- string(i)
			}
			done <- true
		}()
	*/
	go func() {
		<-done
		<-done
		close(c)
	}()
	for n := range c {
		fmt.Println(n)

	}

}

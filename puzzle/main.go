package main

import "fmt"
import "math/rand"

var odds, evens float32

func compute(lim int, c chan int) {
	tickets := [18]int{3, 3, 3, 3, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1}
	var odds int
	for i := 0; i < lim; i++ {
		tmap := make(map[int]bool)

		sum := 0
		for n := 0; n < 3; n++ {
			for true {
				tick := rand.Intn(18)

				if tmap[tick] == false {
					tmap[tick] = true
					sum += tickets[tick]
					break
				}
			}
		}
		if sum%2 == 1 {
			odds++
		}
	}
	c <- odds
}

func main() {
	rand.Seed(9825245)
	limit := 200000
	procs := 1000
	c := make(chan int)
	for i := 0; i < procs; i++ {
		go compute(limit, c)
	}

	var odds int64
	for i := 0; i < procs; i++ {
		odds += int64(<-c)
	}
	evens := int64(procs*limit) - odds

	fmt.Printf("Odds: %d, Evens: %d", odds, evens)
}

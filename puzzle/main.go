package main

import "fmt"
import "math/rand"

var odds, evens float32

func main() {
	rand.Seed(9825245)
	checkmap := make(map[int]int)
	tickets := [18]int{3, 3, 3, 3, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1}
	for i := 0; i < 200000; i++ {
		tmap := make(map[int]bool)

		sum := 0
		for n := 0; n < 3; n++ {
			for true {
				tick := rand.Intn(18)
				checkmap[tick]++
				if tmap[tick] == false {
					tmap[tick] = true
					sum += tickets[tick]

					break
				}
			}
		}
		fmt.Println(sum)
		if sum%2 != 0 {
			odds++
		} else {
			evens++
		}

	}
	fmt.Print(checkmap)
	p := odds / (odds + evens)
	fmt.Printf("Odds: %f, Evens: %f , Prob %f \n", odds, evens, p)
}

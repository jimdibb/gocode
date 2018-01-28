package main

import "fmt"
import "sort"

type people []string

/*
func (p people) Len () int {
  return len(p)
}

func (p people) Less (i , j int) bool {
  return p[i] < p[j]
}

func (p people) Swap (i, j int) {
  p[i], p[j] = p[j], p[i]

}
*/
// all these functions above are the interface that Sort needs, but
// if you use sort.StringSlice, it provides the interface

func main() {
	var x = 12
	studyGroup := sort.StringSlice(people{"Zeno", "John", "Al", "Jenny"})
	studyGroup2 := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
	fmt.Println(x)
	sort.Sort(sort.Reverse(sort.StringSlice(studyGroup2)))
	fmt.Println(studyGroup2)

}

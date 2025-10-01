package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	s = append(s, 4, 5, 6, 7)
	s2 := s
	fmt.Println(s2)
	t := []int{8, 9}
	s = append(s, t...)
	fmt.Println(s)
	/*fmt.Println("Length: ", len(s))
	fmt.Println("Capacity: ", cap(s))*/
}

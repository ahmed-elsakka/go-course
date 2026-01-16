package main

import "fmt"

func main() {
	str1 := "Go"
	str2 := "Go"
	str3 := "Golang"

	fmt.Println(str1 == str2) // true
	fmt.Println(str1 == str3) // false
	fmt.Println(str1 != str3) // true
}

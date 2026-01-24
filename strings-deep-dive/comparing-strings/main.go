package main

import "fmt"

func main() {
	str1 := "Go"
	str2 := "Go"
	str3 := "Golang"

	fmt.Println(str1 == str2) // true
	fmt.Println(str1 == str3) // false
	fmt.Println(str1 != str3) // true

	fmt.Println("Go" == "go") // false

	str4 := "apple"
	str5 := "apple"

	fmt.Println(str4 <= str5) // true
	fmt.Println(str4 >= str5) // false
}

package main

import "fmt"

func main() {
	str := "Go"

	for i := 0; i < len(str); i++ {
		//fmt.Println(str[i])
		fmt.Printf("%c\n", str[i])
	}

	for i, r := range str {
		fmt.Printf("index=%d rune=%c\n", i, r)
	}

}

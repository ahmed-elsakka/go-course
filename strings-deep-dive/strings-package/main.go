package main

import (
	"fmt"
	"strings"
)

func main() {
	rawInput := "   john.doe@EXAMPLE.com,   Jane DOE, Go Developer  "
	fmt.Println("Original input:")
	fmt.Println(rawInput)

	trimmedInput := strings.TrimSpace(rawInput)
	fmt.Println("After TrimSpace:")
	fmt.Println(trimmedInput)

	parts := strings.Split(trimmedInput, ",")
	fmt.Println("After splitting:")
	for idx, part := range parts {
		fmt.Printf("[%d] %q\n", idx, part)
	}
	fmt.Println()

	for idx, part := range parts {
		part = strings.TrimSpace(part)
		part = strings.ToLower(part)
		parts[idx] = part
	}
	fmt.Println("After trimming and normalizing:")
	fmt.Println(parts)
	fmt.Println()

	email := parts[0]
	if strings.Contains(email, "@") {
		domain := strings.Split(email, "@")[1]
		fmt.Println("Email domain: ", domain)
	}
	fmt.Println()

	final := strings.Join(parts, " | ")
	fmt.Println("Final after joining: ")
	fmt.Println(final)
	fmt.Println()

}

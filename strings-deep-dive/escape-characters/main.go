// Go Strings: Escape Characters Demonstration
// This example shows how escape sequences work in interpreted string literals.

package main

import "fmt"

func main() {
	fmt.Println("=== Escape Characters in Go ===")

	// 1️⃣ Newline (\n)
	fmt.Println("\n1. Newline (\\n):")
	fmt.Println("Hello\nWorld") // Moves to the next line

	// 2️⃣ Tab (\t)
	fmt.Println("\n2. Tab (\\t):")
	fmt.Println("A\tB\tC") // Adds horizontal tabs between words

	// 3️⃣ Double Quote (\")
	fmt.Println("\n3. Double Quote (\\\"):")
	fmt.Println("He said, \"Go is awesome!\"") // Embeds double quotes in a string

	// 4️⃣ Backslash (\\)
	fmt.Println("\n4. Backslash (\\\\):")
	fmt.Println("This is a backslash: \\\\") // Prints a single backslash

	// 5️⃣ Carriage Return (\r)
	fmt.Println("\n5. Carriage Return (\\r):")
	fmt.Print("Hello\rWorld\n") // 'World' overwrites 'Hello'

	// 6️⃣ Unicode Characters (\u and \U)
	fmt.Println("\n6. Unicode characters (\\u and \\U):")
	fmt.Println("Lowercase a: \u0061")       // 'a'
	fmt.Println("Chinese character: \u4F60") // '你'
	fmt.Println("Emoji: \U0001F600")         // 😀 (Unicode code point U+1F600)

	// 7️⃣ Hexadecimal (\xNN)
	fmt.Println("\n7. Hexadecimal byte (\\xNN):")
	fmt.Println("Letter A: \x41") // 0x41 = 65 = 'A'

	// 8️⃣ Raw String (no escapes)
	fmt.Println("\n8. Raw string (using backticks):")
	raw := `This is a raw string:
  \n  (No newline)
  \t  (No tab)
  \u4F60 (No Unicode parsing)`
	fmt.Println(raw)

	// 9️⃣ Combining escapes
	fmt.Println("\n9. Combining escape characters:")
	fmt.Println("Line1\n\tIndented Line2\n\t\tIndented Line3")

	fmt.Println("\n=== End of Demo ===")
}

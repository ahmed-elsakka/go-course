package main

func main() {
	user := User{name: "Sara", age: 21, email: "sara@email.com"}

	user.updateEmail("example@email.com")
	user.process()
}

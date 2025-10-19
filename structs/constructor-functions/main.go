package main

import "fmt"

func main() {
	//user := User{name: "Sara", age: 21, email: "sara@email.com"}
	user, err := NewUser("Sara", 21, "sara@email.com")
	if err != nil {
		fmt.Println("Couldn't create user: ", err)
		return
	}

	if err = user.updateEmail("example@email.com"); err != nil {
		fmt.Println("Couldn't update the email: ", err)
		return
	}
	user.process()
}

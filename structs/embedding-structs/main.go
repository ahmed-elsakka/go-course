package main

import "fmt"

func main() {
	address := Address{street: "123 main st.", city: "City", country: "Country"}
	user, err := NewUser("Sara", 21, "sara@email.com", address)
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

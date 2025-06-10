package auth

import (
	"fgo24-go-auth-flow/utils"
	"fmt"
)

func Register() {
	fmt.Println("=== Register ===")

	newUser:= &utils.User{}

	fmt.Print("Enter your email: ")
	fmt.Scanln(&newUser.Email)

	if utils.IsEmailTaken(newUser.Email) {
		fmt.Println("Email is already registered. Please login or use another email.")
		return
	}

	fmt.Print("Enter your password: ")
	fmt.Scanln(&newUser.Password)

	utils.Users = append(utils.Users, *newUser)

	fmt.Println("Account created.")
}

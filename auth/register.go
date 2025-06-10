package auth

import (
	"fmt"
	"fgo24-go-auth-flow/utils"
)

func Register() {
	fmt.Println("=== Register ===")
	fmt.Print("Enter your email: ")
	fmt.Scanln(&utils.DataUser.Email)

	fmt.Print("Enter your password: ")
	fmt.Scanln(&utils.DataUser.Password)

	fmt.Println("Account created.\n")
}

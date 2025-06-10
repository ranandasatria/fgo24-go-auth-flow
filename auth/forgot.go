package auth

import (
	"fmt"
	"fgo24-go-auth-flow/utils"
)

func Forgot() {
	fmt.Println("=== Forgot Password ===")


	var email string
	fmt.Print("Enter your registered email: ")
	fmt.Scanln(&email)

	found := false
	for i := range utils.Users {
		if utils.Users[i].Email == email {
			found = true

			fmt.Print("Enter your new password: ")
			fmt.Scanln(&utils.Users[i].Password)

			fmt.Println("Password updated.")
		}
	}

	if !found {
		fmt.Println("Email not found. Please register or try again.")
	}
}

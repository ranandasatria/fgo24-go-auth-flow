package auth

import (
	"fgo24-go-auth-flow/utils"
	"fmt"
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
			var newPassword string
			fmt.Scanln(&newPassword)

			utils.Users[i].Password = hashPassword(newPassword)

			fmt.Println("Password updated.")
			break
		}
	}

	if !found {
		fmt.Println("Email not found. Please register or try again.")
	}
}

package auth

import (
	"fmt"
	"os"
	"fgo24-go-auth-flow/utils"
)

func Login() {
	fmt.Println("=== Login ===")

	fmt.Print("Enter your email: ")
	fmt.Scanln(&utils.DataLogin.Email)

	if utils.DataLogin.Email != utils.DataUser.Email {
		fmt.Println("Email not found.")
		fmt.Println("1. Register")
		fmt.Println("2. Try again")
		fmt.Println("3. Exit")
		var choice int
		fmt.Print("Choose: ")
		fmt.Scanln(&choice)

		if choice == 1 {
			Register()
		} else if choice == 2 {
			Login()
		} else {
			os.Exit(0)
		}
		return
	}

	fmt.Print("Enter your password: ")
	fmt.Scanln(&utils.DataLogin.Password)

	if utils.DataLogin.Password != utils.DataUser.Password {
		fmt.Println("Wrong password.")
		fmt.Println("1. Forgot")
		fmt.Println("2. Try again")
		fmt.Println("3. Exit")
		var choice int
		fmt.Print("Choose: ")
		fmt.Scanln(&choice)

		if choice == 1 {
			Forgot()
		} else if choice == 2 {
			Login()
		} else {
			os.Exit(0)
		}
		return
	}

	fmt.Println("Login success!")
}

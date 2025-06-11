package auth

import (
	"fgo24-go-auth-flow/utils"
	"fmt"
	"os"
)

func Login() {
	fmt.Println("=== Login ===")

	newLogin := &utils.User{}

	fmt.Print("Enter your email: ")
	fmt.Scanln(&newLogin.Email)

	var foundUser *utils.User
	for i := range utils.Users {
		if utils.Users[i].Email == newLogin.Email {
			foundUser = &utils.Users[i]
			break
		}
	}

	if foundUser == nil {
		fmt.Println("Email not found.")
		fmt.Println("1. Register")
		fmt.Println("2. Try again")
		fmt.Println("3. Exit")
		var choice int
		fmt.Print("Choose: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			Register()
		case 2:
			Login()
		default:
			os.Exit(0)
		}
		return
	}

	fmt.Print("Enter your password: ")
	fmt.Scanln(&newLogin.Password)

	hashedInput := hashPassword(newLogin.Password)

	if hashedInput != foundUser.Password {
		fmt.Println("Wrong password.")
		fmt.Println("1. Forgot password")
		fmt.Println("2. Try again")
		fmt.Println("3. Exit")
		var choice int
		fmt.Print("Choose: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			Forgot()
		case 2:
			Login()
		default:
			os.Exit(0)
		}
		return
	}

	fmt.Println("Login success!")
}

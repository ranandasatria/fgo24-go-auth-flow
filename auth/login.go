package auth

import (
	"fmt"
	"os"
	"fgo24-go-auth-flow/utils"
)

func Login() {
	fmt.Println("=== Login ===")

	newLogin:= &utils.User{}

	fmt.Print("Enter your email: ")
	fmt.Scanln(&newLogin.Email)

	var foundUser *utils.User
	for i := range utils.Users{
		if utils.Users[i].Email == newLogin.Email{
			foundUser = &utils.Users[i]
			break
		}
	}

	if foundUser == nil{
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
	fmt.Scanln(&newLogin.Password)

	if newLogin.Password != foundUser.Password {
		fmt.Println("Wrong password.")
		fmt.Println("1. Forgot password")
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

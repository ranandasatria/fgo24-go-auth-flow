package auth

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
)

func hashPassword(password string) string{
	hash:= md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}

func Auth() {
	for {
		fmt.Println("======")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Forgot Password")
		fmt.Println("0. Exit")

		var choice int
		fmt.Print("Choose menu: ")
		fmt.Scanln(&choice)


		if choice == 1 {
			Register()
		} else if choice == 2 {
			Login()
		} else if choice == 3 {
			Forgot()
		} else if choice == 0 {
			fmt.Println("Goodbye!")
			os.Exit(0)
		} else {
			fmt.Println("Invalid choice.")
		}
	}
}

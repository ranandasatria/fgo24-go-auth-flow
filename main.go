package main

import (
	"fgo24-go-auth-flow/utils"
	"fmt"
)

type User struct {
	Email string
	Password string
}

var users []User

func main() {
	for {
		var menu int
		utils.Menu()
		fmt.Print("Pilih: ")
		fmt.Scanf("%d\n", &menu)

		if menu == 1 {
			var email, password string

			fmt.Print("Masukkan email: ")
			fmt.Scanf("%s\n", &email)
			fmt.Print(email)

		
			if isEmailTaken(email) {
				fmt.Println("Email sudah terdaftar.")
				continue
			}

			fmt.Print("Masukkan password: ")
			fmt.Scanf("%s\n", &password)


			users = append(users, User{Email: email, Password: password})
			fmt.Println("Registrasi berhasil!")
			fmt.Println(users)

		} else if menu == 2 {
		
		} else if menu == 0 {
			fmt.Println("Keluar dari program.")
			break
		} else {
			fmt.Println("Pilihan tidak tersedia.")
		}
	}
}


func isEmailTaken(email string) bool {
	for _, u := range users {
		if u.Email == email {
			return true
		}
	}
	return false
}

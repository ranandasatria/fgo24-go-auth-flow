package utils

type User struct {
	Email string
	Password string
}

var Users []User	


func IsEmailTaken(email string) bool {
	for _, u := range Users {
		if u.Email == email {
			return true
		}
	}
	return false
}


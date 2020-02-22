package utils

import "golang.org/x/crypto/bcrypt"

// ComparePassword to compare hashing password with password
func ComparePassword(hashedPass string, pass []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), pass)
	if err != nil {
		return false
	}
	return true
}

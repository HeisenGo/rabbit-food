package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the given password using bcrypt.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

// CheckPasswordHash compares a hashed password with its plain-text version.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	// ErrWrongPassword
	return err == nil
}

func main() {
	pass := "09191437774371"
	rr, e := HashPassword(pass)
	g := CheckPasswordHash("444444", "$2a$10$q2CfxDFReHa0rbwGKuERj.TsKSfV2P4siTlt3ed9WYbRoRZGUqtZS")
	fmt.Println(rr, e, g)
}

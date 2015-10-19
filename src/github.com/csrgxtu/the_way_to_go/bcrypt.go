package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	passwd := "archer"

	crypted, _ := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	fmt.Println(string(crypted))
	crypteda, _ := bcrypt.GenerateFromPassword([]byte("Archer"), bcrypt.DefaultCost)
	fmt.Println(string(crypteda))
	if err := bcrypt.CompareHashAndPassword(crypted, []byte("Archer")); err != nil {
		fmt.Println("not equal")
	}
}

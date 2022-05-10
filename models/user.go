package models

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/marcio-olmedo-cavalini/financial-transactions-go-webapp/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateRandomPassword() string {
	min := 1
	max := 999999
	rand.Seed(time.Now().UnixNano())
	passwd := rand.Intn(max-min) + min
	passwdString := fmt.Sprintf("%06d", passwd)
	return passwdString
}

func CreateUser(user User) {
	database.DB.Create(&user)
}

func ListUsers() []User {
	var usuarios []User
	database.DB.Where("nome <> 'Admin'").Find(&usuarios)
	return usuarios
}

/*
func main() {
    password := "secret"
    hash, _ := HashPassword(password) // ignore error for the sake of simplicity

    fmt.Println("Password:", password)
    fmt.Println("Hash:    ", hash)

    match := CheckPasswordHash(password, hash)
    fmt.Println("Match:   ", match)
}
*/

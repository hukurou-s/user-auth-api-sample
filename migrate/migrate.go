package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"

	"github.com/hukurou-s/user-auth-api-sample/domain"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "user=hoge dbname=user-auth-sample-db password='poge' sslmode=disable")
	defer db.Close()

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Successful connection")
	}

	db.AutoMigrate(&domain.User{})

	sampleUser := domain.User{
		Name:     "Hoge",
		Email:    "hoge@example.com",
		Password: toHashPassword("pogepoge"),
	}
	db.Create(&sampleUser)
}


func toHashPassword(pass string) string {
	converted, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(converted)
}

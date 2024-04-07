package main

import (
	"fmt"
	"github.com/ghp3000/validator"
)

type User struct {
	User     string `validate:"min=6,max=32"`
	Password string `validate:"min=12,max=32"`
}

func main() {
	err := validator.SetLanguage(validator.LangZh)
	if err != nil {
		fmt.Println(err)
	}
	u := User{User: "admin", Password: "123456"}
	err = validator.Struct(u)
	if err != nil {
		fmt.Println(err)
	}
}

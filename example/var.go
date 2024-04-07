package main

import (
	"fmt"
	"github.com/ghp3000/validator"
)

func main() {
	i := 10
	validator.SetLanguage(validator.LangEn)
	if err := validator.Var(&i, "min=1,max=9"); err != nil {
		fmt.Println(err)
	}
}

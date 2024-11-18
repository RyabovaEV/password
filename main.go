package main

import (
	"fmt"
	"password/account"
)

func main() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	myAccount, err := account.NewAccountWithTimeStemp(login, password, url)
	if err != nil {
		fmt.Print("Не указан логин или не верный формат URL")
		return
	}
	myAccount.OutputPass()
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}

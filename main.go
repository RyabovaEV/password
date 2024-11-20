// package main менеджер паролей
package main

import (
	"fmt"
	"password/account"
	"password/files"
	"password/output"

	"github.com/fatih/color"
)

func main() {

	fmt.Println("__Менеджер паролей__")
	vault := account.NewVault(files.NewJsonDB("data.json"))
	//vault := account.NewVault(cloud.NewCloudDB("http://d.ry"))

Menu:
	for {
		variant := promptData([]string{
			"1 - создать аккаунт",
			"2 - найти аккаунт",
			"3 - удалить аккаунт",
			"4 - выход",
		})
		//choiceMenu := choiceMenu()
		switch variant {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		default:
			break Menu
		}
	}

}

func findAccount(vault *account.VaultWithDB) {
	url := promptData([]string{"Введите url для поиска"})
	accounts := vault.FindAccountsByURL(url)
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.OutputPass()
	}
}

func deleteAccount(vault *account.VaultWithDB) {
	url := promptData([]string{"Введите url для удаления"})
	isDeleted := vault.DeleteAccountsByURL(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		output.PrintError("Не найдено")
	}

}

func createAccount(vault *account.VaultWithDB) {
	login := promptData([]string{"Введите логин"})
	password := promptData([]string{"Введите пароль"})
	url := promptData([]string{"Введите URL"})

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Не указан логин или не верный формат URL")
		//fmt.Print("Не указан логин или не верный формат URL")
		return
	}
	//file, err := myAccount.ToBytes()
	vault.AddAccount(*myAccount)
}

func promptData[T any](actions []T) string {
	for idx, value := range actions {
		if idx+1 != len(actions) {
			fmt.Println(value)
		} else {
			fmt.Print(value, ": ")
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}

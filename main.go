// package main менеджер паролей
package main

import (
	"fmt"
	"password/account"

	"github.com/fatih/color"
)

func main() {

	fmt.Println("__Менеджер паролей__")
	vault := account.NewVault()
Menu:
	for {
		choiceMenu := choiceMenu()
		switch choiceMenu {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			break Menu
		}
	}

}

func choiceMenu() int {
	var choiceMenu int
	fmt.Println("1 - создать аккаунт")
	fmt.Println("2 - найти аккаунт")
	fmt.Println("3 - удалить аккаунт")
	fmt.Println("4 - выход")
	fmt.Scanln(&choiceMenu)
	return choiceMenu
}

func findAccount(vault *account.Vault) {
	url := promptData("Введите url для поиска: ")
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.OutputPass()
	}
}

func deleteAccount(vault *account.Vault) {
	url := promptData("Введите url для удаления: ")
	isDeleted := vault.DeleteAccountsByUrl(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		color.Red("Не найдено")
	}

}

func createAccount(vault *account.Vault) {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Print("Не указан логин или не верный формат URL")
		return
	}
	//file, err := myAccount.ToBytes()
	vault.AddAccount(*myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}

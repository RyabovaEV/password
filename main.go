// package main менеджер паролей
package main

import (
	"fmt"
	"password/account"
	"password/files"
	"password/output"
	"strings"

	"github.com/fatih/color"
)

var menu = map[string]func(*account.VaultWithDB){
	"1": createAccount,
	"2": findAccountByURL,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

var menuVariants = []string{
	"1 - создать аккаунт",
	"2 - найти аккаунт по URL",
	"3 - найти аккаунт по логину",
	"4 - удалить аккаунт",
	"5 - выход",
}

// menuCounter функция замыкания
func menuCounter() func() {
	i := 0
	return func() {
		i++
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("__Менеджер паролей__")
	vault := account.NewVault(files.NewJsonDB("data.json"))
	counter := menuCounter()
	//vault := account.NewVault(cloud.NewCloudDB("http://d.ry"))
Menu:
	for {
		counter()
		variant := promptData(menuVariants...)
		menuFunc := menu[variant]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
	}
}

func findAccountByURL(vault *account.VaultWithDB) {
	url := promptData("Введите url для поиска")
	accounts := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.URL, str)
	})
	outputRezult(&accounts)
}

func findAccountByLogin(vault *account.VaultWithDB) {
	login := promptData("Введите логин для поиска")
	accounts := vault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})
	outputRezult(&accounts)
}

func outputRezult(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range *accounts {
		account.OutputPass()
	}
}

func deleteAccount(vault *account.VaultWithDB) {
	url := promptData("Введите url для удаления")
	isDeleted := vault.DeleteAccountsByURL(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		output.PrintError("Не найдено")
	}

}

func createAccount(vault *account.VaultWithDB) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Не указан логин или не верный формат URL")
		//fmt.Print("Не указан логин или не верный формат URL")
		return
	}
	//file, err := myAccount.ToBytes()
	vault.AddAccount(*myAccount)
}

func promptData(actions ...string) string {
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

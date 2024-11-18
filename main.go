package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

type account struct {
	login    string
	password string
	url      string
}

type accountWithTimeStemp struct {
	createAt time.Time
	updateAt time.Time
	account
}

func (acc account) outputPass() {
	fmt.Println(acc)
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	arrayPass := make([]rune, n, n)

	for i := range arrayPass {
		arrayPass[i] = letters[rand.IntN(len(letters))]
	}
	acc.password = string(arrayPass)
}

/* func newAccount(login, password, urlString string) (*account, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")

	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &account{
		login:    login,
		password: password,
		url:      urlString,
	}
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
} */

func newAccountWithTimeStemp(login, password, urlString string) (*accountWithTimeStemp, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &accountWithTimeStemp{
		createAt: time.Now(),
		updateAt: time.Now(),
		account: account{
			login:    login,
			password: password,
			url:      urlString,
		},
	}
	if password == "" {
		newAcc.account.generatePassword(12)
	}
	return newAcc, nil
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890*!")

func main() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	MyAccount, err := newAccountWithTimeStemp(login, password, url)
	if err != nil {
		fmt.Print("Не указан логин или не верный формат URL")
		return
	}
	MyAccount.outputPass()
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}

func outputPass(acc *account) {
	fmt.Println(*acc)
	fmt.Println(acc.login, acc.password, acc.url)
}

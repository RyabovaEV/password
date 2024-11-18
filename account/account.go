package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890*!")

// Account структура аккаунта
type Account struct {
	login    string
	password string
	url      string
}

// AccountWithTimeStemp доп параметры для аккаунта
type AccountWithTimeStemp struct {
	createAt time.Time
	updateAt time.Time
	Account
}

func (acc *Account) generatePassword(n int) {
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

// NewAccountWithTimeStemp создание аккаунта со структурой AccountWithTimeStemp
func NewAccountWithTimeStemp(login, password, urlString string) (*AccountWithTimeStemp, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &AccountWithTimeStemp{
		createAt: time.Now(),
		updateAt: time.Now(),
		Account: Account{
			login:    login,
			password: password,
			url:      urlString,
		},
	}
	if password == "" {
		newAcc.Account.generatePassword(12)
	}
	return newAcc, nil
}

// OutputPass вывод созданного аккаунта на экран
func (acc *Account) OutputPass() {
	fmt.Println(*acc)
	color.Cyan(acc.login)
	fmt.Println(acc.login, acc.password, acc.url)
}

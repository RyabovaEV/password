// Package account создание пакета
package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890*!")

// Account структура аккаунта
type Account struct {
	Login    string    `json:"login"`
	Password string    `json:"password"`
	URL      string    `json:"url"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

func (acc *Account) generatePassword(n int) {
	arrayPass := make([]rune, n, n)

	for i := range arrayPass {
		arrayPass[i] = letters[rand.IntN(len(letters))]
	}
	acc.Password = string(arrayPass)
}

// NewAccount создание аккаунта
func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")

	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &Account{
		Login:    login,
		Password: password,
		URL:      urlString,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}

// OutputPass вывод созданного аккаунта на экран
func (acc *Account) OutputPass() {
	color.Cyan(acc.Login)
	color.Cyan(acc.Password)
	color.Cyan(acc.URL)
}

package account

import (
	"encoding/json"
	"password/encrypter"
	"password/output"
	"strings"
	"time"

	"github.com/fatih/color"
)

// ByteReader интерфейс чтения
type ByteReader interface {
	Read() ([]byte, error)
}

// ByteWriter интерфейс записи
type ByteWriter interface {
	Write([]byte)
}

// Db интерфейс для работы с БД
type Db interface {
	ByteReader
	ByteWriter
}

// Vault структура контейнера
type Vault struct {
	Accounts []Account `json:"accounts"`
	UpdateAt time.Time `json:"updateAt"`
}

// VaultWithDB структура контейнера с учетом БД
type VaultWithDB struct {
	Vault
	db  Db
	enc encrypter.Encrypter
}

// NewVault создание контейнера для хранения паролей
func NewVault(db Db, enc encrypter.Encrypter) *VaultWithDB {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDB{
			Vault: Vault{
				Accounts: []Account{},
				UpdateAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}

	data := enc.Decrypt(file)
	var vault Vault
	err = json.Unmarshal(data, &vault)
	color.HiCyan("Найдено %d aккaунтов", len(vault.Accounts))
	if err != nil {
		output.PrintError("Не удалось разобрать файл data.vault")
		return &VaultWithDB{
			Vault: Vault{
				Accounts: []Account{},
				UpdateAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	return &VaultWithDB{
		Vault: vault,
		db:    db,
		enc:   enc,
	}
}

// AddAccount добавление аккаунта в контейнер
func (vault *VaultWithDB) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

// ToBytes преобразование в JSON
func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// FindAccounts поиск аккаунта по URL
func (vault *VaultWithDB) FindAccounts(url string, checker func(Account, string) bool) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {
		isMatched := checker(account, url)
		//isMatched := strings.Contains(account.URL, url)
		if isMatched {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

// DeleteAccountsByURL удаление аккаунта по URL
func (vault *VaultWithDB) DeleteAccountsByURL(url string) bool {
	var accounts []Account
	isDeleted := false
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.URL, url)
		if !isMatched {
			accounts = append(accounts, account)
			continue
		}
		isDeleted = true
	}
	vault.Accounts = accounts
	vault.save()
	return isDeleted
}

func (vault *VaultWithDB) save() {
	//vault.Accounts = accounts
	vault.UpdateAt = time.Now()
	data, err := vault.ToBytes()
	encData := vault.enc.Encrypt(data)
	if err != nil {
		output.PrintError("Не удалось преобразовать")
		//color.Red("")
	}
	vault.db.Write(encData)
	//db.Write(data)
}

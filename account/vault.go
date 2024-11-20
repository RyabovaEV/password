package account

import (
	"encoding/json"
	"password/output"
	"strings"
	"time"
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
	db Db
}

// NewVault создание контейнера для хранения паролей
func NewVault(db Db) *VaultWithDB {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDB{
			Vault: Vault{
				Accounts: []Account{},
				UpdateAt: time.Now(),
			},
			db: db,
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		output.PrintError("Не удалось разобрать файл data.json")
		return &VaultWithDB{
			Vault: Vault{
				Accounts: []Account{},
				UpdateAt: time.Now(),
			},
			db: db,
		}
	}
	return &VaultWithDB{
		Vault: vault,
		db:    db,
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

// FindAccountsByURL поиск аккаунта по URL
func (vault *VaultWithDB) FindAccountsByURL(url string) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.URL, url)
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
	if err != nil {
		output.PrintError("Не удалось преобразовать")
		//color.Red("")
	}
	vault.db.Write(data)
	//db.Write(data)
}

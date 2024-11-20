// Package files пакет для работы с файлами
package files

import (
	"fmt"
	"os"
	"password/output"

	"github.com/fatih/color"
)

// JsonDB структра файла json
type JsonDB struct {
	filename string
}

// NewJsonDB сщздание нового файла
func NewJsonDB(name string) *JsonDB {
	return &JsonDB{
		filename: name,
	}
}

// Write запись файла
func (db *JsonDB) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		output.PrintError(err)
		fmt.Println(err)
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		output.PrintError(err)
		return
	}
	color.Green("Запись успешна")
}

// Read чтение файла
func (db *JsonDB) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

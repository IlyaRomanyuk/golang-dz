package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FileJsonDb struct {
	fileName string
}

func NewFileJsonDb(name string) *FileJsonDb {
	return &FileJsonDb{
		fileName: name,
	}
}

func (db *FileJsonDb) Read() ([]byte, error) {
	isJSON := isJSONFile(db.fileName)

	if !isJSON {
		fmt.Println("Ошибка в расширении файла, ожидается json")
		return nil, errors.New("ожидается json")
	}

	dataFile, err := os.ReadFile(db.fileName)

	if err != nil {
		fmt.Println("Ошибка чтения файла", err.Error())
		return nil, err
	}

	return dataFile, nil
}

func isJSONFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".json"
}

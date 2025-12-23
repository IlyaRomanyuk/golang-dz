package storage

import (
	"fmt"
	"os"
)

type Db interface {
	Read() ([]byte, error)
}

type StorageWithDb struct {
	db Db
}

func NewStorage(db Db) *StorageWithDb {
	return &StorageWithDb{
		db,
	}
}

func (db *StorageWithDb) ReadBinList() []byte {
	dataFile, err := db.db.Read()

	if err != nil {
		fmt.Println("Ошибка чтения файла", err.Error())
		return nil
	}

	return dataFile
}

func (db *StorageWithDb) SaveBinList(data []byte, fileName string) bool {
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Ошибка создания файла", err.Error())
		return false
	}

	defer file.Close()

	_, err = file.Write(data)

	if err != nil {
		fmt.Println("Ошибка записи в файл", err.Error())
		return false
	}

	return true
}

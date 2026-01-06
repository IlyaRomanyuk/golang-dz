package storage

import (
	"fmt"
)

type Db interface {
	Read() ([]byte, error)
	Write([]byte) bool
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

func (db *StorageWithDb) SaveBinList(data []byte) bool {
	isFileSaved := db.db.Write(data)

	return isFileSaved
}

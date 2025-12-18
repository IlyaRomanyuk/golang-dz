package storage

import (
	"fmt"
	"os"
)

func SaveBinList(data []byte, fileName string) bool {
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

func ReadBinList(fileName string) []byte {
	dataFile, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Ошибка чтения файла", err.Error())
		return nil
	}

	return dataFile
}

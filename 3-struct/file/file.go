package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadBinList(fileName string) []byte {
	isJSON := isJSONFile(fileName)

	if !isJSON {
		fmt.Println("Ошибка в расширении файла, ожидается json")
		return nil
	}

	dataFile, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Ошибка чтения файла", err.Error())
		return nil
	}

	return dataFile
}

func isJSONFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	fmt.Println(ext, "-----ext-----")
	return ext == ".json"
}

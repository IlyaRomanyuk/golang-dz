package main

import (
	"fmt"
	"time"

	"github.com/IlyaRomanyuk/golang-dz/3-struct/bins"
	"github.com/IlyaRomanyuk/golang-dz/3-struct/file"
	"github.com/IlyaRomanyuk/golang-dz/3-struct/storage"
)

func main() {
	var fileJsonDb = file.NewFileJsonDb("storage.json")
	var storage = storage.NewStorage(fileJsonDb)

	var binList bins.BinList
	binList = append(binList, bins.NewBin("2", true, "brest", time.Now()))
	byteData := bins.ToByteSlice(binList)

	isSaveFile := storage.SaveBinList(byteData)

	if isSaveFile {
		fmt.Println("Файл успешно сохранен")
	}

	readedFile := storage.ReadBinList()

	fmt.Println(string(readedFile), "binList in string json")
}

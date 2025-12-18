package main

import (
	"fmt"
	"time"

	"github.com/IlyaRomanyuk/golang-dz/3-struct/bins"
	"github.com/IlyaRomanyuk/golang-dz/3-struct/storage"
)

func main() {
	var fileName string = "storage.json"
	var binList bins.BinList
	var bin1 = bins.NewBin("1", true, "test", time.Now())

	binList = append(binList, bin1)
	byteData := bins.ToByteSlice(binList)

	isSaveFile := storage.SaveBinList(byteData, fileName)

	if isSaveFile {
		fmt.Println("Файл успешно сохранен")
	}

	readedFile := storage.ReadBinList(fileName)

	fmt.Println(string(readedFile), "binList in string json")
}

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/IlyaRomanyuk/golang-dz/3-struct/api"
	"github.com/IlyaRomanyuk/golang-dz/3-struct/bins"
	"github.com/IlyaRomanyuk/golang-dz/3-struct/config"
	"github.com/IlyaRomanyuk/golang-dz/3-struct/file"
	"github.com/IlyaRomanyuk/golang-dz/3-struct/storage"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}

	config := config.NewConfig()
	api := api.NewApi(config)

	fmt.Println(api.Config.GetKey())

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

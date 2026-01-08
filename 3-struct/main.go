package main

import (
	"bytes"
	"errors"
	"flag"
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

	create := flag.Bool("create", false, "Создание")
	update := flag.Bool("update", false, "Обновление")
	delete := flag.Bool("delete", false, "Удаление")
	get := flag.Bool("get", false, "Получение бина по id")
	list := flag.Bool("list", false, "Список бинов")
	fileName := flag.String("file", "storage.json", "Название файла")
	name := flag.String("name", "test", "Название бина")
	id := flag.String("id", "", "Id бина")

	flag.Parse()

	config := config.NewConfig()
	api := api.NewApi(config)
	storage := storage.NewStorage(file.NewFileJsonDb(*fileName))

	var binList bins.BinList
	binList = append(binList, bins.NewBin("2", true, "brest", time.Now()))
	byteData := bins.ToByteSlice(binList)

	isSaveFile := storage.SaveBinList(byteData)

	if isSaveFile {
		fmt.Println("Файл успешно сохранен")
	}

	readedFile := storage.ReadBinList()

	switch {
	case *create:
		api.CreateBin(*name, bytes.NewReader(readedFile))
		return
	case *update:
		api.UpdateBin(*id, bytes.NewReader(readedFile))
		return
	case *delete:
		api.DeleteBin(*id)
		return
	case *get:
		api.GetBin(*id)
		return
	case *list:
		api.GetList()
		return
	default:
		errors.New("Неизвестная команда")
	}
}

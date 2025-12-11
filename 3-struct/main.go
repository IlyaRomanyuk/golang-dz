package main

import (
	"fmt"
	"time"

	"github.com/IlyaRomanyuk/golang-dz/3-struct/api"
	"github.com/IlyaRomanyuk/golang-dz/3-struct/bins"
	"github.com/IlyaRomanyuk/golang-dz/3-struct/file"
	"github.com/IlyaRomanyuk/golang-dz/3-struct/storage"
)

func main() {
	var binList bins.BinList
	var bin1 = bins.NewBin("1", true, "test", time.Now())

	binList = append(binList, bin1)

	file.GetFile()
	api.GetApiList()
	storage.GetStorage()

	fmt.Println(binList, "binList")
}

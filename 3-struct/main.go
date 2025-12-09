package main

import (
	"fmt"
	"time"
)

type BinStruct struct {
	id        string
	private   bool
	name      string
	createdAt time.Time
}

type BinList []BinStruct

func newBin(id string, private bool, name string, createdAt time.Time) BinStruct {
	var bin BinStruct = BinStruct{
		id,
		private,
		name,
		createdAt,
	}

	return bin
}

func main() {
	var binList BinList
	var bin1 BinStruct = newBin("1", true, "test", time.Now())
	var bin2 BinStruct = newBin("2", false, "brest", time.Now())

	binList = append(binList, bin1)
	binList = append(binList, bin2)

	fmt.Println(binList, "---binList---")
}

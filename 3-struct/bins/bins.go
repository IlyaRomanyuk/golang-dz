package bins

import (
	"encoding/json"
	"fmt"
	"time"
)

type BinStruct struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type BinList []BinStruct

func NewBin(id string, private bool, name string, createdAt time.Time) BinStruct {
	var bin BinStruct = BinStruct{
		Id:        id,
		Private:   private,
		Name:      name,
		CreatedAt: createdAt,
	}

	return bin
}

func ToByteSlice(binList BinList) []byte {
	byteData, err := json.Marshal(binList)

	if err != nil {
		fmt.Println("Ошибка сериализации", err.Error())
		return nil
	}

	return byteData
}

package bins

import "time"

type BinStruct struct {
	id        string
	private   bool
	name      string
	createdAt time.Time
}

type BinList []BinStruct

func NewBin(id string, private bool, name string, createdAt time.Time) BinStruct {
	var bin BinStruct = BinStruct{
		id,
		private,
		name,
		createdAt,
	}

	return bin
}

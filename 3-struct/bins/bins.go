package bins

import (
	"encoding/json"
	"fmt"
	"struct/files"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	bins []*Bin
}

func NewBin(id string, private bool, name string) *Bin {
	return &Bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
}

func NewBinList(bins []*Bin) *BinList {
	return &BinList{
		bins: bins,
	}
}

func (bin *Bin) SaveBin(filePath string) {
	file, err := json.Marshal(bin)
	if err != nil {
		fmt.Println(err)
	}
	files.WriteFile(file, filePath)
}

func (bin *Bin) ReadBinsList(filePath string) []byte {
	file, err := files.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	return file
}

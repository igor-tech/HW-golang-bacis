package bins

import (
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}

type BinList struct {
	Bins []Bin `json:"bins"`
}

func NewBin(id string, private bool, name string) *Bin {
	return &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}

func (bl *BinList) AddBin(id string, private bool, name string) {
	bin := NewBin(id, private, name)
	bl.Bins = append(bl.Bins, *bin)
}

func NewBinList() *BinList {
	return &BinList{
		Bins: []Bin{},
	}
}

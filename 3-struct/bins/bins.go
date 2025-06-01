package bins

import (
	"errors"
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

func NewBin(id string, private bool, name string) (*Bin, error) {
	if id == "" || name == "" {
		return nil, errors.New("id and name must not be empty")
	}
	return &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}, nil
}

func (bl *BinList) AddBin(id string, private bool, name string) error {
	bin, err := NewBin(id, private, name)
	if err != nil {
		return err
	}
	bl.Bins = append(bl.Bins, *bin)
	return nil
}

func NewBinList() *BinList {
	return &BinList{
		Bins: []Bin{},
	}
}

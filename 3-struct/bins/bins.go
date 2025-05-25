package bins

import (
	"encoding/json"
	"fmt"
	"struct/files"
	"time"
)

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

type BinList struct {
	Bins []*Bin
}

func NewBin(id string, private bool, name string) *Bin {
	return &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}

func NewBinList(bins []*Bin) *BinList {
	return &BinList{
		Bins: bins,
	}
}

func (bin *Bin) SaveBin(filePath string) error {
	data, err := json.Marshal(bin)
	if err != nil {
		return fmt.Errorf("failed to marshal bin: %w", err)
	}
	err = files.WriteFile(data, filePath)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func ReadBinsList(filePath string) (*BinList, error) {
	data, err := files.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	var bins BinList
	err = json.Unmarshal(data, &bins)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal bins list: %w", err)
	}
	return &bins, nil
}

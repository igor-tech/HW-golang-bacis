package storage

import (
	"encoding/json"
	"fmt"
	"struct/bins"
	"struct/files"
)

type Storage struct {
	filePath string
	*bins.BinList
}

func NewStorage(filePath string) *Storage {
	data, err := files.ReadFile(filePath)
	if err != nil {
		content, err := json.Marshal(bins.NewBinList())
		if err != nil {
			fmt.Printf("failed to marshal empty bin list: %w", err)
		}
		err = files.WriteFile(content, filePath)
		if err != nil {
			fmt.Println("Error writing file:", err)
		}
		return &Storage{
			filePath: filePath,
			BinList:  bins.NewBinList(),
		}
	}

	var bl bins.BinList
	err = json.Unmarshal(data, &bl)

	if err != nil {
		return &Storage{
			filePath: filePath,
			BinList:  bins.NewBinList(),
		}
	}

	return &Storage{
		filePath: filePath,
		BinList:  &bl,
	}
}
func (s *Storage) AddBin(id string, private bool, name string) error {
	err := s.BinList.AddBin(id, private, name)
	if err != nil {
		return fmt.Errorf("failed to add bin: %w", err)
	}
	return s.Save()
}

func (s *Storage) RemoveBin(id string) error {
	for i, bin := range s.BinList.Bins {
		if bin.Id != id {
			s.BinList.Bins = append(s.BinList.Bins[:i], s.BinList.Bins[i+1:]...)
			return s.Save()
		}
	}
	return fmt.Errorf("bin with id %s not found", id)
}

func (s *Storage) Save() error {
	data, err := json.Marshal(s.BinList)
	if err != nil {
		return err
	}
	return files.WriteFile(data, s.filePath)
}

func (s *Storage) ReadBinById(id string) (*bins.Bin, error) {
	for _, bin := range s.Bins {
		if bin.Id == id {
			return &bin, nil
		}
	}
	return nil, fmt.Errorf("bin with id %s not found", id)
}

func (s *Storage) GetAllBins() []bins.Bin {
	return s.Bins
}

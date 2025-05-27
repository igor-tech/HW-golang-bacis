package main

import (
	"fmt"
	"struct/storage"
)

func main() {
	binList := storage.NewStorage("data.json")
	err := binList.AddBin("1", true, "Мои данные")
	err = binList.AddBin("2", true, "Не мои данные")
	if err != nil {
		fmt.Println(err)
	}
	bin, _ := binList.ReadBinById("1")
	fmt.Println(bin)
	fmt.Println(binList.GetAllBins())
}

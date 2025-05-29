package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"struct/storage"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// При работе с api будем передавать ключ
	//appConfig := config.NewConfig("KEY")

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

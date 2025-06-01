package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"log"
	"os"
	"struct/api"
	"struct/config"
	"struct/files"
	"struct/storage"
)

const BinApiBaseUrl = "https://api.jsonbin.io/v3/b"

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "A utility for managing JSON bins via JSON Bin API")
		fmt.Fprintln(os.Stderr, "Flags:")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nExamples:")
		fmt.Fprintln(os.Stderr, "  Create a bin:  -create -file data.json -name mybin")
		fmt.Fprintln(os.Stderr, "  Update a bin:  -update -file data.json -id <bin_id>")
		fmt.Fprintln(os.Stderr, "  Delete a bin:  -delete -id <bin_id>")
		fmt.Fprintln(os.Stderr, "  Get a bin:     -get -id <bin_id>")
		fmt.Fprintln(os.Stderr, "  List bins:     -list")
	}
}

func main() {
	isCreate := flag.Bool("create", false, "create a new bin")
	isUpdate := flag.Bool("update", false, "update a bin")
	isDelete := flag.Bool("delete", false, "delete a bin")
	isGet := flag.Bool("get", false, "get a bin")
	isList := flag.Bool("list", false, "list a bins")
	pathFile := flag.String("file", "", "path to a file")
	id := flag.String("id", "", "a bin id")
	name := flag.String("name", "", "name of bin")

	flag.Parse()

	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		log.Fatal("Error: .env file not found")
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	appConfig, err := config.NewConfig("JSON_BIN_KEY")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	binStorage := storage.NewStorage("data.json")
	binApi, err := api.NewBinApi(*appConfig, *binStorage, BinApiBaseUrl)
	if err != nil {
		log.Fatalf("Error initializing bin API: %v", err)
	}

	if err := validateFlags(*isCreate, *isUpdate, *isDelete, *isGet, *isList, *pathFile, *id, *name); err != nil {
		fmt.Println(err)
		flag.Usage()
		return
	}

	if *isCreate {
		if err := handleCreate(binApi, binStorage, *pathFile, *name); err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
	}

	if *isUpdate {
		if err := handleUpdate(binApi, *pathFile, *id); err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
	}

	if *isDelete {
		if err := handleDelete(binApi, binStorage, *id); err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
	}

	if *isGet {
		if err := handleGet(binApi, *id); err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
	}

	if *isList {
		hanglePrintList(binStorage)
	}

}

func validateFlags(isCreate, isUpdate, isDelete, isGet, isList bool, pathFile, id, name string) error {
	flags := []bool{isCreate, isUpdate, isDelete, isGet, isList}
	count := 0
	for _, f := range flags {
		if f {
			count++
		}
	}
	if count > 1 {
		return errors.New("only one operation (create, update, delete, get, list) can be specified")
	}
	if count == 0 {
		return errors.New("no operation specified")
	}
	if isCreate && (name == "" || pathFile == "") {
		return errors.New("file and name are required for create")
	}
	if isUpdate && (pathFile == "" || id == "") {
		return errors.New("file and id are required for update")
	}
	if isDelete && id == "" {
		return errors.New("id is required for delete")
	}
	if isGet && id == "" {
		return errors.New("id is required for get")
	}

	return nil
}

func handleCreate(binApi *api.BinApi, binStorage *storage.Storage, pathFile, name string) error {
	file, err := files.ReadFile(pathFile)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	metaData, err := binApi.CreateBin(file)
	if err != nil {
		return fmt.Errorf("failed to create bin in api: %w", err)
	}
	if err := binStorage.AddBin(metaData.ID, metaData.Private, name); err != nil {
		return fmt.Errorf("failed to add bin to storage: %w", err)
	}
	fmt.Println(color.GreenString("Bin created successfully, ID: %s", metaData.ID))
	return nil
}

func handleUpdate(binApi *api.BinApi, pathFile, id string) error {
	file, err := files.ReadFile(pathFile)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	err = binApi.UpdateBin(file, id)
	if err != nil {
		return fmt.Errorf("failed to update bin in api: %w", err)
	}
	color.GreenString("Bin updated successfully, ID: %s", id)
	return nil
}

func handleDelete(binApi *api.BinApi, binStorage *storage.Storage, id string) error {
	err := binApi.DeleteBin(id)
	if err != nil {
		return fmt.Errorf("failed to delete bin in api: %w", err)
	}
	err = binStorage.RemoveBin(id)
	if err != nil {
		return fmt.Errorf("failed to remove bin from storage: %w", err)
	}
	color.GreenString("Bin deleted successfully, ID: %s", id)
	return nil
}

func handleGet(binApi *api.BinApi, id string) error {
	bin, err := binApi.GetById(id)
	if err != nil {
		return fmt.Errorf("failed to get bin in api: %w", err)
	}
	bytes, err := json.MarshalIndent(bin, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to formatting json: %w", err)
	}
	fmt.Println(string(bytes))
	return nil
}

func hanglePrintList(binStorage *storage.Storage) {
	bins := binStorage.GetAllBins()

	for index, bin := range bins {
		color.Cyan("%d. id: %s, name: %s", index+1, bin.Id, bin.Name)

	}
}

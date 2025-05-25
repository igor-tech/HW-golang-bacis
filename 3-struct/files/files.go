package files

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadFile(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func IsJSON(file []byte) bool {
	var jsonFile map[string]interface{}
	err := json.Unmarshal(file, &jsonFile)
	if err != nil {
		return false
	}
	return true
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = file.Write(content)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Запись успешна")
}

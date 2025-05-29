package files

// Для чтения и записи файлов
import (
	"encoding/json"
	"fmt"
	"os"
)

// Тут мы читаем файл и в дальнейшем его передаем дальше для отправки в JSON BIN

func ReadFile(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	if !IsJSON(file) {
		return nil, fmt.Errorf("file %s is not JSON", path)
	}

	return file, nil
}

func IsJSON(file []byte) bool {
	if len(file) == 0 {
		return false
	}
	var jsonFile interface{}
	return json.Unmarshal(file, &jsonFile) == nil
}

func WriteFile(content []byte, name string) error {
	file, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("failed to close file: %v", err)
		}
	}()

	_, err = file.Write(content)

	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

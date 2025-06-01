package files

import (
	"encoding/json"
	"fmt"
	"os"
)

// ReadFile читает содержимое файла по указанному пути и проверяет, является ли оно валидным JSON.
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

// IsJSON проверяет, является ли содержимое файла валидным JSON.
func IsJSON(file []byte) bool {
	return len(file) > 0 && json.Valid(file)
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

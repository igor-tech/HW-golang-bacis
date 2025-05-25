package files

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadFile(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	return file, nil
}

func IsJSON(file []byte) bool {
	var jsonFile interface{}
	return json.Unmarshal(file, &jsonFile) == nil
}

func WriteFile(content []byte, name string) error {
	file, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = file.Write(content)

	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

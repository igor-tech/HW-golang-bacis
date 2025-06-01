package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"struct/config"
	"struct/storage"
)

type CreateMetaData struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	Private   bool   `json:"private"`
}
type UpdateMetaData struct {
	ID      string `json:"parentId"`
	Private bool   `json:"private"`
}
type DeleteMetaData struct {
	ID              string `json:"id"`
	VersionsDeleted bool   `json:"versionsDeleted"`
}

type CreateBinResponse struct {
	Record   map[string]interface{} `json:"record"`
	Metadata CreateMetaData         `json:"metadata"`
}

type UpdateBinResponse struct {
	Record   map[string]interface{} `json:"record"`
	Metadata UpdateMetaData         `json:"metadata"`
}

type DeleteBinResponse struct {
	Message  string         `json:"message"`
	Metadata DeleteMetaData `json:"metadata"`
}

type BinApi struct {
	config.Config
	storage.Storage
	baseUrl string
	client  *http.Client
}

// NewBinApi Создает новый binApi с указанным конфигом, хранилищем и базовым урл адресом
func NewBinApi(config config.Config, storage storage.Storage, baseUrl string) (*BinApi, error) {
	if config.Key == "" {
		return nil, fmt.Errorf("API key is empty")
	}

	return &BinApi{
		Config:  config,
		Storage: storage,
		baseUrl: baseUrl,
		client:  &http.Client{},
	}, nil
}

// DoRequest выполняет HTTP-запрос с указанным методом, URL и содержимым, возвращая тело ответа.
func (binApi *BinApi) DoRequest(method string, path string, content []byte) ([]byte, error) {
	if method == "POST" || method == "PUT" && len(content) == 0 {
		return nil, fmt.Errorf("content is empty")
	}
	client := binApi.client
	req, err := http.NewRequest(method, path, bytes.NewBuffer(content))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Add("X-Master-Key", binApi.Config.GetKey())
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("request failed with status %s: %s", resp.Status, string(body))
	}

	return io.ReadAll(resp.Body)
}

// CreateBin создает новый JSON Bin с содержимым файла и возвращает его метаданные.
func (binApi *BinApi) CreateBin(file []byte) (*CreateMetaData, error) {
	body, err := binApi.DoRequest("POST", binApi.baseUrl, file)

	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}

	var binResponse CreateBinResponse
	err = json.Unmarshal(body, &binResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body: " + err.Error())
	}

	return &binResponse.Metadata, nil
}

// UpdateBin обновляет указанный JSON Bin
func (binApi *BinApi) UpdateBin(file []byte, id string) error {
	body, err := binApi.DoRequest("PUT", binApi.baseUrl+"/"+id, file)

	if err != nil {
		return fmt.Errorf("failed to read body: %w", err)
	}

	var binResponse UpdateBinResponse
	err = json.Unmarshal(body, &binResponse)
	if err != nil {
		return fmt.Errorf("failed to unmarshal body: " + err.Error())
	}
	return nil
}

// DeleteBin удаляет указанный JSON Bin
func (binApi *BinApi) DeleteBin(id string) error {
	_, err := binApi.DoRequest("DELETE", binApi.baseUrl+"/"+id, nil)
	if err != nil {
		return fmt.Errorf("failed to delete bin: %w", err)
	}
	return nil
}

// GetById получает JSON Bin по его идентификатору и возвращает его содержимое.
func (binApi *BinApi) GetById(id string) (*map[string]interface{}, error) {
	body, err := binApi.DoRequest("GET", binApi.baseUrl+"/"+id, nil)
	if err != nil {
		return nil, err
	}
	var binResponse CreateBinResponse
	err = json.Unmarshal(body, &binResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body: " + err.Error())
	}

	return &binResponse.Record, nil
}

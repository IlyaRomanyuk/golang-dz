package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Config interface {
	GetKey() string
	GetUrl() string
}

type ApiSetting struct {
	config Config
}

func NewApi(config Config) *ApiSetting {
	return &ApiSetting{
		config: config,
	}
}

func printResult(id, name, createdAt string, private bool) {
	fmt.Println("---")
	fmt.Printf("  ID записи: %s\n", id)
	fmt.Printf("  Имя: %s\n", name)
	fmt.Printf("  Приватная: %v\n", private)
	fmt.Printf("  Создано: %s\n", createdAt)
	fmt.Println("---")
}

func (api *ApiSetting) createRequest(method, path string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, api.config.GetUrl()+path, body)

	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Master-Key", "$2a$10$KnjDbi2h1XbQiLdkfxtDH.ng.ksFEEtv7cn378anHlk6/RtBq8Mtu")

	return request, nil
}

func (api *ApiSetting) DeleteBin(id string) error {
	request, err := api.createRequest("DELETE", "/b/"+id, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	return nil
}

func (api *ApiSetting) GetBin(id string) error {
	request, err := api.createRequest("GET", "/b/"+id, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var meta MetaData
	err = json.Unmarshal(body, &meta)
	if err != nil {
		return fmt.Errorf("ошибка парсинга JSON: %v", err)
	}

	printResult(meta.Metadata.Id, meta.Metadata.Name, meta.Metadata.CreatedAt, meta.Metadata.Private)
	return nil
}

func (api *ApiSetting) CreateBin(binName string, body *bytes.Reader) error {
	request, err := api.createRequest("POST", "/b", body)
	if err != nil {
		return err
	}
	request.Header.Set("X-Bin-Name", binName)

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return err
	}
	defer response.Body.Close()
	return nil
}

func (api *ApiSetting) UpdateBin(id string, body *bytes.Reader) error {
	request, err := api.createRequest("PUT", "/b/"+id, body)
	if err != nil {
		return err
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return err
	}
	defer response.Body.Close()
	return nil
}

func (api *ApiSetting) GetList() error {
	request, err := api.createRequest("GET", "/c/uncategorized/bins/", nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var bins []BinItem
	err = json.Unmarshal(body, &bins)
	if err != nil {
		return fmt.Errorf("ошибка парсинга JSON: %v", err)
	}

	for _, bin := range bins {
		printResult(bin.Record, bin.SnippetMeta.Name, bin.CreatedAt, bin.Private)
	}

	return nil
}

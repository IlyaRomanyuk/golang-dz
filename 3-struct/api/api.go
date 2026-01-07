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

type SnippetMeta struct {
	Name string `json:"name"`
}
type BinItem struct {
	Private     bool        `json:"private"`
	SnippetMeta SnippetMeta `json:"snippetMeta"`
	Record      string      `json:"record"`
	CreatedAt   string      `json:"createdAt"`
}

type MetaData struct {
	Metadata MetaDataItem `json:"metadata"`
}
type MetaDataItem struct {
	Private   bool   `json:"private"`
	Name      string `json:"name"`
	Id        string `json:"id"`
	CreatedAt string `json:"createdAt"`
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

func (api *ApiSetting) DeleteBin(id string) error {
	request, err := http.NewRequest("DELETE", api.config.GetUrl()+"/b/"+id, nil)

	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Master-Key", "$2a$10$KnjDbi2h1XbQiLdkfxtDH.ng.ksFEEtv7cn378anHlk6/RtBq8Mtu")

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
	request, err := http.NewRequest("GET", api.config.GetUrl()+"/b/"+id, nil)

	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Master-Key", "$2a$10$KnjDbi2h1XbQiLdkfxtDH.ng.ksFEEtv7cn378anHlk6/RtBq8Mtu")

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
	request, err := http.NewRequest("POST", api.config.GetUrl()+"/b", body)

	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Master-Key", "$2a$10$KnjDbi2h1XbQiLdkfxtDH.ng.ksFEEtv7cn378anHlk6/RtBq8Mtu")
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
	request, err := http.NewRequest("PUT", api.config.GetUrl()+"/b/"+id, body)

	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Master-Key", "$2a$10$KnjDbi2h1XbQiLdkfxtDH.ng.ksFEEtv7cn378anHlk6/RtBq8Mtu")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return err
	}
	defer response.Body.Close()
	return nil
}

func (api *ApiSetting) GetList() error {
	request, err := http.NewRequest("GET", api.config.GetUrl()+"/c/uncategorized/bins/", nil)

	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Master-Key", "$2a$10$KnjDbi2h1XbQiLdkfxtDH.ng.ksFEEtv7cn378anHlk6/RtBq8Mtu")

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

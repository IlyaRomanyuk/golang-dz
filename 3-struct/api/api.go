package api

import (
	"fmt"
)

type Config interface {
	GetKey() string
}

type ApiSetting struct {
	Config Config
}

func NewApi(config Config) *ApiSetting {
	return &ApiSetting{
		Config: config,
	}
}

func GetApiList() {
	fmt.Println("Some api list")
}

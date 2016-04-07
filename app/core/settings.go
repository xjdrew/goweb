package core

import (
	"encoding/json"
	"io/ioutil"
)

type DatabaseSettings struct {
	Hosts    string `json:"hosts"`
	Database string `json:"database"`
}

type Settings struct {
	Secret       string           `json:"secret"`
	PublicPath   string           `json:"public_path"`
	TemplatePath string           `json:"template_path"`
	Database     DatabaseSettings `json:"database"`
}

func (settings *Settings) Load(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, settings)
	return err
}

package json_utils

import (
	"encoding/json"

	"github.com/ghodss/yaml"
)

func Marshal(data interface{}) ([]byte, error) {
	bytes, err := json.Marshal(data)
	return bytes, err
}

func Unmarshal(dataStr string, data interface{}) error {
	err := json.Unmarshal([]byte(dataStr), data)
	return err
}

func YamlMarshal(data interface{}) ([]byte, error) {
	bytes, err := yaml.Marshal(data)
	return bytes, err
}

func YamlUnmarshal(dataStr string, data interface{}) error {
	err := yaml.Unmarshal([]byte(dataStr), data)
	return err
}

package json_utils

import "encoding/json"

func Marshal(data interface{}) ([]byte, error) {
	bytes, err := json.Marshal(data)
	return bytes, err
}

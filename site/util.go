package site

import (
	"encoding/json"
	"os"
)

func ReadJSON(fileName string, result interface{}) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(result)
}

package site

import (
	"log"
	"encoding/json"
	"os"
)

func ReadJSON(fileName string, result interface{}) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer func(){
		if err := file.Close(); err != nil {
			log.Println(err)
		}
	}()
	return json.NewDecoder(file).Decode(result)
}

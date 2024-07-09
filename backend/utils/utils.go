package utils

// The definition of logs is in main.go

import (
	"encoding/json"
	"log"
	"os"
)

func ReadModelfile() string {
	dat, err := os.ReadFile("./models/Modelfile")
	Check(err)
	return string(dat)
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func DecodeJson(json_data []byte) map[string]interface{} {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(json_data), &result)
	if err != nil {
		log.Println("Error in the decoding of json", err)
		return nil
	}
	return result
}

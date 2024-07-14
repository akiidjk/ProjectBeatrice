package utils

// The definition of logs is in main.go

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
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

func PrintProgressBar(iteration, total int, prefix, suffix string, length int, fill string) {
	percent := float64(iteration) / float64(total)
	filledLength := int(length * iteration / total)
	end := ">"

	if iteration == total {
		end = "="
	}
	bar := strings.Repeat(fill, filledLength) + end + strings.Repeat("-", (length-filledLength))
	fmt.Printf("\r%s [%s] %f%% %s", prefix, bar, percent, suffix)
	if iteration == total {
		fmt.Println()
	}
}

// func ReadStreamData(reader *bufio.Reader) {
// 	for {
// 		line, error := reader.ReadBytes('\n')
// 		if error != nil {
// 			log.Println(error.Error())
// 			break
// 		}

// 		result := DecodeJson(line)
// 		if result != nil {
// 			log.Println(result)
// 		}
// 	}
// }

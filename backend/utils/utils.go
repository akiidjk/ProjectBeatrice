package utils

// The definition of logs is in main.go

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"
)

func ReadModelfile() string {
	dat, err := os.ReadFile("./models/Modelfile")
	Check(err)
	return string(dat)
}

func Check(e error) {
	if e != nil {
		fmt.Print(e.Error())
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

func PrintProgressBar(iteration, total float64, prefix, suffix string, length int, fill string) {
	percent := iteration / total
	filledLength := int(float64(length) * iteration / total)
	end := ">"

	var color string

	if iteration == total {
		end = "="
	}
	bar := strings.Repeat(fill, filledLength) + end + strings.Repeat("-", (length-filledLength))

	if percent < 70 && percent >= 40 {
		color = Yellow
	} else if percent < 40 {
		color = Red
	} else {
		color = Green
	}

	fmt.Printf(Blue+"\r%s "+color+" [%s] "+White+" %.1f%% %s", prefix, bar, percent*100, suffix)
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

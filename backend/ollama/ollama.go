package ollama

import (
	"backend/config"
	"backend/logger"
	"backend/utils"
	"bufio"
	"bytes"
	"fmt"
	"net/http"
)

func CheckApi() int {
	resp, error := http.Get(config.URL_API + "api/ps")
	if error != nil {
		logger.Error(error.Error())
		return 0
	}
	return resp.StatusCode
}

func CreateModel() {
	logger.Info("Creating model...")

	var model = utils.ReadModelfile()

	userData := []byte(`{"name": "beatrice","modelfile": "` + model + `"}`)

	request, error := http.NewRequest("POST", config.URL_API+"api/create", bytes.NewBuffer(userData))
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	if error != nil {
		logger.Error(error.Error())
		utils.Check(error)
	}

	client := &http.Client{}
	response, error := client.Do(request)

	if error != nil {
		logger.Error(error.Error())
		utils.Check(error)
	}

	reader := bufio.NewReader(response.Body)
	for {
		line, error := reader.ReadBytes('\n')
		if error != nil {
			logger.Error(error.Error())
			break
		}

		result := utils.DecodeJson(line)
		if result != nil {
			logger.Info(result["status"].(string))
		}
	}

	defer response.Body.Close()
}

func SendMessage(message string) int {
	logger.Info("Sending message...")

	data := []byte(`{
  "model": "beatrice",
  "messages": [
    {
      "role": "user",
      "content": "` + message + `"
    }
  ],
	"stream": true
}`)

	logger.Debug(string(data))

	request, error := http.NewRequest("POST", config.URL_API+"api/chat", bytes.NewBuffer(data))
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	if error != nil {
		logger.Error(error.Error())
		utils.Check(error)
	}

	client := &http.Client{}
	response, error := client.Do(request)

	if error != nil {
		logger.Error(error.Error())
		utils.Check(error)
	}

	reader := bufio.NewReader(response.Body)
	for {
		line, error := reader.ReadBytes('\n')
		if error != nil {
			logger.Error(error.Error())
			break
		}

		result := utils.DecodeJson(line)
		if result != nil {
			message, ok := result["message"].(map[string]interface{})
			if ok {
				content, ok := message["content"].(string)
				if ok {
					// logger.Info(content)
					fmt.Print(content)
				}
			}
		}
	}

	defer response.Body.Close()

	return response.StatusCode
}

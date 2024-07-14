package ollama

import (
	"backend/config"
	"backend/logger"
	"backend/utils"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var messages []string

func CheckApi() int {
	resp, error := http.Get(config.API_URL + "api/ps")
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

	request, error := http.NewRequest("POST", config.API_URL+"api/create", bytes.NewBuffer(userData))
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

	//Read stream
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
			if result["total"] != nil {
				utils.PrintProgressBar(int(result["completed"].(float64)), int(result["total"].(float64)), "Pulling model", "Model downloaded", 50, "=")
			}
		}
	}

	defer response.Body.Close()
}

func SendMessage(c *gin.Context, user_message string) int {
	logger.Info("Sending message...")

	logger.Debug("Message: " + user_message)

	user_input := []byte(`{"role": "user","content": "` + user_message + `"}`)
	messages = append(messages, string(user_input))

	data := []byte(`{
	  "model": "beatrice",
	  "messages": [` + strings.Join(messages, ",") + `],
		"stream": true,
		"json": true
}`)

	logger.Debug(string(data))

	request, err := http.NewRequest("POST", config.API_URL+"api/chat", bytes.NewBuffer(data))
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	if err != nil {
		logger.Error(err.Error())
		utils.Check(err)
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		logger.Error(err.Error())
		utils.Check(err)
	}

	reader := bufio.NewReader(response.Body)
	var model_message string = ""
	logger.Info("Streaming data")

	type StreamMessage struct {
		Index   int
		Content string
	}

	chanStream := make(chan StreamMessage, 10)

	go func() {
		c.Stream(func(w io.Writer) bool {
			if msg, ok := <-chanStream; ok {
				c.SSEvent("message", gin.H{
					"index":   msg.Index,
					"content": msg.Content,
				})
				return true
			}
			return false
		})
	}()

	index := 0

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			logger.Error(err.Error())
			break
		}
		result := utils.DecodeJson(line)

		if result != nil {
			message, ok := result["message"].(map[string]interface{})
			if ok {
				content, ok := message["content"].(string)
				if ok {
					fmt.Print(content)
					model_message += content

					chanStream <- StreamMessage{Index: index, Content: content}
					index++
				}
			}
		}
	}

	close(chanStream)

	model_output := []byte(`{"role": "assistant","content": "` + strings.ReplaceAll(model_message, "\n", "") + `"}`)
	messages = append(messages, string(model_output))

	return 0
}

func ShowInfoModel(model string, verbose string) (string, int) {
	logger.Info("Getting info model...")

	data := []byte(`{"name": "` + model + `","verbose": ` + verbose + `}`)

	logger.Debug("Data: " + string(data))

	request, error := http.NewRequest("POST", config.API_URL+"api/show", bytes.NewBuffer(data))
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

	logger.Info("Model info")
	bodyBytes, _ := io.ReadAll(response.Body)
	bodyString := string(bodyBytes)

	return bodyString, response.StatusCode
}

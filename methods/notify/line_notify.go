package notify

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"pfms/logs"

	"github.com/patcharp/golib/requests"
	"github.com/spf13/viper"
)

type NotifyHeader struct {
	Message string `json:"message"`
}

type responseNotify struct {
	Response string
}

func LineNotify(message string) string {

	// Create a new buffer to hold the multipart form data
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Add the message field to the form
	err := writer.WriteField("message", message)
	if err != nil {
		logs.Error(err)
	}

	// Close the writer to finalize the form
	writer.Close()

	// Set the headers including the Authorization token and Content-Type
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %v", viper.GetString("line.token")),
		"Content-Type":  writer.FormDataContentType(),
	}

	url := "https://notify-api.line.me/api/notify"

	// Make the POST request
	resp, err := requests.Post(url, headers, &requestBody, 20)
	if err != nil {
		logs.Error(err)
	}
	resNotify := new(responseNotify)
	if resp.Code == 200 {
		resNotify.Response = "send Notify success"
	}

	return resNotify.Response
}

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

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	err := writer.WriteField("message", message)
	if err != nil {
		logs.Error(err)
	}

	writer.Close()

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %v", viper.GetString("line.token")),
		"Content-Type":  writer.FormDataContentType(),
	}

	url := viper.GetString("notify.line_url")

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

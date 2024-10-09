package script

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"pfms/logs"
	"pfms/utility"

	"github.com/gofiber/fiber/v2"
	"github.com/patcharp/golib/requests"
	"github.com/spf13/viper"
)

type NotifyHeader struct {
	Message string `json:"message"`
}

type responseNotify struct {
	Response string
}

func TestNotify(c *fiber.Ctx) error {

	reqMessage := new(NotifyHeader)

	if err := c.BodyParser(reqMessage); err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	// Create a new buffer to hold the multipart form data
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Add the message field to the form
	err := writer.WriteField("message", reqMessage.Message)
	if err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	// Close the writer to finalize the form
	writer.Close()

	// Set the headers including the Authorization token and Content-Type
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %v", viper.GetString("line.token")),
		"Content-Type":  writer.FormDataContentType(),
	}

	url := viper.GetString("notify.line_url")

	// Make the POST request
	resp, err := requests.Post(url, headers, &requestBody, 20)
	if err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}
	resNotify := new(responseNotify)
	if resp.Code == 200 {
		resNotify.Response = "send Notify success"
	}

	return utility.ResponseSuccess(c, resNotify)
}

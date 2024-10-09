package script

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"pfms/logs"
	"pfms/models"
	"pfms/utility"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func TestDiscordNotify(c *fiber.Ctx) error {

	reqMessage := new(NotifyHeader)

	if err := c.BodyParser(reqMessage); err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	webhookURL := viper.GetString("notify.discord_url")

	message := new(models.WebhookMessage)
	message.Content = reqMessage.Message

	messageBody, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error marshalling message:", err)
		return utility.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(messageBody))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return utility.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}
	defer resp.Body.Close()

	resNotify := new(responseNotify)
	if resp.StatusCode == http.StatusNoContent {
		resNotify.Response = "send Notify success"
		fmt.Println("Notification sent successfully!")
	} else {
		fmt.Printf("Failed to send notification. Status code: %d\n", resp.StatusCode)
		return utility.ResponseSuccess(c, fmt.Sprintf("Failed to send notification. Status code: %d", resp.StatusCode))
	}

	return utility.ResponseSuccess(c, resNotify)
}

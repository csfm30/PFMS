package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"pfms/models"

	"github.com/spf13/viper"
)

func DiscordNotify(message string) string {

	webhookURL := viper.GetString("notify.discord_url")

	reqMessage := models.WebhookMessage{
		Content: message,
	}

	messageBody, err := json.Marshal(reqMessage)
	if err != nil {
		fmt.Println("Error marshalling message:", err)
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(messageBody))
	if err != nil {
		fmt.Println("Error sending request:", err)
	}
	defer resp.Body.Close()

	resNotify := new(responseNotify)
	if resp.StatusCode == http.StatusNoContent {
		resNotify.Response = "send Notify success"
		fmt.Println("Notification sent successfully!")
	} else {
		fmt.Printf("Failed to send notification. Status code: %d\n", resp.StatusCode)
	}

	return resNotify.Response
}

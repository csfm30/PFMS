package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"pfms/models"
)

func DiscordNotify(message string) string {

	webhookURL := "https://discord.com/api/webhooks/1293419811128475698/r879xi1q3qMVTBNQQT8ybbu0E_Mw2VkZ2euUItmQWRYL0y_2EM7tdyC5xJdscLp6j154"

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

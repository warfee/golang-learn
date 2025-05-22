package tasks

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "github.com/hibiken/asynq"
)

const TypeSlackSendMessage = "slack:send-message"

type SlackSendMessagePayload struct {
    WebhookURL string `json:"webhook"`
    Message    string `json:"message"`
}

func NewSlackSendMessageTask(webhookURL, message string) (*asynq.Task, error) {
    payload, err := json.Marshal(SlackSendMessagePayload{
        WebhookURL: webhookURL,
        Message:    message,
    })
    if err != nil {
        return nil, err
    }
    return asynq.NewTask(TypeSlackSendMessage, payload), nil
}

func HandleSlackSendMessageTask(ctx context.Context, t *asynq.Task) error {
    var p SlackSendMessagePayload
    if err := json.Unmarshal(t.Payload(), &p); err != nil {
        return fmt.Errorf("unmarshal error: %v", err)
    }

    jsonData, err := json.Marshal(map[string]string{
        "text": p.Message,
    })
    if err != nil {
        return fmt.Errorf("failed to marshal JSON: %v", err)
    }

    req, err := http.NewRequest("POST", p.WebhookURL, bytes.NewBuffer(jsonData))
    if err != nil {
        return fmt.Errorf("failed to create request: %v", err)
    }
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return fmt.Errorf("failed to send request: %v", err)
    }
    defer resp.Body.Close()

    log.Printf("Sent to Slack: webhook=%s, message=%s, status=%s", p.WebhookURL, p.Message, resp.Status)
    return nil
}
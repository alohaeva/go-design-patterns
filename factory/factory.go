package factory

import "fmt"

type Message interface {
	GetContent() string
}

type EmailMessage struct {
	Email   string
	Content string
}

type SlackMessage struct {
	Channel string
	Content string
}

func (s *SlackMessage) GetContent() string {
	return s.Content
}

func (e *EmailMessage) GetContent() string {
	return e.Content
}

type NotificationService interface {
	SendMessage(message Message)
}

type EmailService struct{}

type SlackService struct{}

func (p *EmailService) SendMessage(message Message) {
	emailMessage, ok := message.(*EmailMessage)
	if !ok {
		fmt.Println("Invalid message type for EmailService")
		return
	}
	fmt.Printf("Sending email to %s with content: %s\n", emailMessage.Email, emailMessage.GetContent())
}

func (p *SlackService) SendMessage(message Message) {
	slackMessage, ok := message.(*SlackMessage)
	if !ok {
		fmt.Println("Invalid message type for SlackService")
		return
	}
	fmt.Printf("Sending Slack message to channel %s with content: %s\n", slackMessage.Channel, slackMessage.GetContent())
}

func NotificationFactory(notificationType string) (NotificationService, error) {
	switch notificationType {
	case "email":
		return &EmailService{}, nil
	case "slack":
		return &SlackService{}, nil
	default:
		return nil, fmt.Errorf("unknown notification type: %s", notificationType)
	}
}

func TestFactory(notificationType string) {
	notificationService, err := NotificationFactory(notificationType)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if notificationType == "email" {
		emailMessage := &EmailMessage{
			Email:   "test@example.com",
			Content: "Hello via email!",
		}
		notificationService.SendMessage(emailMessage)
	}

	if notificationType == "slack" {
		slackMessage := &SlackMessage{
			Channel: "#general",
			Content: "Hello via Slack!",
		}
		notificationService.SendMessage(slackMessage)
	}
}

package firebasepush

import (
	"context"
	"fmt"
	"time"

	fcm "firebase.google.com/go/v4/messaging"
)

func (fc *FirebaseClient) SendPushNotification(fcmTokenDevice, title, body string) error {
	// logrus.Info("🟨 SendPushNotification")
	if fc == nil || fc.Client == nil {
		return fmt.Errorf("💔 error FirebaseClient or its Client field is nil")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	message := &fcm.Message{
		Token: fcmTokenDevice,
		Notification: &fcm.Notification{
			Title: title,
			Body:  body,
		},
	}

	response, err := fc.Client.Send(ctx, message)
	if err != nil {
		// Логирование ошибки для диагностики
		fmt.Printf("💔 error sending message: %v\n", err)
		return fmt.Errorf("error sending message: %v", err)
	}

	fmt.Printf("Successfully sent message: %s\n", response)
	return nil
}

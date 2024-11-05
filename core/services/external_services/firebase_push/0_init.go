package firebasepush

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"sync"
	"time"

	fbase "firebase.google.com/go/v4"
	fcm "firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

var (
	onceFirebaseModule     sync.Once
	firebaseModuleInstance *FirebaseClient
)

type FirebaseClient struct {
	Client *fcm.Client
}

type FirebaseConfig struct {
	FcmServerToken *string
}

func GetFirebaseModuleInstance(firebaseConfigObj *FirebaseConfig) *FirebaseClient {
	// logrus.Info("🟨 GetFirebaseModuleInstance")
	if firebaseConfigObj == nil {
		log.Fatalf("firebaseConfigObj is nil")
	}

	onceFirebaseModule.Do(func() {
		firebaseClient, err := getClientFirebase(firebaseConfigObj.FcmServerToken)
		if err != nil {
			log.Fatalf("error getting Firebase client: %v", err)
		}
		firebaseModuleInstance = &FirebaseClient{
			Client: firebaseClient,
		}
	})
	return firebaseModuleInstance
}

// CreateClient инициализирует и возвращает клиента Firebase для отправки уведомлений.
func getClientFirebase(fcmServerToken *string) (*fcm.Client, error) {
	// logrus.Info("🟨 getClientFirebase")
	if fcmServerToken == nil {
		return nil, fmt.Errorf("fcmServerToken is nil")
	}
	log.Println("🔰🔰 Initializing Firebase client...: ")
	key, err := base64.StdEncoding.DecodeString(*fcmServerToken)
	if err != nil {
		return nil, fmt.Errorf("error decoding base64: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	opts := []option.ClientOption{option.WithCredentialsJSON(key)}
	app, err := fbase.NewApp(ctx, &fbase.Config{}, opts...)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	// Получение клиента для отправки уведомлений
	client, err := app.Messaging(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting Messaging client: %v", err)
	}

	return client, nil
}

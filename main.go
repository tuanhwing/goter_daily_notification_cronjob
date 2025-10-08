package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"tuanhwing/goter_daily_notification_cronjob/resources"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

var fcmClient *messaging.Client

// Initialize access messaging firebase service
func NewFirebaseMessaging() *messaging.Client {
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_ADMIN_FILE"))
	app, err := firebase.NewApp(context.TODO(), nil, opt)
	if err != nil {
		panic(err)
	}

	messaging, err := app.Messaging(context.Background())
	if err != nil {
		panic(err)
	}
	return messaging
}

func main() {
	log.Print("Started")

	//Load env
	env := ".env"
	err := godotenv.Load(env)
	if err != nil {
		fmt.Printf("Error loading %s file", env)
	}

	topic := os.Getenv("BROADCAST_TOPIC")
	fcmClient := NewFirebaseMessaging()
	notification := resources.GetRandomNotification()

	result, err := fcmClient.Send(context.Background(), &messaging.Message{
		Notification: &messaging.Notification{
			Title: notification.Title,
			Body:  notification.Body,
		},
		Topic: topic,
	})

	if err != nil {
		log.Println("FCM send notfication error: " + err.Error())
	}
	log.Println("FCM send notfication: " + result)

	log.Print("End")
}

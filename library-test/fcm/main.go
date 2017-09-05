package main

import (
	"context"
	"fmt"
	"log"
	"roshar/lib/fcm"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

const (
	apiKey   = "AAAA_1hXGc4:APA91bEfh2o71HrOIYIASMaxQ9Rr0fSujrzDKH4Bf-qfkzd7bX-l-YvhxQPn-MvgbqxFlk4IvCmgkseYGrqO_nepGfCeRmiLlsuPNXRv1koz_l6FX8xam5YClXUXOuDWULVywmqDZ2RrtKm5DLQXuwFCNup8ED8PnQ"
	senderID = "1096698763726"
	token    = "enDPvjRlzyY:APA91bGm6QXDkWI-XppAkMHcSWbp8iNDclLmb7qKCYYPklxExQtWo7oh6QtTk0UsNGXk5F06nrCEspQL2Rs0-wDQfFC8hvT68XgHAsArf8ZGOUxYpj2YXIeDvV0LXgLOg-0kPLaZQWrrwR_HiujJ-38t_Ac5cC8gsw"
	key      = "APA91bFROe0M1KeFpe36I99CkvEU5BUyAORqpNifo2HjjcSYIf_4O0zrCD1ECMorBJjLgVemKX8hywR_a4PkP6jz7ayAkX_YmdPfZZGSOUgdwgbNXcdZ_UeyqJ1onMm4HUDjR0mD5J9w"
)

func main() {
	FCMOfficialPackage()
}

func FCMOfficialPackage() {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}

	fcm, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	m := messaging.Message{
		Data: map[string]string{
			"temp":     "58.0F",
			"humidity": "47",
			"wind":     "18mph",
		},
		Token: token,
	}
	resp, err := fcm.SendDryRun(ctx, &m)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Message verified:", resp)
}

func manageDeviceGroup() {
	// Create a FCM client to send the message.
	client, err := fcm.NewClient(apiKey, fcm.WithEndpoint(fcm.NotificationEndpoint), fcm.WithSenderID(senderID))
	if err != nil {
		panic(err)
	}

	deviceGroupMsg := &fcm.DeviceGroupMsg{
		Operation:       "add",
		KeyName:         "1-test@mail.com",
		RegistrationIDs: []string{token},
		Key:             key,
	}

	// Send the message and receive the response without retries.
	response, err := client.ManageDeviceGroup(context.Background(), deviceGroupMsg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", response)
}

func sendMsg() {
	// Create a FCM client to send the message.
	client, err := fcm.NewClient(apiKey)
	if err != nil {
		panic(err)
	}

	msg := &fcm.Message{
		Token: key,
		Data: map[string]interface{}{
			"msg": "This is a test message to Firebase",
		},
	}

	// Send the message and receive the response without retries.
	response, err := client.Send(context.Background(), msg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", response)
}

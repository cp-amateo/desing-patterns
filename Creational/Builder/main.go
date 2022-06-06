package main

import "fmt"

func main() {
	notificationBuilder := newNotificationBuilder()

	notificationBuilder.SetTitle("New tittle")
	notificationBuilder.SetSubtitle("Subtitle")
	notificationBuilder.SetIcon("someIcon.png")
	notificationBuilder.SetImage("anImage.jpg")
	notificationBuilder.SetPriority(1)
	notificationBuilder.SetMessage("a super coll message")

	notification, err := notificationBuilder.Build()

	if err != nil {
		fmt.Println("Error creating the notification:", err)
	} else {
		fmt.Printf("Notification: %+v", notification)
	}
}

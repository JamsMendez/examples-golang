package builder

import (
	"fmt"
	"log"
)

func Run() {
	nBuilder := newNotificationBuilder()

	nBuilder.SetTitle("New Notification")
	nBuilder.SetIcon("icon.png")
	nBuilder.SetMessage("basic notification")
	nBuilder.SetPriority(1)
	nBuilder.SetNotType("alert")

	notification, err := nBuilder.Build()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s %s success", notification.GetTitle(), notification.GetMessage())
}

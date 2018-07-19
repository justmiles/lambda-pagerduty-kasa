package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/justmiles/lambda-pagerduty-kasa/SmartPlug"
)

var (
	deviceName = os.Getenv("DEVICE_NAME")
)

func Handler(request Request) error {
	fmt.Println(request)

	device := smartplug.GetDeviceByAlias(deviceName)
	device.On()

	for _, message := range request.Messages {
		fmt.Println(message.Event)
		switch message.Event {
		case "incident.resolve":
			device.Off()
		case "incident.acknowledge":
			device.Off()
		case "incident.trigger":
			device.On()
		}
	}

	return nil
}

func main() {
	lambda.Start(Handler)
}

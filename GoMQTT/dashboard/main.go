// 3) The dashboard (subscriber). Subscribes to "sensors/greenhouse/#"
// using a wildcard to capture both temperature readings and status updates.
// The message handler processes each topic differently — temperature
// readings are displayed, status changes trigger alerts.
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	pahomqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	opts := pahomqtt.NewClientOptions().
		AddBroker("tcp://localhost:1883").
		SetClientID("dashboard-01").
		SetCleanSession(false)

	opts.SetOnConnectHandler(func(c pahomqtt.Client) {
		log.Println("Dashboard connected to broker")

		// Subscribe to all greenhouse sensor topics using wildcard
		token := c.Subscribe("sensors/greenhouse/#", 1, nil)
		token.Wait()
		log.Println("Subscribed to sensors/greenhouse/#")
	})

	opts.SetDefaultPublishHandler(func(c pahomqtt.Client, msg pahomqtt.Message) {
		topic := msg.Topic()
		payload := string(msg.Payload())

		switch topic {
		case "sensors/greenhouse/temperature":
			fmt.Printf("[Temperature] %s C\n", payload)
		case "sensors/greenhouse/status":
			if payload == "offline" {
				fmt.Printf("[ALERT] Sensor went OFFLINE!\n")
			} else {
				fmt.Printf("[Status] Sensor is %s\n", payload)
			}
		default:
			fmt.Printf("[%s] %s\n", topic, payload)
		}
	})

	opts.SetConnectionLostHandler(func(c pahomqtt.Client, err error) {
		log.Printf("Dashboard connection lost: %v", err)
	})

	client := pahomqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	log.Println("Dashboard shutting down...")
	client.Disconnect(250)
}
package main

import (
	"context"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"os/signal"
	"syscall"
	"time"

	pahomqtt "github.com/eclipse/paho.mqtt.golang"
)

func main () {
	opts := pahomqtt.NewClientOptions().
	AddBroker("tcp://localhost:1883").
	SetClientID("sensor-greenhouse-01").
	SetCleanSession(false).
	SetWill(
		"sensors/greenhouse/status", //topic
		"offline", // payload
		1, // Qos 1
		true, // retain
	)

	opts.SetOnConnectHandler(func (c pahomqtt.Client) {
		log.Println("Sensor connected to broker")
	})

	opts.SetConnectionLostHandler(func(c pahomqtt.Client, err error) {
		log.Println("Sensor connection lost: ", err)
	})

	client := pahomqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	client.Publish("sensors/greenhouse/status", 1, true, "online")
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <- ctx.Done():
					return
			case <- ticker.C:
				temp := 20.0 + rand.Float64() * 10.0
				payload := fmt.Sprintf("%.1f", temp)

				token := client.Publish(
					"sensors/greenhouse/temperature",
					1,
					true,
					payload,
				)
				token.Wait()
				log.Printf("Publish: %s C", payload)
			}
		}
	}()

	sig := make(chan os.Signal , 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<- sig


	log.Println("Sensor shutting down")
	cancel()
	client.Disconnect(250)
}

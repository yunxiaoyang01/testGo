package main

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"log"
)
func main()  {
	fmt.Println("Pulsar Consumer")

	//实例化Pulsar client
	client,err := pulsar.NewClient(pulsar.ClientOptions{
		URL:"pulsar://106.13.39.6:6650", // xx.xx.xx.xx代表Pulsar IP
	})

	if err != nil {
		log.Fatal(err)
	}

	//使用client对象实例化consumer
	consumer,err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:"my-topic",
		SubscriptionName:"sub-demo",
	})

	if err != nil {
		log.Fatal(err)
	}

	defer consumer.Close()

	ctx := context.Background()

	//无限循环监听topic
	for {
		msg,err := consumer.Receive(ctx)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Received message : %v",string(msg.Payload()))
		}

		consumer.Ack(msg)

	}
}

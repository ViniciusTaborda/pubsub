package main

import (
	"fmt"
	"pubsub/msg"
	"pubsub/pub"
	"pubsub/sub"
	"pubsub/uuid"
	"time"
)

func main() {
	fmt.Println("Starting PubSub Example...")

	publisher := pub.NewChPublisher(uuid.MustNewUUID())

	exampleTopic := "example-topic"

	subscriber1 := sub.NewChSubscriber(
		exampleTopic,
		uuid.MustNewUUID(),
		make(chan any),
		(time.Second * 5),
	)

	subscriber2 := sub.NewChSubscriber(
		exampleTopic,
		uuid.MustNewUUID(),
		make(chan any),
		(time.Second * 5),
	)

	publisher.CreateTopic(exampleTopic)

	publisher.Subscribe(subscriber1, exampleTopic)
	publisher.Subscribe(subscriber2, exampleTopic)

	go subscriber1.Listen(publisher.GetWaitGroup())
	go subscriber2.Listen(publisher.GetWaitGroup())

	publisher.Publish(&msg.GenericMessageHolder{
		Topic: exampleTopic,
		Body:  "Hello, World",
		Id:    uuid.MustNewUUID(),
	}, exampleTopic)

	publisher.GetWaitGroup().Wait()
}

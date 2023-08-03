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
	exampleTopic2 := "example-topic2"

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
	publisher.CreateTopic(exampleTopic2)

	publisher.Subscribe(subscriber1, exampleTopic)
	publisher.Subscribe(subscriber2, exampleTopic2)

	go subscriber1.Listen(publisher.GetWaitGroup())
	go subscriber2.Listen(publisher.GetWaitGroup())

	publisher.Publish(&msg.GenericMessageHolder{
		Topic: exampleTopic,
		Body:  "Hello, World",
		Id:    uuid.MustNewUUID(),
	}, exampleTopic)

	publisher.Publish(&msg.GenericMessageHolder{
		Topic: exampleTopic2,
		Body:  "Hello, World 2",
		Id:    uuid.MustNewUUID(),
	}, exampleTopic2)

	publisher.GetWaitGroup().Wait()
}

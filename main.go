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
	fmt.Println("Made by: Vinicius Taborda")
	fmt.Println("")

	// Start by Creating the Publisher
	publisher := pub.NewChPublisher(uuid.MustNewUUID())

	//Topics have to be created before hand
	firstTopic := "first-topic"
	secondTopic := "second-topic"

	publisher.CreateTopic(firstTopic)
	publisher.CreateTopic(secondTopic)

	subscriber1 := sub.NewChSubscriber(
		firstTopic,
		uuid.MustNewUUID(),
		make(chan any),
		(time.Second * 5),
		3,
	)
	subscriber2 := sub.NewChSubscriber(
		firstTopic,
		uuid.MustNewUUID(),
		make(chan any),
		(time.Second * 5),
		3,
	)

	//Subscribe each subscriber to a topic
	publisher.Subscribe(subscriber1, firstTopic)
	publisher.Subscribe(subscriber2, secondTopic)

	//Subscribers start to listen for messages
	go subscriber1.Listen(publisher.GetWaitGroup())
	go subscriber2.Listen(publisher.GetWaitGroup())

	//Publish 3 messages, 2 for the first-topic and 1 for the second-topic
	publisher.Publish(&msg.GenericMessageHolder{
		Topic: firstTopic,
		Body:  "Hello, World",
		Id:    uuid.MustNewUUID(),
	}, firstTopic)

	publisher.Publish(&msg.GenericMessageHolder{
		Topic: firstTopic,
		Body:  "Hello, World",
		Id:    uuid.MustNewUUID(),
	}, firstTopic)

	publisher.Publish(&msg.GenericMessageHolder{
		Topic: secondTopic,
		Body:  "Hello, World 2",
		Id:    uuid.MustNewUUID(),
	}, secondTopic)

	// Wait for subscribers to finish listening
	// Console output should show that the it was received messages for first-topic
	// and only one for second topic
	publisher.Done()
}

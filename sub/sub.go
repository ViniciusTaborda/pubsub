package sub

import (
	"fmt"
	"pubsub/msg"
	"sync"
	"time"
)

type ChannelSubscriber struct {
	Topic     string
	Id        string
	MessageCh chan any
	Closed    bool
	TimeOut   time.Duration
}

func NewChSubscriber(topic, id string, messageCh chan any, timeout time.Duration) Subscriber {

	return &ChannelSubscriber{
		Topic:     topic,
		Id:        id,
		MessageCh: messageCh,
		TimeOut:   timeout,
	}

}

func (cs *ChannelSubscriber) Listen(waitGroup *sync.WaitGroup) any {

	defer waitGroup.Done()

	for {
		select {
		case message := <-cs.MessageCh:
			fmt.Println(message.(msg.MessageHolder))
			return message
		case <-cs.GetTimeOut():
			return nil
		}
	}

}

func (cs *ChannelSubscriber) GetTopic() string {

	return cs.Topic
}

func (cs *ChannelSubscriber) GetId() string {

	return cs.Id
}

func (cs *ChannelSubscriber) GetChannel() chan any {

	return cs.MessageCh
}

func (cs *ChannelSubscriber) GetTimeOut() <-chan time.Time {

	return time.After(cs.TimeOut)
}

func (cs *ChannelSubscriber) IsClosed() bool {

	return cs.Closed
}

func (cs *ChannelSubscriber) Close() {
	cs.Closed = true
}

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
	Attempts  int
}

func NewChSubscriber(topic, id string, messageCh chan any, timeout time.Duration, attempts int) Subscriber {

	return &ChannelSubscriber{
		Topic:     topic,
		Id:        id,
		MessageCh: messageCh,
		TimeOut:   timeout,
		Attempts:  attempts,
	}

}

func (cs *ChannelSubscriber) Listen(waitGroup *sync.WaitGroup) any {
	defer waitGroup.Done()

	for i := 1; i < cs.Attempts+1; i++ {
		select {
		case message := <-cs.MessageCh:
			fmt.Printf("Received: %s\n", message.(msg.MessageHolder))
		case <-cs.GetTimeOut():
			fmt.Printf(
				"%s - %d Attempt - Nothing received... \n",
				cs.Id, i)
		}
	}

	return nil

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

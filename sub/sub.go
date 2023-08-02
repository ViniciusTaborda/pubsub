package sub

import (
	"fmt"
	"pubsub/msg"
	"time"
)

type Subscriber interface {
	ListenMessages() any
	GetTopic() string
	GetId() string
	GetChannel() chan any
	GetTimeOut() time.Duration
	IsClosed() bool
	Close()
}

type ChannelSubscriber struct {
	Topic     string
	Id        string
	MessageCh chan any
	Closed    bool
	TimeOut   <-chan time.Time
}

func (cs *ChannelSubscriber) ListenMessages() any {

	for {
		select {
		case message := <-cs.MessageCh:
			fmt.Println(message.(msg.MessageHolder))
			return message
		case <-cs.TimeOut:
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

	return cs.TimeOut
}

func (cs *ChannelSubscriber) IsClosed() bool {

	return cs.Closed
}

func (cs *ChannelSubscriber) Close() {
	cs.Closed = true
}

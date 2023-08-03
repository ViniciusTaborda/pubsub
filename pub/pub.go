package pub

import (
	"fmt"
	"pubsub/errs"
	"pubsub/msg"
	"pubsub/sub"
	"sync"
)

type ChannelPublisher struct {
	Id          string
	Subscribers *sync.Map
	WaitGroup   *sync.WaitGroup
}

func NewChPublisher(id string) Publisher {

	return &ChannelPublisher{
		Id:          id,
		Subscribers: &sync.Map{},
		WaitGroup:   &sync.WaitGroup{},
	}

}

func (cp *ChannelPublisher) Publish(message msg.MessageHolder, topic string) error {

	subs := cp.GetSubsByTopic(topic)

	subsSlice := subs.([]any)

	for _, s := range subsSlice {

		subscriber := s.(sub.Subscriber)

		if subscriber.IsClosed() {
			continue
		}

		//Non-blocking send
		select {
		case subscriber.GetChannel() <- message:
			fmt.Printf("Value sent to %s\n", subscriber.GetId())
		default:
		}
	}

	return nil
}

func (cp *ChannelPublisher) GetSubsByTopic(topic string) any {

	subs, ok := cp.Subscribers.Load(topic)
	if !ok {
		return []any{}
	}

	return subs
}

func (cp *ChannelPublisher) Subscribe(subscriber sub.Subscriber, topic string) error {

	if subscriber.IsClosed() {
		return errs.ErrClosedSubscriber
	}

	subs, ok := cp.Subscribers.Load(topic)
	if !ok {
		return errs.ErrInvalidTopic
	}

	updatedList := append(subs.([]any), subscriber)
	cp.Subscribers.Store(topic, updatedList)

	cp.GetWaitGroup().Add(1)

	return nil
}

func (cp *ChannelPublisher) GetTopics() []string {

	var topics []string
	cp.Subscribers.Range(func(key, value any) bool {
		topics = append(topics, key.(string))
		return true
	})

	return topics

}

func (cp *ChannelPublisher) CreateTopic(topic string) {

	_, ok := cp.Subscribers.Load(topic)
	if !ok {
		cp.Subscribers.Store(topic, []any{})
	}
}

func (cp *ChannelPublisher) GetId() string {

	return cp.Id

}

func (cp *ChannelPublisher) GetWaitGroup() *sync.WaitGroup {

	return cp.WaitGroup

}

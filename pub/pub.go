package pub

import (
	"pubsub/errs"
	"pubsub/msg"
	"pubsub/sub"
	"sync"
)

type Publisher interface {
	Publish(any, string) error
	GetSubsByTopic(string) []sub.Subscriber
	AddTopic(string)
	GetTopics() []string
	GetId() string
	Subscribe(sub.Subscriber, string) error
}

type ChannelPublisher struct {
	Id          string
	Subscribers sync.Map
}

func (cp *ChannelPublisher) Publish(message msg.MessageHolder, topic string) error {

	subs := cp.GetSubsByTopic(topic)

	for _, sub := range subs {

		if sub.IsClosed() {
			continue
		}

		sub.GetChannel() <- message

	}

	return nil
}
func (cp *ChannelPublisher) GetSubsByTopic(topic string) []sub.Subscriber {

	subs, ok := cp.Subscribers.Load(topic)
	if !ok {
		return []sub.Subscriber{}
	}

	subsList, ok := subs.([]sub.Subscriber)
	if !ok {
		return []sub.Subscriber{}
	}

	return subsList
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

func (cp *ChannelPublisher) GetId() string {

	return cp.Id

}

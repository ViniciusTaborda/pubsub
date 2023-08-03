package pub

import (
	"pubsub/msg"
	"pubsub/sub"
	"sync"
)

type Publisher interface {
	Publish(msg.MessageHolder, string) error
	GetSubsByTopic(string) any
	CreateTopic(string)
	GetTopics() []string
	GetId() string
	Subscribe(sub.Subscriber, string) error
	GetWaitGroup() *sync.WaitGroup
	Done()
}

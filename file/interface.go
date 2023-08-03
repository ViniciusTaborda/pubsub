package file

import (
	"pubsub/msg"
)

type MessageWriter interface {
	Write(message msg.MessageHolder, publisherID string, subscriberID string) error
}

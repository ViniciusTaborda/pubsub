package file

import (
	"pubsub/msg"
)

type MessageWriter interface {
	Write(msg.MessageHolder, string, string, string)
}

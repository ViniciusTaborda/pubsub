package file

import (
	"pubsub/msg"
)

type MessageWriterMock struct {
	WriteMock func(msg.MessageHolder, string, string, string)
}

func NewMessageWriterMock() MessageWriter {
	return &MessageWriterMock{}

}

func (mwm *MessageWriterMock) Write(message msg.MessageHolder, publisherID, subscriberID, topic string) {
	mwm.WriteMock(message, publisherID, subscriberID, topic)
}

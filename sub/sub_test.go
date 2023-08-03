package sub

import (
	"pubsub/msg"
	"sync"
	"testing"
	"time"
)

func TestChannelSubscriberFields(t *testing.T) {
	topic := "testTopic"
	id := "testId"
	messageCh := make(chan any)
	timeout := time.Second
	attempts := 3
	cs := NewChSubscriber(topic, id, messageCh, timeout, attempts)

	testCases := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "GetTopic returns correct topic",
			test: func(t *testing.T) {
				if cs.GetTopic() != topic {
					t.Errorf("Expected topic to be %s, got %s", topic, cs.GetTopic())
				}
			},
		},
		{
			name: "GetId returns correct id",
			test: func(t *testing.T) {
				if cs.GetId() != id {
					t.Errorf("Expected id to be %s, got %s", id, cs.GetId())
				}
			},
		},
		{
			name: "GetChannel returns correct channel",
			test: func(t *testing.T) {
				if cs.GetChannel() != messageCh {
					t.Errorf("Expected channel to be %v, got %v", messageCh, cs.GetChannel())
				}
			},
		},
		{
			name: "IsClosed returns correct flag",
			test: func(t *testing.T) {
				if cs.IsClosed() != false {
					t.Errorf("Expected clsoed flag to be %v, got %v", false, cs.IsClosed())
				}
			},
		},
		{
			name: "Close closes the subscriber",
			test: func(t *testing.T) {
				cs.Close()

				if cs.IsClosed() != true {
					t.Errorf("Expected clsoed flag to be %v, got %v", true, cs.IsClosed())
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, tc.test)
	}
}

func TestListenReceivesMessage(t *testing.T) {
	topic := "testTopic"
	id := "testId"
	messageCh := make(chan any)
	timeout := time.Second
	attempts := 3
	cs := NewChSubscriber(topic, id, messageCh, timeout, attempts)

	var wg sync.WaitGroup
	wg.Add(1)
	go cs.Listen(&wg)

	//Replicating the beahaviour of a publisher
	message := &msg.GenericMessageHolder{
		Topic: topic,
		Body:  "test message",
		Id:    id,
	}
	messageCh <- message

	wg.Wait()

}

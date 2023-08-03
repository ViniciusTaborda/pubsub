package pub

import (
	"pubsub/errs"
	"pubsub/msg"
	"pubsub/sub"
	"testing"
	"time"
)

func TestSubscribe(t *testing.T) {
	cp := NewChPublisher("test")
	topic := "testTopic"

	testCases := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "Subscribe with closed subscriber",
			test: func(t *testing.T) {
				subscriber := sub.NewChSubscriber("sub1", topic, make(chan any), time.Hour)

				subscriber.Close()
				err := cp.Subscribe(subscriber, topic)
				if err != errs.ErrClosedSubscriber {
					t.Errorf("Expected error to be ErrClosedSubscriber, got %v", err)
				}
			},
		},
		{
			name: "Subscribe with invalid topic",
			test: func(t *testing.T) {
				subscriber := sub.NewChSubscriber("sub1", topic, make(chan any), time.Hour)

				err := cp.Subscribe(subscriber, topic)
				if err != errs.ErrInvalidTopic {
					t.Errorf("Expected error to be ErrInvalidTopic, got %v", err)
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, tc.test)
	}
}

func TestGetSubsByTopic(t *testing.T) {
	topic := "testTopic"

	testCases := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "No subscribers for topic",
			test: func(t *testing.T) {
				cp := NewChPublisher("test")

				subs := cp.GetSubsByTopic(topic)
				if len(subs.([]any)) != 0 {
					t.Errorf("Expected 0 subscribers, got %d", len(subs.([]any)))
				}
			},
		},
		{
			name: "Adding one subscriber for topic",
			test: func(t *testing.T) {
				cp := NewChPublisher("test")

				subscriber := sub.NewChSubscriber("sub1", topic, make(chan any), time.Millisecond)

				cp.CreateTopic(topic)
				cp.Subscribe(subscriber, topic)

				subs := cp.GetSubsByTopic(topic)
				if len(subs.([]any)) != 1 {
					t.Errorf("Expected 1 subscriber, got %d", len(subs.([]any)))
				}
			},
		},
		{
			name: "Adding multiple subscribers for topic",
			test: func(t *testing.T) {
				cp := NewChPublisher("test")

				subscriber1 := sub.NewChSubscriber("sub1", topic, make(chan any), time.Millisecond)
				subscriber2 := sub.NewChSubscriber("sub2", topic, make(chan any), time.Millisecond)

				cp.CreateTopic(topic)
				cp.Subscribe(subscriber1, topic)
				cp.Subscribe(subscriber2, topic)

				subs := cp.GetSubsByTopic(topic)
				if len(subs.([]any)) != 2 {
					t.Errorf("Expected 2 subscribers, got %d", len(subs.([]any)))
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, tc.test)
	}
}

func TestGetTopics(t *testing.T) {
	testCases := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "No subscribers for topic",
			test: func(t *testing.T) {
				cp := NewChPublisher("test")

				subs := cp.GetTopics()
				if len(subs) != 0 {
					t.Errorf("Expected 0 topics, got %d", len(subs))
				}
			},
		},
		{
			name: "Adding one topic",
			test: func(t *testing.T) {
				cp := NewChPublisher("test")
				topic := "testTopic"

				subscriber := sub.NewChSubscriber("sub1", topic, make(chan any), time.Millisecond)

				cp.CreateTopic(topic)
				cp.Subscribe(subscriber, topic)

				subs := cp.GetTopics()
				if len(subs) != 1 {
					t.Errorf("Expected 1 topic, got %d", len(subs))
				}
			},
		},
		{
			name: "Adding multiple topics",
			test: func(t *testing.T) {
				cp := NewChPublisher("test")
				topic := "testTopic"
				topic2 := "testTopic2"

				subscriber1 := sub.NewChSubscriber("sub1", topic, make(chan any), time.Millisecond)
				subscriber2 := sub.NewChSubscriber("sub2", topic, make(chan any), time.Millisecond)

				cp.CreateTopic(topic)
				cp.CreateTopic(topic2)
				cp.Subscribe(subscriber1, topic)
				cp.Subscribe(subscriber2, topic2)

				subs := cp.GetTopics()
				if len(subs) != 2 {
					t.Errorf("Expected 2 subscribers, got %d", len(subs))
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, tc.test)
	}
}

func TestCreateTopic(t *testing.T) {
	testCases := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "No topics",
			test: func(t *testing.T) {
				cp := NewChPublisher("test")

				subs := cp.GetTopics()
				if len(subs) != 0 {
					t.Errorf("Expected 0 topics, got %d", len(subs))
				}
			},
		},
		{
			name: "Create one topic",
			test: func(t *testing.T) {
				cp := NewChPublisher("test")
				topic := "testTopic"

				cp.CreateTopic(topic)

				subs := cp.GetTopics()
				if len(subs) != 1 {
					t.Errorf("Expected 1 topic, got %d", len(subs))
				}
			},
		},
		{
			name: "Adding multiple topics",
			test: func(t *testing.T) {
				cp := NewChPublisher("test")
				topic := "testTopic"
				topic2 := "testTopic2"

				cp.CreateTopic(topic)
				cp.CreateTopic(topic2)

				subs := cp.GetTopics()
				if len(subs) != 2 {
					t.Errorf("Expected 2 subscribers, got %d", len(subs))
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, tc.test)
	}
}

func TestGetId(t *testing.T) {
	testCases := []struct {
		name string
		id   string
	}{
		{name: "GetId with id 'test1'", id: "test1"},
		{name: "GetId with id empty", id: ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cp := NewChPublisher(tc.id)
			if cp.GetId() != tc.id {
				t.Errorf("Expected id to be %s, got %s", tc.id, cp.GetId())
			}
		})
	}
}

func TestGetWaitGroup(t *testing.T) {

	cp := NewChPublisher("foo")
	if cp.GetWaitGroup() == nil {
		t.Error("Expected wg to be nil, got not nil")
	}
}

func TestPublishToClosedSubscriber(t *testing.T) {
	cp := NewChPublisher("test")
	topic := "testTopic"

	cp.CreateTopic(topic)

	subscriber1 := sub.NewChSubscriber("sub1", topic, make(chan any), time.Second*1)

	cp.Subscribe(subscriber1, topic)

	message := &msg.GenericMessageHolder{Body: "test message"}

	subscriber1.Close()

	cp.Publish(message, topic)

	select {
	case <-subscriber1.GetChannel():
		t.Errorf("Expected to not receive message, got no message")
	default:
	}

}

func TestPublishToSubscriber(t *testing.T) {
	cp := NewChPublisher("test")
	topic := "testTopic"

	cp.CreateTopic(topic)

	subscriber1 := sub.NewChSubscriber("sub1", topic, make(chan any), time.Second*1)

	cp.Subscribe(subscriber1, topic)

	message := &msg.GenericMessageHolder{Body: "test message"}

	go subscriber1.Listen(cp.GetWaitGroup())

	cp.Publish(message, topic)

	cp.GetWaitGroup().Wait()

}

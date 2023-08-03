package sub

import (
	"sync"
	"time"
)

type Subscriber interface {
	Listen(*sync.WaitGroup) any
	GetTopic() string
	GetId() string
	GetChannel() chan any
	GetTimeOut() <-chan time.Time
	IsClosed() bool
	Close()
}

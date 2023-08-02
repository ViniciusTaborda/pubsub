package errs

import "errors"

var (
	ErrClosedSubscriber = errors.New("closed subscriber")
	ErrInvalidTopic     = errors.New("invalid topic")
)

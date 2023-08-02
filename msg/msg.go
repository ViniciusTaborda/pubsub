package msg

import "fmt"

type MessageHolder interface {
	GetBody() any
	String() any
}

type GenericMessageHolder struct {
	topic string
	body  any
	id    string
}

func (gmh *GenericMessageHolder) GetBody() any {
	return gmh.body
}

func (gmh *GenericMessageHolder) String() any {
	return fmt.Sprintf("%s - %s - %s \n", gmh.id, gmh.topic, gmh.body)
}

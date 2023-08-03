package msg

import "fmt"

type GenericMessageHolder struct {
	Topic string
	Body  any
	Id    string
}

func (gmh *GenericMessageHolder) GetBody() any {
	return gmh.Body
}

func (gmh *GenericMessageHolder) String() any {
	return fmt.Sprintf("%s - %s - %s \n", gmh.Id, gmh.Topic, gmh.Body)
}

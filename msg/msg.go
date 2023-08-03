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

func (gmh *GenericMessageHolder) GetStringBody() string {
	return gmh.Body.(string)
}

func (gmh *GenericMessageHolder) String() string {
	return fmt.Sprintf("%s - %s - %s \n", gmh.Id, gmh.Topic, gmh.Body)
}

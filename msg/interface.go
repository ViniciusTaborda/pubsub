package msg

type MessageHolder interface {
	GetBody() any
	GetStringBody() string
	String() string
}

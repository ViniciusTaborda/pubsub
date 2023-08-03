package msg

import (
	"testing"
)

func TestGenericMessageHolder(t *testing.T) {
	testCases := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "GetBody returns the correct body",
			test: func(t *testing.T) {
				message := &GenericMessageHolder{Body: "test body"}
				if message.GetBody() != "test body" {
					t.Errorf("Expected body to be 'test body', got '%v'", message.GetBody())
				}
			},
		},
		{
			name: "GetStringBody returns the correct string body",
			test: func(t *testing.T) {
				message := &GenericMessageHolder{Body: "test body"}
				if message.GetStringBody() != "test body" {
					t.Errorf("Expected string body to be 'test body', got '%s'", message.GetStringBody())
				}
			},
		},
		{
			name: "String returns the correct string representation",
			test: func(t *testing.T) {
				message := &GenericMessageHolder{Id: "1", Topic: "test topic", Body: "test body"}
				expected := "1 - test topic - test body"
				if message.String() != expected {
					t.Errorf("Expected string representation to be '%s', got '%s'", expected, message.String())
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, tc.test)
	}
}
